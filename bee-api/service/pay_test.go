package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPaySrv_calRechargeGetAmount(t *testing.T) {
	srv := GetPaySrv()
	ctx := context.Background()
	payTotal := decimal.NewFromInt(10)

	Convey("赠送金额计算测试", t, func() {

		patchers := gomonkey.NewPatches()
		defer patchers.Reset()

		patchers.ApplyMethodSeq(srv, "GetSendRuleByAmount", []gomonkey.OutputCell{
			{
				Values: []interface{}{
					nil, nil,
				},
				Times: 1,
			},
			{
				Values: []interface{}{
					&model.RechargeSendRule{
						Confine: decimal.NewFromInt(10),
						Send:    decimal.NewFromInt(10),
						Loop:    false,
					}, nil,
				},
				Times: 1,
			},

			{
				Values: []interface{}{
					&model.RechargeSendRule{
						Confine: decimal.NewFromInt(10),
						Send:    decimal.NewFromInt(5),
						Loop:    true,
					}, nil,
				},
				Times: 3,
			},
		})

		got, err := srv.calRechargeGetAmount(ctx, payTotal, 100)
		So(err, ShouldBeNil)
		So(got, ShouldEqual, payTotal)

		got, err = srv.calRechargeGetAmount(ctx, decimal.NewFromInt(20), 100)
		So(err, ShouldBeNil)
		So(got.String(), ShouldEqual, decimal.NewFromInt(30).String())

		got, err = srv.calRechargeGetAmount(ctx, decimal.NewFromInt(10), 100)
		So(err, ShouldBeNil)
		So(got.String(), ShouldEqual, decimal.NewFromInt(15).String())

		got, err = srv.calRechargeGetAmount(ctx, decimal.NewFromInt(20), 100)
		So(err, ShouldBeNil)
		So(got.String(), ShouldEqual, decimal.NewFromInt(30).String())

		got, err = srv.calRechargeGetAmount(ctx, decimal.NewFromInt(29), 100)
		So(err, ShouldBeNil)
		So(got.String(), ShouldEqual, decimal.NewFromInt(39).String())
	})

}

func TestPaySrv_dealPayNotify(t *testing.T) {
	config.InitConfig()
	ctx := GetTestContext()
	Convey("支付回调测试", t, func() {
		beePayLog := &model.BeePayLog{
			BaseModel:  *kit.GetInsertBaseModel(ctx),
			Money:      decimal.NewFromInt(10),
			NextAction: "",
			OrderNo:    util.GetRandInt64(),
			PayGate:    enum.PayGateWXAPP,
			Remark:     "",
			Status:     enum.PayLogStatusUnPaid,
			Uid:        kit.GetUid(ctx),
		}
		err := db.GetDB().Create(beePayLog).Error
		So(err, ShouldBeNil)

		// 获取当前余额
		balance, err := GetBalanceSrv().GetAmount(ctx, kit.GetUid(ctx))
		So(err, ShouldBeNil)

		err = GetPaySrv().dealPayNotify(ctx, "127.0.0.1", &wechat.V3DecryptPayResult{
			OutTradeNo:    beePayLog.OrderNo,
			TransactionId: util.GetUUIDStr(),
			Amount: &wechat.Amount{
				Total:         0,
				PayerTotal:    int(beePayLog.Money.Sub(decimal.NewFromInt(10)).Mul(decimal.NewFromInt(100)).IntPart()),
				DiscountTotal: 0,
				Currency:      "",
				PayerCurrency: "",
			},
		})
		So(err, ShouldNotBeNil)

		err = GetPaySrv().dealPayNotify(ctx, "127.0.0.1", &wechat.V3DecryptPayResult{
			OutTradeNo:    beePayLog.OrderNo,
			TransactionId: util.GetUUIDStr(),
			Amount: &wechat.Amount{
				Total:         0,
				PayerTotal:    int(beePayLog.Money.Mul(decimal.NewFromInt(100)).IntPart()),
				DiscountTotal: 0,
				Currency:      "",
				PayerCurrency: "",
			},
		})
		So(err, ShouldBeNil)

		balanceNew, err := GetBalanceSrv().GetAmount(ctx, kit.GetUid(ctx))
		So(err, ShouldBeNil)
		So(balanceNew.Balance.String(), ShouldEqual, balance.Balance.Add(beePayLog.Money).String())
	})
}
