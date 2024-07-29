package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/samber/lo"
	"sync"
)

type FreightTplSrv struct {
	BaseSrv
}

var freightTplSrvOnce sync.Once
var freightTplSrvInstance *FreightTplSrv

func GetFreightTplSrv() *FreightTplSrv {
	freightTplSrvOnce.Do(func() {
		freightTplSrvInstance = &FreightTplSrv{}
	})
	return freightTplSrvInstance
}

func (srv *FreightTplSrv) GetBeeLogisticsDtoByIds(c context.Context, logisticsIds []int64) ([]*model.BeeLogistics, error) {
	var tpls []*model.BeeLogistics
	err := db.GetDB().Where("id in ?", logisticsIds).Find(&tpls).Error
	if err != nil {
		return nil, err
	}

	lo.ForEach(tpls, func(item *model.BeeLogistics, _ int) {
		var details []*model.BeeLogisticsDetail
		var freeShipping []*model.BeeShopLogisticsFreeShipping
		_ = util.UnmarshalJson(item.DetailsJsonStr, &details)
		_ = util.UnmarshalJson(item.FreeShippingSetting, &freeShipping)
		item.Details = details
		item.FreeShippings = freeShipping
	})
	return tpls, nil
}

func (srv *FreightTplSrv) GetBeeLogistics(c context.Context, logisticsId int64, regionId string) (*proto.GoodsLogistics, error) {
	tpls, err := srv.GetBeeLogisticsDtoByIds(c, []int64{logisticsId})
	if err != nil {
		return nil, err
	}
	if len(tpls) == 0 {
		return nil, nil
	}
	tpl := tpls[0]
	resp := &proto.GoodsLogistics{
		IsFree:  tpl.IsFree,
		FeeType: tpl.FeeType,
		Details: make([]*proto.GoodsLogisticsDetail, 0),
	}
	var defaultDetail *proto.GoodsLogisticsDetail
	resp.Details = lo.FilterMap(tpl.Details, func(item *model.BeeLogisticsDetail, _ int) (*proto.GoodsLogisticsDetail, bool) {
		if item.Default {
			defaultDetail = &proto.GoodsLogisticsDetail{
				FirstAmount: item.FirstAmount,
				FirstNumber: item.FirstNumber,
				AddAmount:   item.AddAmount,
				AddNumber:   item.AddNumber,
				Type:        item.Type,
			}
			return nil, false
		}
		if lo.Contains(item.CityIds, regionId) {
			return &proto.GoodsLogisticsDetail{
				FirstAmount: item.FirstAmount,
				FirstNumber: item.FirstNumber,
				AddAmount:   item.AddAmount,
				AddNumber:   item.AddNumber,
				Type:        item.Type,
			}, true
		}
		return nil, false
	})
	if len(resp.Details) <= 0 {
		resp.Details = append(resp.Details, defaultDetail)
	}
	resp.FillData()
	return resp, nil
}
