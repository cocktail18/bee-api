package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeShopGoodsPropService struct{}

// CreateBeeShopGoodsProp 创建商品属性记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsPropService *BeeShopGoodsPropService) CreateBeeShopGoodsProp(beeShopGoodsProp *bee.BeeShopGoodsProp) (err error) {
	beeShopGoodsProp.DateAdd = utils.NowPtr()
	beeShopGoodsProp.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeShopGoodsProp).Error
	return err
}

// DeleteBeeShopGoodsProp 删除商品属性记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsPropService *BeeShopGoodsPropService) DeleteBeeShopGoodsProp(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeShopGoodsProp{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeShopGoodsPropByIds 批量删除商品属性记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsPropService *BeeShopGoodsPropService) DeleteBeeShopGoodsPropByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeShopGoodsProp{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeShopGoodsProp 更新商品属性记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsPropService *BeeShopGoodsPropService) UpdateBeeShopGoodsProp(beeShopGoodsProp bee.BeeShopGoodsProp, shopUserId int) (err error) {
	beeShopGoodsProp.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeShopGoodsProp{}).Where("id = ? and user_id = ?", beeShopGoodsProp.Id, shopUserId).Updates(&beeShopGoodsProp).Error
	return err
}

// GetBeeShopGoodsProp 根据id获取商品属性记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsPropService *BeeShopGoodsPropService) GetBeeShopGoodsProp(id string, shopUserId int) (beeShopGoodsProp bee.BeeShopGoodsProp, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeShopGoodsProp).Error
	return
}

// GetBeeShopGoodsPropInfoList 分页获取商品属性记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsPropService *BeeShopGoodsPropService) GetBeeShopGoodsPropInfoList(info beeReq.BeeShopGoodsPropSearch, shopUserId int) (list []bee.BeeShopGoodsProp, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeShopGoodsProp{})
	db = db.Where("user_id = ?", shopUserId)
	var beeShopGoodsProps []bee.BeeShopGoodsProp
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Id != nil {
		db = db.Where("id = ?", info.Id)
	}
	if info.PropertyId != nil {
		db = db.Where("property_id = ?", info.PropertyId)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeShopGoodsProps).Error
	return beeShopGoodsProps, total, err
}
