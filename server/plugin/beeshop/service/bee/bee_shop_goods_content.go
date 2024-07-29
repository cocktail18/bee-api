package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeShopGoodsContentService struct{}

// CreateBeeShopGoodsContent 创建商品详情描述记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsContentService *BeeShopGoodsContentService) CreateBeeShopGoodsContent(beeShopGoodsContent *bee.BeeShopGoodsContent) (err error) {
	beeShopGoodsContent.DateAdd = utils.NowPtr()
	beeShopGoodsContent.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeShopGoodsContent).Error
	return err
}

// DeleteBeeShopGoodsContent 删除商品详情描述记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsContentService *BeeShopGoodsContentService) DeleteBeeShopGoodsContent(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopGoodsContent{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeShopGoodsContentByIds 批量删除商品详情描述记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsContentService *BeeShopGoodsContentService) DeleteBeeShopGoodsContentByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopGoodsContent{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeShopGoodsContent 更新商品详情描述记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsContentService *BeeShopGoodsContentService) UpdateBeeShopGoodsContent(beeShopGoodsContent bee.BeeShopGoodsContent, shopUserId int) (err error) {
	beeShopGoodsContent.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopGoodsContent{}).Where("id = ? and user_id = ?", beeShopGoodsContent.Id, shopUserId).Updates(&beeShopGoodsContent).Error
	return err
}

// GetBeeShopGoodsContent 根据id获取商品详情描述记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsContentService *BeeShopGoodsContentService) GetBeeShopGoodsContent(id string, shopUserId int) (beeShopGoodsContent bee.BeeShopGoodsContent, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeShopGoodsContent).Error
	return
}

// GetBeeShopGoodsContentInfoList 分页获取商品详情描述记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsContentService *BeeShopGoodsContentService) GetBeeShopGoodsContentInfoList(info beeReq.BeeShopGoodsContentSearch, shopUserId int) (list []bee.BeeShopGoodsContent, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopGoodsContent{})
	db = db.Where("user_id = ?", shopUserId)
	var beeShopGoodsContents []bee.BeeShopGoodsContent
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.GoodsId != nil {
		db = db.Where("goods_id = ?", info.GoodsId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeShopGoodsContents).Error
	return beeShopGoodsContents, total, err
}
