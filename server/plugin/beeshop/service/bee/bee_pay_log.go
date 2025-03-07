package bee

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/dto"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeePayLogService struct{}

// CreateBeePayLog 创建支付流水记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePayLogService *BeePayLogService) CreateBeePayLog(beePayLog *bee.BeePayLog) (err error) {
	beePayLog.DateAdd = utils.NowPtr()
	beePayLog.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beePayLog).Error
	return err
}

// DeleteBeePayLog 删除支付流水记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePayLogService *BeePayLogService) DeleteBeePayLog(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeePayLog{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeePayLogByIds 批量删除支付流水记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePayLogService *BeePayLogService) DeleteBeePayLogByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeePayLog{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeePayLog 更新支付流水记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePayLogService *BeePayLogService) UpdateBeePayLog(beePayLog bee.BeePayLog, shopUserId int) (err error) {
	beePayLog.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeePayLog{}).Where("id = ? and user_id = ?", beePayLog.Id, shopUserId).Updates(&beePayLog).Error
	return err
}

// GetBeePayLog 根据id获取支付流水记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePayLogService *BeePayLogService) GetBeePayLog(id string, shopUserId int) (beePayLog bee.BeePayLog, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beePayLog).Error
	return
}

// GetBeePayLogInfoList 分页获取支付流水记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePayLogService *BeePayLogService) GetBeePayLogInfoList(info beeReq.BeePayLogSearch, shopUserId int) (list []dto.BeePayLogInfoDto, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&dto.BeePayLogInfoDto{}).Joins(" left join bee_order a on a.order_number = bee_pay_log.order_no left join bee_shop_info b on b.id = a.shop_id")
	db = db.Where("bee_pay_log.user_id = ?", shopUserId)
	var beePayLogs []dto.BeePayLogInfoDto
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Distinct != "" {
		db.Select(fmt.Sprintf("distinct(%s)", info.Distinct))
	}
	if info.StartDateAdd != nil && info.EndDateAdd != nil {
		db = db.Where("bee_pay_log.date_add BETWEEN ? AND ? ", info.StartDateAdd, info.EndDateAdd)
	}
	if info.StartDateUpdate != nil && info.EndDateUpdate != nil {
		db = db.Where("bee_pay_log.date_update BETWEEN ? AND ? ", info.StartDateUpdate, info.EndDateUpdate)
	}
	if info.StartMoney != nil && info.EndMoney != nil {
		db = db.Where("bee_pay_log.money BETWEEN ? AND ? ", info.StartMoney, info.EndMoney)
	}
	if info.OrderNo != "" {
		db = db.Where("bee_pay_log.order_no = ?", info.OrderNo)
	}
	if info.Status != "" {
		db = db.Where("bee_pay_log.status = ?", info.Status)
	}
	if info.Uid != nil {
		db = db.Where("bee_pay_log.uid = ?", info.Uid)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["date_add"] = true
	orderMap["date_update"] = true
	orderMap["money"] = true
	orderMap["status"] = true
	orderMap["uid"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Select("bee_pay_log.*, b.name as shopName").Debug()
	}

	err = db.Find(&beePayLogs).Error
	return beePayLogs, total, err
}

// GetBeePayTotal 获取支付流水总数
func (beePayLogService *BeePayLogService) GetBeePayTotal(info beeReq.BeePayLogSearch, shopUserId int) (total float64, err error) {
	// 创建db
	db := GetBeeDB().Model(&bee.BeePayLog{})
	db = db.Where("user_id = ?", shopUserId)
	if info.StartDateAdd != nil && info.EndDateAdd != nil {
		db = db.Where("date_add BETWEEN ? AND ? ", info.StartDateAdd, info.EndDateAdd)
	}
	var totalDb = struct {
		Total float64 `gorm:"total"`
	}{}
	if err = db.Select("sum(" + info.Sum + ") as total").Take(&totalDb).Error; err != nil {
		return 0, err
	}
	total = totalDb.Total
	return total, nil
}
