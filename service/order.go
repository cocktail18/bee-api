package service

import (
	"context"
	"database/sql"
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
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

func (s OrderSrv) Close(userId int64, orderId int64, remark string) error {
	var orderInfo model.BeeOrder
	if err := db.GetDB().Where("id = ? and uid = ?", orderId, userId).Take(&orderInfo).Error; err != nil {
		return err
	}
	if orderInfo.Status != enum.OrderStatusUnPaid {
		return errors.New("订单状态错误")
	}
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ? and uid = ?", orderId, userId).Updates(map[string]interface{}{
			"status":     enum.OrderStatusClose,
			"updated_at": time.Now(),
			"remark":     remark,
		}).Error; err != nil {
			return err
		}
		//@todo 退换优惠券、积分
		return nil
	})
}

func (s OrderSrv) Create(c *gin.Context, orderGoodsList []*proto.BeeOrderGoods, remark string,
	peisongType string, canHx bool, shopIdZt int64, couponId int64, extJsonStr string, calculate bool) (*proto.CreateOrderResp, error) {
	//@todo
	//goodsJsonStr: [{"propertyChildIds":",142479:1205135,142480:1205137,142481:1205139","goodsId":1262301,"number":1,"logisticsType":0,"inviter_id":0},{"propertyChildIds":",142479:1205135,142480:1205137,142481:1205140","goodsId":1262306,"number":2,"logisticsType":0,"inviter_id":0},{"goodsId":1262300,"number":2,"logisticsType":0,"inviter_id":0}]
	//remark:
	//peisongType: kd
	//isCanHx: true
	//shopIdZt: 56387
	//shopNameZt: 龙楼店（文昌市龙楼镇）
	//extJsonStr: {}
	//calculate: true
	uid := kit.GetUid(c)
	userId := kit.GetUserId(c)
	// 获取优惠券
	couponList, err := GetCouponSrv().GetMyCouponListByStatus(uid, enum.CouponStatusNormal)
	if err != nil {
		return nil, errors.Wrap(err, "获取优惠券失败")
	}
	var couponInfo *model.BeeUserCoupon
	if couponId > 0 {
		couponInfo, err = GetCouponSrv().GetUserCoupon(uid, couponId)
		if err != nil {
			return nil, errors.Wrap(err, "获取优惠券失败")
		}
		if !couponInfo.CanUse() {
			return nil, errors.New("优惠券暂时不可用")
		}
	}
	// 获取商品
	resp := &proto.CreateOrderResp{}
	resp.CouponUserList = make([]*proto.UserCouponResp, 0, 10)
	for _, coupon := range couponList {
		resp.CouponUserList = append(resp.CouponUserList, proto.NewUserCouponResp(coupon))
	}
	orderId := strings.ReplaceAll(carbon.Now().ToShortDateTimeNanoString(), ".", "")
	now := carbon.Now()
	for i, goods := range orderGoodsList {
		goodsInfo, skuInfo, err := GetGoodsSrv().GetGoodsWithSku(c, goods.GoodsID, goods.PropertyChildIds)
		if err != nil {
			return nil, errors.Wrap(err, "获取商品详情失败")
		}
		orderGoodsList[i].OrderID = orderId
		orderGoodsList[i].GoodsIDStr = cast.ToString(goods.GoodsID)
		orderGoodsList[i].GoodsName = goodsInfo.Name
		orderGoodsList[i].Property = skuInfo.PropertyChildNames
		orderGoodsList[i].Pic = goodsInfo.Pic
		orderGoodsList[i].AfterSale = goodsInfo.AfterSale
		orderGoodsList[i].Amount = s.callAmount(goods, goodsInfo, skuInfo)
		orderGoodsList[i].AmountCoupon = s.callAmount(goods, goodsInfo, skuInfo)
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
		//PriceID          int64           `json:"priceId"`
		//Purchase         bool            `json:"purchase"`
		//RefundStatus     int             `json:"refundStatus"`
		//Score            decimal.Decimal `json:"score"`
		//ShopID           int64           `json:"shopId"`
		//Status           int             `json:"status"`
		//Type             int             `json:"type"`
		//UID              int64           `json:"uid"`
		//Unit             string          `json:"unit"`
		//UserID           int64           `json:"uid"`

	}
	return resp, nil
}

func (s OrderSrv) callAmount(goods *proto.BeeOrderGoods, goodsInfo *model.BeeShopGoods, skuInfo *model.BeeShopGoodsSku) decimal.Decimal {
	//@todo 拼团之类的
	return decimal.NewFromInt(goods.Number).Mul(skuInfo.Price)
}

func (s OrderSrv) List(c context.Context, userId int64, req *proto.ListOrderReq) (*proto.ListOrderResp, error) {
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

	dbIns := db.GetDB().Where("uid = ?", userId).Where("deleted_at is null")
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
		OrderGoodsMap: lo.KeyBy(orderGoodsList, func(item *model.BeeOrderGoods) string {
			return cast.ToString(item.OrderId)
		}),
	}, nil
}

func (s OrderSrv) Delete(ctx context.Context, uid int64, orderId int64) error {
	var orderInfo model.BeeOrder
	if err := db.GetDB().Where("id = ? and uid = ? and deleted_at is null", orderId, uid).Take(&orderInfo).Error; err != nil {
		return err
	}
	//@todo 检查订单状态
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ? and uid = ?", orderId, uid).Updates(map[string]interface{}{
			"deleted_at": sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		}).Error; err != nil {
			return err
		}
		//@todo 退换优惠券、积分
		return nil
	})
}

func (s OrderSrv) Delivery(c *gin.Context, orderId int64) error {
	var orderInfo model.BeeOrder
	uid := kit.GetUid(c)
	if err := db.GetDB().Where("id = ? and uid = ? and deleted_at is null", orderId, uid).Take(&orderInfo).Error; err != nil {
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

func (s OrderSrv) Reputation(c *gin.Context, req *proto.BeeOrderReputationReq) error {
	orderId := req.OrderId
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		for _, reputationItem := range req.Reputations {
			item := model.BeeOrderReputation{
				BaseModel: common.BaseModel{
					UserId:     kit.GetUserId(c),
					IsDeleted:  false,
					DateAdd:    common.JsonTime(time.Now()),
					DateUpdate: common.JsonTime(time.Now()),
				},
				Uid:        kit.GetUid(c),
				OrderId:    orderId,
				Reputation: reputationItem.Reputation,
				Remark:     reputationItem.Remark,
				PicsStr:    strings.Join(reputationItem.Pics, ","),
				Pics:       reputationItem.Pics,
			}
			if err := db.GetDB().Create(&item).Error; err != nil {
				return err
			}
		}
		return nil
	})

}

func (s OrderSrv) Hx(c *gin.Context, number string) error {
	var order model.BeeOrder
	if err := db.GetDB().Where("hx_number = ? and is_deleted = 0", number).Take(&order).Error; err != nil {
		return err
	}
	if !order.IsCanHx || order.IsEnd || order.Status != enum.OrderStatusPaid {
		return errors.New("订单不能被核销掉")
	}
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("order_id = ?", order.Id).Updates(map[string]interface{}{
			"status":      enum.OrderGoodsStatusConfirmShipped,
			"date_update": time.Now(),
		}).Error; err != nil {
			return err
		}
		return tx.Where("id = ?", order.Id).Updates(map[string]interface{}{
			"status":      enum.OrderStatusConfirmShipped,
			"date_update": time.Now(),
		}).Error
	})
}
