package service

import (
	"context"
	"fmt"
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/spf13/cast"
	"github.com/xhd2015/xgo/runtime/mock"
	"time"

	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOrderSrv_CreateOrder(t *testing.T) {
	config.InitConfig()
	ctx := GetTestContext()
	srv := GetOrderSrv()
	Convey("金额计算测试", t, func() {

		couponVal := float64(4)
		freightTotal := float64(10)
		mock.Patch(srv.calFreight, func(context.Context, *proto.GoodsLogistics, LogisticsItem) decimal.Decimal {
			return decimal.NewFromFloat(freightTotal)
		})

		goodsJsonStr := `[{"propertyChildIds":",211245:1686933,211246:1686935,211247:1686938","goodsId":1808573,"number":2,"logisticsType":0,"inviter_id":0,"goodsTimesDay":"","goodsTimesItem":""}]`
		createOrderReq := &proto.CreateOrderReq{
			PeisongType:  "kd",
			IsCanHx:      "true",
			ShopIdZt:     391592,
			ShopNameZt:   "龙楼店（文昌市龙楼镇）",
			ExtJsonStr:   "{}",
			DadaShopNo:   "4217-1478883",
			Lat:          "23.12463",
			Lng:          "113.36199",
			Calculate:    "true",
			GoodsJsonStr: goodsJsonStr,
		}
		resp, err := srv.CreateOrder(ctx, "127.0.0.1", createOrderReq)
		fmt.Println(util.ToJsonWithoutErr(resp, ""))
		So(err, ShouldBeNil)
		So(resp.AmountLogistics.String(), ShouldEqual, decimal.NewFromFloat(freightTotal).String())
		So(resp.GoodsNumber, ShouldEqual, 2)

		amount := resp.Amount
		amountReal := resp.AmountReal
		amountTotal := resp.AmountTotle
		amountTotalOrigin := resp.AmountTotleOriginal

		// 运费券测试
		userCoupon := NewFixCoupon(ctx, couponVal, true)
		mock.Patch(GetCouponSrv().GetUserCouponByIds, func(c context.Context, userId int64, ids []int64) ([]*model.BeeUserCoupon, error) {
			return []*model.BeeUserCoupon{
				userCoupon,
			}, nil
		})
		createOrderReq.CouponId = cast.ToString(userCoupon.Id)
		resp, err = srv.CreateOrder(ctx, "127.0.0.1", createOrderReq)
		fmt.Println(util.ToJsonWithoutErr(resp, ""))
		So(err, ShouldBeNil)
		So(resp.AmountLogisticsCoupons.String(), ShouldEqual, decimal.NewFromFloat(couponVal).String())
		So(resp.AmountLogisticsReal.String(), ShouldEqual, decimal.NewFromFloat(freightTotal-couponVal).String())
		So(resp.AmountReal.String(), ShouldEqual, amountReal.Sub(decimal.NewFromFloat(couponVal)).String())
		So(resp.Amount.String(), ShouldEqual, amount.String())
		So(resp.AmountTotleOriginal.String(), ShouldEqual, amountTotalOrigin.String())
		So(resp.AmountTotle.String(), ShouldEqual, amountTotal.Sub(decimal.NewFromFloat(couponVal)).String())
		So(resp.AmountTotle.String(), ShouldEqual, resp.Amount.Add(decimal.NewFromFloat(freightTotal)).Sub(decimal.NewFromFloat(couponVal)).String())

		// 普通优惠券测试
		userCoupon = NewFixCoupon(ctx, 4, false)
		mock.Patch(GetCouponSrv().GetUserCouponByIds, func(c context.Context, userId int64, ids []int64) ([]*model.BeeUserCoupon, error) {
			return []*model.BeeUserCoupon{
				userCoupon,
			}, nil
		})
		createOrderReq.CouponId = cast.ToString(userCoupon.Id)
		resp, err = srv.CreateOrder(ctx, "127.0.0.1", createOrderReq)
		So(err, ShouldBeNil)
		So(resp.AmountCoupons.String(), ShouldEqual, decimal.NewFromFloat(couponVal).String())
		So(resp.AmountReal.String(), ShouldEqual, amountReal.Sub(decimal.NewFromFloat(couponVal)).String())
		So(resp.Amount.String(), ShouldEqual, amount.String())
		So(resp.AmountTotleOriginal.String(), ShouldEqual, amountTotalOrigin.String())
		So(resp.AmountTotle.String(), ShouldEqual, amountTotal.Sub(decimal.NewFromFloat(couponVal)).String())
		So(resp.AmountTotle.String(), ShouldEqual, resp.Amount.Add(decimal.NewFromFloat(freightTotal)).Sub(decimal.NewFromFloat(couponVal)).String())

	})
}

func TestOrderSrv_PayOrder(t *testing.T) {
	config.InitConfig()
	ctx := GetTestContext()
	srv := GetOrderSrv()
	Convey("支付订单测试", t, func() {
		goodsJsonStr := `[{"propertyChildIds":",211245:1686933,211246:1686935,211247:1686938","goodsId":1808573,"number":2,"logisticsType":0,"inviter_id":0,"goodsTimesDay":"","goodsTimesItem":""}]`
		createOrderReq := &proto.CreateOrderReq{
			PeisongType:  "kd",
			IsCanHx:      "true",
			ShopIdZt:     391592,
			ShopNameZt:   "龙楼店（文昌市龙楼镇）",
			ExtJsonStr:   "{}",
			DadaShopNo:   "4217-1478883",
			Lat:          "23.12463",
			Lng:          "113.36199",
			Calculate:    "false",
			GoodsJsonStr: goodsJsonStr,
		}
		resp, err := srv.CreateOrder(ctx, "127.0.0.1", createOrderReq)
		So(err, ShouldBeNil)
		err = srv.PaySuccess(ctx, "ip", &model.BeePayLog{}, cast.ToString(resp.OrderId), "", resp.AmountReal)
		So(err, ShouldBeNil)
		orderInfo, err := srv.GetOrderByOrderId(ctx, resp.OrderId)
		So(err, ShouldBeNil)
		So(orderInfo.Status, ShouldEqual, enum.OrderStatusPaid)
		So(orderInfo.IsPay, ShouldEqual, true)
	})
}

func NewFixCoupon(c context.Context, money float64, onlyFreight bool) *model.BeeUserCoupon {
	return &model.BeeUserCoupon{
		BaseModel: common.BaseModel{
			Id: time.Now().UnixMilli(),
		},
		Uid:           kit.GetUid(c),
		DateStart:     common.JsonTime(time.Now()),
		ExpiryMillis:  time.Now().Add(time.Hour * 1).UnixMilli(),
		Money:         decimal.NewFromFloat(money),
		MoneyHreshold: decimal.Zero,
		MoneyMin:      decimal.Zero,
		MoneyMax:      decimal.Zero,
		MoneyType:     enum.CouponMoneyTypeFixed,
		Name:          "test",
		OnlyFreight:   onlyFreight,
		Pid:           0,
		ShopId:        0,
		Source:        "",
		Status:        enum.CouponStatusNormal,
	}
}

func NewRatioCoupon(c context.Context, ratio float64, onlyFreight bool) *model.BeeUserCoupon {
	return &model.BeeUserCoupon{
		BaseModel: common.BaseModel{
			Id: time.Now().UnixMilli(),
		},
		Uid:           kit.GetUid(c),
		DateStart:     common.JsonTime(time.Now()),
		ExpiryMillis:  time.Now().Add(time.Hour * 1).UnixMilli(),
		Money:         decimal.NewFromFloat(ratio),
		MoneyHreshold: decimal.Zero,
		MoneyMin:      decimal.Zero,
		MoneyMax:      decimal.Zero,
		MoneyType:     enum.CouponMoneyTypeRatio,
		Name:          "test",
		OnlyFreight:   onlyFreight,
		Pid:           0,
		ShopId:        0,
		Source:        "",
		Status:        enum.CouponStatusNormal,
	}
}
