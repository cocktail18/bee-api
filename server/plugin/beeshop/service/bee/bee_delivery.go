package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeDeliveryService struct{}

// CreateBeeDelivery 创建beeDelivery表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeDeliveryService *BeeDeliveryService) CreateBeeDelivery(beeDelivery *bee.BeeDelivery) (err error) {
	beeDelivery.DateAdd = utils.NowPtr()
	beeDelivery.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeDelivery).Error
	return err
}

// DeleteBeeDelivery 删除beeDelivery表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeDeliveryService *BeeDeliveryService) DeleteBeeDelivery(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeDelivery{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeDeliveryByIds 批量删除beeDelivery表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeDeliveryService *BeeDeliveryService) DeleteBeeDeliveryByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeDelivery{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeDelivery 更新beeDelivery表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeDeliveryService *BeeDeliveryService) UpdateBeeDelivery(beeDelivery bee.BeeDelivery, shopUserId int) (err error) {
	beeDelivery.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeDelivery{}).Where("id = ? and user_id = ?", beeDelivery.Id, shopUserId).Updates(&beeDelivery).Error
	return err
}

// GetBeeDelivery 根据id获取beeDelivery表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeDeliveryService *BeeDeliveryService) GetBeeDelivery(id string, shopUserId int) (beeDelivery bee.BeeDelivery, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeDelivery).Error
	return
}

// GetBeeDeliveryInfoList 分页获取beeDelivery表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeDeliveryService *BeeDeliveryService) GetBeeDeliveryInfoList(info beeReq.BeeDeliverySearch, shopUserId int) (list []bee.BeeDelivery, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeDelivery{})
	db = db.Where("user_id = ? and is_deleted = 0", shopUserId)
	var beeDeliverys []bee.BeeDelivery
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
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

	err = db.Find(&beeDeliverys).Error
	return beeDeliverys, total, err
}
