package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeUserAddressService struct{}

// CreateBeeUserAddress 创建用户地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserAddressService *BeeUserAddressService) CreateBeeUserAddress(beeUserAddress *bee.BeeUserAddress) (err error) {
	beeUserAddress.DateAdd = utils.NowPtr()
	beeUserAddress.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeUserAddress).Error
	return err
}

// DeleteBeeUserAddress 删除用户地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserAddressService *BeeUserAddressService) DeleteBeeUserAddress(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeUserAddress{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeUserAddressByIds 批量删除用户地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserAddressService *BeeUserAddressService) DeleteBeeUserAddressByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeUserAddress{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeUserAddress 更新用户地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserAddressService *BeeUserAddressService) UpdateBeeUserAddress(beeUserAddress bee.BeeUserAddress, shopUserId int) (err error) {
	beeUserAddress.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeUserAddress{}).Where("id = ? and user_id = ?", beeUserAddress.Id, shopUserId).Updates(&beeUserAddress).Error
	return err
}

// GetBeeUserAddress 根据id获取用户地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserAddressService *BeeUserAddressService) GetBeeUserAddress(id string, shopUserId int) (beeUserAddress bee.BeeUserAddress, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeUserAddress).Error
	return
}

// GetBeeUserAddressInfoList 分页获取用户地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserAddressService *BeeUserAddressService) GetBeeUserAddressInfoList(info beeReq.BeeUserAddressSearch, shopUserId int) (list []bee.BeeUserAddress, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeUserAddress{})
	db = db.Where("user_id = ?", shopUserId)
	var beeUserAddresss []bee.BeeUserAddress
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.Uid != nil {
		db = db.Where("uid = ?", info.Uid)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeUserAddresss).Error
	return beeUserAddresss, total, err
}
