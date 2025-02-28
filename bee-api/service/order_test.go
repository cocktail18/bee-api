package service

import (
	"context"
	"fmt"
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/model/sys"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/spf13/cast"
	"time"

	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOrderSrv_CreateOrderDada(t *testing.T) {
	config.InitConfig()
	ctx := GetTestContext()
	srv := GetOrderSrv()
	patchers := gomonkey.NewPatches()
	defer patchers.Reset()
	patchers.ApplyMethodFunc(GetDeliverySrv(), "GetNotifyUrl", func(ctx context.Context, t enum.DeliveryType) (string, error) {
		return "", nil
	})
	Convey("金额计算测试", t, func() {

		couponVal := float64(2)
		freightTotal := float64(10)
		deliverFee := decimal.NewFromFloat(3.53)
		calFreightPatches := gomonkey.ApplyPrivateMethod(srv, "calFreight", func(context.Context, *proto.GoodsLogistics, LogisticsItem) decimal.Decimal {
			return decimal.NewFromFloat(freightTotal)
		})
		defer calFreightPatches.Reset()
		calDeliverFeePatches := gomonkey.ApplyPrivateMethod(srv, "calDeliveryFee", func(ctx context.Context, distance float64, total decimal.Decimal) (decimal.Decimal, error) {
			return deliverFee, nil
		})
		defer calDeliverFeePatches.Reset()
		//mock.Patch(srv.calFreight, func(context.Context, *proto.GoodsLogistics, LogisticsItem) decimal.Decimal {
		//	return decimal.NewFromFloat(freightTotal)
		//})
		addressPatches := gomonkey.ApplyMethodFunc(GetAddressSrv(), "GetAddressDto", func(c context.Context, userId int64, id int64) (*model.BeeUserAddress, error) {
			return &model.BeeUserAddress{
				BaseModel: common.BaseModel{
					Id: 1,
				},
			}, nil
		})
		defer addressPatches.Reset()

		deliveryPatches := gomonkey.ApplyMethodFunc(GetDeliverySrv(), "QueryDeliveryFee", func(ctx context.Context, t enum.DeliveryType, req *proto.QueryDeliverFeeReq) (*proto.QueryDeliverFeeResult, error) {
			return &proto.QueryDeliverFeeResult{
				DeliverFee: deliverFee,
			}, nil
		})
		defer deliveryPatches.Reset()

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
			DistrictId:   "1",
		}
		resp, err := srv.CreateOrder(ctx, "127.0.0.1", createOrderReq)
		fmt.Println(util.ToJsonWithoutErr(resp, ""))
		So(err, ShouldBeNil)
		So(resp.AmountLogistics.String(), ShouldEqual, deliverFee.String())
		So(resp.GoodsNumber, ShouldEqual, 2)

		amount := resp.Amount
		amountReal := resp.AmountReal
		amountTotalOrigin := resp.AmountTotleOriginal

		// 运费券测试
		userCoupon := NewFixCoupon(ctx, couponVal, true)
		//mock.Patch(GetCouponSrv().GetUserCouponByIds, func(c context.Context, userId int64, ids []int64) ([]*model.BeeUserCoupon, error) {
		//	return []*model.BeeUserCoupon{
		//		userCoupon,
		//	}, nil
		//})
		getUserCouponByIdsPatches := gomonkey.ApplyMethodFunc(GetCouponSrv(), "GetUserCouponByIds", func(c context.Context, userId int64, ids []int64) ([]*model.BeeUserCoupon, error) {
			return []*model.BeeUserCoupon{
				userCoupon,
			}, nil
		})

		createOrderReq.CouponId = cast.ToString(userCoupon.Id)
		resp, err = srv.CreateOrder(ctx, "127.0.0.1", createOrderReq)
		fmt.Println(util.ToJsonWithoutErr(resp, ""))
		So(err, ShouldBeNil)
		So(resp.AmountLogisticsCoupons.String(), ShouldEqual, decimal.NewFromFloat(couponVal).String())
		So(resp.AmountLogisticsReal.String(), ShouldEqual, deliverFee.Sub(decimal.NewFromFloat(couponVal)).String())
		So(resp.AmountReal.String(), ShouldEqual, amountReal.Sub(decimal.NewFromFloat(couponVal)).String())
		So(resp.Amount.String(), ShouldEqual, amount.String())
		So(resp.AmountTotleOriginal.String(), ShouldEqual, amountTotalOrigin.String())

		// 普通优惠券测试
		userCoupon = NewFixCoupon(ctx, couponVal, false)
		//mock.Patch(GetCouponSrv().GetUserCouponByIds, func(c context.Context, userId int64, ids []int64) ([]*model.BeeUserCoupon, error) {
		//	return []*model.BeeUserCoupon{
		//		userCoupon,
		//	}, nil
		//})
		getUserCouponByIdsPatches.Reset()
		getUserCouponByIdsPatches = gomonkey.ApplyMethodFunc(GetCouponSrv(), "GetUserCouponByIds", func(c context.Context, userId int64, ids []int64) ([]*model.BeeUserCoupon, error) {
			return []*model.BeeUserCoupon{
				userCoupon,
			}, nil
		})
		defer getUserCouponByIdsPatches.Reset()
		createOrderReq.CouponId = cast.ToString(userCoupon.Id)
		resp, err = srv.CreateOrder(ctx, "127.0.0.1", createOrderReq)
		So(err, ShouldBeNil)
		So(resp.AmountCoupons.String(), ShouldEqual, decimal.NewFromFloat(couponVal).String())
		So(resp.AmountReal.String(), ShouldEqual, amountReal.Sub(decimal.NewFromFloat(couponVal)).String())
		So(resp.Amount.String(), ShouldEqual, amount.String())
		So(resp.AmountTotleOriginal.String(), ShouldEqual, amountTotalOrigin.String())

	})
}

func TestOrderSrv_CreateOrderZQ(t *testing.T) {
	// 自提订单不需要运费
	config.InitConfig()
	ctx := GetTestContext()
	srv := GetOrderSrv()
	Convey("金额计算测试", t, func() {

		couponVal := float64(4)
		freightTotal := float64(10)
		calFreightPatches := gomonkey.ApplyPrivateMethod(srv, "calFreight", func(context.Context, *proto.GoodsLogistics, LogisticsItem) decimal.Decimal {
			return decimal.NewFromFloat(freightTotal)
		})
		defer calFreightPatches.Reset()
		//mock.Patch(srv.calFreight, func(context.Context, *proto.GoodsLogistics, LogisticsItem) decimal.Decimal {
		//	return decimal.NewFromFloat(freightTotal)
		//})

		goodsJsonStr := `[{"propertyChildIds":",211245:1686933,211246:1686935,211247:1686938","goodsId":1808573,"number":2,"logisticsType":0,"inviter_id":0,"goodsTimesDay":"","goodsTimesItem":""}]`
		createOrderReq := &proto.CreateOrderReq{
			PeisongType:  "zq",
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
		So(resp.AmountLogistics.String(), ShouldEqual, decimal.NewFromFloat(0).String())
		So(resp.GoodsNumber, ShouldEqual, 2)

		amount := resp.Amount
		amountReal := resp.AmountReal
		amountTotalOrigin := resp.AmountTotleOriginal

		// 普通优惠券测试
		userCoupon := NewFixCoupon(ctx, 4, false)
		getUserCouponByIdsPatches := gomonkey.ApplyMethodFunc(GetCouponSrv(), "GetUserCouponByIds", func(c context.Context, userId int64, ids []int64) ([]*model.BeeUserCoupon, error) {
			return []*model.BeeUserCoupon{
				userCoupon,
			}, nil
		})
		defer getUserCouponByIdsPatches.Reset()
		createOrderReq.CouponId = cast.ToString(userCoupon.Id)
		resp, err = srv.CreateOrder(ctx, "127.0.0.1", createOrderReq)
		So(err, ShouldBeNil)
		So(resp.AmountCoupons.String(), ShouldEqual, decimal.NewFromFloat(couponVal).String())
		So(resp.AmountReal.String(), ShouldEqual, amountReal.Sub(decimal.NewFromFloat(couponVal)).String())
		So(resp.Amount.String(), ShouldEqual, amount.String())
		So(resp.AmountTotleOriginal.String(), ShouldEqual, amountTotalOrigin.String())
	})
}

func TestOrderSrv_PayOrder(t *testing.T) {
	config.InitConfig()
	ctx := GetTestContext()
	srv := GetOrderSrv()
	Convey("支付订单测试", t, func() {
		goodsJsonStr := `[{"propertyChildIds":",211245:1686933,211246:1686935,211247:1686938","goodsId":1808573,"number":2,"logisticsType":0,"inviter_id":0,"goodsTimesDay":"","goodsTimesItem":""}]`
		createOrderReq := &proto.CreateOrderReq{
			PeisongType:  "zq",
			IsCanHx:      "true",
			ShopIdZt:     391592,
			ShopNameZt:   "98咖啡（未来社区）",
			ExtJsonStr:   "{}",
			DadaShopNo:   "4217-1478883",
			Lat:          "23.12463",
			Lng:          "113.36199",
			Calculate:    "false",
			GoodsJsonStr: goodsJsonStr,
		}
		resp, err := srv.CreateOrder(ctx, "127.0.0.1", createOrderReq)
		So(err, ShouldBeNil)
		payLog := &model.BeePayLog{
			BaseModel:  common.BaseModel{},
			Money:      decimal.Decimal{},
			NextAction: "",
			OrderNo:    "",
			PayGate:    "",
			PayGateStr: "",
			Remark:     "",
			Status:     0,
			StatusStr:  "",
			Uid:        kit.GetUid(ctx),
		}
		db.GetDB().Create(payLog)
		err = srv.PayOrderByBalance(ctx, "ip", payLog, cast.ToString(resp.Id), util.GetRandInt64(), resp.AmountReal)
		So(err, ShouldBeNil)
		orderInfo, err := srv.GetOrderByOrderId(ctx, resp.Id)
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
func GetTestContext() context.Context {
	ctx := context.Background()
	sysUserInfo := &sys.SysUserModel{}
	db.GetDB().Order("id asc").First(sysUserInfo)
	ctx = context.WithValue(ctx, string(enum.CtxKeySysUser), sysUserInfo)

	userInfo := &model.BeeUser{}
	db.GetDB().Order("id asc").First(userInfo)
	ctx = context.WithValue(ctx, enum.UserInfoKey, userInfo)
	return ctx
}
