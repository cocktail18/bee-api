package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeUserCouponService struct{}

// CreateBeeUserCoupon 创建用户优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserCouponService *BeeUserCouponService) CreateBeeUserCoupon(beeUserCoupon *bee.BeeUserCoupon) (err error) {
	beeUserCoupon.DateAdd = utils.NowPtr()
	beeUserCoupon.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeUserCoupon).Error
	return err
}

// DeleteBeeUserCoupon 删除用户优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserCouponService *BeeUserCouponService) DeleteBeeUserCoupon(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeUserCoupon{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeUserCouponByIds 批量删除用户优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserCouponService *BeeUserCouponService) DeleteBeeUserCouponByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeUserCoupon{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeUserCoupon 更新用户优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserCouponService *BeeUserCouponService) UpdateBeeUserCoupon(beeUserCoupon bee.BeeUserCoupon, shopUserId int) (err error) {
	beeUserCoupon.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeUserCoupon{}).Where("id = ? and user_id = ?", beeUserCoupon.Id, shopUserId).Updates(&beeUserCoupon).Error
	return err
}

// GetBeeUserCoupon 根据id获取用户优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserCouponService *BeeUserCouponService) GetBeeUserCoupon(id string, shopUserId int) (beeUserCoupon bee.BeeUserCoupon, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeUserCoupon).Error
	return
}

// GetBeeUserCouponInfoList 分页获取用户优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserCouponService *BeeUserCouponService) GetBeeUserCouponInfoList(info beeReq.BeeUserCouponSearch, shopUserId int) (list []bee.BeeUserCoupon, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeUserCoupon{})
	db = db.Where("user_id = ?", shopUserId)
	var beeUserCoupons []bee.BeeUserCoupon
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeUserCoupons).Error
	return beeUserCoupons, total, err
}
