package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeUserMapperService struct{}

// CreateBeeUserMapper 创建beeUserMapper表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserMapperService *BeeUserMapperService) CreateBeeUserMapper(beeUserMapper *bee.BeeUserMapper) (err error) {
	beeUserMapper.DateAdd = utils.NowPtr()
	beeUserMapper.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeUserMapper).Error
	return err
}

// DeleteBeeUserMapper 删除beeUserMapper表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserMapperService *BeeUserMapperService) DeleteBeeUserMapper(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUserMapper{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeUserMapperByIds 批量删除beeUserMapper表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserMapperService *BeeUserMapperService) DeleteBeeUserMapperByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUserMapper{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeUserMapper 更新beeUserMapper表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserMapperService *BeeUserMapperService) UpdateBeeUserMapper(beeUserMapper bee.BeeUserMapper, shopUserId int) (err error) {
	beeUserMapper.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUserMapper{}).Where("id = ? and user_id = ?", beeUserMapper.Id, shopUserId).Updates(&beeUserMapper).Error
	return err
}

// GetBeeUserMapper 根据id获取beeUserMapper表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserMapperService *BeeUserMapperService) GetBeeUserMapper(id string, shopUserId int) (beeUserMapper bee.BeeUserMapper, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeUserMapper).Error
	return
}

// GetBeeUserMapperInfoList 分页获取beeUserMapper表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserMapperService *BeeUserMapperService) GetBeeUserMapperInfoList(info beeReq.BeeUserMapperSearch, shopUserId int) (list []bee.BeeUserMapper, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeUserMapper{})
	db = db.Where("user_id = ?", shopUserId)
	var beeUserMappers []bee.BeeUserMapper
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Uid != nil {
		db = db.Where("uid = ?", info.Uid)
	}
	if info.Source != nil {
		db = db.Where("source = ?", info.Source)
	}
	if info.OpenId != "" {
		db = db.Where("open_id = ?", info.OpenId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeUserMappers).Error
	return beeUserMappers, total, err
}
