package service

import (
	"context"
	"fmt"
	"gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/bitly/go-simplejson"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

type PaySrv struct {
	BaseSrv
}

var paySrvOnce sync.Once
var paySrvInstance *PaySrv

func GetPaySrv() *PaySrv {
	paySrvOnce.Do(func() {
		paySrvInstance = &PaySrv{}
	})
	return paySrvInstance
}

// WxNotify 微信支付回调网关
func (fee *PaySrv) WxNotify(c context.Context, ip string, req *wechat.V3NotifyReq) (*proto.WxPayNotifyResp, error) {
	if req.EventType != "TRANSACTION.SUCCESS" {
		return nil, errors.New("暂不支持的回调事件！")
	}
	if req.ResourceType != "encrypt-resource" {
		return nil, errors.New("暂不支持的回调事件！")
	}
	// 获取配置
	var wxPayConfig model.BeeWxPayConfig
	if err := db.GetDB().Where("user_id = ? and is_deleted = 0").Take(&wxPayConfig).Error; err != nil {
		return nil, errors.Wrap(err, "获取微信配置失败！")
	}

	payClient, err := fee.GetWechatPayClient(c, &WxPayConfig{
		MchId:           wxPayConfig.MchId,
		Secret:          wxPayConfig.AppSecret,
		Token:           wxPayConfig.Token,
		ReturnUrl:       "",
		NotifyUrl:       "",
		PrivateCertPath: wxPayConfig.PrivateCert,
	})
	if err != nil {
		return nil, errors.Wrap(err, "获取微信支付客户端失败！")
	}
	publicMap := payClient.WxPublicKeyMap()
	for _, key := range publicMap {
		logrus.Debugf("微信回调 publicMap key:%v", key)
	}
	logrus.Debugf("微信回调 v.SignInfo.HeaderSerial:%v", req.SignInfo.HeaderSerial)
	err = req.VerifySignByPKMap(publicMap)
	if err != nil {
		return nil, errors.Wrap(err, "验证微信请求失败！")
	}
	var payResult = &wechat.V3DecryptPayResult{}
	err = wechat.V3DecryptNotifyCipherTextToStruct(req.Resource.Ciphertext, req.Resource.Nonce, req.Resource.AssociatedData, string(payClient.ApiV3Key), payResult)
	if err != nil {
		return nil, errors.Wrap(err, "回调数据解析失败")
	}
	logrus.Infof("微信支付回调内容：%v", util.ToJsonWithoutErr(payResult, ""))
	if payResult.Mchid != wxPayConfig.MchId {
		return nil, fmt.Errorf("商户号不一致:%v %v！", payResult.Mchid, wxPayConfig.AppId)
	}
	if payResult.TradeState != "SUCCESS" {
		logrus.Infof("未付款成功：%v %v", payResult.OutTradeNo, payResult.TradeState)
		return nil, nil
	}
	if payResult.Amount.Currency != "CNY" {
		return nil, fmt.Errorf("不是使用人民币付款:%v %v！", payResult.OutTradeNo, payResult.Amount.Currency)
	}

	var payLog = &model.BeePayLog{}
	if err := db.GetDB().Where("id = ? and is_deleted = 0", payResult.OutTradeNo).Take(payLog).Error; err != nil {
		return nil, errors.Wrap(err, "获取支付记录失败！")
	}
	if payLog.Status != enum.PayLogStatusUnPaid {
		logrus.Infof("重复支付：%v %v", payResult.OutTradeNo, payLog.Status)
		return nil, nil
	}
	nextAction := payLog.NextAction
	var nextActionJson *simplejson.Json
	var payType enum.PayNextActionType
	if len(nextAction) > 0 {
		nextActionJson, err = simplejson.NewJson([]byte(nextAction))
		if err != nil {
			return nil, errors.New("nextAction err")
		}
		payType = enum.PayNextActionType(nextActionJson.Get("type").MustInt())
	} else {
		payType = enum.PayNextActionTypeRecharge
	}
	switch payType {
	case enum.PayNextActionTypeRecharge:
		_, err = GetBalanceSrv().OperAmount(c, payLog.Uid, enum.BalanceTypeBalance, decimal.NewFromInt(int64(payResult.Amount.PayerTotal)), "recharge_"+payResult.OutTradeNo, "充值", func(tx *gorm.DB) error {
			updateLogRes := tx.Model(&model.BeePayLog{}).Where("id = ?", payLog.Id).Updates(map[string]interface{}{
				"status":      enum.PayLogStatusPaid,
				"date_update": time.Now(),
			})
			if updateLogRes.Error != nil {
				return err
			}
			if updateLogRes.RowsAffected != 1 {
				return errors.New("操作冲突")
			}
			return nil
		})
	case enum.PayNextActionTypePayOrder:
		err = GetOrderSrv().PaySuccess(c, ip, payLog, nextActionJson.Get("id").MustString(), payResult.TransactionId, decimal.NewFromInt(int64(payResult.Amount.PayerTotal)), func(tx *gorm.DB) error {
			updateLogRes := tx.Model(&model.BeePayLog{}).Where("id = ?", payLog.Id).Updates(map[string]interface{}{
				"status":      enum.PayLogStatusPaid,
				"date_update": time.Now(),
			})
			if updateLogRes.Error != nil {
				return err
			}
			if updateLogRes.RowsAffected != 1 {
				return errors.New("操作冲突")
			}
			return nil
		})
	case enum.PayNextActionTypePayDirect:
		moneyTotal, err := nextActionJson.Get("money").Float64()
		if err != nil {
			return nil, errors.Wrap(err, "获取金额失败")
		}
		err = fee.payBill(c, decimal.NewFromFloat(moneyTotal), payLog)
	}
	if err != nil {
		return nil, errors.Wrap(err, "处理支付结果失败~")
	}
	return nil, nil
}

func (fee *PaySrv) payBill(c context.Context, moneyTotal decimal.Decimal, payLog *model.BeePayLog) error {
	//入账
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		// 扣余额
		moneyAmount := moneyTotal.Sub(payLog.Money)
		if moneyAmount.LessThan(decimal.Zero) {
			return errors.New("扣除的余额不能是负数")
		}
		_, err := GetBalanceSrv().OperAmountByTx(c, tx, payLog.Uid, enum.BalanceTypeBalance, moneyAmount.Mul(decimal.NewFromInt(-1)), "pay_direct_"+payLog.OrderNo, payLog.Remark)
		if err != nil {
			return err
		}
		updateLogRes := tx.Model(&model.BeePayLog{}).Where("id = ?", payLog.Id).Updates(map[string]interface{}{
			"status":      enum.PayLogStatusPaid,
			"date_update": time.Now(),
		})
		if updateLogRes.Error != nil {
			return err
		}
		if updateLogRes.RowsAffected != 1 {
			return errors.New("操作冲突")
		}
		return nil
	})
}

func (fee *PaySrv) GetWechatPayClient(ctx context.Context, cfg *WxPayConfig) (*wechat.ClientV3, error) {
	// NewClientV3 初始化微信客户端 v3
	// mchid：商户ID 或者服务商模式的 sp_mchid
	// serialNo：商户证书的证书序列号
	// apiV3Key：apiV3Key，商户平台获取
	// privateKey：私钥 apiclient_key.pem 读取后的内容
	var (
		err          error
		wechatClient *wechat.ClientV3
		privateKey   []byte
	)

	privateKey, err = ioutil.ReadFile(cfg.PrivateCertPath)
	if err != nil {
		return nil, errors.Wrap(err, "读取私钥失败")
	}

	wechatClient, err = wechat.NewClientV3(cfg.MchId, cfg.Token, cfg.Secret, string(privateKey))
	if err != nil {
		return nil, errors.Wrap(err, "初始化微信客户端失败")
	}

	// 设置微信平台API证书和序列号（推荐开启自动验签，无需手动设置证书公钥等信息）
	//client.SetPlatformCert([]byte(""), "")

	// 启用自动同步返回验签，并定时更新微信平台API证书（开启自动验签时，无需单独设置微信平台API证书和序列号）
	err = wechatClient.AutoVerifySign()
	if err != nil {
		return nil, errors.Wrap(err, "初始化微信客户端失败")
	}

	// 打开Debug开关，输出日志，默认是关闭的
	wechatClient.DebugSwitch = gopay.DebugOff
	return wechatClient, nil
}

func (fee *PaySrv) GetWxAppPayInfo(c context.Context, appHost string, money decimal.Decimal, remark string, nextAction string, name string) (*proto.GetWxPayInfoRes, error) {
	// 获取配置
	var wxPayConfig model.BeeWxPayConfig
	var payType = enum.PayNextActionTypeRecharge
	var nextActionJson *simplejson.Json
	var payOrderId string
	var err error
	if err = db.GetDB().Where("user_id = ? and is_deleted = 0", kit.GetUserId(c)).Take(&wxPayConfig).Error; err != nil {
		return nil, errors.Wrap(err, "获取微信配置失败！")
	}

	if len(nextAction) > 0 {
		nextActionJson, err = simplejson.NewJson([]byte(nextAction))
		if err != nil {
			return nil, errors.New("nextAction err")
		}
		payType = enum.PayNextActionType(nextActionJson.Get("type").MustInt())
	} else {
		payType = enum.PayNextActionTypeRecharge
	}
	switch payType {
	case enum.PayNextActionTypePayOrder:
		orderId := nextActionJson.Get("id").MustString()
		orderInfo, err := GetOrderSrv().GetOrderByOrderId(c, cast.ToInt64(orderId))
		if err != nil {
			return nil, err
		}
		if orderInfo.Status != enum.OrderStatusUnPaid {
			return nil, errors.New("订单状态错误")
		}
	}
	payOrderId = util.GetRandInt64()
	notifyUrl := strings.TrimRight(config.AppConfigIns.App.PayNotifyUrl, "/")
	if notifyUrl == "" {
		notifyUrl = "https://" + appHost
	}
	notifyUrl = notifyUrl + "/" + kit.GetDomain(c) + "/notify/wx/pay"
	wxPayClient, err := fee.GetWechatPayClient(c, &WxPayConfig{
		MchId:           wxPayConfig.MchId,
		Secret:          wxPayConfig.AppSecret,
		Token:           wxPayConfig.Token,
		ReturnUrl:       "",
		NotifyUrl:       notifyUrl,
		PrivateCertPath: wxPayConfig.PrivateCert,
	})
	if err != nil {
		return nil, errors.Wrap(err, "获取微信支付客户端失败！")
	}
	userOpenId, err := GetUserSrv().GetUserWxOpenId(c)
	if err != nil {
		return nil, err
	}

	wxResp, err := wxPayClient.V3TransactionJsapi(c, gopay.BodyMap{
		"mchid":        wxPayConfig.MchId,
		"out_trade_no": payOrderId,
		"appid":        wxPayConfig.AppId,
		"description":  "",
		"notify_url":   notifyUrl,
		"amount": map[string]interface{}{
			"total":    money.Mul(decimal.NewFromInt(100)).StringFixed(0),
			"currency": "CNY",
		},
		"payer": map[string]interface{}{
			"openid": userOpenId,
		},
	})
	if err != nil {
		return nil, err
	}
	if wxResp.Code != 0 {
		return nil, errors.New("微信请求失败：" + wxResp.Error)
	}
	beePayLog := &model.BeePayLog{
		BaseModel:  *kit.GetInsertBaseModel(c),
		Money:      money,
		NextAction: nextAction,
		OrderNo:    payOrderId,
		PayGate:    enum.PayGateWXAPP,
		Remark:     "",
		Status:     enum.PayLogStatusUnPaid,
		Uid:        kit.GetUid(c),
	}
	err = db.GetDB().Create(beePayLog).Error
	if err != nil {
		return nil, errors.Wrap(err, "微信下单失败")
	}
	jsapiSignInfo, err := wxPayClient.PaySignOfApplet(wxPayConfig.AppId, wxResp.Response.PrepayId)
	if err != nil {
		return nil, errors.Wrap(err, "获取微信支付签名失败")
	}
	return &proto.GetWxPayInfoRes{
		TimeStamp:  jsapiSignInfo.TimeStamp,
		OutTradeId: cast.ToString(beePayLog.Id),
		Package:    jsapiSignInfo.Package,
		PaySign:    jsapiSignInfo.PaySign,
		Appid:      wxPayConfig.AppId,
		Sign:       jsapiSignInfo.PaySign,
		SignType:   jsapiSignInfo.SignType,
		PrepayId:   payOrderId,
		NonceStr:   jsapiSignInfo.NonceStr,
	}, nil
}

func (fee *PaySrv) PayBill(c context.Context, shopId int64, money decimal.Decimal, pwd string, calculate bool, smsCode string) (*proto.PayBillRes, error) {
	//@todo 检查密码、短信验证码
	//@todo 优惠
	if calculate {
		return &proto.PayBillRes{
			RealMoney: money,
		}, nil
	}
	balance, err := GetBalanceSrv().GetAmount(kit.GetUid(c))
	if err != nil {
		return nil, err
	}
	if balance.Balance.LessThan(money) {
		return nil, errors.New("余额不足")
	}
	balanceNew, err := GetBalanceSrv().OperAmount(c, kit.GetUid(c), enum.BalanceTypeBalance, money, "pay_bill"+util.GetRandInt64(), "优惠买单", func(tx *gorm.DB) error {
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &proto.PayBillRes{
		RealMoney: money,
		Amount:    money,
		Balance:   balanceNew.Balance,
		Freeze:    balanceNew.Freeze,
		Remark:    "优惠买单",
	}, nil
}

type WxPayConfig struct {
	MchId           string `json:"mch_id"`            //商户号
	Secret          string `json:"secret"`            //秘钥,对应微信支付是, APIv3Key，商户平台获取
	Token           string `json:"token"`             //token, 商户证书的证书序列号
	ReturnUrl       string `json:"return_url"`        //返回地址
	NotifyUrl       string `json:"notify_url"`        //回调地址
	PrivateCertPath string `json:"private_cert_path"` //商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
}
