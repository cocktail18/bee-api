package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeConfigService struct{}

// CreateBeeConfig 创建公用配置表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeConfigService *BeeConfigService) CreateBeeConfig(beeConfig *bee.BeeConfig) (err error) {
	beeConfig.DateAdd = utils.NowPtr()
	beeConfig.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeConfig).Error
	return err
}

// DeleteBeeConfig 删除公用配置表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeConfigService *BeeConfigService) DeleteBeeConfig(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeConfig{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeConfigByIds 批量删除公用配置表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeConfigService *BeeConfigService) DeleteBeeConfigByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeConfig{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeConfig 更新公用配置表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeConfigService *BeeConfigService) UpdateBeeConfig(beeConfig bee.BeeConfig, shopUserId int) (err error) {
	beeConfig.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeConfig{}).Where("id = ? and user_id = ?", beeConfig.Id, shopUserId).Updates(&beeConfig).Error
	return err
}

// GetBeeConfig 根据id获取公用配置表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeConfigService *BeeConfigService) GetBeeConfig(id string, shopUserId int) (beeConfig bee.BeeConfig, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeConfig).Error
	return
}

// GetBeeConfigInfoList 分页获取公用配置表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeConfigService *BeeConfigService) GetBeeConfigInfoList(info beeReq.BeeConfigSearch, shopUserId int) (list []bee.BeeConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeConfig{})
	db = db.Where("user_id = ?", shopUserId)
	var beeConfigs []bee.BeeConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeConfigs).Error
	return beeConfigs, total, err
}
