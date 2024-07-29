package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeRegionService struct{}

// CreateBeeRegion 创建地址库记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeRegionService *BeeRegionService) CreateBeeRegion(beeRegion *bee.BeeRegion) (err error) {
	beeRegion.DateAdd = utils.NowPtr()
	beeRegion.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeRegion).Error
	return err
}

// DeleteBeeRegion 删除地址库记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeRegionService *BeeRegionService) DeleteBeeRegion(id string) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeRegion{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeRegionByIds 批量删除地址库记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeRegionService *BeeRegionService) DeleteBeeRegionByIds(ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeRegion{}).Where("id = ?", ids).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeRegion 更新地址库记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeRegionService *BeeRegionService) UpdateBeeRegion(beeRegion bee.BeeRegion) (err error) {
	beeRegion.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeRegion{}).Where("id = ?", beeRegion.Id).Updates(&beeRegion).Error
	return err
}

// GetBeeRegion 根据id获取地址库记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeRegionService *BeeRegionService) GetBeeRegion(id string) (beeRegion bee.BeeRegion, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ?", id).First(&beeRegion).Error
	return
}

// GetBeeRegionInfoList 分页获取地址库记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeRegionService *BeeRegionService) GetBeeRegionInfoList(info beeReq.BeeRegionSearch) (list []bee.BeeRegion, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeRegion{})
	var beeRegions []bee.BeeRegion
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.IsDeleted != nil {
		db = db.Where("is_deleted = ?", info.IsDeleted)
	}
	if info.StartDateAdd != nil && info.EndDateAdd != nil {
		db = db.Where("date_add BETWEEN ? AND ? ", info.StartDateAdd, info.EndDateAdd)
	}
	if info.StartDateUpdate != nil && info.EndDateUpdate != nil {
		db = db.Where("date_update BETWEEN ? AND ? ", info.StartDateUpdate, info.EndDateUpdate)
	}
	if info.FirstLetter != "" {
		db = db.Where("first_letter LIKE ?", "%"+info.FirstLetter+"%")
	}
	if info.Jianpin != "" {
		db = db.Where("jianpin LIKE ?", "%"+info.Jianpin+"%")
	}
	if info.Level != nil {
		db = db.Where("level = ?", info.Level)
	}
	if info.LevelList != nil {
		db = db.Where("`level` in ?", info.LevelList)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.NameEn != "" {
		db = db.Where("name_en LIKE ?", "%"+info.NameEn+"%")
	}
	if info.Pinyin != "" {
		db = db.Where("pinyin LIKE ?", "%"+info.Pinyin+"%")
	}
	if info.Pid != "" {
		db = db.Where("pid = ?", info.Pid)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["pid"] = true
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

	err = db.Find(&beeRegions).Error
	return beeRegions, total, err
}
