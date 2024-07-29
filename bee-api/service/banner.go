package service

import (
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/model"
	"sync"
)

type BannerSrv struct {
	BaseSrv
}

var bannerSrvOnce sync.Once
var bannerSrvInstance *BannerSrv

func GetBannerSrv() *BannerSrv {
	bannerSrvOnce.Do(func() {
		bannerSrvInstance = &BannerSrv{}
	})
	return bannerSrvInstance
}

func (srv *BannerSrv) GetBannerList() ([]*model.BeeBanner, error) {
	var list []*model.BeeBanner
	err := db.GetDB().Where("status", 0).Order("paixu desc").Find(&list).Error
	return list, err
}
