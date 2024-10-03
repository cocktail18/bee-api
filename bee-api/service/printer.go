package service

import (
	"bytes"
	"context"
	"encoding/json"
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/printer"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"sync"
	"text/template"
	"time"
)

type PrinterSrv struct {
	BaseSrv
}

var printerSrvOnce sync.Once
var printerSrvInstance *PrinterSrv

func GetPrinterSrv() *PrinterSrv {
	printerSrvOnce.Do(func() {
		printerSrvInstance = &PrinterSrv{}
	})
	return printerSrvInstance
}

func (srv *PrinterSrv) printOrder(ctx context.Context, item *model.BeeOrderPrintLog) error {
	//获取打印机信息
	var printers []*model.BeePrinter
	if err := db.GetDB().Where("user_id = ? and `condition` = ? and is_deleted=0 ", item.UserId, item.Condition).Find(&printers).Error; err != nil {
		return err
	}
	if len(printers) == 0 {
		logger.GetLogger().Debug("未配置打印机")
		return nil
	}
	orderDto, err := GetOrderSrv().GetOrderDetailByOrderId(ctx, item.OrderId)
	if err != nil {
		return err
	}
	userBalance, err := GetBalanceSrv().GetAmount(ctx, orderDto.Uid)
	if err != nil {
		return err
	}

	userInfo, err := GetUserSrv().GetUserInfoByUid(ctx, orderDto.Uid)
	if err != nil {
		return err
	}
	var extJson map[string]interface{}
	if orderDto.ExtJsonStr != "" {
		_ = json.Unmarshal([]byte(orderDto.ExtJsonStr), &extJson)
	}
	var nowStr = time.Now().Format("2006-01-02 15:04:05")
	for _, _printer := range printers {
		if _printer.ShopId > 0 && orderDto.ShopId != _printer.ShopId {
			logger.GetLogger().Debug("订单不属于该打印机", zap.Any("printer", _printer), zap.Any("order", orderDto))
			continue
		}

		content, err := srv.buildContent(ctx, _printer.Template, map[string]interface{}{
			"order":     orderDto,
			"beeOrder":  orderDto.OrderDto,
			"userCash":  userBalance,
			"user":      userInfo,
			"logistics": orderDto.OrderLogistics,
			"goods":     orderDto.OrderGoods,
			"extJson":   extJson,
			"nowStr":    nowStr,
		})
		if err != nil {
			return err
		}
		if err := printer.GetPrinter(_printer).Print(_printer, _printer.Voice, content); err != nil {
			return err
		}
	}
	return nil
}

func (srv *PrinterSrv) TestPrinter(ctx context.Context, _printer *model.BeePrinter) error {
	orderDto := &proto.OrderDetailDto{
		OrderDto: proto.OrderDto{
			BeeOrder: model.BeeOrder{
				BaseModel: common.BaseModel{
					Id:         10000,
					UserId:     1,
					IsDeleted:  false,
					DateAdd:    common.JsonTime(time.Now()),
					DateUpdate: common.JsonTime(time.Now()),
				},
				Amount:            decimal.NewFromFloat(100.00),
				AmountCard:        decimal.NewFromFloat(100.00),
				AmountCoupons:     decimal.NewFromFloat(100.00),
				AmountLogistics:   decimal.NewFromFloat(100.00),
				AmountBalance:     decimal.NewFromFloat(100.00),
				AmountReal:        decimal.NewFromFloat(100.00),
				AmountRefundTotal: decimal.NewFromFloat(100.00),
				AmountTax:         decimal.NewFromFloat(100.00),
				AmountTaxGst:      decimal.NewFromFloat(100.00),
				AmountTaxService:  decimal.NewFromFloat(100.00),
				AutoDeliverStatus: 0,
				DateClose:         common.JsonTime(time.Now()),
				DatePay:           common.JsonTime(time.Now()),
				GoodsNumber:       1,
				HasRefund:         false,
				HxNumber:          util.GetRandInt64(),
				Ip:                "127.0.0.1",
				IsCanHx:           true,
				IsDelUser:         false,
				IsEnd:             true,
				IsHasBenefit:      false,
				IsNeedLogistics:   false,
				IsPay:             true,
				IsScoreOrder:      false,
				IsSuccessPingtuan: false,
				OrderNumber:       util.GetRandInt64(),
				OrderType:         enum.OrderTypeNormal,
				Pid:               0,
				Qudanhao:          "10",
				RefundStatus:      enum.OrderRefundStatusNone,
				Remark:            "备注测试",
				Score:             0,
				ScoreDeduction:    0,
				ShopId:            10,
				ShopIdZt:          10,
				ShopNameZt:        "小蜜蜂店铺",
				Status:            0,
				Trips:             decimal.NewFromFloat(10.00),
				Type:              0,
				Uid:               0,
				ExtJsonStr:        "",
			},
			DifferHours: 0,
		},
		OrderGoods: []*model.BeeOrderGoods{
			&model.BeeOrderGoods{
				BaseModel:        common.BaseModel{},
				AfterSale:        "",
				HadPayAmount:     decimal.Decimal{},
				Amount:           decimal.NewFromInt(20),
				AmountCoupon:     decimal.NewFromInt(10),
				AmountSingle:     decimal.NewFromInt(20),
				AmountSingleBase: decimal.NewFromInt(20),
				BuyRewardEnd:     false,
				CategoryId:       0,
				CyTableStatus:    0,
				FxType:           0,
				GoodsId:          0,
				GoodsName:        "商品名",
				GoodsSubName:     "",
				IsScoreOrder:     false,
				Number:           1,
				NumberNoFahuo:    10,
				OrderId:          10000,
				Persion:          0,
				Pic:              "",
				PriceId:          0,
				Property:         "",
				PropertyChildIds: "",
				Purchase:         false,
				RefundStatus:     0,
				Score:            decimal.Decimal{},
				ShopId:           0,
				Status:           0,
				Type:             0,
				Uid:              0,
				Unit:             "",
			},
		},
		OrderLogistics: nil,
	}
	userBalance := &model.BeeUserAmount{
		BaseModel: common.BaseModel{
			Id:         10,
			DateAdd:    common.JsonTime(time.Now()),
			DateUpdate: common.JsonTime(time.Now()),
		},
		Uid:               10,
		Balance:           decimal.NewFromFloat(10.00),
		Freeze:            decimal.NewFromFloat(10.00),
		FxCommisionPaying: decimal.NewFromFloat(10.00),
		Growth:            decimal.NewFromFloat(10.00),
		Score:             decimal.NewFromFloat(10.00),
		ScoreLastRound:    decimal.NewFromFloat(10.00),
		TotalPayAmount:    decimal.NewFromFloat(10.00),
		TotalPayNumber:    decimal.NewFromFloat(10.00),
		TotalScore:        decimal.NewFromFloat(10.00),
		TotalWithdraw:     decimal.NewFromFloat(10.00),
		TotalConsumed:     decimal.NewFromFloat(10.00),
	}
	userInfo := &model.BeeUser{
		BaseModel: common.BaseModel{
			Id:         10,
			DateAdd:    common.JsonTime(time.Now()),
			DateUpdate: common.JsonTime(time.Now()),
		},
		ShowUid:    10,
		AvatarUrl:  "",
		Birthday:   "1990-01-01",
		CardNumber: "100",
		City:       "广州市",
		DateLogin:  common.JsonTime(time.Now()),
		Gender:     0,
		IpAdd:      "127.0.0.1",
		IpLogin:    "127.0.0.1",
		Nick:       "mike",
		Province:   "广东省",
		VipLevel:   1,
		Source:     enum.BeeUserSourceWx,
		Status:     enum.BeeUserStatusNormal,
		Mobile:     "10086",
		SessionKey: "",
	}
	content, err := srv.buildContent(ctx, _printer.Template, map[string]interface{}{
		"order":     orderDto,
		"beeOrder":  orderDto.OrderDto,
		"userCash":  userBalance,
		"user":      userInfo,
		"logistics": orderDto.OrderLogistics,
		"goods":     orderDto.OrderGoods,
		"extJson": map[string]interface{}{
			"ext": 1,
		},
		"nowStr": time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return err
	}
	if err := printer.GetPrinter(_printer).Print(_printer, _printer.Voice, content); err != nil {
		return err
	}
	return nil
}

func (srv *PrinterSrv) buildContent(ctx context.Context, tpl string, data any) (string, error) {
	t := template.New("tmp")
	tplInstance, err := t.Parse(tpl)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)
	if err := tplInstance.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (srv *PrinterSrv) StartDaemon(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				var items []*model.BeeOrderPrintLog
				if err := db.GetDB().Where("status = ? and is_deleted=0 and last_try_unix < ?",
					enum.OrderPrinterStatusWaiting, time.Now().Add(time.Second*-10).Unix()).Find(&items).Error; err != nil {
					logger.GetLogger().Error("查询待打印订单失败", zap.Error(err))
					continue
				}
				for _, item := range items {
					ctx := context.Background()
					if err := srv.printOrder(ctx, item); err != nil {
						logger.GetLogger().Error("打印订单失败", zap.Error(err))
						db.GetDB().Model(&model.BeeOrderPrintLog{}).Where("id = ?", item.Id).Updates(map[string]interface{}{
							"last_try_unix": time.Now().Unix(),
							"try_times":     item.TryTimes + 1,
							"err_msg":       lo.Substring(err.Error(), 0, 100),
							"status":        util.IF(item.TryTimes+1 >= 3, enum.OrderPrinterStatusErr, enum.OrderPrinterStatusWaiting),
						})
					} else {
						logger.GetLogger().Info("打印订单成功", zap.Any("item", item))
						db.GetDB().Model(&model.BeeOrderPrintLog{}).Where("id = ?", item.Id).Updates(map[string]interface{}{
							"last_try_unix": time.Now().Unix(),
							"try_times":     item.TryTimes + 1,
							"status":        enum.OrderPrinterStatusDone,
						})
					}
				}
			}
		}
	}()
}
