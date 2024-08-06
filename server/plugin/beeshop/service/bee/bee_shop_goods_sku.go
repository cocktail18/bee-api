package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeShopGoodsSkuService struct{}

// CreateBeeShopGoodsSku 创建商品sku记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsSkuService *BeeShopGoodsSkuService) CreateBeeShopGoodsSku(beeShopGoodsSku *bee.BeeShopGoodsSku) (err error) {
	beeShopGoodsSku.DateAdd = utils.NowPtr()
	beeShopGoodsSku.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeShopGoodsSku).Error
	return err
}

// DeleteBeeShopGoodsSku 删除商品sku记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsSkuService *BeeShopGoodsSkuService) DeleteBeeShopGoodsSku(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeShopGoodsSku{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeShopGoodsSkuByIds 批量删除商品sku记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsSkuService *BeeShopGoodsSkuService) DeleteBeeShopGoodsSkuByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeShopGoodsSku{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeShopGoodsSku 更新商品sku记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsSkuService *BeeShopGoodsSkuService) UpdateBeeShopGoodsSku(beeShopGoodsSku bee.BeeShopGoodsSku, shopUserId int) (err error) {
	beeShopGoodsSku.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeShopGoodsSku{}).Where("id = ? and user_id = ?", beeShopGoodsSku.Id, shopUserId).Updates(&beeShopGoodsSku).Error
	return err
}

// GetBeeShopGoodsSku 根据id获取商品sku记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsSkuService *BeeShopGoodsSkuService) GetBeeShopGoodsSku(id string, shopUserId int) (beeShopGoodsSku bee.BeeShopGoodsSku, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeShopGoodsSku).Error
	return
}

// GetBeeShopGoodsSkuInfoList 分页获取商品sku记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsSkuService *BeeShopGoodsSkuService) GetBeeShopGoodsSkuInfoList(info beeReq.BeeShopGoodsSkuSearch, shopUserId int) (list []bee.BeeShopGoodsSku, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeShopGoodsSku{})
	db = db.Where("user_id = ?", shopUserId)
	var beeShopGoodsSkus []bee.BeeShopGoodsSku
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Id != nil {
		db = db.Where("id = ?", info.Id)
	}
	if info.GoodsId != nil {
		db = db.Where("goods_id = ?", info.GoodsId)
	}
	if info.Code != "" {
		db = db.Where("code = ?", info.Code)
	}
	if info.StartPrice != nil && info.EndPrice != nil {
		db = db.Where("price BETWEEN ? AND ? ", info.StartPrice, info.EndPrice)
	}
	if info.StartStores != nil && info.EndStores != nil {
		db = db.Where("stores BETWEEN ? AND ? ", info.StartStores, info.EndStores)
	}
	if info.StartWeight != nil && info.EndWeight != nil {
		db = db.Where("weight BETWEEN ? AND ? ", info.StartWeight, info.EndWeight)
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

	err = db.Find(&beeShopGoodsSkus).Error
	return beeShopGoodsSkus, total, err
}
