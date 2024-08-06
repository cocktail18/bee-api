package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeWxConfigService struct{}

// CreateBeeWxConfig 创建微信配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeWxConfigService *BeeWxConfigService) CreateBeeWxConfig(beeWxConfig *bee.BeeWxConfig) (err error) {
	beeWxConfig.DateAdd = utils.NowPtr()
	beeWxConfig.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeWxConfig).Error
	return err
}

// DeleteBeeWxConfig 删除微信配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeWxConfigService *BeeWxConfigService) DeleteBeeWxConfig(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeWxConfig{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeWxConfigByIds 批量删除微信配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeWxConfigService *BeeWxConfigService) DeleteBeeWxConfigByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeWxConfig{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeWxConfig 更新微信配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeWxConfigService *BeeWxConfigService) UpdateBeeWxConfig(beeWxConfig bee.BeeWxConfig, shopUserId int) (err error) {
	beeWxConfig.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeWxConfig{}).Where("id = ? and user_id = ?", beeWxConfig.Id, shopUserId).Updates(&beeWxConfig).Error
	return err
}

// GetBeeWxConfig 根据id获取微信配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeWxConfigService *BeeWxConfigService) GetBeeWxConfig(id string, shopUserId int) (beeWxConfig bee.BeeWxConfig, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeWxConfig).Error
	return
}

// GetBeeWxConfigInfoList 分页获取微信配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeWxConfigService *BeeWxConfigService) GetBeeWxConfigInfoList(info beeReq.BeeWxConfigSearch, shopUserId int) (list []bee.BeeWxConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeWxConfig{})
	db = db.Where("user_id = ?", shopUserId)
	var beeWxConfigs []bee.BeeWxConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeWxConfigs).Error
	return beeWxConfigs, total, err
}
