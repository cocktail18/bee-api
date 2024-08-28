package service

import (
	"context"
	dadasdk "dada/gosdk"
	"database/sql"
	"encoding/json"
	"fmt"
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/sys"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/bitly/go-simplejson"
	"github.com/golang-module/carbon/v2"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math"
	"strings"
	"sync"
	"time"
)

type OrderSrv struct {
	BaseSrv
}

var orderSrvOnce sync.Once
var orderSrvInstance *OrderSrv

func GetOrderSrv() *OrderSrv {
	orderSrvOnce.Do(func() {
		orderSrvInstance = &OrderSrv{}
	})
	return orderSrvInstance
}

type LogisticsItem struct {
	LogisticsId int64
	Amount      decimal.Decimal
	Num         decimal.Decimal
	Weight      decimal.Decimal
}

func (s *OrderSrv) Close(c context.Context, orderId int64, remark string) error {
	var orderInfo model.BeeOrder
	if err := db.GetDB().Where("id = ? and uid = ? and is_deleted = 0", orderId, kit.GetUid(c)).Take(&orderInfo).Error; err != nil {
		return err
	}
	if orderInfo.Status != enum.OrderStatusUnPaid {
		return errors.New("订单状态错误")
	}
	orderCoupons := make([]*model.BeeOrderCoupon, 0)
	if err := db.GetDB().Where("order_id = ?", orderId).Find(&orderCoupons).Error; err != nil {
		return err
	}
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.BeeOrder{}).Where("id = ? and uid = ?", orderId, kit.GetUid(c)).Updates(map[string]interface{}{
			"status":      enum.OrderStatusClose,
			"date_update": time.Now(),
			"remark":      remark,
		}).Error; err != nil {
			return err
		}
		orderLog := &model.BeeOrderLog{
			BaseModel: *kit.GetInsertBaseModel(c),
			OrderId:   orderId,
			Type:      enum.OrderLogTypeClose,
		}
		if err := tx.Create(orderLog).Error; err != nil {
			return err
		}
		//@todo 退换优惠券、积分
		if orderInfo.Status == enum.OrderStatusUnPaid {
			couponIds := lo.Map(orderCoupons, func(item *model.BeeOrderCoupon, _ int) int64 {
				return item.CouponId
			})
			if len(couponIds) > 0 {
				if err := tx.Model(&model.BeeOrderCoupon{}).Where("coupon_id in ? and uid = ? and is_deleted = 0", couponIds, kit.GetUid(c)).Updates(map[string]interface{}{
					"status":      enum.OrderStatusClose,
					"date_update": time.Now(),
				}).Error; err != nil {
					return err
				}
				if err := tx.Model(&model.BeeUserCoupon{}).Where("id in ? and is_deleted = 0", couponIds).Updates(map[string]interface{}{
					"status":      enum.CouponStatusNormal,
					"date_update": time.Now(),
				}).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}

// CreateOrder 生成订单
// token: 2139988c3b3611ef857c0a2c448f23fc
// goodsJsonStr: [{"propertyChildIds":",211245:1686933,211246:1686935,211247:1686938","goodsId":1808569,"number":1,"logisticsType":0,"inviter_id":0,"goodsTimesDay":"","goodsTimesItem":""}]
// remark: 订单备注测试
// peisongType: zq
// isCanHx: true
// shopIdZt: 391592
// shopNameZt: 龙楼店（文昌市龙楼镇）
// extJsonStr: {"联系电话":"18283848592","取餐时间":"09:30"}
func (s *OrderSrv) CreateOrder(c context.Context, ip string, req *proto.CreateOrderReq) (*proto.CreateOrderResp, error) {
	var (
		shopInfo               *model.BeeShopInfo
		userInfo               *model.BeeUser
		err                    error
		amount                 = decimal.Zero //商品总额
		amountTotal            = decimal.Zero //不包括运费、优惠券、积分
		payTotal               = decimal.Zero //包括运费、优惠券、积分
		amountLogistics        = decimal.Zero
		amountLogisticsReal    = decimal.Zero
		amountLogisticsCoupons = decimal.Zero
		amountCoupon           = decimal.Zero
		weightTotal            = decimal.Zero
		amountBalance          = decimal.Zero
		trips                  = decimal.Zero //小费、打包费
		isNeedLogistics        = true
		needPeisong            = true
		queryFeeRes            *proto.QueryDeliverFeeResult
		dateClose              = time.Now().Add(time.Minute * 30)
	)

	if req.PeisongType == "zq" {
		isNeedLogistics = false
		needPeisong = false
	} else if req.PeisongType == "kd" {
		needPeisong = true
		isNeedLogistics = false
	} else if !cast.ToBool(req.Calculate) {
		return nil, errors.New("暂不支持配送")
	}
	_ = needPeisong

	userInfo, err = GetUserSrv().GetUserInfoByUid(c, kit.GetUid(c))
	if err != nil {
		return nil, err
	}
	_ = userInfo

	if req.ShopIdZt > 0 {
		shopInfo, err = GetShopSrv().GetShopInfo(c, req.ShopIdZt, cast.ToFloat64(req.Lat), cast.ToFloat64(req.Lng))
		if err != nil {
			return nil, err
		}
	} else {
		shopList, err := GetShopSrv().List(c, &proto.ListShopReq{
			Latitude:  cast.ToFloat64(req.Lat),
			Longitude: cast.ToFloat64(req.Lng),
		})
		if err != nil {
			return nil, err
		}
		minDistance := math.MaxFloat64
		for _, _info := range shopList {
			info := _info
			if info.Distance < minDistance {
				minDistance = info.Distance
				shopInfo = info
			}
		}
	}

	orderGoodsList := make([]*proto.BeeOrderGoods, 0, 10)
	if err := json.Unmarshal([]byte(req.GoodsJsonStr), &orderGoodsList); err != nil {
		return nil, errors.Wrap(err, "解析商品信息失败")
	}
	uid := kit.GetUid(c)
	userId := kit.GetUserId(c)
	couponIds := make([]int64, 0)
	if "" != req.CouponId {
		couponIds = lo.Map(strings.Split(req.CouponId, ","), func(item string, _ int) int64 {
			return cast.ToInt64(item)
		})
	}

	// 获取收获地址
	var address *model.BeeUserAddress

	// 获取商品
	resp := &proto.CreateOrderResp{}
	orderId := strings.ReplaceAll(carbon.Now().ToShortDateTimeNanoString(), ".", "")
	orderNo := util.GetRandInt64()
	now := carbon.Now()

	var logisticsId2item = make(map[int64]LogisticsItem)
	var goodsNumTotal = int64(0)
	for i, goods := range orderGoodsList {
		goodsInfo, skuInfo, err := GetGoodsSrv().GetGoodsWithSku(c, goods.GoodsID, goods.PropertyChildIds)
		if err != nil {
			return nil, errors.Wrap(err, "获取商品详情失败")
		}
		goodsNumTotal = goodsNumTotal + goods.Number
		if _, ok := logisticsId2item[goodsInfo.LogisticsId]; !ok {
			logisticsId2item[goodsInfo.LogisticsId] = LogisticsItem{
				LogisticsId: goodsInfo.LogisticsId,
				Amount:      decimal.Zero,
				Num:         decimal.Zero,
				Weight:      decimal.Zero,
			}
		}
		logisticsItem := logisticsId2item[goodsInfo.LogisticsId]
		skuAmount := s.callAmount(goods, goodsInfo, skuInfo)
		weight := decimal.Zero
		if !skuInfo.Weight.IsZero() {
			weight = skuInfo.Weight.Mul(decimal.NewFromInt(goods.Number))
		} else {
			weight = goodsInfo.Weight.Mul(decimal.NewFromInt(goods.Number))
		}

		logisticsItem.Amount = logisticsItem.Amount.Add(skuAmount)
		logisticsItem.Weight = logisticsItem.Weight.Add(weight)
		logisticsItem.Num = logisticsItem.Num.Add(decimal.NewFromInt(goods.Number))
		weightTotal = weightTotal.Add(weight)
		amount = amount.Add(skuAmount)
		orderGoodsList[i].OrderID = orderId
		orderGoodsList[i].GoodsIDStr = cast.ToString(goods.GoodsID)
		orderGoodsList[i].GoodsName = goodsInfo.Name
		orderGoodsList[i].Property = skuInfo.PropertyChildNames
		orderGoodsList[i].Pic = goodsInfo.Pic
		orderGoodsList[i].AfterSale = goodsInfo.AfterSale
		orderGoodsList[i].Amount = s.callAmount(goods, goodsInfo, skuInfo)
		orderGoodsList[i].AmountCoupon = decimal.Zero
		orderGoodsList[i].AmountSingle = skuInfo.Price
		orderGoodsList[i].AmountSingleBase = skuInfo.Price
		orderGoodsList[i].CategoryID = goodsInfo.CategoryId
		orderGoodsList[i].CyTableStatus = enum.CyTableStatusConfirm
		orderGoodsList[i].DateAdd = now.ToDateTimeString()
		orderGoodsList[i].FxType = skuInfo.FxType
		orderGoodsList[i].InviterId = goods.InviterId
		orderGoodsList[i].LogisticsType = goods.LogisticsType
		orderGoodsList[i].Number = goods.Number
		orderGoodsList[i].NumberNoFahuo = goods.NumberNoFahuo
		orderGoodsList[i].Pic = goodsInfo.Pic
		orderGoodsList[i].Purchase = false
		orderGoodsList[i].UserID = userId
		orderGoodsList[i].UID = uid
		orderGoodsList[i].Unit = goodsInfo.Unit
		orderGoodsList[i].Weight = weight
	}
	amountTotal = amount

	if address != nil && address.Id > 0 && isNeedLogistics {
		for logisticsId, item := range logisticsId2item {
			goodsLogistic, err := GetFreightTplSrv().GetBeeLogistics(c, logisticsId, address.CityId)
			if err != nil {
				return nil, errors.Wrap(err, "获取配送地址失败")
			}
			_amountLogistics := s.calFreight(c, goodsLogistic, item)
			amountLogistics = amountLogistics.Add(_amountLogistics)
		}
	}

	// 计算配送费
	distance := shopInfo.Distance
	_ = distance

	if needPeisong {
		// 配送是直接使用传过来的地址
		if address == nil {
			address = &model.BeeUserAddress{
				CityId:     req.CityId,
				DistrictId: req.DistrictId,
				ProvinceId: req.ProvinceId,
				Address:    util.IF(req.Address == "", "天上人间888号", req.Address),
				LinkMan:    util.IF(req.LinkMan == "", "张三", req.LinkMan),
				Mobile:     util.IF(req.Mobile == "", "10086", req.Mobile),
				Latitude:   cast.ToFloat64(req.Lat),
				Longitude:  cast.ToFloat64(req.Lng),
			}

			if !cast.ToBool(req.Calculate) {
				if req.Address == "" {
					return nil, errors.New("请填写地址")
				}
				if req.LinkMan == "" {
					return nil, errors.New("请填写联系人")
				}
				if req.Mobile == "" {
					return nil, errors.New("请填写手机号")
				}
			}
		}

		if req.Lat == "" || req.Lng == "" {
			return nil, errors.New("获取地址失败")
		}
		if weightTotal.IsZero() {
			return nil, errors.New("sku重量不能为空")
		}
		// 计算运费
		if address != nil {
			var notifyUrl string
			notifyUrl, err = GetDeliverySrv().GetNotifyUrl(c, enum.DeliveryStrMap[shopInfo.ExpressType])
			if err != nil {
				return nil, errors.Wrap(err, "获取回调地址失败，请检查配送配置")
			}

			queryFeeRes, err = GetDeliverySrv().QueryDeliveryFee(c, enum.DeliveryStrMap[shopInfo.ExpressType], &proto.QueryDeliverFeeReq{
				ShopNo:          shopInfo.DadaShopNo,
				OriginId:        util.GetRandInt64(),
				CargoPrice:      amount,
				IsPrepay:        0,
				ReceiverName:    address.LinkMan,
				ReceiverAddress: address.Address,
				Callback:        notifyUrl,
				CargoWeight:     cast.ToFloat64(weightTotal.StringFixed(2)),
				ReceiverLat:     cast.ToFloat64(req.Lat),
				ReceiverLng:     cast.ToFloat64(req.Lng),
				ReceiverPhone:   address.Mobile,
			})
			if err != nil {
				return nil, errors.Wrap(err, "获取配送费失败，请检查配送配置")
			}
			var deliveryFee decimal.Decimal
			deliveryFee, err = s.calDeliveryFee(c, cast.ToInt64(req.PeisongFeeId), queryFeeRes.Distance, queryFeeRes.DeliverFee)
			if err != nil {
				return nil, errors.Wrap(err, "获取配送费失败")
			}
			amountLogistics = amountLogistics.Add(deliveryFee)
		}

	}
	amountTotal = amountTotal.Add(amountLogistics)
	if req.Trips != "" {
		trips, err = decimal.NewFromString(req.Trips)
		if err != nil {
			return nil, errors.Wrap(err, "小费错误")
		}
		if trips.LessThan(decimal.Zero) {
			return nil, errors.New("小费不能小于0")
		}
		amountTotal = amountTotal.Add(trips)
	}

	amountLogisticsReal = amountLogistics
	// 获取优惠券
	couponList, err := GetCouponSrv().GetMyCouponListByStatus(uid, enum.CouponStatusNormal)
	if err != nil {
		return nil, errors.Wrap(err, "获取优惠券失败")
	}
	var couponInfos = make([]*model.BeeUserCoupon, 0)
	payTotal = amountTotal
	if len(couponIds) > 0 {
		couponInfos, err = GetCouponSrv().GetUserCouponByIds(c, uid, couponIds)
		if err != nil {
			return nil, errors.Wrap(err, "获取优惠券失败")
		}
		//用户自己排序吧
		//sort.Slice(couponInfos, func(i, j int) bool {
		//	if couponInfos[i].OnlyFreight != couponInfos[j].OnlyFreight {
		//		return couponInfos[i].OnlyFreight
		//	}
		//	return couponInfos[i].MoneyHreshold.GreaterThan(couponInfos[j].MoneyHreshold)
		//})
		for _, couponInfo := range couponInfos {
			if !couponInfo.CanUse(payTotal) {
				return nil, errors.New("优惠券暂时不可用")
			}
			if couponInfo.OnlyFreight && amountLogisticsReal.Equal(decimal.Zero) {
				return nil, errors.New("不需要支付运费，优惠券暂时不可用")
			}

			moneyCal := decimal.Zero
			if couponInfo.MoneyType == enum.CouponMoneyTypeFixed {
				// 固定值
				moneyCal = couponInfo.Money
			} else if couponInfo.MoneyType == enum.CouponMoneyTypeRatio {
				if couponInfo.OnlyFreight {
					moneyCal = amountLogistics.Mul(couponInfo.Money)
				} else {
					moneyCal = payTotal.Mul(couponInfo.Money)
				}

				if moneyCal.LessThan(couponInfo.MoneyMin) && !couponInfo.MoneyMin.Equal(decimal.Zero) {
					moneyCal = couponInfo.MoneyMin
				} else if moneyCal.GreaterThan(couponInfo.MoneyMax) && !couponInfo.MoneyMax.Equal(decimal.Zero) {
					moneyCal = couponInfo.MoneyMax
				}
			} else {
				logger.GetLogger().Error("优惠券类型不可用",
					zap.Any("couponId", couponInfo.Id),
					zap.Any("moneyType", couponInfo.MoneyType))
				return nil, errors.New("优惠券暂时不可用")
			}

			if couponInfo.OnlyFreight {
				if moneyCal.GreaterThan(amountLogisticsReal) {
					moneyCal = amountLogisticsReal
				}
				amountLogisticsReal = amountLogisticsReal.Sub(moneyCal)
				amountLogisticsCoupons = amountLogisticsCoupons.Add(moneyCal)
			} else {
				if moneyCal.GreaterThan(payTotal) {
					moneyCal = payTotal
				}
				amountCoupon = amountCoupon.Add(moneyCal)
			}
			payTotal = payTotal.Sub(moneyCal)
		}
	}

	// 计算金额
	resp.GoodsNumber = goodsNumTotal
	resp.Amount = amount
	resp.AmountTotle = amountTotal
	resp.AmountTotleOriginal = amountTotal
	resp.CouponId = couponIds
	resp.CouponUserList = make([]*proto.UserCouponResp, 0, 10)
	resp.AmountReal = payTotal
	resp.AmountLogistics = amountLogistics
	resp.AmountLogisticsCoupons = amountLogisticsCoupons
	resp.AmountLogisticsReal = amountLogisticsReal
	resp.AmountCoupons = amountCoupon

	amountReal := resp.AmountReal //扣除余额的值
	for _, coupon := range couponList {
		if !coupon.CanUse(amount) {
			continue
		}
		resp.CouponUserList = append(resp.CouponUserList, proto.NewUserCouponResp(coupon))
	}
	if cast.ToBool(req.Calculate) {
		return resp, nil
	}
	qudanHao, err := s.getQuDanHao(c, shopInfo, "default")
	if err != nil {
		return nil, err
	}
	// 开始拼装保存数据

	order := &model.BeeOrder{
		BaseModel:         *kit.GetInsertBaseModel(c),
		Amount:            amount,
		AmountCard:        decimal.Zero,
		AmountCoupons:     amountCoupon,
		AmountLogistics:   amountLogistics,
		AmountBalance:     amountBalance,
		AmountReal:        amountReal,
		AmountRefundTotal: decimal.Zero,
		AmountTax:         decimal.Zero,
		AmountTaxGst:      decimal.Zero,
		AmountTaxService:  decimal.Zero,
		AutoDeliverStatus: util.IF(!shopInfo.GoodsNeedCheck && needPeisong, enum.AutoDeliverStatusOn, enum.AutoDeliverStatusOff),
		GoodsNumber:       goodsNumTotal,
		HasRefund:         false,
		HxNumber:          util.GetRandInt64(),
		Ip:                ip,
		IsCanHx:           cast.ToBool(req.IsCanHx),
		IsDelUser:         false,
		IsEnd:             false,
		IsHasBenefit:      false,
		IsNeedLogistics:   isNeedLogistics || needPeisong,
		IsPay:             false,
		IsScoreOrder:      false,
		IsSuccessPingtuan: false,
		OrderNumber:       orderNo,
		OrderType:         enum.OrderTypeNormal,
		Pid:               0,
		Qudanhao:          qudanHao,
		RefundStatus:      enum.OrderRefundStatusNone,
		Remark:            req.Remark,
		Score:             0,
		ScoreDeduction:    0,
		ShopId:            shopInfo.Id,
		ShopIdZt:          req.ShopIdZt,
		ShopNameZt:        req.ShopNameZt,
		ExtJsonStr:        req.ExtJsonStr,
		PeisongType:       req.PeisongType,
		Status:            enum.OrderStatusUnPaid,
		Trips:             trips,
		Type:              enum.OrderTypeNormal,
		Uid:               kit.GetUid(c),
		DateClose:         common.JsonTime(dateClose),
	}
	err = db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		for _, goodsInfo := range orderGoodsList {
			beeOrderGoods := &model.BeeOrderGoods{
				BaseModel:        *kit.GetInsertBaseModel(c),
				AfterSale:        goodsInfo.AfterSale,
				Amount:           goodsInfo.Amount,
				AmountCoupon:     goodsInfo.AmountCoupon,
				AmountSingle:     goodsInfo.AmountSingle,
				AmountSingleBase: goodsInfo.AmountSingleBase,
				BuyRewardEnd:     goodsInfo.BuyRewardEnd,
				CategoryId:       goodsInfo.CategoryID,
				CyTableStatus:    goodsInfo.CyTableStatus,
				FxType:           goodsInfo.FxType,
				GoodsId:          goodsInfo.GoodsID,
				GoodsName:        goodsInfo.GoodsName,
				//GoodsSubName     :goodsInfo.GoodsSubName     ,
				IsScoreOrder:     goodsInfo.IsScoreOrder,
				Number:           goodsInfo.Number,
				NumberNoFahuo:    goodsInfo.NumberNoFahuo,
				OrderId:          order.Id,
				Persion:          goodsInfo.Persion,
				Pic:              goodsInfo.Pic,
				PriceId:          goodsInfo.PriceID,
				Property:         goodsInfo.Property,
				PropertyChildIds: goodsInfo.PropertyChildIds,
				Purchase:         goodsInfo.Purchase,
				RefundStatus:     goodsInfo.RefundStatus,
				Score:            goodsInfo.Score,
				ShopId:           goodsInfo.ShopID,
				Status:           goodsInfo.Status,
				Type:             goodsInfo.Type,
				Uid:              kit.GetUid(c),
				Unit:             goodsInfo.Unit,
				Weight:           goodsInfo.Weight,
			}
			if err := tx.Create(beeOrderGoods).Error; err != nil {
				return err
			}
		}
		if address != nil {
			orderLogistics := &model.BeeOrderLogistics{
				BeeUserAddress: *address,
				OrderId:        order.Id,
				DadaShopNo:     req.DadaShopNo,
				ShopId:         shopInfo.Id,
			}
			orderLogistics.UserId = kit.GetUserId(c)
			if err := tx.Create(orderLogistics).Error; err != nil {
				return err
			}
		}

		orderLog := &model.BeeOrderLog{
			BaseModel: *kit.GetInsertBaseModel(c),
			OrderId:   order.Id,
			Type:      enum.OrderLogTypeCreate,
		}
		if err := tx.Create(orderLog).Error; err != nil {
			return err
		}

		for _, coupon := range couponInfos {
			if err := tx.Create(&model.BeeOrderCoupon{
				BaseModel: *kit.GetInsertBaseModel(c),
				OrderId:   order.Id,
				CouponId:  coupon.Id,
				Uid:       kit.GetUid(c),
			}).Error; err != nil {
				return err
			}

			// 更新用户优惠券状态
			if err := GetCouponSrv().UseCoupon(c, tx, kit.GetUid(c), coupon.Id); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	resp.Id = order.Id
	resp.HxNumber = order.HxNumber
	resp.IsPay = order.IsPay
	resp.OrderNumber = order.OrderNumber
	resp.NearbyCloseMillis = dateClose.UnixMilli()
	resp.Status = order.Status
	s.afterCreateOrder(c, shopInfo, order)
	return resp, nil
}

func (s *OrderSrv) afterCreateOrder(c context.Context, shopInfo *model.BeeShopInfo, orderInfo *model.BeeOrder) {

}

func (s *OrderSrv) getQuDanHao(c context.Context, shopInfo *model.BeeShopInfo, t string) (string, error) {
	var item = &model.BeeOrderQuDanHao{}
	if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Set("gorm:query_option", "FOR UPDATE").Where("shop_id = ? and `type`=?  and is_deleted = 0", shopInfo.Id, t).
			Assign(model.BeeOrderQuDanHao{
				BaseModel: *kit.GetInsertBaseModel(c),
				ShopId:    shopInfo.Id,
				Type:      t,
				Num:       0,
			}).
			Order("id desc").FirstOrCreate(item).Error; err != nil {
			return err
		}
		item.Num++
		if item.Num == 1000 {
			item.Num = 1
		}
		return tx.Save(item).Error
	}); err != nil {
		return "", err
	}

	return cast.ToString(item.Num), nil
}

func (s *OrderSrv) callAmount(goods *proto.BeeOrderGoods, goodsInfo *model.BeeShopGoods, skuInfo *model.BeeShopGoodsSku) decimal.Decimal {
	//@todo 拼团之类的
	return decimal.NewFromInt(goods.Number).Mul(skuInfo.Price)
}

func (s *OrderSrv) calFreight(c context.Context, logistic *proto.GoodsLogistics, item LogisticsItem) decimal.Decimal {
	amount := decimal.Zero
	if len(logistic.Details) <= 0 {
		return amount
	}
	detail := logistic.Details[0]
	switch logistic.FeeType {
	case enum.BeeLogisticsTypeAmount:
		if item.Amount.LessThanOrEqual(detail.FirstNumber) {
			amount = detail.FirstAmount
		} else {
			amount = detail.FirstAmount
			addNum := item.Amount.Div(detail.AddNumber).Ceil()
			amount = amount.Add(detail.AddAmount.Mul(addNum))
		}
	case enum.BeeLogisticsTypeNum:
		if item.Num.LessThanOrEqual(detail.FirstNumber) {
			amount = detail.FirstAmount
		} else {
			amount = detail.FirstAmount
			addNum := item.Num.Div(detail.AddNumber).Ceil()
			amount = amount.Add(detail.AddAmount.Mul(addNum))
		}
	case enum.BeeLogisticsTypeWeight:
		if item.Weight.LessThanOrEqual(detail.FirstNumber) {
			amount = detail.FirstAmount
		} else {
			amount = detail.FirstAmount
			addNum := item.Weight.Div(detail.AddNumber).Ceil()
			amount = amount.Add(detail.AddAmount.Mul(addNum))
		}
	}
	return amount
}
func (s *OrderSrv) List(c context.Context, req *proto.ListOrderReq) (*proto.ListOrderResp, error) {
	var orderList []*model.BeeOrder
	var cnt int64
	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10000
	}

	dbIns := db.GetDB().Where("uid = ?", kit.GetUid(c)).Where("is_deleted = 0 and is_del_user = 0")
	if req.DateAddBegin != "" {
		dbIns = dbIns.Where("date_add >= ?", req.DateAddBegin)
	}
	if req.DateAddEnd != "" {
		dbIns = dbIns.Where("date_add <= ?", req.DateAddEnd)
	}
	if req.Status != nil {
		dbIns = dbIns.Where("status = ?", req.Status)
	}
	if req.StatusBatch != "" {
		dbIns = dbIns.Where("status in ?", strings.Split(req.StatusBatch, ","))
	}
	if err := dbIns.Model(&model.BeeOrder{}).Count(&cnt).Error; err != nil {
		return nil, err
	}
	if err := dbIns.Offset((page - 1) * pageSize).Order("id desc").Limit(pageSize).Find(&orderList).Error; err != nil {
		return nil, err
	}
	orderIds := lo.Map(orderList, func(item *model.BeeOrder, index int) int64 {
		return item.Id
	})
	var orderGoodsList []*model.BeeOrderGoods
	if err := db.GetDB().Where("order_id in ?", orderIds).Find(&orderGoodsList).Error; err != nil {
		return nil, err
	}
	lo.ForEach(orderGoodsList, func(item *model.BeeOrderGoods, _ int) {
		item.Pic = GetDfsSrv().FillFileUrl(item.Pic)
	})
	var orderLogisticsList []*model.BeeOrderLogistics
	if err := db.GetDB().Where("order_id in ?", orderIds).Find(&orderLogisticsList).Error; err != nil {
		return nil, err
	}
	return &proto.ListOrderResp{
		TotalRow:  cnt,
		TotalPage: int64(math.Ceil(float64(cnt / int64(pageSize)))),
		OrderList: lo.Map(orderList, func(item *model.BeeOrder, _ int) *proto.OrderDto {
			return proto.BeeOrder2Dto(item)
		}),
		LogisticsMap: lo.KeyBy(orderLogisticsList, func(item *model.BeeOrderLogistics) string {
			return cast.ToString(item.OrderId)
		}),
		OrderGoodsMap: lo.GroupBy(orderGoodsList, func(item *model.BeeOrderGoods) string {
			return cast.ToString(item.OrderId)
		}),
	}, nil
}

func (s *OrderSrv) afterOrderDeleteOrClose(c context.Context, tx *gorm.DB, orderInfo *model.BeeOrder) error {
	if orderInfo.Status == enum.OrderStatusUnPaid { // 退优惠券
		var orderCoupons = make([]*model.BeeOrderCoupon, 0)
		if err := db.GetDB().Where("order_id = ? and is_deleted = 0", orderInfo.Id).Find(&orderCoupons).Error; err != nil {
			return err
		}
		couponIds := lo.Map(orderCoupons, func(item *model.BeeOrderCoupon, _ int) int64 {
			return item.CouponId
		})
		if len(couponIds) > 0 {
			if err := tx.Model(&model.BeeOrderCoupon{}).Where("coupon_id in ? and uid = ?", couponIds, orderInfo.Uid).Updates(map[string]interface{}{
				"is_deleted": 1,
				"date_delete": sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
			}).Error; err != nil {
				return err
			}
			if err := tx.Model(&model.BeeUserCoupon{}).Where("id in ?", couponIds).Updates(map[string]interface{}{
				"status":      enum.CouponStatusNormal,
				"date_update": time.Now(),
			}).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *OrderSrv) Delete(c context.Context, orderId int64) error {
	uid := kit.GetUid(c)
	var orderInfo model.BeeOrder

	if err := db.GetDB().Where("id = ? and uid = ? and is_deleted = 0 and is_del_user = 0", orderId, uid).Take(&orderInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	if orderInfo.Status == enum.OrderStatusShipped || orderInfo.Status == enum.OrderStatusPaid {
		return errors.New("订单发货中不能删除")
	}
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.BeeOrder{}).Where("id = ? and uid = ? and is_deleted = 0 and is_del_user = 0", orderId, uid).Updates(map[string]interface{}{
			"is_del_user": 1,
			//"is_deleted":  1,
			"date_delete": sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		}).Error; err != nil {
			return err
		}
		if err := s.afterOrderDeleteOrClose(c, tx, &orderInfo); err != nil {
			return err
		}
		return nil
	})
}

func (s *OrderSrv) Delivery(c context.Context, orderId int64) error {
	var orderInfo model.BeeOrder
	uid := kit.GetUid(c)
	if err := db.GetDB().Where("id = ? and uid = ? and is_deleted = 0", orderId, uid).Take(&orderInfo).Error; err != nil {
		return err
	}
	if orderInfo.Status != enum.OrderStatusPaid && orderInfo.Status != enum.OrderStatusShipped {
		return errors.New("订单状态错误")
	}
	return db.GetDB().Where("id = ?", orderInfo.Id).Updates(map[string]interface{}{
		"status":      enum.OrderStatusConfirmShipped,
		"date_update": time.Now(),
	}).Error
}

func (s *OrderSrv) Reputation(c context.Context, req *proto.BeeOrderReputationReq) error {
	orderId := cast.ToInt64(req.OrderId)
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		for _, reputationItem := range req.Reputations {
			item := model.BeeOrderReputation{
				BaseModel:  *kit.GetInsertBaseModel(c),
				Uid:        kit.GetUid(c),
				OrderId:    orderId,
				Reputation: reputationItem.Reputation,
				Remark:     reputationItem.Remark,
				PicsStr:    strings.Join(reputationItem.Pics, ","),
				Pics:       reputationItem.Pics,
			}
			if err := tx.Create(&item).Error; err != nil {
				return err
			}

			if err := tx.Create(&model.BeeOrderLog{
				BaseModel: *kit.GetInsertBaseModel(c),
				OrderId:   orderId,
				Type:      enum.OrderLogTypeReputation,
			}).Error; err != nil {
				return err
			}
		}
		return nil
	})

}

func (s *OrderSrv) Hx(c context.Context, number string) error {
	var order model.BeeOrder
	if err := db.GetDB().Where("hx_number = ? and is_deleted = 0", number).Take(&order).Error; err != nil {
		return err
	}
	if !order.IsCanHx || order.IsEnd || (order.Status != enum.OrderStatusPaid && order.Status != enum.OrderStatusShipped && order.Status != enum.OrderStatusConfirmShipped) {
		return errors.New("订单不能被核销掉")
	}
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.BeeOrderGoods{}).Where("order_id = ?", order.Id).Updates(map[string]interface{}{
			"status":      enum.OrderGoodsStatusConfirmShipped,
			"date_update": time.Now(),
		}).Error; err != nil {
			return err
		}
		if err := tx.Create(&model.BeeOrderLog{
			BaseModel: *kit.GetInsertBaseModel(c),
			OrderId:   order.Id,
			Type:      enum.OrderLogTypeHx,
		}).Error; err != nil {
			return err
		}
		return tx.Model(&model.BeeOrder{}).Where("id = ?", order.Id).Updates(map[string]interface{}{
			"status":      enum.OrderStatusConfirmShipped,
			"date_update": time.Now(),
		}).Error
	})
}

func (s *OrderSrv) afterPaySuccess(c context.Context, tx *gorm.DB, userInfo *model.BeeUser, orderInfo *proto.OrderDto) error {
	if !userInfo.IsVirtual {
		if err := GetUserSrv().IncrUserLevelAmount(c, tx, orderInfo.Uid, orderInfo.AmountReal); err != nil {
			return errors.Wrap(err, "增加用户等级信息失败")
		}
	}
	orderLog := &model.BeeOrderLog{
		BaseModel: *kit.GetInsertBaseModel(c),
		OrderId:   orderInfo.Id,
		Type:      enum.OrderLogTypePay,
	}
	if err := tx.Create(orderLog).Error; err != nil {
		return err
	}
	if err := tx.Create(&model.BeeOrderPrintLog{
		BaseModel: *kit.GetInsertBaseModel(c),
		OrderId:   orderInfo.Id,
		Condition: enum.PrinterConditionAfterPaySuccess,
		Status:    enum.OrderPrinterStatusWaiting,
	}).Error; err != nil {
		return err
	}

	//@todo 积分奖励
	return nil
}

func (s *OrderSrv) PayOrderOffline(c context.Context, orderId string, extraTx ...func(tx *gorm.DB) error) error {
	orderInfo, err := s.GetOrderByOrderId(c, cast.ToInt64(orderId))
	if err != nil {
		return errors.Wrap(err, "获取订单信息失败！")
	}
	if orderInfo.IsPay { //已支付了
		return nil
	}
	if orderInfo.IsDeleted || orderInfo.IsDelUser {
		return errors.New("订单已删除")
	}
	if orderInfo.Status != enum.OrderStatusUnPaid {
		return errors.New("不是未支付状态")
	}
	payLog := &model.BeePayLog{
		BaseModel:  *kit.GetInsertBaseModel(c),
		Money:      orderInfo.AmountReal,
		NextAction: "",
		OrderNo:    util.GetRandInt64(),
		PayGate:    enum.PayGateBalance,
		Remark:     "",
		Status:     enum.PayLogStatusUnPaid,
		Uid:        orderInfo.Uid,
	}
	if err = db.GetDB().Create(payLog).Error; err != nil {
		return err
	}
	return s.PayOrderByBalance(c, "", payLog, orderId, util.GetRandInt64(), orderInfo.AmountReal, extraTx...)
}

// PayOrderByBalance 支付成功
// amount 标志第三方支付金额，如果是用余额支付，请用0
func (s *OrderSrv) PayOrderByBalance(c context.Context, ip string, payLog *model.BeePayLog, orderId string, thirdId string, amount decimal.Decimal, extraTx ...func(tx *gorm.DB) error) error {
	orderInfo, err := s.GetOrderByOrderId(c, cast.ToInt64(orderId))
	if err != nil {
		return errors.Wrap(err, "获取订单信息失败！")
	}
	if orderInfo.IsPay { //已支付了
		return nil
	}
	if orderInfo.IsDeleted || orderInfo.IsDelUser {
		return errors.New("订单已删除")
	}
	if orderInfo.Status != enum.OrderStatusUnPaid {
		return errors.New("不是未支付状态")
	}

	userInfo, err := GetUserSrv().GetUserInfoByUid(c, orderInfo.Uid)
	if err != nil {
		return errors.Wrap(err, "获取用户信息失败")
	}

	// 获取用户余额
	userAmount, err := GetBalanceSrv().GetAmount(c, orderInfo.Uid)
	if err != nil {
		return errors.Wrap(err, "获取用户余额失败")
	}

	if !amount.Equal(orderInfo.AmountReal) && orderInfo.AmountReal.GreaterThan(userAmount.Balance.Add(amount)) {
		return fmt.Errorf("金额不正确，应该为：%v 实际为：%v", orderInfo.AmountReal, userAmount.Balance.Add(amount))
	}
	amountBalance := orderInfo.AmountReal.Sub(amount)
	// 更新支付状态

	err = db.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error
		// 扣除余额
		if amountBalance.GreaterThan(decimal.Zero) {
			_, err = GetBalanceSrv().OperAmountByTx(c, tx, orderInfo.Uid, enum.BalanceTypeBalance, amountBalance.Neg(), "pay"+orderInfo.OrderNumber, "订单支付")
			if err != nil {
				return errors.Wrap(err, "扣除余额失败")
			}
		}

		if err = s.paySuccess(c, tx, orderInfo, userInfo, amount, ip); err != nil {
			return err
		}

		for _, fun := range extraTx {
			if err := fun(tx); err != nil {
				return err
			}
		}
		payLogUpdateRs := db.GetDB().Model(&model.BeePayLog{}).Where("id = ?", payLog.Id).Updates(map[string]interface{}{
			"status":      enum.PayLogStatusPaid,
			"date_update": time.Now(),
		})
		if payLogUpdateRs.Error != nil {
			return errors.Wrap(payLogUpdateRs.Error, "更新支付信息失败")
		}
		if payLogUpdateRs.RowsAffected != 1 {
			return errors.New("更新冲突")
		}
		return nil
	})
	if err == nil {
		// 通知快递
		ctx2 := kit.WithSysUser(context.Background(), kit.GetUserInfo(c))
		go func() {
			if orderInfo.AutoDeliverStatus == enum.AutoDeliverStatusOn {
				logger.GetLogger().Info("支付成功，生成配送订单")
				if err2 := s.createDeliveryByOrderInfo(ctx2, orderInfo.Id); err2 != nil {
					logger.GetLogger().Error("生成配送订单失败", zap.Any("orderId", orderInfo.Id), zap.Error(err2))
				}
			}
		}()
	}
	return err
}

func (s *OrderSrv) paySuccess(c context.Context, tx *gorm.DB, orderInfo *proto.OrderDto, userInfo *model.BeeUser, amountBalance decimal.Decimal, ip string) error {
	var (
		err     error
		payTime = time.Now()
	)
	err = tx.Model(&model.BeeOrder{}).Where("id = ? and status = ?", orderInfo.Id, enum.OrderStatusUnPaid).
		Updates(map[string]interface{}{
			"status":         enum.OrderStatusPaid,
			"amount_balance": amountBalance,
			"is_pay":         true,
			"ip":             ip,
			"date_pay":       payTime,
			"date_update":    payTime,
		}).Error
	if err != nil {
		return errors.Wrap(err, "更新订单信息失败")
	}

	err = tx.Model(&model.BeeOrderGoods{}).Where("order_id = ?", orderInfo.Id).
		Updates(map[string]interface{}{
			"purchase":    true,
			"date_update": payTime,
		}).Error
	if err != nil {
		return errors.Wrap(err, "更新订单信息失败")
	}

	if err = s.afterPaySuccess(c, tx, userInfo, orderInfo); err != nil {
		return err
	}
	return nil
}

func (s *OrderSrv) GetOrderByOrderNo(c context.Context, orderNo string) (*proto.OrderDto, error) {
	item := &model.BeeOrder{}
	if err := db.GetDB().Where("order_number = ?", orderNo).Take(item).Error; err != nil {
		return nil, err
	}
	return proto.BeeOrder2Dto(item), nil
}

func (s *OrderSrv) GetOrderByOrderId(c context.Context, orderId int64) (*proto.OrderDto, error) {
	item := &model.BeeOrder{}
	if err := db.GetDB().Where("id = ?", orderId).Take(item).Error; err != nil {
		return nil, err
	}
	return proto.BeeOrder2Dto(item), nil
}

func (s *OrderSrv) GetOrderDetailByOrderId(c context.Context, orderId int64) (*proto.OrderDetailDto, error) {
	orderInfo, err := s.GetOrderByOrderId(c, orderId)
	if err != nil {
		return nil, err
	}
	//获取地址
	var logistics model.BeeOrderLogistics
	var orderGoods = make([]*model.BeeOrderGoods, 0)
	if err = db.GetDB().Where("order_id = ?", orderInfo.Id).Take(&logistics).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	// 获取商品
	if err = db.GetDB().Where("order_id = ?", orderInfo.Id).Find(&orderGoods).Error; err != nil {
		return nil, err
	}
	return &proto.OrderDetailDto{
		OrderDto:       *orderInfo,
		OrderGoods:     orderGoods,
		OrderLogistics: &logistics,
	}, nil
}

func (s *OrderSrv) PayByBalance(c context.Context, ip, orderId string, code string, pwd string) error {
	orderNumArr := strings.Split(orderId, ",")
	// @todo 短信验证码检查
	amount, err := GetBalanceSrv().GetAmount(c, kit.GetUid(c))
	if err != nil {
		return err
	}
	if amount.Pwd != "" && amount.Pwd != util.Md5(amount.GetPwdEncode(pwd)) {
		return errors.New("密码错误")
	}
	if len(orderNumArr) > 1 {
		return errors.New("目前不支持同时支付多个订单")
	}
	// check余额
	orderInfo, err := s.GetOrderByOrderId(c, cast.ToInt64(orderNumArr[0]))
	if err != nil {
		return errors.Wrap(err, "获取订单信息失败")
	}
	if orderInfo.AmountReal.GreaterThan(amount.Balance) {
		return errors.New("余额不足")
	}
	payLog := &model.BeePayLog{
		BaseModel:  *kit.GetInsertBaseModel(c),
		Money:      orderInfo.AmountReal,
		NextAction: "",
		OrderNo:    util.GetRandInt64(),
		PayGate:    enum.PayGateBalance,
		Remark:     "",
		Status:     enum.PayLogStatusUnPaid,
		Uid:        kit.GetUid(c),
	}
	if err = db.GetDB().Create(payLog).Error; err != nil {
		return err
	}
	return s.PayOrderByBalance(c, ip, payLog, cast.ToString(orderInfo.Id), "balance_"+util.GetRandInt64(), decimal.Zero)
}

func (s *OrderSrv) Detail(c context.Context, orderId int64, hxNumber string) (*proto.GetOrderDetailResp, error) {
	if orderId <= 0 && hxNumber == "" {
		return nil, enum.ErrParamError
	}
	var (
		orderInfo  = &model.BeeOrder{}
		logistics  = &model.BeeOrderLogistics{}
		orderGoods = make([]*model.BeeOrderGoods, 0)
		err        error
		logs       = make([]*model.BeeOrderLog, 0)
	)
	if orderId > 0 {
		err = db.GetDB().Where("id = ? and uid = ?", orderId, kit.GetUid(c)).Take(orderInfo).Error
	} else {
		err = db.GetDB().Where("hx_number = ? and uid = ?", hxNumber, kit.GetUid(c)).Take(orderInfo).Error
	}
	if err != nil {
		return nil, err
	}
	// 获取日志
	if err = db.GetDB().Where("order_id = ?", orderInfo.Id).Find(&logs).Error; err != nil {
		return nil, err
	}
	//获取地址
	if err = db.GetDB().Where("order_id = ?", orderInfo.Id).Take(&logistics).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	// 获取商品
	if err = db.GetDB().Where("order_id = ?", orderInfo.Id).Find(&orderGoods).Error; err != nil {
		return nil, err
	}

	resp := &proto.GetOrderDetailResp{
		OrderInfo: proto.BeeOrder2Dto(orderInfo),
		Goods: lo.Map(orderGoods, func(item *model.BeeOrderGoods, _ int) *proto.BeeOrderGoods {
			return proto.BeeOrderGoodsToDto(item)
		}),
		Logs: logs,
	}
	if logistics.Id > 0 {
		resp.Logistics = logistics
	}
	return resp, nil
}

func (s *OrderSrv) StartDaemon(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(time.Second * 1)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.notifySupplierToDelivery(ctx)
				s.closeOrderExceedDaemon(ctx)
				s.autoConfirmShippedOrderDaemon(ctx)
			}
		}
	}()
}

// 通知第三方进行配送
func (s *OrderSrv) notifySupplierToDelivery(ctx context.Context) {
	items := make([]*model.BeeOrderPeisong, 0)
	if err := db.GetDB().Where("status = ? and is_deleted = 0 and retry_times < 3 and last_retry_unix < ?", enum.OrderPaisongStatusNone,
		time.Now().Add(-1*time.Second*30).Unix()).Find(&items).Error; err != nil {
		return
	}
	lo.ForEach(items, func(item *model.BeeOrderPeisong, _ int) {
		var (
			ctx2 context.Context
			err  error
		)
		defer func() {
			if err != nil {
				errMsg := err.Error()
				if len(errMsg) > 500 {
					errMsg = errMsg[:500] + "..."
				}

				if err2 := db.GetDB().Model(&model.BeeOrderPeisong{}).Where("id = ?", item.Id).Updates(map[string]interface{}{
					"retry_times":     item.RetryTimes + 1,
					"last_retry_unix": time.Now().Unix(),
					"is_cancel":       util.IF(item.RetryTimes+1 >= 3, 1, 0),
					"err_msg":         errMsg,
				}).Where("id = ? and status = ?", item.Id, item.Status).Error; err2 != nil {
					logger.GetLogger().Error("更新订单配送信息失败", zap.Error(err), zap.Any("item", item))
				}
			}
		}()
		ctx2, err = sys.GetUserSrv().GetContext(ctx, item.UserId)
		if err != nil {
			logger.GetLogger().Error("获取系统用户信息失败", zap.Error(err), zap.Any("user_id", item.UserId))
			return
		}

		err = GetDeliverySrv().AddOrder(ctx2, item.Type, item.PeisongOrderId)
		if err == nil {
			logger.GetLogger().Info("通知供应商配送成功", zap.Any("orderId", item.OrderId), zap.Any("peisongOrderId", item.PeisongOrderId))
			err = db.GetDB().Model(&model.BeeOrderPeisong{}).Where("id = ? and status = ?", item.Id, item.Status).Updates(map[string]interface{}{
				"last_retry_unix": time.Now().Unix(),
				"status":          enum.OrderPaisongStatusWaiting,
				"date_update":     time.Now(),
			}).Error
			return
		}
		logger.GetLogger().Error("添加配送订单失败", zap.Error(err), zap.Any("item", item))
		var dadaErr *dadasdk.BussError
		if errors.As(err, &dadaErr) {
			switch dadaErr.Code {
			case 2005, 2011, 2062, 2064: //达达订单不存在,数据异常, 重新下单
				logger.GetLogger().Info("达达订单不存在,重新下单", zap.Any("orderId", item.OrderId))
				err = s.NotifyDelivery(ctx2, item.OrderId)
				if err == nil {
					if err2 := db.GetDB().Model(&model.BeeOrderPeisong{}).Updates(map[string]interface{}{
						"status":          item.RetryTimes + 1,
						"is_cancel":       1,
						"err_msg":         fmt.Sprintf("达达订单不存在:%d", dadaErr.Code),
						"last_retry_unix": time.Now().Unix(),
					}).Where("id = ? and status = ?", item.Id, item.Status).Error; err2 != nil {
						logger.GetLogger().Error("更新订单配送信息失败", zap.Error(err), zap.Any("item", item))
					}
				}
				return
			case 2012, 2050, 2051, 2052, 2065, 2066: //处理中
				logger.GetLogger().Info("达达订单处理中，不需要额外处理", zap.Any("orderId", item.OrderId))
				err = db.GetDB().Model(&model.BeeOrderPeisong{}).Updates(map[string]interface{}{
					"retry_times":     item.RetryTimes + 1,
					"last_retry_unix": time.Now().Unix(),
					"status":          enum.OrderPaisongStatusWaiting,
				}).Where("id = ? and status = ?", item.Id, item.Status).Error
				return
			}
		}
	})
}

func (s *OrderSrv) GetPeisongOrderInfoByPeisongOrderNo(c context.Context, peisongOrderNo string) (*model.BeeOrderPeisong, error) {
	var item model.BeeOrderPeisong
	if err := db.GetDB().Where("peisong_order_no = ?", peisongOrderNo).Take(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
func (s *OrderSrv) GetPeisongOrderInfoById(c context.Context, id int64) (*model.BeeOrderPeisong, error) {
	var item model.BeeOrderPeisong
	if err := db.GetDB().Where("id = ?", id).Take(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *OrderSrv) UpdatePeisongOrderInfoStatus(c context.Context, info *model.BeeOrderPeisong) error {
	return db.GetDB().Model(&model.BeeOrderPeisong{}).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", info.Id).Updates(map[string]interface{}{
			"status":       info.Status,
			"third_status": info.ThirdStatus,
			"date_update":  time.Now(),
		}).Error; err != nil {
			return err
		}
		return nil
	})
}

// NotifyDelivery 通知配送
func (s *OrderSrv) NotifyDelivery(ctx context.Context, orderId int64) error {
	orderInfo, err := s.GetOrderByOrderId(ctx, orderId)
	if err != nil {
		return err
	}
	if orderInfo.Status == enum.OrderStatusClose ||
		orderInfo.Status == enum.OrderStatusConfirmShipped ||
		orderInfo.Status == enum.OrderStatusHadComment {
		return errors.New("订单已结束，不能通知配送")
	}
	// 获取最新一条配送信息
	curOrderPeisong := &model.BeeOrderPeisong{}
	err = db.GetDB().Where("order_id = ? and is_deleted = 0 and is_cancel = 0", orderId).Order("id desc").Take(curOrderPeisong).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 创建单
		return s.createDeliveryByOrderInfo(ctx, orderId)
	} else if err != nil {
		return err
	} else {
		switch curOrderPeisong.Status {
		case enum.OrderPaisongStatusArrive:
			return errors.New("订单已送达")
		case enum.OrderPaisongStatusNone, enum.OrderPaisongStatusWaiting, enum.OrderPaisongStatusPickup,
			enum.OrderPaisongStatusDelivering, enum.OrderPaisongStatusAssign, enum.OrderPaisongStatusReturning:
			return errors.New("配送中，不能重复通知配送")
		default:
			// 可以重新发单
			return s.createDeliveryByOrderInfo(ctx, orderId)
		}
	}
}

func (s *OrderSrv) createDeliveryByOrderInfo(ctx context.Context, orderId int64) error {
	// 获取收货地址
	orderDto, err := s.GetOrderDetailByOrderId(ctx, orderId)
	if err != nil {
		return err
	}
	orderLogistics := orderDto.OrderLogistics
	mobile := orderLogistics.Mobile
	if orderDto.ExtJsonStr != "" {
		extInfo, err := simplejson.NewJson([]byte(orderDto.ExtJsonStr))
		if err != nil {
			return err
		}
		_mobile, err := extInfo.Get("联系电话").String()
		if err == nil && _mobile != "" {
			mobile = _mobile
		}
	}
	shopInfo, err := GetShopSrv().GetShopInfo(ctx, orderDto.ShopId, 0, 0)
	if err != nil {
		return errors.Wrap(err, "获取商店信息失败")
	}
	// 计算运费
	var notifyUrl string
	notifyUrl, err = GetDeliverySrv().GetNotifyUrl(ctx, enum.DeliveryStrMap[shopInfo.ExpressType])
	if err != nil {
		return errors.Wrap(err, "获取回调地址失败")
	}

	weightTotal := decimal.Zero
	for _, orderGoods := range orderDto.OrderGoods {
		weightTotal = weightTotal.Add(orderGoods.Weight)
	}
	if weightTotal.IsZero() {
		weightTotal = decimal.NewFromFloat(0.1)
	}
	peisongOrderNo := util.GetRandInt64()
	amountTotal := orderDto.Amount
	deliveryReq := &proto.QueryDeliverFeeReq{
		ShopNo:          shopInfo.DadaShopNo,
		OriginId:        peisongOrderNo,
		CargoPrice:      amountTotal,
		IsPrepay:        0,
		ReceiverName:    orderLogistics.LinkMan,
		ReceiverAddress: orderLogistics.Address,
		Callback:        notifyUrl,
		CargoWeight:     cast.ToFloat64(weightTotal.StringFixed(2)),
		ReceiverLat:     orderLogistics.Latitude,
		ReceiverLng:     orderLogistics.Longitude,
		ReceiverPhone:   mobile,
		Info:            orderDto.ExtJsonToString(),
		ReceiverTel:     "",
	}
	deliveryFee, err := GetDeliverySrv().QueryDeliveryFee(ctx, enum.DeliveryStrMap[shopInfo.ExpressType], deliveryReq)
	if err != nil {
		return errors.Wrap(err, "获取配送费失败")
	}
	orderPeison := &model.BeeOrderPeisong{
		BaseModel:      *kit.GetInsertBaseModelWithUserId(orderDto.UserId),
		OrderId:        orderId,
		PeisongOrderNo: peisongOrderNo,
		Type:           enum.DeliveryStrMap[shopInfo.ExpressType],
		PeisongOrderId: deliveryFee.DeliveryNo,
		Money:          deliveryFee.DeliverFee,
		RealMoney:      deliveryFee.Fee,
		DeductFee:      decimal.Zero,
		Status:         enum.OrderPaisongStatusNone,
		IsCancel:       false,
		ErrMsg:         "",
		ThirdStatus:    "",
		RetryTimes:     0,
		LastRetryUnix:  0,
		ReqData:        util.ToJsonWithoutErr(deliveryReq, ""),
	}
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(orderPeison).Error; err != nil {
			return err
		}

		if err := tx.Create(&model.BeeOrderLog{
			BaseModel: *kit.GetInsertBaseModelWithUserId(orderDto.UserId),
			OrderId:   orderId,
			Type:      enum.OrderLogTypeDeliver,
			TypeStr:   enum.OrderLogTypeMap[enum.OrderLogTypeDeliver],
			Remark:    "发起配送订单-" + deliveryFee.DeliveryNo,
		}).Error; err != nil {
			return err
		}
		logger.GetLogger().Info("保存配送订单成功", zap.Any("orderId", orderId), zap.Any("deliveryFee", deliveryFee))
		return nil
	})
}

// CancelDelivery 取消配送
func (s *OrderSrv) CancelDelivery(ctx context.Context, peisongId int64, reasonId int, reason string) error {
	item, err := s.GetPeisongOrderInfoById(ctx, peisongId)
	if err != nil {
		return err
	}
	if item.Status == enum.OrderPaisongStatusFinish {
		return errors.New("订单已结束，不能取消配送")
	}
	if item.Status == enum.OrderPaisongStatusArrive {
		return errors.New("订单已送达，不能取消配送")
	}
	if item.Status == enum.OrderPaisongStatusCancel {
		return errors.New("订单已取消配送")
	}
	if item.ReqData == "" {
		return errors.New("请求数据为空，不能重试")
	}
	//@todo 兼容不同配送商
	res, err := GetDeliverySrv().CancelOrder(ctx, item.Type, &proto.CancelDeliverOrderReq{
		OrderId:        item.PeisongOrderNo,
		CancelReasonId: cast.ToInt64(reasonId),
		CancelReason:   reason,
	})

	if err != nil {
		var dadaErr *dadasdk.BussError
		if errors.As(err, &dadaErr) {
			switch dadaErr.Code {
			case 2011:
				//订单不存在
				logger.GetLogger().Warn("达达订单不存在", zap.Any("orderId", item.PeisongOrderId))
				err = nil
				res = &proto.CancelDeliverOrderRes{
					DeductFee: decimal.Zero,
				}
			default:
				return err
			}
		} else {
			return err
		}
	}
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&model.BeeOrderLog{
			BaseModel: *kit.GetInsertBaseModelWithUserId(item.UserId),
			OrderId:   item.OrderId,
			Type:      enum.OrderLogTypeDeliver,
			TypeStr:   enum.OrderLogTypeMap[enum.OrderLogTypeDeliver],
			Remark:    "取消配送-" + item.PeisongOrderId,
		}).Error; err != nil {
			return err
		}
		return tx.Model(&model.BeeOrderPeisong{}).Where("id = ?", item.Id).Updates(map[string]interface{}{
			"is_cancel":   1,
			"date_update": time.Now(),
			"deduct_fee":  res.DeductFee,
		}).Error
	})

}

// 运费计算，注意，distance 跟 total 是第三方返回的距离跟运费
func (s *OrderSrv) calDeliveryFee(ctx context.Context, peisongFeedId int64, distance float64, total decimal.Decimal) (decimal.Decimal, error) {
	peisongConfigList := make([]*model.BeePeiSong, 0)
	if err := db.GetDB().Where("is_deleted = 0 and user_id = ?", kit.GetUserId(ctx)).Order("distance asc").
		Find(&peisongConfigList).Error; err != nil {
		return decimal.Zero, err
	}
	if len(peisongConfigList) == 0 {
		return decimal.Zero, errors.New("配送配置为空")
	}
	for _, item := range peisongConfigList {
		if item.Distance.GreaterThanOrEqual(decimal.NewFromFloat(distance)) || item.Id == peisongFeedId {
			fee := decimal.Zero
			if item.Fwf1Type == enum.BeePeiSongFeeTypeFixed { //固定金额
				fee = fee.Add(item.Fwf1Number)
			} else {
				fee = fee.Add(total.Mul(item.Fwf1Number))
			}
			fee = fee.Add(item.Fwf1Min)
			return fee, nil
		}
	}
	return decimal.Zero, errors.New("超出配送配置范围")
}

func (s *OrderSrv) autoConfirmShippedOrderDaemon(ctx context.Context) {
	var orderList []*model.BeeOrder
	if err := db.GetDB().Where("status = ? and date_close < ?", enum.OrderStatusShipped, time.Now()).Limit(100).Find(&orderList).Error; err != nil {
		logger.GetLogger().Error("获取超时未确认的订单失败", zap.Error(err))
		return
	}
	for _, order := range orderList {
		logger.GetLogger().Info("关闭超时未确认的订单", zap.Any("orderId", order.Id))
		if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
			rs := tx.Model(&model.BeeOrder{}).Where("id = ? and status = ?", order.Id, order.Status).Updates(map[string]interface{}{
				"status":      enum.OrderStatusConfirmShipped,
				"is_end":      1,
				"date_update": time.Now(),
			})
			if rs.Error != nil {
				return rs.Error
			}
			if rs.RowsAffected == 0 {
				return errors.New("订单状态已改变")
			}
			if err := tx.Create(&model.BeeOrderLog{
				BaseModel: *kit.GetInsertBaseModelWithUserId(order.UserId),
				OrderId:   order.Id,
				Type:      enum.OrderLogTypeConfirmDeliver,
				TypeStr:   enum.OrderLogTypeMap[enum.OrderLogTypeConfirmDeliver],
				Remark:    "订单超时未确认，自动确认",
			}).Error; err != nil {
				return err
			}
			return nil
		}); err != nil {
			logger.GetLogger().Error("关闭超时未确认的订单失败", zap.Error(err), zap.Any("orderId", order.Id))
		}
	}
}
func (s *OrderSrv) closeOrderExceedDaemon(ctx context.Context) {
	var orderList []*model.BeeOrder
	if err := db.GetDB().Where("status = ? and date_close < ?", enum.OrderStatusUnPaid, time.Now()).Limit(100).Find(&orderList).Error; err != nil {
		logger.GetLogger().Error("获取超时未支付的订单失败", zap.Error(err))
		return
	}
	for _, order := range orderList {
		logger.GetLogger().Info("关闭超时未支付的订单", zap.Any("orderId", order.Id))
		if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
			rs := tx.Model(&model.BeeOrder{}).Where("id = ? and status = ?", order.Id, order.Status).Updates(map[string]interface{}{
				"status":      enum.OrderStatusClose,
				"is_end":      1,
				"date_update": time.Now(),
			})
			if rs.Error != nil {
				return rs.Error
			}
			if rs.RowsAffected == 0 {
				return errors.New("订单状态已改变")
			}
			if err := tx.Create(&model.BeeOrderLog{
				BaseModel: *kit.GetInsertBaseModelWithUserId(order.UserId),
				OrderId:   order.Id,
				Type:      enum.OrderLogTypeClose,
				TypeStr:   enum.OrderLogTypeMap[enum.OrderLogTypeClose],
				Remark:    "订单超时未支付，自动关闭",
			}).Error; err != nil {
				return err
			}

			if err := s.afterOrderDeleteOrClose(ctx, tx, order); err != nil {
				return err
			}
			return nil
		}); err != nil {
			logger.GetLogger().Error("关闭超时未支付的订单失败", zap.Error(err), zap.Any("orderId", order.Id))
		}
	}
}

func (s *OrderSrv) ShippedBeeOrder(ctx context.Context, orderId int64) error {
	// 获取收货地址
	orderDto, err := s.GetOrderDetailByOrderId(ctx, orderId)
	if err != nil {
		return err
	}
	if !orderDto.IsPay || orderDto.Status != enum.OrderStatusPaid {
		return errors.New("订单未支付")
	}
	// 判断是否需要配送
	if orderDto.IsNeedLogistics {
		beeOrderPeisong := &model.BeeOrderPeisong{}
		err = db.GetDB().Where("order_id = ? and is_cancel = 0", orderId).Take(beeOrderPeisong).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := s.createDeliveryByOrderInfo(ctx, orderId); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&model.BeeOrderLog{
			BaseModel: *kit.GetInsertBaseModelWithUserId(orderDto.UserId),
			OrderId:   orderId,
			Type:      enum.OrderLogTypeDeliver,
			TypeStr:   enum.OrderLogTypeMap[enum.OrderLogTypeDeliver],
			Remark:    "订单已发货",
		}).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.BeeOrder{}).Where("id = ? and status = ?", orderId, enum.OrderStatusPaid).Updates(map[string]interface{}{
			"status":      enum.OrderStatusShipped,
			"date_update": time.Now(),
			"date_close":  time.Now().Add(time.Hour * 72), //@todo 配置收货确认时间
		}).Error; err != nil {
			logger.GetLogger().Error("更新订单状态为已发货失败", zap.Error(err), zap.Any("orderId", orderId))
			return err
		}
		return nil
	})
}

func (s *OrderSrv) GetPeisongDetail(ctx context.Context, peisongId int64) (*proto.QueryDeliveryResult, error) {
	var info = &model.BeeOrderPeisong{}
	if err := db.GetDB().Where("id = ?", peisongId).Take(info).Error; err != nil {
		return nil, err
	}
	return GetDeliverySrv().QueryDetail(ctx, info.Type, info.PeisongOrderNo)
}
