package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeRechargeSendRuleService struct{}

// CreateBeeRechargeSendRule 创建beeRechargeSendRule表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeRechargeSendRuleService *BeeRechargeSendRuleService) CreateBeeRechargeSendRule(beeRechargeSendRule *bee.BeeRechargeSendRule) (err error) {
	beeRechargeSendRule.DateAdd = utils.NowPtr()
	beeRechargeSendRule.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeRechargeSendRule).Error
	return err
}

// DeleteBeeRechargeSendRule 删除beeRechargeSendRule表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeRechargeSendRuleService *BeeRechargeSendRuleService) DeleteBeeRechargeSendRule(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeRechargeSendRule{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeRechargeSendRuleByIds 批量删除beeRechargeSendRule表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeRechargeSendRuleService *BeeRechargeSendRuleService) DeleteBeeRechargeSendRuleByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeRechargeSendRule{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeRechargeSendRule 更新beeRechargeSendRule表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeRechargeSendRuleService *BeeRechargeSendRuleService) UpdateBeeRechargeSendRule(beeRechargeSendRule bee.BeeRechargeSendRule, shopUserId int) (err error) {
	beeRechargeSendRule.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeRechargeSendRule{}).Where("id = ? and user_id = ?", beeRechargeSendRule.Id, shopUserId).Updates(&beeRechargeSendRule).Error
	return err
}

// GetBeeRechargeSendRule 根据id获取beeRechargeSendRule表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeRechargeSendRuleService *BeeRechargeSendRuleService) GetBeeRechargeSendRule(id string, shopUserId int) (beeRechargeSendRule bee.BeeRechargeSendRule, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeRechargeSendRule).Error
	return
}

// GetBeeRechargeSendRuleInfoList 分页获取beeRechargeSendRule表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeRechargeSendRuleService *BeeRechargeSendRuleService) GetBeeRechargeSendRuleInfoList(info beeReq.BeeRechargeSendRuleSearch, shopUserId int) (list []bee.BeeRechargeSendRule, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeRechargeSendRule{})
	db = db.Where("user_id = ?", shopUserId)
	var beeRechargeSendRules []bee.BeeRechargeSendRule
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeRechargeSendRules).Error
	return beeRechargeSendRules, total, err
}
