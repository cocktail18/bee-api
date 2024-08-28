package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeUserBalanceLogService struct{}

// CreateBeeUserBalanceLog 创建用户消费记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserBalanceLogService *BeeUserBalanceLogService) CreateBeeUserBalanceLog(beeUserBalanceLog *bee.BeeUserBalanceLog) (err error) {
	beeUserBalanceLog.DateAdd = utils.NowPtr()
	beeUserBalanceLog.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeUserBalanceLog).Error
	return err
}

// DeleteBeeUserBalanceLog 删除用户消费记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserBalanceLogService *BeeUserBalanceLogService) DeleteBeeUserBalanceLog(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeUserBalanceLog{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeUserBalanceLogByIds 批量删除用户消费记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserBalanceLogService *BeeUserBalanceLogService) DeleteBeeUserBalanceLogByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeUserBalanceLog{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeUserBalanceLog 更新用户消费记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserBalanceLogService *BeeUserBalanceLogService) UpdateBeeUserBalanceLog(beeUserBalanceLog bee.BeeUserBalanceLog, shopUserId int) (err error) {
	beeUserBalanceLog.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeUserBalanceLog{}).Where("id = ? and user_id = ?", beeUserBalanceLog.Id, shopUserId).Updates(&beeUserBalanceLog).Error
	return err
}

// GetBeeUserBalanceLog 根据id获取用户消费记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserBalanceLogService *BeeUserBalanceLogService) GetBeeUserBalanceLog(id string, shopUserId int) (beeUserBalanceLog bee.BeeUserBalanceLog, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeUserBalanceLog).Error
	return
}

// GetBeeUserBalanceLogInfoList 分页获取用户消费记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserBalanceLogService *BeeUserBalanceLogService) GetBeeUserBalanceLogInfoList(info beeReq.BeeUserBalanceLogSearch, shopUserId int) (list []bee.BeeUserBalanceLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeUserBalanceLog{})
	db = db.Where("user_id = ?", shopUserId)
	var beeUserBalanceLogs []bee.BeeUserBalanceLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Uid != nil {
		db = db.Where("uid = ?", info.Uid)
	}
	if info.StartDateAdd != nil && info.EndDateAdd != nil {
		db = db.Where("date_add BETWEEN ? AND ? ", info.StartDateAdd, info.EndDateAdd)
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

	err = db.Find(&beeUserBalanceLogs).Error
	return beeUserBalanceLogs, total, err
}
