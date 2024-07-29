package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeOrderLogService struct{}

// CreateBeeOrderLog 创建订单日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderLogService *BeeOrderLogService) CreateBeeOrderLog(beeOrderLog *bee.BeeOrderLog) (err error) {
	beeOrderLog.DateAdd = utils.NowPtr()
	beeOrderLog.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeOrderLog).Error
	return err
}

// DeleteBeeOrderLog 删除订单日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderLogService *BeeOrderLogService) DeleteBeeOrderLog(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeOrderLog{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeOrderLogByIds 批量删除订单日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderLogService *BeeOrderLogService) DeleteBeeOrderLogByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeOrderLog{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeOrderLog 更新订单日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderLogService *BeeOrderLogService) UpdateBeeOrderLog(beeOrderLog bee.BeeOrderLog, shopUserId int) (err error) {
	beeOrderLog.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeOrderLog{}).Where("id = ? and user_id = ?", beeOrderLog.Id, shopUserId).Updates(&beeOrderLog).Error
	return err
}

// GetBeeOrderLog 根据id获取订单日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderLogService *BeeOrderLogService) GetBeeOrderLog(id string, shopUserId int) (beeOrderLog bee.BeeOrderLog, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeOrderLog).Error
	return
}

// GetBeeOrderLogInfoList 分页获取订单日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderLogService *BeeOrderLogService) GetBeeOrderLogInfoList(info beeReq.BeeOrderLogSearch, shopUserId int) (list []bee.BeeOrderLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeOrderLog{})
	db = db.Where("user_id = ?", shopUserId)
	var beeOrderLogs []bee.BeeOrderLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderId != nil {
		db = db.Where("order_id = ?", info.OrderId)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeOrderLogs).Error
	return beeOrderLogs, total, err
}
