package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeShopGoodsService struct{}

// CreateBeeShopGoods 创建商品表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsService *BeeShopGoodsService) CreateBeeShopGoods(beeShopGoods *bee.BeeShopGoods) (err error) {
	beeShopGoods.DateAdd = utils.NowPtr()
	beeShopGoods.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeShopGoods).Error
	return err
}

// DeleteBeeShopGoods 删除商品表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsService *BeeShopGoodsService) DeleteBeeShopGoods(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeShopGoods{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeShopGoodsByIds 批量删除商品表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsService *BeeShopGoodsService) DeleteBeeShopGoodsByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeShopGoods{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeShopGoods 更新商品表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsService *BeeShopGoodsService) UpdateBeeShopGoods(beeShopGoods bee.BeeShopGoods, shopUserId int) (err error) {
	beeShopGoods.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeShopGoods{}).Where("id = ? and user_id = ?", beeShopGoods.Id, shopUserId).Updates(&beeShopGoods).Error
	return err
}

// GetBeeShopGoods 根据id获取商品表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsService *BeeShopGoodsService) GetBeeShopGoods(id string, shopUserId int) (beeShopGoods bee.BeeShopGoods, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeShopGoods).Error
	return
}

// GetBeeShopGoodsInfoList 分页获取商品表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsService *BeeShopGoodsService) GetBeeShopGoodsInfoList(info beeReq.BeeShopGoodsSearch, shopUserId int) (list []bee.BeeShopGoods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeShopGoods{})
	db = db.Where("user_id = ?", shopUserId)
	var beeShopGoodss []bee.BeeShopGoods
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.BarCode != "" {
		db = db.Where("bar_code LIKE ?", "%"+info.BarCode+"%")
	}
	if info.CategoryId != nil {
		db = db.Where("category_id = ?", info.CategoryId)
	}
	if info.LogisticsId != nil {
		db = db.Where("logistics_id = ?", info.LogisticsId)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.ShopId != nil {
		db = db.Where("shop_id = ?", info.ShopId)
	}
	if info.StartSellBeginTime != nil && info.EndSellBeginTime != nil {
		db = db.Where("sell_begin_time BETWEEN ? AND ? ", info.StartSellBeginTime, info.EndSellBeginTime)
	}
	if info.StartSellEndTime != nil && info.EndSellEndTime != nil {
		db = db.Where("sell_end_time BETWEEN ? AND ? ", info.StartSellEndTime, info.EndSellEndTime)
	}
	if info.StartDateAdd != nil && info.EndDateAdd != nil {
		db = db.Where("date_add BETWEEN ? AND ? ", info.StartDateAdd, info.EndDateAdd)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["date_add"] = true
	orderMap["paixu"] = true
	orderMap["shop_id"] = true
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

	err = db.Find(&beeShopGoodss).Error
	return beeShopGoodss, total, err
}
