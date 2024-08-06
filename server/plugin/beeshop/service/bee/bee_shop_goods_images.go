package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeShopGoodsImagesService struct{}

// CreateBeeShopGoodsImages 创建商品图记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsImagesService *BeeShopGoodsImagesService) CreateBeeShopGoodsImages(beeShopGoodsImages *bee.BeeShopGoodsImages) (err error) {
	beeShopGoodsImages.DateAdd = utils.NowPtr()
	beeShopGoodsImages.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeShopGoodsImages).Error
	return err
}

// DeleteBeeShopGoodsImages 删除商品图记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsImagesService *BeeShopGoodsImagesService) DeleteBeeShopGoodsImages(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeShopGoodsImages{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeShopGoodsImagesByIds 批量删除商品图记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsImagesService *BeeShopGoodsImagesService) DeleteBeeShopGoodsImagesByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeShopGoodsImages{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeShopGoodsImages 更新商品图记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsImagesService *BeeShopGoodsImagesService) UpdateBeeShopGoodsImages(beeShopGoodsImages bee.BeeShopGoodsImages, shopUserId int) (err error) {
	beeShopGoodsImages.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeShopGoodsImages{}).Where("id = ? and user_id = ?", beeShopGoodsImages.Id, shopUserId).Updates(&beeShopGoodsImages).Error
	return err
}

// GetBeeShopGoodsImages 根据id获取商品图记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsImagesService *BeeShopGoodsImagesService) GetBeeShopGoodsImages(id string, shopUserId int) (beeShopGoodsImages bee.BeeShopGoodsImages, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeShopGoodsImages).Error
	return
}

// GetBeeShopGoodsImagesInfoList 分页获取商品图记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsImagesService *BeeShopGoodsImagesService) GetBeeShopGoodsImagesInfoList(info beeReq.BeeShopGoodsImagesSearch, shopUserId int) (list []bee.BeeShopGoodsImages, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeShopGoodsImages{})
	db = db.Where("user_id = ?", shopUserId)
	var beeShopGoodsImagess []bee.BeeShopGoodsImages
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.GoodsId != nil {
		db = db.Where("goods_id = ?", info.GoodsId)
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

	err = db.Find(&beeShopGoodsImagess).Error
	return beeShopGoodsImagess, total, err
}
