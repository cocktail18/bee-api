package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeOrderLogisticsService struct{}

// CreateBeeOrderLogistics 创建用户订单地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderLogisticsService *BeeOrderLogisticsService) CreateBeeOrderLogistics(beeOrderLogistics *bee.BeeOrderLogistics) (err error) {
	beeOrderLogistics.DateAdd = utils.NowPtr()
	beeOrderLogistics.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeOrderLogistics).Error
	return err
}

// DeleteBeeOrderLogistics 删除用户订单地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderLogisticsService *BeeOrderLogisticsService) DeleteBeeOrderLogistics(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeOrderLogistics{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeOrderLogisticsByIds 批量删除用户订单地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderLogisticsService *BeeOrderLogisticsService) DeleteBeeOrderLogisticsByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeOrderLogistics{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeOrderLogistics 更新用户订单地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderLogisticsService *BeeOrderLogisticsService) UpdateBeeOrderLogistics(beeOrderLogistics bee.BeeOrderLogistics, shopUserId int) (err error) {
	beeOrderLogistics.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeOrderLogistics{}).Where("id = ? and user_id = ?", beeOrderLogistics.Id, shopUserId).Updates(&beeOrderLogistics).Error
	return err
}

// GetBeeOrderLogistics 根据id获取用户订单地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderLogisticsService *BeeOrderLogisticsService) GetBeeOrderLogistics(id string, shopUserId int) (beeOrderLogistics bee.BeeOrderLogistics, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeOrderLogistics).Error
	return
}

// GetBeeOrderLogisticsInfoList 分页获取用户订单地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderLogisticsService *BeeOrderLogisticsService) GetBeeOrderLogisticsInfoList(info beeReq.BeeOrderLogisticsSearch, shopUserId int) (list []bee.BeeOrderLogistics, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeOrderLogistics{})
	db = db.Where("user_id = ?", shopUserId)
	var beeOrderLogisticss []bee.BeeOrderLogistics
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Uid != nil {
		db = db.Where("uid = ?", info.Uid)
	}
	if info.OrderId != nil {
		db = db.Where("order_id = ?", info.OrderId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeOrderLogisticss).Error
	return beeOrderLogisticss, total, err
}
