package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/golang-module/carbon/v2"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"math"
	"sort"
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

func (s OrderSrv) Close(c context.Context, orderId int64, remark string) error {
	var orderInfo model.BeeOrder
	if err := db.GetDB().Where("id = ? and uid = ?", orderId, kit.GetUid(c)).Take(&orderInfo).Error; err != nil {
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
				if err := tx.Model(&model.BeeOrderCoupon{}).Where("coupon_id in ? and uid = ?", couponIds, kit.GetUid(c)).Updates(map[string]interface{}{
					"status":      enum.OrderStatusClose,
					"date_update": time.Now(),
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
	})
}

// Create 生成订单
// token: 2139988c3b3611ef857c0a2c448f23fc
// goodsJsonStr: [{"propertyChildIds":",211245:1686933,211246:1686935,211247:1686938","goodsId":1808569,"number":1,"logisticsType":0,"inviter_id":0,"goodsTimesDay":"","goodsTimesItem":""}]
// remark: 订单备注测试
// peisongType: zq
// isCanHx: true
// shopIdZt: 391592
// shopNameZt: 龙楼店（文昌市龙楼镇）
// extJsonStr: {"联系电话":"18283848592","取餐时间":"09:30"}
func (s OrderSrv) Create(c context.Context, ip string, req *proto.CreateOrderReq) (*proto.CreateOrderResp, error) {
	var (
		shopInfo *model.BeeShopInfo
		err      error
	)
	needPeisong := true
	if req.PeisongType == "zq" {
		needPeisong = false
	}

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
	var address = &model.BeeUserAddress{}
	if req.DistrictId != "" {
		address, err = GetAddressSrv().GetAddressDto(c, kit.GetUid(c), cast.ToInt64(req.DistrictId))
		if err != nil {
			return nil, errors.Wrap(err, "获取收获地址失败")
		}
	}

	// 获取商品
	resp := &proto.CreateOrderResp{}
	orderId := strings.ReplaceAll(carbon.Now().ToShortDateTimeNanoString(), ".", "")
	now := carbon.Now()
	var amount = decimal.Zero
	var amountTotal = decimal.Zero
	var amountLogistics = decimal.Zero
	var amountCoupon = decimal.Zero
	var weightTotal = decimal.Zero

	type LogisticsItem struct {
		LogisticsId int64
		Amount      decimal.Decimal
		Num         decimal.Decimal
		Weight      decimal.Decimal
	}
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
		weight := skuInfo.Weight.Mul(decimal.NewFromInt(goods.Number))
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
	}
	//@todo 计算运费
	amountTotal = amount
	//logisticsIds := lo.Keys(logisticsId2item)
	//GetFreightTplSrv().GetBeeLogisticsDtoByIds(c, logisticsIds)
	//for _, logisticsItem := range logisticsId2item {
	//
	//}

	// 计算配送费
	distance := shopInfo.Distance
	_ = distance
	if needPeisong {
		peiSong, err := GetFeeSrv().GetPeiSongById(c, cast.ToInt64(req.PeisongFeeId))
		if err != nil {
			return nil, err
		}
		amountLogistics = amountLogistics.Add(peiSong.Fwf1Min)
	}

	// 获取优惠券
	couponList, err := GetCouponSrv().GetMyCouponListByStatus(uid, enum.CouponStatusNormal)
	if err != nil {
		return nil, errors.Wrap(err, "获取优惠券失败")
	}
	var couponInfos = make([]*model.BeeUserCoupon, 0)
	if len(couponIds) > 0 {
		couponInfos, err = GetCouponSrv().GetUserCouponByIds(c, uid, couponIds)
		if err != nil {
			return nil, errors.Wrap(err, "获取优惠券失败")
		}
		sort.Slice(couponInfos, func(i, j int) bool {
			return couponInfos[i].MoneyHreshold.GreaterThan(couponInfos[j].MoneyHreshold)
		})
		for _, couponInfo := range couponInfos {

			if !couponInfo.CanUse(amountTotal.Sub(amountCoupon)) {
				return nil, errors.New("优惠券暂时不可用")
			}
			if couponInfo.OnlyFreight {
				//仅扣运费 @todo
				return nil, errors.New("优惠券暂时不可用")
			}
			if couponInfo.MoneyType == enum.CouponMoneyTypeFixed {
				// 固定值
				amountCoupon.Add(couponInfo.Money)
			} else {
				logrus.Errorf("优惠券【%d】类型不可用", couponInfo.Id)
				return nil, errors.New("优惠券暂时不可用")
			}
		}
	}

	// 计算金额
	resp.GoodsNumber = goodsNumTotal
	resp.AmountTotleOriginal = amount
	resp.AmountTotle = amountTotal.Sub(amountCoupon)
	resp.CouponId = couponIds
	resp.CouponUserList = make([]*proto.UserCouponResp, 0, 10)
	resp.AmountReal = resp.AmountTotle //@todo 扣除余额
	amountReal := resp.AmountReal      //@todo
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
	err = db.GetDB().Transaction(func(tx *gorm.DB) error {

		order := &model.BeeOrder{
			BaseModel:         *kit.GetInsertBaseModel(c),
			Amount:            amount,
			AmountCard:        decimal.Zero,
			AmountCoupons:     amountCoupon,
			AmountLogistics:   amountLogistics,
			AmountBalance:     decimal.Zero, //@todo
			AmountReal:        amountReal,
			AmountRefundTotal: decimal.Zero,
			AmountTax:         decimal.Zero,
			AmountTaxGst:      decimal.Zero,
			AmountTaxService:  decimal.Zero,
			AutoDeliverStatus: 0,
			GoodsNumber:       goodsNumTotal,
			HasRefund:         false,
			HxNumber:          util.GetRandInt64(),
			Ip:                ip,
			IsCanHx:           cast.ToBool(req.IsCanHx),
			IsDelUser:         false,
			IsEnd:             false,
			IsHasBenefit:      false,
			IsNeedLogistics:   true,
			IsPay:             false,
			IsScoreOrder:      false,
			IsSuccessPingtuan: false,
			OrderNumber:       util.GetRandInt64(),
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
			Status:            enum.OrderStatusUnPaid,
			Trips:             decimal.Zero,
			Type:              enum.OrderTypeNormal,
			Uid:               kit.GetUid(c),
		}
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
			}
			if err := tx.Create(beeOrderGoods).Error; err != nil {
				return err
			}
		}
		if needPeisong {
			orderLogistics := &model.BeeOrderLogistics{
				BeeUserAddress: *address,
				OrderId:        order.Id,
			}
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
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s OrderSrv) getQuDanHao(c context.Context, shopInfo *model.BeeShopInfo, t string) (string, error) {
	var item = &model.BeeOrderQuDanHao{}
	if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Set("gorm:query_option", "FOR UPDATE").Where("shop_id = ? and `type`=?", shopInfo.Id, t).
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

func (s OrderSrv) callAmount(goods *proto.BeeOrderGoods, goodsInfo *model.BeeShopGoods, skuInfo *model.BeeShopGoodsSku) decimal.Decimal {
	//@todo 拼团之类的
	return decimal.NewFromInt(goods.Number).Mul(skuInfo.Price)
}

func (s OrderSrv) calFreight(c context.Context, beeLogistics *model.BeeLogistics, address *model.BeeUserAddress,
	amountTotal decimal.Decimal, weight decimal.Decimal, goodsNum decimal.Decimal) decimal.Decimal {
	amount := decimal.Zero
	return amount
	//@todo
	sort.Slice(beeLogistics.Details, func(i, j int) bool {
		if beeLogistics.Details[i].AddNumber != beeLogistics.Details[j].AddNumber {
			return beeLogistics.Details[i].AddNumber.GreaterThan(beeLogistics.Details[j].AddNumber)
		}
		return beeLogistics.Details[i].AddAmount.GreaterThan(beeLogistics.Details[j].AddAmount)
	})
	for _, detail := range beeLogistics.Details {
		if detail.AddNumber.LessThanOrEqual(goodsNum) || detail.AddAmount.LessThanOrEqual(amountTotal) {

		}
	}
	return amount
}
func (s OrderSrv) List(c context.Context, req *proto.ListOrderReq) (*proto.ListOrderResp, error) {
	var orderList []*model.BeeOrder
	var cnt int64
	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	dbIns := db.GetDB().Where("uid = ?", kit.GetUid(c)).Where("is_deleted = 0")
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
	if err := dbIns.Offset((page - 1) * pageSize).Limit(pageSize).Find(&orderList).Error; err != nil {
		return nil, err
	}
	orderIds := lo.Map(orderList, func(item *model.BeeOrder, index int) int64 {
		return item.Id
	})
	var orderGoodsList []*model.BeeOrderGoods
	if err := db.GetDB().Where("order_id in ?", orderIds).Find(&orderGoodsList).Error; err != nil {
		return nil, err
	}
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

func (s OrderSrv) Delete(c context.Context, orderId int64) error {
	uid := kit.GetUid(c)
	var orderInfo model.BeeOrder
	var orderCoupons = make([]*model.BeeOrderCoupon, 0)
	if err := db.GetDB().Where("id = ? and uid = ? and is_deleted = 0", orderId, uid).Take(&orderInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	if err := db.GetDB().Where("order_id = ?", orderId).Find(&orderCoupons).Error; err != nil {
		return err
	}
	//@todo 检查订单状态
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.BeeOrder{}).Where("id = ? and uid = ? and is_deleted = 0", orderId, uid).Updates(map[string]interface{}{
			"is_deleted": 1,
			"date_delete": sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		}).Error; err != nil {
			return err
		}
		if orderInfo.Status == enum.OrderStatusUnPaid {
			couponIds := lo.Map(orderCoupons, func(item *model.BeeOrderCoupon, _ int) int64 {
				return item.CouponId
			})
			if len(couponIds) > 0 {
				if err := tx.Model(&model.BeeOrderCoupon{}).Where("coupon_id in ? and uid = ?", couponIds, uid).Updates(map[string]interface{}{
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
		//@todo 退积分
		return nil
	})
}

func (s OrderSrv) Delivery(c context.Context, orderId int64) error {
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

func (s OrderSrv) Reputation(c context.Context, req *proto.BeeOrderReputationReq) error {
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

func (s OrderSrv) Hx(c context.Context, number string) error {
	var order model.BeeOrder
	if err := db.GetDB().Where("hx_number = ? and is_deleted = 0", number).Take(&order).Error; err != nil {
		return err
	}
	if !order.IsCanHx || order.IsEnd || order.Status != enum.OrderStatusPaid {
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

func (s OrderSrv) PaySuccess(c context.Context, ip string, payLog *model.BeePayLog, orderId string, thirdId string, amount decimal.Decimal, extraTx ...func(tx *gorm.DB) error) error {
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
	if !amount.Equal(orderInfo.AmountReal) {
		return fmt.Errorf("金额不正确，应该为：%v 实际为：%v", orderInfo.AmountReal, amount)
	}

	userInfo, err := GetUserSrv().GetUserInfoByUid(c, orderInfo.Uid)
	if err != nil {
		return errors.Wrap(err, "获取用户信息失败")
	}
	//@todo 部分支付

	// 更新支付状态
	payTime := time.Now()
	err = db.GetDB().Transaction(func(tx *gorm.DB) error {
		updateLogRes := tx.Where("id = ?", payLog.Id).Updates(map[string]interface{}{
			"status":      enum.PayLogStatusPaid,
			"date_update": time.Now(),
		})
		if updateLogRes.Error != nil {
			return err
		}
		if updateLogRes.RowsAffected != 1 {
			return errors.New("操作冲突")
		}

		err = tx.Model(&model.BeeOrder{}).Where("order_id = ? and status = ?", orderInfo.Id, enum.OrderStatusUnPaid).
			Updates(map[string]interface{}{
				"status":      enum.OrderStatusPaid,
				"is_pay":      1,
				"ip":          ip,
				"date_pay":    payTime,
				"date_update": payTime,
			}).Error
		if err != nil {
			return errors.Wrap(err, "更新订单信息失败")
		}

		if !userInfo.IsVirtual {
			err = GetUserSrv().IncrUserLevelAmount(c, tx, orderInfo.Uid, orderInfo.AmountReal)
			if err != nil {
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

		tx.Where("id = ?")
		for _, fun := range extraTx {
			if err := fun(tx); err != nil {
				return err
			}
		}
		//@todo 积分奖励
		return nil
	})

	return err
}

func (s OrderSrv) GetOrderByOrderNo(c context.Context, orderNo string) (*proto.OrderDto, error) {
	item := &model.BeeOrder{}
	if err := db.GetDB().Where("order_number = ?", orderNo).Take(item).Error; err != nil {
		return nil, err
	}
	return proto.BeeOrder2Dto(item), nil
}

func (s OrderSrv) GetOrderByOrderId(c context.Context, orderId int64) (*proto.OrderDto, error) {
	item := &model.BeeOrder{}
	if err := db.GetDB().Where("id = ?", orderId).Take(item).Error; err != nil {
		return nil, err
	}
	return proto.BeeOrder2Dto(item), nil
}

func (s OrderSrv) PayByBalance(c context.Context, ip, orderId string, code string, pwd string) error {
	orderNumArr := strings.Split(orderId, ",")
	// @todo 短信验证码检查
	amount, err := GetBalanceSrv().GetAmount(kit.GetUid(c))
	if err != nil {
		return err
	}
	if amount.Pwd != util.Md5(amount.GetPwdEncode(pwd)) {
		return errors.New("密码错误")
	}
	if len(orderNumArr) > 1 {
		return errors.New("目前不支持同时支付多个订单")
	}
	// check余额
	orderInfo, err := s.GetOrderByOrderNo(c, orderNumArr[0])
	if err != nil {
		return err
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
	return s.PaySuccess(c, ip, payLog, orderInfo.OrderNumber, "balance_"+util.GetRandInt64(), orderInfo.AmountReal, func(tx *gorm.DB) error {
		// 扣余额
		rs := tx.Where("uid = ? and balance >= ?", kit.GetUid(c), orderInfo.AmountReal).Updates(map[string]interface{}{
			"balance":     gorm.Expr("balance - ?", orderInfo.AmountReal),
			"date_update": time.Now(),
		})
		if rs.Error != nil {
			return err
		}
		if rs.RowsAffected != 1 {
			return errors.New("扣除余额失败")
		}
		return nil
	})
}

func (s OrderSrv) Detail(c context.Context, orderId int64, hxNumber string) (*proto.GetOrderDetailResp, error) {
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
		Logistics: logistics,
		Logs:      logs,
	}
	return resp, nil
}
