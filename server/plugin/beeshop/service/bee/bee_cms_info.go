package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeCmsInfoService struct{}

// CreateBeeCmsInfo 创建cms信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCmsInfoService *BeeCmsInfoService) CreateBeeCmsInfo(beeCmsInfo *bee.BeeCmsInfo) (err error) {
	beeCmsInfo.DateAdd = utils.NowPtr()
	beeCmsInfo.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeCmsInfo).Error
	return err
}

// DeleteBeeCmsInfo 删除cms信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCmsInfoService *BeeCmsInfoService) DeleteBeeCmsInfo(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeCmsInfo{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeCmsInfoByIds 批量删除cms信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCmsInfoService *BeeCmsInfoService) DeleteBeeCmsInfoByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeCmsInfo{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeCmsInfo 更新cms信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCmsInfoService *BeeCmsInfoService) UpdateBeeCmsInfo(beeCmsInfo bee.BeeCmsInfo, shopUserId int) (err error) {
	beeCmsInfo.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeCmsInfo{}).Where("id = ? and user_id = ?", beeCmsInfo.Id, shopUserId).Updates(&beeCmsInfo).Error
	return err
}

// GetBeeCmsInfo 根据id获取cms信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCmsInfoService *BeeCmsInfoService) GetBeeCmsInfo(id string, shopUserId int) (beeCmsInfo bee.BeeCmsInfo, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeCmsInfo).Error
	return
}

// GetBeeCmsInfoInfoList 分页获取cms信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCmsInfoService *BeeCmsInfoService) GetBeeCmsInfoInfoList(info beeReq.BeeCmsInfoSearch, shopUserId int) (list []bee.BeeCmsInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeCmsInfo{})
	db = db.Where("user_id = ?", shopUserId)
	var beeCmsInfos []bee.BeeCmsInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
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

	err = db.Find(&beeCmsInfos).Error
	return beeCmsInfos, total, err
}
