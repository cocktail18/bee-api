package service

import (
	"context"
	"fmt"
	"gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/bitly/go-simplejson"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

type PaySrv struct {
	BaseSrv
	clientCache *sync.Map
}

var paySrvOnce sync.Once
var paySrvInstance *PaySrv

func GetPaySrv() *PaySrv {
	paySrvOnce.Do(func() {
		paySrvInstance = &PaySrv{
			clientCache: &sync.Map{},
		}
	})
	return paySrvInstance
}
func (fee *PaySrv) getWxPayNotifyUrl(c context.Context, wxPayConfig *model.BeeWxPayConfig) string {
	host := util.IF(wxPayConfig.NotifyUrl == "", config.GetHost(), wxPayConfig.NotifyUrl)
	return strings.TrimRight(host, "/") + "/" + kit.GetDomain(c) + "/notify/wx/pay"
}

// WxNotify 微信支付回调网关
func (fee *PaySrv) WxNotify(c context.Context, ip string, req *wechat.V3NotifyReq) error {
	if req.EventType != "TRANSACTION.SUCCESS" {
		return errors.New("暂不支持的回调事件！")
	}
	if req.ResourceType != "encrypt-resource" {
		return errors.New("暂不支持的回调事件！")
	}
	// 获取配置
	var wxPayConfig model.BeeWxPayConfig
	if err := db.GetDB().Where("user_id = ? and is_deleted = 0").Take(&wxPayConfig).Error; err != nil {
		return errors.Wrap(err, "获取微信配置失败！")
	}

	payClient, err := fee.GetWechatPayClient(c, &WxPayConfig{
		MchId:           wxPayConfig.MchId,
		Secret:          wxPayConfig.AppSecret,
		Token:           wxPayConfig.Token,
		ReturnUrl:       "",
		NotifyUrl:       fee.getWxPayNotifyUrl(c, &wxPayConfig),
		PrivateCertPath: wxPayConfig.PrivateCert,
		Debug:           wxPayConfig.Debug,
	})
	if err != nil {
		return errors.Wrap(err, "获取微信支付客户端失败！")
	}
	publicMap := payClient.WxPublicKeyMap()
	for _, key := range publicMap {
		logger.GetLogger().Debug("微信回调", zap.Any("key", key))
	}
	logger.GetLogger().Debug("微信回调", zap.Any("req.SignInfo.HeaderSerial", req.SignInfo.HeaderSerial))
	err = req.VerifySignByPKMap(publicMap)
	if err != nil {
		return errors.Wrap(err, "验证微信请求失败！")
	}
	var payResult = &wechat.V3DecryptPayResult{}
	err = wechat.V3DecryptNotifyCipherTextToStruct(req.Resource.Ciphertext, req.Resource.Nonce, req.Resource.AssociatedData, string(payClient.ApiV3Key), payResult)
	if err != nil {
		return errors.Wrap(err, "回调数据解析失败")
	}
	logger.GetLogger().Info("微信回调", zap.Any("payResult", util.ToJsonWithoutErr(payResult, "")))
	if payResult.Mchid != wxPayConfig.MchId {
		return fmt.Errorf("商户号不一致:%v %v！", payResult.Mchid, wxPayConfig.AppId)
	}
	if payResult.TradeState != "SUCCESS" {
		logger.GetLogger().Warn("未付款成功",
			zap.Any("payResult", util.ToJsonWithoutErr(payResult, "")))
		return nil
	}
	if payResult.Amount.Currency != "CNY" {
		return fmt.Errorf("不是使用人民币付款:%v %v！", payResult.OutTradeNo, payResult.Amount.Currency)
	}

	if err = fee.dealPayNotify(c, ip, payResult); err != nil {
		return errors.Wrap(err, "处理支付结果失败")
	}
	return nil
}

func (fee *PaySrv) dealPayNotify(c context.Context, ip string, payResult *wechat.V3DecryptPayResult) error {
	var err error
	var payLog = &model.BeePayLog{}
	var payerTotal = decimal.NewFromInt(int64(payResult.Amount.PayerTotal)).Div(decimal.NewFromInt(100))
	if err := db.GetDB().Where("order_no = ? and is_deleted = 0", payResult.OutTradeNo).Take(payLog).Error; err != nil {
		return errors.Wrap(err, "获取支付记录失败！")
	}
	if payLog.Status != enum.PayLogStatusUnPaid {
		logger.GetLogger().Warn("重复支付",
			zap.Any("payResult", util.ToJsonWithoutErr(payResult, "")),
			zap.Any("payLog", util.ToJsonWithoutErr(payLog, "")),
		)
		return nil
	}
	if payerTotal.LessThan(payLog.Money) {
		return fmt.Errorf("支付金额错误:%v %v 应该为 %v！", payResult.OutTradeNo, payerTotal, payLog.Money)
	}
	nextAction := payLog.NextAction
	var nextActionJson *simplejson.Json
	var payType enum.PayNextActionType
	if len(nextAction) > 0 {
		nextActionJson, err = simplejson.NewJson([]byte(nextAction))
		if err != nil {
			return errors.New("nextAction err")
		}
		payType = enum.PayNextActionType(nextActionJson.Get("type").MustInt())
	} else {
		payType = enum.PayNextActionTypeRecharge
	}

	switch payType {
	case enum.PayNextActionTypeRecharge:
		// 计算充值优惠
		var shouldGetTotal decimal.Decimal
		shouldGetTotal, err = fee.calRechargeGetAmount(c, payerTotal, payLog.Uid)
		if err != nil {
			return errors.Wrap(err, "计算充值优惠失败！")
		}
		_, err = GetBalanceSrv().OperAmount(c, payLog.Uid, enum.BalanceTypeBalance, shouldGetTotal, "recharge_"+payResult.OutTradeNo, "充值", func(tx *gorm.DB) error {
			updateLogRes := tx.Model(&model.BeePayLog{}).Where("id = ?", payLog.Id).Updates(map[string]interface{}{
				"status":         enum.PayLogStatusPaid,
				"third_order_no": payResult.TransactionId,
				"date_update":    time.Now(),
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
		err = GetOrderSrv().PayOrderByBalance(c, ip, payLog, nextActionJson.Get("id").MustString(), payResult.TransactionId, payerTotal)
	case enum.PayNextActionTypePayDirect:
		moneyTotal, err := nextActionJson.Get("money").Float64()
		if err != nil {
			return errors.Wrap(err, "获取金额失败")
		}
		err = fee.payBill(c, decimal.NewFromFloat(moneyTotal), payLog, payResult.OutTradeNo)
	}
	if err != nil {
		return errors.Wrap(err, "处理支付结果失败~")
	}
	return nil
}

func (fee *PaySrv) GetSendRuleByAmount(ctx context.Context, payTotal decimal.Decimal, userId int64) (*model.RechargeSendRule, error) {
	var item = &model.RechargeSendRule{}
	if err := db.GetDB().Where("confine <= ? and is_deleted = 0 and user_id = ?", payTotal, userId).Order("confine desc").Take(item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { //没有优惠
			return nil, nil
		}
		return nil, err
	}
	return item, nil
}

func (fee *PaySrv) RechargeSendRule(c context.Context) ([]*model.RechargeSendRule, error) {
	var list []*model.RechargeSendRule
	if err := db.GetDB().Where("user_id = ?", kit.GetUserId(c)).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (fee *PaySrv) calRechargeGetAmount(ctx context.Context, payTotal decimal.Decimal, userId int64) (decimal.Decimal, error) {
	item, err := fee.GetSendRuleByAmount(ctx, payTotal, userId)
	if err != nil {
		return decimal.Zero, err
	}
	if item == nil {
		return payTotal, nil
	}
	total := payTotal
	if item.Loop {
		total = total.Add(decimal.NewFromInt(total.Div(item.Confine).IntPart()).Mul(item.Send))
	} else {
		total = total.Add(item.Send)
	}
	return total, nil
}

func (fee *PaySrv) payBill(c context.Context, moneyTotal decimal.Decimal, payLog *model.BeePayLog, transactionId string) error {
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
			"status":         enum.PayLogStatusPaid,
			"date_update":    time.Now(),
			"third_order_no": transactionId,
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
		cfgKey       = util.Md5(util.ToJsonWithoutErr(cfg, ""))
	)
	if c, ok := fee.clientCache.Load(cfgKey); ok {
		return c.(*wechat.ClientV3), nil
	}

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
	if cfg.Debug {
		wechatClient.DebugSwitch = gopay.DebugOn
	}
	fee.clientCache.Store(cfgKey, wechatClient)
	return wechatClient, nil
}

func (fee *PaySrv) GetWxAppPayInfo(c context.Context, money decimal.Decimal, remark string, nextAction string, name string) (*proto.GetWxPayInfoRes, error) {
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

	wxPayClient, err := fee.GetWechatPayClient(c, &WxPayConfig{
		MchId:           wxPayConfig.MchId,
		Secret:          wxPayConfig.AppSecret,
		Token:           wxPayConfig.Token,
		ReturnUrl:       "",
		NotifyUrl:       fee.getWxPayNotifyUrl(c, &wxPayConfig),
		PrivateCertPath: wxPayConfig.PrivateCert,
		Debug:           wxPayConfig.Debug,
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
		"description":  util.IF(remark != "", remark, "订单支付"),
		"notify_url":   fee.getWxPayNotifyUrl(c, &wxPayConfig),
		"amount": map[string]interface{}{
			"total":    money.Mul(decimal.NewFromInt(100)).StringFixed(0),
			"currency": "CNY",
		},
		"time_expire": time.Now().Add(time.Hour * 1).Format(time.RFC3339), // 限制一小时内支付
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
	balance, err := GetBalanceSrv().GetAmount(c, kit.GetUid(c))
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
	MchId           string `json:"mchId"`           //商户号
	Secret          string `json:"secret"`          //秘钥,对应微信支付是, APIv3Key，商户平台获取
	Token           string `json:"token"`           //token, 商户证书的证书序列号
	ReturnUrl       string `json:"returnUrl"`       //返回地址
	NotifyUrl       string `json:"notifyUrl"`       //回调地址
	PrivateCertPath string `json:"privateCertPath"` //商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
	Debug           bool   `json:"debug"`
}
