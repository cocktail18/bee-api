package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeOrderPeisongLogService struct{}

// CreateBeeOrderPeisongLog 创建beeOrderPeisongLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongLogService *BeeOrderPeisongLogService) CreateBeeOrderPeisongLog(beeOrderPeisongLog *bee.BeeOrderPeisongLog) (err error) {
	beeOrderPeisongLog.DateAdd = utils.NowPtr()
	beeOrderPeisongLog.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeOrderPeisongLog).Error
	return err
}

// DeleteBeeOrderPeisongLog 删除beeOrderPeisongLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongLogService *BeeOrderPeisongLogService) DeleteBeeOrderPeisongLog(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeOrderPeisongLog{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeOrderPeisongLogByIds 批量删除beeOrderPeisongLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongLogService *BeeOrderPeisongLogService) DeleteBeeOrderPeisongLogByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeOrderPeisongLog{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeOrderPeisongLog 更新beeOrderPeisongLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongLogService *BeeOrderPeisongLogService) UpdateBeeOrderPeisongLog(beeOrderPeisongLog bee.BeeOrderPeisongLog, shopUserId int) (err error) {
	beeOrderPeisongLog.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeOrderPeisongLog{}).Where("id = ? and user_id = ?", beeOrderPeisongLog.Id, shopUserId).Updates(&beeOrderPeisongLog).Error
	return err
}

// GetBeeOrderPeisongLog 根据id获取beeOrderPeisongLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongLogService *BeeOrderPeisongLogService) GetBeeOrderPeisongLog(id string, shopUserId int) (beeOrderPeisongLog bee.BeeOrderPeisongLog, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeOrderPeisongLog).Error
	return
}

// GetBeeOrderPeisongLogInfoList 分页获取beeOrderPeisongLog表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongLogService *BeeOrderPeisongLogService) GetBeeOrderPeisongLogInfoList(info beeReq.BeeOrderPeisongLogSearch, shopUserId int) (list []bee.BeeOrderPeisongLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeOrderPeisongLog{})
	db = db.Where("user_id = ?", shopUserId)
	var beeOrderPeisongLogs []bee.BeeOrderPeisongLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderId != nil {
		db = db.Where("order_id = ?", info.OrderId)
	}
	if info.PeisongOrderNo != "" {
		db = db.Where("peisong_order_no = ?", info.PeisongOrderNo)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeOrderPeisongLogs).Error
	return beeOrderPeisongLogs, total, err
}
