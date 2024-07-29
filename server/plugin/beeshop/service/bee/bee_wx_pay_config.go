package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeWxPayConfigService struct{}

// CreateBeeWxPayConfig 创建微信支付配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeWxPayConfigService *BeeWxPayConfigService) CreateBeeWxPayConfig(beeWxPayConfig *bee.BeeWxPayConfig) (err error) {
	beeWxPayConfig.DateAdd = utils.NowPtr()
	beeWxPayConfig.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeWxPayConfig).Error
	return err
}

// DeleteBeeWxPayConfig 删除微信支付配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeWxPayConfigService *BeeWxPayConfigService) DeleteBeeWxPayConfig(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeWxPayConfig{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeWxPayConfigByIds 批量删除微信支付配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeWxPayConfigService *BeeWxPayConfigService) DeleteBeeWxPayConfigByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeWxPayConfig{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeWxPayConfig 更新微信支付配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeWxPayConfigService *BeeWxPayConfigService) UpdateBeeWxPayConfig(beeWxPayConfig bee.BeeWxPayConfig, shopUserId int) (err error) {
	beeWxPayConfig.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeWxPayConfig{}).Where("id = ? and user_id = ?", beeWxPayConfig.Id, shopUserId).Updates(&beeWxPayConfig).Error
	return err
}

// GetBeeWxPayConfig 根据id获取微信支付配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeWxPayConfigService *BeeWxPayConfigService) GetBeeWxPayConfig(id string, shopUserId int) (beeWxPayConfig bee.BeeWxPayConfig, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeWxPayConfig).Error
	return
}

// GetBeeWxPayConfigInfoList 分页获取微信支付配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeWxPayConfigService *BeeWxPayConfigService) GetBeeWxPayConfigInfoList(info beeReq.BeeWxPayConfigSearch, shopUserId int) (list []bee.BeeWxPayConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeWxPayConfig{})
	db = db.Where("user_id = ?", shopUserId)
	var beeWxPayConfigs []bee.BeeWxPayConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeWxPayConfigs).Error
	return beeWxPayConfigs, total, err
}
