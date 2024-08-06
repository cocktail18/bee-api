package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeShopGoodsCategoryService struct{}

// CreateBeeShopGoodsCategory 创建商品分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsCategoryService *BeeShopGoodsCategoryService) CreateBeeShopGoodsCategory(beeShopGoodsCategory *bee.BeeShopGoodsCategory) (err error) {
	beeShopGoodsCategory.DateAdd = utils.NowPtr()
	beeShopGoodsCategory.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeShopGoodsCategory).Error
	return err
}

// DeleteBeeShopGoodsCategory 删除商品分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsCategoryService *BeeShopGoodsCategoryService) DeleteBeeShopGoodsCategory(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeShopGoodsCategory{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeShopGoodsCategoryByIds 批量删除商品分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsCategoryService *BeeShopGoodsCategoryService) DeleteBeeShopGoodsCategoryByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeShopGoodsCategory{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeShopGoodsCategory 更新商品分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsCategoryService *BeeShopGoodsCategoryService) UpdateBeeShopGoodsCategory(beeShopGoodsCategory bee.BeeShopGoodsCategory, shopUserId int) (err error) {
	beeShopGoodsCategory.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeShopGoodsCategory{}).Where("id = ? and user_id = ?", beeShopGoodsCategory.Id, shopUserId).Updates(&beeShopGoodsCategory).Error
	return err
}

// GetBeeShopGoodsCategory 根据id获取商品分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsCategoryService *BeeShopGoodsCategoryService) GetBeeShopGoodsCategory(id string, shopUserId int) (beeShopGoodsCategory bee.BeeShopGoodsCategory, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeShopGoodsCategory).Error
	return
}

// GetBeeShopGoodsCategoryInfoList 分页获取商品分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsCategoryService *BeeShopGoodsCategoryService) GetBeeShopGoodsCategoryInfoList(info beeReq.BeeShopGoodsCategorySearch, shopUserId int) (list []bee.BeeShopGoodsCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeShopGoodsCategory{})
	db = db.Where("user_id = ?", shopUserId)
	var beeShopGoodsCategorys []bee.BeeShopGoodsCategory
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeShopGoodsCategorys).Error
	return beeShopGoodsCategorys, total, err
}
