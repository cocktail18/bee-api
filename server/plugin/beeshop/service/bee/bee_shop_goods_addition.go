package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeShopGoodsAdditionService struct{}

// CreateBeeShopGoodsAddition 创建商品额外信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsAdditionService *BeeShopGoodsAdditionService) CreateBeeShopGoodsAddition(beeShopGoodsAddition *bee.BeeShopGoodsAddition) (err error) {
	beeShopGoodsAddition.DateAdd = utils.NowPtr()
	beeShopGoodsAddition.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeShopGoodsAddition).Error
	return err
}

// DeleteBeeShopGoodsAddition 删除商品额外信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsAdditionService *BeeShopGoodsAdditionService) DeleteBeeShopGoodsAddition(id string, shopUserId int) (err error) {
	err = GetBeeDB().Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeShopGoodsAdditionByIds 批量删除商品额外信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsAdditionService *BeeShopGoodsAdditionService) DeleteBeeShopGoodsAdditionByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeShopGoodsAddition 更新商品额外信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsAdditionService *BeeShopGoodsAdditionService) UpdateBeeShopGoodsAddition(beeShopGoodsAddition bee.BeeShopGoodsAddition, shopUserId int) (err error) {
	beeShopGoodsAddition.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Where("id = ? and user_id = ?", beeShopGoodsAddition.Id, shopUserId).Updates(&beeShopGoodsAddition).Error
	return err
}

// GetBeeShopGoodsAddition 根据id获取商品额外信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsAdditionService *BeeShopGoodsAdditionService) GetBeeShopGoodsAddition(id string, shopUserId int) (beeShopGoodsAddition bee.BeeShopGoodsAddition, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeShopGoodsAddition).Error
	return
}

// GetBeeShopGoodsAdditionInfoList 分页获取商品额外信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsAdditionService *BeeShopGoodsAdditionService) GetBeeShopGoodsAdditionInfoList(info beeReq.BeeShopGoodsAdditionSearch, shopUserId int) (list []bee.BeeShopGoodsAddition, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeShopGoodsAddition{})
	db = db.Where("user_id = ?", shopUserId)
	var beeShopGoodsAdditions []bee.BeeShopGoodsAddition
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Id != nil {
		db = db.Where("id = ?", info.Id)
	}
	if info.GoodsId != nil {
		db = db.Where("goods_id = ?", info.GoodsId)
	}
	if info.Pid != nil {
		db = db.Where("pid = ?", info.Pid)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["date_add"] = true
	orderMap["goods_id"] = true
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

	err = db.Find(&beeShopGoodsAdditions).Error
	return beeShopGoodsAdditions, total, err
}
