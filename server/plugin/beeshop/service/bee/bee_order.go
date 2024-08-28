package bee

import (
	"context"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
	"github.com/spf13/cast"
)

type BeeOrderService struct{}

// CreateBeeOrder 创建用户订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderService *BeeOrderService) CreateBeeOrder(beeOrder *bee.BeeOrder) (err error) {
	beeOrder.DateAdd = utils.NowPtr()
	beeOrder.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeOrder).Error
	return err
}

// DeleteBeeOrder 删除用户订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderService *BeeOrderService) DeleteBeeOrder(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeOrder{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeOrderByIds 批量删除用户订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderService *BeeOrderService) DeleteBeeOrderByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeOrder{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// MarkBeeOrderPaid 批量设置为已支付
func (beeOrderService *BeeOrderService) MarkBeeOrderPaid(ids []string, shopUserId int) (err error) {
	for _, id := range ids {
		err = service.GetOrderSrv().PayOrderOffline(context.Background(), id)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarkBeeOrderDone 批量设置为已完成
func (beeOrderService *BeeOrderService) MarkBeeOrderDone(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeOrder{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_paid":     true,
			"status":      enum.OrderStatusConfirmShipped,
			"date_update": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeOrderExtJsonStr 更新extJsonStr字段
func (beeOrderService *BeeOrderService) UpdateBeeOrderExtJsonStr(beeOrder bee.BeeOrder, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeOrder{}).Where("id = ? and user_id = ?", beeOrder.Id, shopUserId).
		Updates(map[string]interface{}{
			"ext_json_str": beeOrder.ExtJsonStr,
			"date_update":  utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeOrderStatus 更新status字段
func (beeOrderService *BeeOrderService) UpdateBeeOrderStatus(beeOrder bee.BeeOrder, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeOrder{}).Where("id = ? and user_id = ?", beeOrder.Id, shopUserId).
		Updates(map[string]interface{}{
			"status":      beeOrder.Status,
			"date_update": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeOrder 更新用户订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderService *BeeOrderService) UpdateBeeOrder(beeOrder bee.BeeOrder, shopUserId int) (err error) {
	beeOrder.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeOrder{}).Where("id = ? and user_id = ?", beeOrder.Id, shopUserId).Updates(&beeOrder).Error
	return err
}

// GetBeeOrder 根据id获取用户订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderService *BeeOrderService) GetBeeOrder(id string, shopUserId int) (beeOrder bee.BeeOrder, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeOrder).Error
	return
}

// GetBeeOrderInfoList 分页获取用户订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderService *BeeOrderService) GetBeeOrderInfoList(info beeReq.BeeOrderSearch, shopUserId int) (list []bee.BeeOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeOrder{})
	db = db.Where("user_id = ?", shopUserId)
	var beeOrders []bee.BeeOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ID != nil {
		db = db.Where("id = ?", info.ID)
	}
	if info.IsDeleted != nil {
		db = db.Where("is_deleted = ?", info.IsDeleted)
	}
	if info.StartDateAdd != nil && info.EndDateAdd != nil {
		db = db.Where("date_add BETWEEN ? AND ? ", info.StartDateAdd, info.EndDateAdd)
	}
	if info.StartDateUpdate != nil && info.EndDateUpdate != nil {
		db = db.Where("date_update BETWEEN ? AND ? ", info.StartDateUpdate, info.EndDateUpdate)
	}
	if info.StartDateDelete != nil && info.EndDateDelete != nil {
		db = db.Where("date_delete BETWEEN ? AND ? ", info.StartDateDelete, info.EndDateDelete)
	}
	if info.StartAmountLogistics != nil && info.EndAmountLogistics != nil {
		db = db.Where("amount_logistics BETWEEN ? AND ? ", info.StartAmountLogistics, info.EndAmountLogistics)
	}
	if info.StartAmountReal != nil && info.EndAmountReal != nil {
		db = db.Where("amount_real BETWEEN ? AND ? ", info.StartAmountReal, info.EndAmountReal)
	}
	if info.AutoDeliverStatus != nil {
		db = db.Where("auto_deliver_status = ?", info.AutoDeliverStatus)
	}
	if info.StartDateClose != nil && info.EndDateClose != nil {
		db = db.Where("date_close BETWEEN ? AND ? ", info.StartDateClose, info.EndDateClose)
	}
	if info.StartDatePay != nil && info.EndDatePay != nil {
		db = db.Where("date_pay BETWEEN ? AND ? ", info.StartDatePay, info.EndDatePay)
	}
	if info.HasRefund != nil {
		db = db.Where("has_refund = ?", info.HasRefund)
	}
	if info.IsEnd != nil {
		db = db.Where("is_end = ?", info.IsEnd)
	}
	if info.Pid != nil {
		db = db.Where("pid = ?", info.Pid)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.StartTrips != nil && info.EndTrips != nil {
		db = db.Where("trips BETWEEN ? AND ? ", info.StartTrips, info.EndTrips)
	}
	if info.Type != nil {
		db = db.Where("`type` = ?", info.Type)
	}
	if info.Uid != nil {
		db = db.Where("uid = ?", info.Uid)
	}
	if info.OrderNumber != nil {
		db = db.Where("order_number = ?", info.OrderNumber)
	}
	if info.HxNumber != nil {
		db = db.Where("hx_number = ?", info.HxNumber)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["date_add"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeOrders).Error
	return beeOrders, total, err
}

func (beeOrderService *BeeOrderService) ShippedBeeOrder(id int64, shopUserId int) error {
	ctx, err := getContext(shopUserId)
	if err != nil {
		return err
	}
	err = service.GetOrderSrv().ShippedBeeOrder(ctx, cast.ToInt64(id))
	if err != nil {
		return err
	}
	return nil
}
