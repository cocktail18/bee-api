package service

import (
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

func (srv *FreightTplSrv) GetBeeLogistics(c *gin.Context, logisticsId int64, regionId string) (*model.BeeLogistics, error) {
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
	//@todo
	return resp, nil
}
