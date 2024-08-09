package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeOrderGoodsService struct{}

// CreateBeeOrderGoods 创建订单商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderGoodsService *BeeOrderGoodsService) CreateBeeOrderGoods(beeOrderGoods *bee.BeeOrderGoods) (err error) {
	beeOrderGoods.DateAdd = utils.NowPtr()
	beeOrderGoods.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeOrderGoods).Error
	return err
}

// DeleteBeeOrderGoods 删除订单商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderGoodsService *BeeOrderGoodsService) DeleteBeeOrderGoods(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeOrderGoods{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeOrderGoodsByIds 批量删除订单商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderGoodsService *BeeOrderGoodsService) DeleteBeeOrderGoodsByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeOrderGoods{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeOrderGoods 更新订单商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderGoodsService *BeeOrderGoodsService) UpdateBeeOrderGoods(beeOrderGoods bee.BeeOrderGoods, shopUserId int) (err error) {
	beeOrderGoods.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeOrderGoods{}).Where("id = ? and user_id = ?", beeOrderGoods.Id, shopUserId).Updates(&beeOrderGoods).Error
	return err
}

// GetBeeOrderGoods 根据id获取订单商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderGoodsService *BeeOrderGoodsService) GetBeeOrderGoods(id string, shopUserId int) (beeOrderGoods bee.BeeOrderGoods, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeOrderGoods).Error
	return
}

// GetBeeOrderGoodsInfoList 分页获取订单商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderGoodsService *BeeOrderGoodsService) GetBeeOrderGoodsInfoList(info beeReq.BeeOrderGoodsSearch, shopUserId int) (list []bee.BeeOrderGoods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeOrderGoods{})
	db = db.Where("user_id = ?", shopUserId)
	var beeOrderGoodss []bee.BeeOrderGoods
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartDateAdd != nil && info.EndDateAdd != nil {
		db = db.Where("date_add BETWEEN ? AND ? ", info.StartDateAdd, info.EndDateAdd)
	}
	if info.StartDateUpdate != nil && info.EndDateUpdate != nil {
		db = db.Where("date_update BETWEEN ? AND ? ", info.StartDateUpdate, info.EndDateUpdate)
	}
	if info.OrderId != nil {
		db = db.Where("order_id = ?", info.OrderId)
	}
	if info.Purchase != "" {
		db = db.Where("purchase = ?", info.Purchase)
	}
	if info.RefundStatus != nil {
		db = db.Where("refund_status = ?", info.RefundStatus)
	}
	if info.Uid != nil {
		db = db.Where("uid = ?", info.Uid)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["date_add"] = true
	orderMap["date_update"] = true
	orderMap["id"] = true
	orderMap["order_id"] = true
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

	err = db.Find(&beeOrderGoodss).Error
	return beeOrderGoodss, total, err
}
