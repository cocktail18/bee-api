package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeUploadFileService struct{}

// CreateBeeUploadFile 创建用户上传文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUploadFileService *BeeUploadFileService) CreateBeeUploadFile(beeUploadFile *bee.BeeUploadFile) (err error) {
	beeUploadFile.DateAdd = utils.NowPtr()
	beeUploadFile.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeUploadFile).Error
	return err
}

// DeleteBeeUploadFile 删除用户上传文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUploadFileService *BeeUploadFileService) DeleteBeeUploadFile(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUploadFile{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeUploadFileByIds 批量删除用户上传文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUploadFileService *BeeUploadFileService) DeleteBeeUploadFileByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUploadFile{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeUploadFile 更新用户上传文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUploadFileService *BeeUploadFileService) UpdateBeeUploadFile(beeUploadFile bee.BeeUploadFile, shopUserId int) (err error) {
	beeUploadFile.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUploadFile{}).Where("id = ? and user_id = ?", beeUploadFile.Id, shopUserId).Updates(&beeUploadFile).Error
	return err
}

// GetBeeUploadFile 根据id获取用户上传文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUploadFileService *BeeUploadFileService) GetBeeUploadFile(id string, shopUserId int) (beeUploadFile bee.BeeUploadFile, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeUploadFile).Error
	return
}

// GetBeeUploadFileInfoList 分页获取用户上传文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUploadFileService *BeeUploadFileService) GetBeeUploadFileInfoList(info beeReq.BeeUploadFileSearch, shopUserId int) (list []bee.BeeUploadFile, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUploadFile{})
	db = db.Where("user_id = ?", shopUserId)
	var beeUploadFiles []bee.BeeUploadFile
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartDateAdd != nil && info.EndDateAdd != nil {
		db = db.Where("date_add BETWEEN ? AND ? ", info.StartDateAdd, info.EndDateAdd)
	}
	if info.StartDateUpdate != nil && info.EndDateUpdate != nil {
		db = db.Where("date_update BETWEEN ? AND ? ", info.StartDateUpdate, info.EndDateUpdate)
	}
	if info.Filename != "" {
		db = db.Where("filename LIKE ?", "%"+info.Filename+"%")
	}
	if info.StartExpireAt != nil && info.EndExpireAt != nil {
		db = db.Where("expire_at BETWEEN ? AND ? ", info.StartExpireAt, info.EndExpireAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeUploadFiles).Error
	return beeUploadFiles, total, err
}
