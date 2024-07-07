package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"gorm.io/gorm"
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
	var details []*model.BeeLogisticsDetail
	err = db.GetDB().Where("logistics_id in ?", logisticsIds).Take(&details).Error
	if err != nil {
		return nil, err
	}
	var exceptions []*model.BeeShopLogisticsException
	err = db.GetDB().Where("logistics_id in ?", logisticsIds).Find(&exceptions).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	var freeShipping []*model.BeeShopLogisticsFreeShipping
	err = db.GetDB().Where("logistics_id in ?", logisticsIds).Find(&freeShipping).Error
	if err != nil {
		return nil, err
	}

	detailById := lo.GroupBy(details, func(item *model.BeeLogisticsDetail) int64 {
		return item.LogisticsId
	})
	exceptionsById := lo.GroupBy(exceptions, func(item *model.BeeShopLogisticsException) int64 {
		return item.LogisticsId
	})
	freeShippingById := lo.GroupBy(freeShipping, func(item *model.BeeShopLogisticsFreeShipping) int64 {
		return item.LogisticsId
	})
	lo.ForEach(tpls, func(item *model.BeeLogistics, _ int) {
		item.Details = detailById[item.Id]
		exceptionsByDetails := lo.GroupBy(exceptionsById[item.Id], func(item *model.BeeShopLogisticsException) int64 {
			return item.DetailId
		})
		lo.ForEach(item.Details, func(item *model.BeeLogisticsDetail, _ int) {
			item.Exceptions = exceptionsByDetails[item.Id]
		})
		item.FreeShippings = freeShippingById[item.Id]
	})
	return tpls, nil
}

func (srv *FreightTplSrv) GetBeeLogistics(c context.Context, logisticsId int64, regionId string) (*model.BeeLogistics, error) {
	var tpl model.BeeLogistics
	err := db.GetDB().Where("id", logisticsId).Take(&tpl).Error
	if err != nil {
		return nil, err
	}
	var detail model.BeeLogisticsDetail
	err = db.GetDB().Where("logistics_id", tpl.Id).Take(&detail).Error
	if err != nil {
		return nil, err
	}
	var exceptions []*model.BeeShopLogisticsException
	err = db.GetDB().Where("logistics_id", tpl.Id).Find(&exceptions).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	var freeShipping []*model.BeeShopLogisticsFreeShipping
	err = db.GetDB().Where("logistics_id", tpl.Id).Find(&freeShipping).Error
	if err != nil {
		return nil, err
	}
	resp := &model.BeeLogistics{}
	resp.FeeType = tpl.FeeType
	resp.IsFree = tpl.IsFree
	resp.FeeTypeStr = enum.BeeGoodsFeeTypeMap[resp.FeeType]
	resp.Details = make([]*model.BeeLogisticsDetail, 0)
	return resp, nil
}
