package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeOrderReputationService struct{}

// CreateBeeOrderReputation 创建商品评价记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderReputationService *BeeOrderReputationService) CreateBeeOrderReputation(beeOrderReputation *bee.BeeOrderReputation) (err error) {
	beeOrderReputation.DateAdd = utils.NowPtr()
	beeOrderReputation.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeOrderReputation).Error
	return err
}

// DeleteBeeOrderReputation 删除商品评价记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderReputationService *BeeOrderReputationService) DeleteBeeOrderReputation(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeOrderReputation{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeOrderReputationByIds 批量删除商品评价记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderReputationService *BeeOrderReputationService) DeleteBeeOrderReputationByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeOrderReputation{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeOrderReputation 更新商品评价记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderReputationService *BeeOrderReputationService) UpdateBeeOrderReputation(beeOrderReputation bee.BeeOrderReputation, shopUserId int) (err error) {
	beeOrderReputation.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeOrderReputation{}).Where("id = ? and user_id = ?", beeOrderReputation.Id, shopUserId).Updates(&beeOrderReputation).Error
	return err
}

// GetBeeOrderReputation 根据id获取商品评价记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderReputationService *BeeOrderReputationService) GetBeeOrderReputation(id string, shopUserId int) (beeOrderReputation bee.BeeOrderReputation, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeOrderReputation).Error
	return
}

// GetBeeOrderReputationInfoList 分页获取商品评价记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderReputationService *BeeOrderReputationService) GetBeeOrderReputationInfoList(info beeReq.BeeOrderReputationSearch, shopUserId int) (list []bee.BeeOrderReputation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeOrderReputation{})
	db = db.Where("user_id = ?", shopUserId)
	var beeOrderReputations []bee.BeeOrderReputation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartDateAdd != nil && info.EndDateAdd != nil {
		db = db.Where("date_add BETWEEN ? AND ? ", info.StartDateAdd, info.EndDateAdd)
	}
	if info.Uid != nil {
		db = db.Where("uid = ?", info.Uid)
	}
	if info.OrderId != "" {
		db = db.Where("order_id = ?", info.OrderId)
	}
	if info.Reputation != nil {
		db = db.Where("reputation = ?", info.Reputation)
	}
	if info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+info.Remark+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["reputation"] = true
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

	err = db.Find(&beeOrderReputations).Error
	return beeOrderReputations, total, err
}
