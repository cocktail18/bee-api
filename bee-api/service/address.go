package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
	"time"
)

type AddressSrv struct {
	BaseSrv
}

var addressSrvOnce sync.Once
var addressSrvInstance *AddressSrv

func GetAddressSrv() *AddressSrv {
	addressSrvOnce.Do(func() {
		addressSrvInstance = &AddressSrv{}
	})
	return addressSrvInstance
}

func (srv *AddressSrv) SetDefault(c context.Context, uid int64, addressId int64) error {
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.BeeUserAddress{}).Where("uid = ? and is_default = 1", uid).
			Updates(map[string]interface{}{"is_default": 0}).Error; err != nil {
			return err
		}
		return tx.Model(&model.BeeUserAddress{}).Where("uid = ? and id = ?", uid, addressId).
			Updates(map[string]interface{}{"is_default": 1}).Error
	})
}

func (srv *AddressSrv) AddAddress(c context.Context, address *model.BeeUserAddress) (*proto.UserAddressResp, error) {
	resp := &proto.UserAddressResp{}
	regionDistrict, err := GetRegionSrv().GetRegion(address.DistrictId)
	if err != nil {
		return nil, err
	}
	regionCity, err := GetRegionSrv().GetRegion(address.CityId)
	if err != nil {
		return nil, err
	}
	regionProvince, err := GetRegionSrv().GetRegion(address.ProvinceId)
	if err != nil {
		return nil, err
	}
	if regionCity != nil {
		address.CityStr = regionCity.Name
	}
	if regionDistrict != nil {
		address.AreaStr = regionDistrict.Name
	}
	if regionProvince != nil {
		address.ProvinceStr = regionProvince.Name
	}
	address.Uid = kit.GetUid(c)
	address.UserId = kit.GetUserId(c)
	address.DateAdd = common.JsonTime(time.Now())
	address.DateUpdate = common.JsonTime(time.Now())
	err = db.GetDB().Create(address).Error
	if err != nil {
		return nil, err
	}

	if address.IsDefault {
		_ = srv.SetDefault(c, address.Uid, address.Id)
	}

	util.CopyStruct(resp, address)
	resp.RegionCity = regionCity
	resp.RegionDistrict = regionDistrict
	resp.RegionProvince = regionProvince
	return resp, nil
}

func (srv *AddressSrv) SaveAddress(c context.Context, address *model.BeeUserAddress) (*proto.UserAddressResp, error) {
	var oldAddress model.BeeUserAddress
	err := db.GetDB().Where("id = ? and uid = ? and is_deleted = 0", address.Id, address.Uid).Take(&oldAddress).Error
	if err != nil {
		return nil, err
	}

	resp := &proto.UserAddressResp{}
	regionDistrict, err := GetRegionSrv().GetRegion(address.DistrictId)
	if err != nil {
		return nil, err
	}
	regionCity, err := GetRegionSrv().GetRegion(address.CityId)
	if err != nil {
		return nil, err
	}
	regionProvince, err := GetRegionSrv().GetRegion(address.ProvinceId)
	if err != nil {
		return nil, err
	}
	if regionCity != nil {
		address.CityStr = regionCity.Name
	}
	if regionDistrict != nil {
		address.AreaStr = regionDistrict.Name
	}
	if regionProvince != nil {
		address.ProvinceStr = regionProvince.Name
	}
	address.DateUpdate = common.JsonTime(time.Now())
	err = db.GetDB().Save(address).Error
	if err != nil {
		return nil, err
	}
	if address.IsDefault {
		_ = srv.SetDefault(c, address.Uid, address.Id)
	}
	util.CopyStruct(resp, address)
	resp.RegionCity = regionCity
	resp.RegionDistrict = regionDistrict
	resp.RegionProvince = regionProvince
	return resp, nil
}

func (srv *AddressSrv) ListAddress(c context.Context) ([]*proto.UserAddressResp, error) {
	var list []*model.BeeUserAddress
	err := db.GetDB().Where("uid = ? and is_deleted = 0", kit.GetUid(c)).Find(&list).Error
	if err != nil {
		return nil, err
	}
	resp := make([]*proto.UserAddressResp, 0, len(list))
	for _, address := range list { //@todo
		regionDistrict, err := GetRegionSrv().GetRegion(address.DistrictId)
		if err != nil {
			return nil, err
		}
		regionCity, err := GetRegionSrv().GetRegion(address.CityId)
		if err != nil {
			return nil, err
		}
		regionProvince, err := GetRegionSrv().GetRegion(address.ProvinceId)
		if err != nil {
			return nil, err
		}
		if regionCity != nil {
			address.CityStr = regionCity.Name
		}
		if regionDistrict != nil {
			address.AreaStr = regionDistrict.Name
		}
		if regionProvince != nil {
			address.ProvinceStr = regionProvince.Name
		}
		item := &proto.UserAddressResp{}
		util.CopyStruct(item, address)
		item.RegionCity = regionCity
		item.RegionDistrict = regionDistrict
		item.RegionProvince = regionProvince
		resp = append(resp, item)
	}
	return resp, nil
}

func (srv *AddressSrv) GetAddressDto(c context.Context, uid int64, id int64) (*model.BeeUserAddress, error) {
	var address model.BeeUserAddress
	err := db.GetDB().Where("id = ? and uid = ? and is_deleted = 0", id, uid).Take(&address).Error
	return &address, err
}

func (srv *AddressSrv) GetAddress(c context.Context, id int64) (*proto.UserAddressDetailResp, error) {
	var address model.BeeUserAddress
	err := db.GetDB().Where("id = ? and uid = ? and is_deleted = 0", id, kit.GetUid(c)).Take(&address).Error
	if err != nil {
		return nil, err
	}
	addressResp := &proto.UserAddressResp{}
	regionDistrict, err := GetRegionSrv().GetRegion(address.DistrictId)
	if err != nil {
		return nil, err
	}
	regionCity, err := GetRegionSrv().GetRegion(address.CityId)
	if err != nil {
		return nil, err
	}
	regionProvince, err := GetRegionSrv().GetRegion(address.ProvinceId)
	if err != nil {
		return nil, err
	}
	if regionCity != nil {
		address.CityStr = regionCity.Name
	}
	if regionDistrict != nil {
		address.AreaStr = regionDistrict.Name
	}
	if regionProvince != nil {
		address.ProvinceStr = regionProvince.Name
	}
	util.CopyStruct(addressResp, &address)
	addressResp.RegionCity = regionCity
	addressResp.RegionDistrict = regionDistrict
	addressResp.RegionProvince = regionProvince
	resp := &proto.UserAddressDetailResp{
		Info:    addressResp,
		ExtJson: make(map[string]interface{}),
	}
	return resp, nil
}

func (srv *AddressSrv) DeleteAddress(userId int64, id int64) error {
	var address model.BeeUserAddress
	err := db.GetDB().Where("id = ? and uid = ?", id, userId).Take(&address).Error
	if err != nil {
		return err
	}
	return db.GetDB().Where("id = ? and uid = ?", id, userId).Delete(&address).Error
}

func (srv *AddressSrv) Default(c context.Context) (*proto.UserAddressDetailResp, error) {
	var address model.BeeUserAddress
	uid := kit.GetUid(c)
	err := db.GetDB().Where("is_default = 1 and uid = ? and is_deleted = 0", uid).Take(&address).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, enum.ErrEmpty
	}
	if err != nil {
		return nil, err
	}
	return srv.GetAddress(c, address.Id)
}
