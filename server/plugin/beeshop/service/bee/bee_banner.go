package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeBannerService struct{}

// CreateBeeBanner 创建商店banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeBannerService *BeeBannerService) CreateBeeBanner(beeBanner *bee.BeeBanner) (err error) {
	beeBanner.DateAdd = utils.NowPtr()
	beeBanner.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeBanner).Error
	return err
}

// DeleteBeeBanner 删除商店banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeBannerService *BeeBannerService) DeleteBeeBanner(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeBanner{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeBannerByIds 批量删除商店banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeBannerService *BeeBannerService) DeleteBeeBannerByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeBanner{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeBanner 更新商店banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeBannerService *BeeBannerService) UpdateBeeBanner(beeBanner bee.BeeBanner, shopUserId int) (err error) {
	beeBanner.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeBanner{}).Where("id = ? and user_id = ?", beeBanner.Id, shopUserId).Updates(&beeBanner).Error
	return err
}

// GetBeeBanner 根据id获取商店banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeBannerService *BeeBannerService) GetBeeBanner(id string, shopUserId int) (beeBanner bee.BeeBanner, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeBanner).Error
	return
}

// GetBeeBannerInfoList 分页获取商店banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeBannerService *BeeBannerService) GetBeeBannerInfoList(info beeReq.BeeBannerSearch, shopUserId int) (list []bee.BeeBanner, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeBanner{})
	db = db.Where("user_id = ?", shopUserId)
	var beeBanners []bee.BeeBanner
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ShopId != nil {
		db = db.Where("shop_id = ?", info.ShopId)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["paixu"] = true
	orderMap["shop_id"] = true
	orderMap["status"] = true
	orderMap["type"] = true
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

	err = db.Find(&beeBanners).Error
	return beeBanners, total, err
}
