package service

import (
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"sync"
)

type CmsSrv struct {
	BaseSrv
}

var cmsSrvOnce sync.Once
var cmsSrvInstance *CmsSrv

func GetCmsSrv() *CmsSrv {
	cmsSrvOnce.Do(func() {
		cmsSrvInstance = &CmsSrv{}
	})
	return cmsSrvInstance
}

func (srv *CmsSrv) GetCmsInfo(key string) (*proto.CmsResp, error) {
	var item model.BeeCmsInfo
	err := db.GetDB().Where(map[string]interface{}{
		"status": 0,
		"key":    key,
	}).Take(&item).Error
	if err != nil {
		return nil, err
	}
	return &proto.CmsResp{Info: &item}, err
}
