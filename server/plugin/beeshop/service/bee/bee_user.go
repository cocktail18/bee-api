package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeUserService struct{}

// CreateBeeUser 创建beeUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserService *BeeUserService) CreateBeeUser(beeUser *bee.BeeUser) (err error) {
	beeUser.DateAdd = utils.NowPtr()
	beeUser.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeUser).Error
	return err
}

// DeleteBeeUser 删除beeUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserService *BeeUserService) DeleteBeeUser(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUser{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeUserByIds 批量删除beeUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserService *BeeUserService) DeleteBeeUserByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUser{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeUser 更新beeUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserService *BeeUserService) UpdateBeeUser(beeUser bee.BeeUser, shopUserId int) (err error) {
	beeUser.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUser{}).Where("id = ? and user_id = ?", beeUser.Id, shopUserId).Updates(&beeUser).Error
	return err
}

// GetBeeUser 根据id获取beeUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserService *BeeUserService) GetBeeUser(id string, shopUserId int) (beeUser bee.BeeUser, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeUser).Error
	return
}

// GetBeeUserInfoList 分页获取beeUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserService *BeeUserService) GetBeeUserInfoList(info beeReq.BeeUserSearch, shopUserId int) (list []bee.BeeUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUser{})
	db = db.Where("user_id = ?", shopUserId)
	var beeUsers []bee.BeeUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Id != nil {
		db = db.Where("id = ?", info.Id)
	}
	if info.ShowUid != nil {
		db = db.Where("show_uid = ?", info.ShowUid)
	}
	if info.CardNumber != "" {
		db = db.Where("card_number = ?", info.CardNumber)
	}
	if info.IsVirtual != nil {
		db = db.Where("is_virtual = ?", info.IsVirtual)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
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

	err = db.Find(&beeUsers).Error
	return beeUsers, total, err
}
