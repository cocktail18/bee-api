package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/samber/lo"
	"sync"
)

type ShopSrv struct {
	BaseSrv
}

var shopSrvOnce sync.Once
var shopSrvInstance *ShopSrv

func GetShopSrv() *ShopSrv {
	shopSrvOnce.Do(func() {
		shopSrvInstance = &ShopSrv{}
	})
	return shopSrvInstance
}

func (srv *ShopSrv) GetShopInfo(c context.Context, id int64, lat, lon float64) (*model.BeeShopInfo, error) {
	var info = &model.BeeShopInfo{}
	err := db.GetDB().Where("id = ? and user_id = ?", id, kit.GetUserId(c)).Take(info).Error
	if err != nil {
		return nil, err
	}
	info.FillData(lat, lon)
	return info, err
}

func (srv *ShopSrv) List(c context.Context, req *proto.ListShopReq) (proto.ListShopResp, error) {
	var infos []*model.BeeShopInfo
	err := db.GetDB().Where("user_id = ? and is_deleted = 0", kit.GetUserId(c)).Find(&infos).Error
	if err != nil {
		return nil, err
	}
	lo.ForEach(infos, func(info *model.BeeShopInfo, _ int) {
		info.FillData(req.Latitude, req.Longitude)
	})
	return infos, nil
}
