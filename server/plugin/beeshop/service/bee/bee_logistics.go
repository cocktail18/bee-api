package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeLogisticsService struct{}

// CreateBeeLogistics 创建运费模版记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeLogisticsService *BeeLogisticsService) CreateBeeLogistics(beeLogistics *bee.BeeLogistics) (err error) {
	beeLogistics.DateAdd = utils.NowPtr()
	beeLogistics.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeLogistics).Error
	return err
}

// DeleteBeeLogistics 删除运费模版记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeLogisticsService *BeeLogisticsService) DeleteBeeLogistics(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeLogistics{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeLogisticsByIds 批量删除运费模版记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeLogisticsService *BeeLogisticsService) DeleteBeeLogisticsByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeLogistics{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeLogistics 更新运费模版记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeLogisticsService *BeeLogisticsService) UpdateBeeLogistics(beeLogistics bee.BeeLogistics, shopUserId int) (err error) {
	beeLogistics.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeLogistics{}).Where("id = ? and user_id = ?", beeLogistics.Id, shopUserId).Updates(&beeLogistics).Error
	return err
}

// GetBeeLogistics 根据id获取运费模版记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeLogisticsService *BeeLogisticsService) GetBeeLogistics(id string, shopUserId int) (beeLogistics bee.BeeLogistics, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeLogistics).Error
	return
}

// GetBeeLogisticsInfoList 分页获取运费模版记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeLogisticsService *BeeLogisticsService) GetBeeLogisticsInfoList(info beeReq.BeeLogisticsSearch, shopUserId int) (list []bee.BeeLogistics, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeLogistics{})
	db = db.Where("user_id = ?", shopUserId)
	var beeLogisticss []bee.BeeLogistics
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Id != nil {
		db = db.Where("id = ?", info.Id)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeLogisticss).Error
	return beeLogisticss, total, err
}
