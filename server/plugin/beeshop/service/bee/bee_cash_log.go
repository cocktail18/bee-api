package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeCashLogService struct{}

// CreateBeeCashLog 创建用户现金消费记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCashLogService *BeeCashLogService) CreateBeeCashLog(beeCashLog *bee.BeeCashLog) (err error) {
	beeCashLog.DateAdd = utils.NowPtr()
	beeCashLog.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeCashLog).Error
	return err
}

// DeleteBeeCashLog 删除用户现金消费记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCashLogService *BeeCashLogService) DeleteBeeCashLog(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeCashLog{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeCashLogByIds 批量删除用户现金消费记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCashLogService *BeeCashLogService) DeleteBeeCashLogByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeCashLog{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeCashLog 更新用户现金消费记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCashLogService *BeeCashLogService) UpdateBeeCashLog(beeCashLog bee.BeeCashLog, shopUserId int) (err error) {
	beeCashLog.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeCashLog{}).Where("id = ? and user_id = ?", beeCashLog.Id, shopUserId).Updates(&beeCashLog).Error
	return err
}

// GetBeeCashLog 根据id获取用户现金消费记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCashLogService *BeeCashLogService) GetBeeCashLog(id string, shopUserId int) (beeCashLog bee.BeeCashLog, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeCashLog).Error
	return
}

// GetBeeCashLogInfoList 分页获取用户现金消费记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCashLogService *BeeCashLogService) GetBeeCashLogInfoList(info beeReq.BeeCashLogSearch, shopUserId int) (list []bee.BeeCashLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeCashLog{})
	db = db.Where("user_id = ?", shopUserId)
	var beeCashLogs []bee.BeeCashLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Uid != nil {
		db = db.Where("uid = ?", info.Uid)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeCashLogs).Error
	return beeCashLogs, total, err
}
