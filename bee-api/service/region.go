package service

import (
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/model"
	"sync"
)

type RegionSrv struct {
	BaseSrv
}

var regionSrvOnce sync.Once
var regionSrvInstance *RegionSrv

func GetRegionSrv() *RegionSrv {
	regionSrvOnce.Do(func() {
		regionSrvInstance = &RegionSrv{}
	})
	return regionSrvInstance
}

func (srv *RegionSrv) GetAllProvince() ([]*model.BeeRegion, error) {
	var list []*model.BeeRegion
	err := db.GetDB().Where("level = ? and is_deleted = 0", 1).Find(&list).Error
	return list, err
}

func (srv *RegionSrv) GetAllChild(pid string) ([]*model.BeeRegion, error) {
	var list []*model.BeeRegion
	err := db.GetDB().Where("pid = ? and is_deleted = 0", pid).Find(&list).Error
	return list, err
}

func (srv *RegionSrv) GetRegion(id string) (*model.BeeRegion, error) {
	if id == "" {
		return nil, nil
	}
	var item model.BeeRegion
	err := db.GetDB().Where("id = ? and is_deleted = 0", id).Take(&item).Error
	return &item, err
}
