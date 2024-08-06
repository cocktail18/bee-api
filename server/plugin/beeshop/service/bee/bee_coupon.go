package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeCouponService struct{}

// CreateBeeCoupon 创建优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCouponService *BeeCouponService) CreateBeeCoupon(beeCoupon *bee.BeeCoupon) (err error) {
	beeCoupon.DateAdd = utils.NowPtr()
	beeCoupon.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeCoupon).Error
	return err
}

// DeleteBeeCoupon 删除优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCouponService *BeeCouponService) DeleteBeeCoupon(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeCoupon{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeCouponByIds 批量删除优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCouponService *BeeCouponService) DeleteBeeCouponByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeCoupon{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeCoupon 更新优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCouponService *BeeCouponService) UpdateBeeCoupon(beeCoupon bee.BeeCoupon, shopUserId int) (err error) {
	beeCoupon.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeCoupon{}).Where("id = ? and user_id = ?", beeCoupon.Id, shopUserId).Updates(&beeCoupon).Error
	return err
}

// GetBeeCoupon 根据id获取优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCouponService *BeeCouponService) GetBeeCoupon(id string, shopUserId int) (beeCoupon bee.BeeCoupon, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeCoupon).Error
	return
}

// GetBeeCouponInfoList 分页获取优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCouponService *BeeCouponService) GetBeeCouponInfoList(info beeReq.BeeCouponSearch, shopUserId int) (list []bee.BeeCoupon, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeCoupon{})
	db = db.Where("user_id = ?", shopUserId)
	var beeCoupons []bee.BeeCoupon
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeCoupons).Error
	return beeCoupons, total, err
}
