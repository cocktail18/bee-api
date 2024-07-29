package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeUserAmountService struct{}

// CreateBeeUserAmount 创建用户资产记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserAmountService *BeeUserAmountService) CreateBeeUserAmount(beeUserAmount *bee.BeeUserAmount) (err error) {
	beeUserAmount.DateAdd = utils.NowPtr()
	beeUserAmount.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeUserAmount).Error
	return err
}

// DeleteBeeUserAmount 删除用户资产记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserAmountService *BeeUserAmountService) DeleteBeeUserAmount(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUserAmount{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeUserAmountByIds 批量删除用户资产记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserAmountService *BeeUserAmountService) DeleteBeeUserAmountByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUserAmount{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeUserAmount 更新用户资产记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserAmountService *BeeUserAmountService) UpdateBeeUserAmount(beeUserAmount bee.BeeUserAmount, shopUserId int) (err error) {
	beeUserAmount.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUserAmount{}).Where("id = ? and user_id = ?", beeUserAmount.Id, shopUserId).Updates(&beeUserAmount).Error
	return err
}

// GetBeeUserAmount 根据id获取用户资产记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserAmountService *BeeUserAmountService) GetBeeUserAmount(id string, shopUserId int) (beeUserAmount bee.BeeUserAmount, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeUserAmount).Error
	return
}

// GetBeeUserAmountInfoList 分页获取用户资产记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserAmountService *BeeUserAmountService) GetBeeUserAmountInfoList(info beeReq.BeeUserAmountSearch, shopUserId int) (list []bee.BeeUserAmount, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUserAmount{})
	db = db.Where("user_id = ?", shopUserId)
	var beeUserAmounts []bee.BeeUserAmount
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartDateAdd != nil && info.EndDateAdd != nil {
		db = db.Where("date_add BETWEEN ? AND ? ", info.StartDateAdd, info.EndDateAdd)
	}
	if info.StartDateUpdate != nil && info.EndDateUpdate != nil {
		db = db.Where("date_update BETWEEN ? AND ? ", info.StartDateUpdate, info.EndDateUpdate)
	}
	if info.Uid != nil {
		db = db.Where("uid = ?", info.Uid)
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
	orderMap["uid"] = true
	orderMap["balance"] = true
	orderMap["freeze"] = true
	orderMap["fx_commision_paying"] = true
	orderMap["growth"] = true
	orderMap["score"] = true
	orderMap["score_last_round"] = true
	orderMap["total_pay_amount"] = true
	orderMap["total_pay_number"] = true
	orderMap["total_score"] = true
	orderMap["total_withdraw"] = true
	orderMap["total_consumed"] = true
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

	err = db.Find(&beeUserAmounts).Error
	return beeUserAmounts, total, err
}
