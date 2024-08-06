package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeNoticeService struct{}

// CreateBeeNotice 创建系统公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeNoticeService *BeeNoticeService) CreateBeeNotice(beeNotice *bee.BeeNotice) (err error) {
	beeNotice.DateAdd = utils.NowPtr()
	beeNotice.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeNotice).Error
	return err
}

// DeleteBeeNotice 删除系统公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeNoticeService *BeeNoticeService) DeleteBeeNotice(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeNotice{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeNoticeByIds 批量删除系统公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeNoticeService *BeeNoticeService) DeleteBeeNoticeByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeNotice{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeNotice 更新系统公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeNoticeService *BeeNoticeService) UpdateBeeNotice(beeNotice bee.BeeNotice, shopUserId int) (err error) {
	beeNotice.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeNotice{}).Where("id = ? and user_id = ?", beeNotice.Id, shopUserId).Updates(&beeNotice).Error
	return err
}

// GetBeeNotice 根据id获取系统公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeNoticeService *BeeNoticeService) GetBeeNotice(id string, shopUserId int) (beeNotice bee.BeeNotice, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeNotice).Error
	return
}

// GetBeeNoticeInfoList 分页获取系统公告记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeNoticeService *BeeNoticeService) GetBeeNoticeInfoList(info beeReq.BeeNoticeSearch, shopUserId int) (list []bee.BeeNotice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeNotice{})
	db = db.Where("user_id = ?", shopUserId)
	var beeNotices []bee.BeeNotice
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeNotices).Error
	return beeNotices, total, err
}
