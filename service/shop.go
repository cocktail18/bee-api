package service

import (
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/gin-gonic/gin"
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

func (srv *ShopSrv) GetShopInfo(id int64) (*model.BeeShopInfo, error) {
	var info = &model.BeeShopInfo{}
	err := db.GetDB().Where("id", id).Take(info).Error
	if err != nil {
		return nil, err
	}
	info.FillData(39.9042, 116.4074)
	return info, err
}

func (srv *ShopSrv) List(c *gin.Context, req *proto.ListShopReq) (proto.ListShopResp, error) {
	var infos []*model.BeeShopInfo
	err := db.GetDB().Where("user_id = ?", kit.GetUserId(c)).Find(&infos).Error
	if err != nil {
		return nil, err
	}
	lo.ForEach(infos, func(info *model.BeeShopInfo, _ int) {
		info.FillData(req.Latitude, req.Longitude)
	})
	return infos, nil
}
