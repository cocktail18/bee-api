package service

import (
	config2 "gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type DfsSrv struct {
	BaseSrv
}

var dfsSrvOnce sync.Once
var dfsSrvInstance *DfsSrv

func GetDfsSrv() *DfsSrv {
	dfsSrvOnce.Do(func() {
		dfsSrvInstance = &DfsSrv{}
	})
	return dfsSrvInstance
}

func (s DfsSrv) SaveUploadedFile(c *gin.Context, domain string, filename, dst string, hours int64) (*proto.UploadFileResp, error) {
	data := model.BeeUploadFile{
		BaseModel: *kit.GetInsertBaseModel(c),
		Domain:    domain,
		Filename:  filename,
		Dst:       dst,
		ExpireAt:  time.Now().Add(time.Hour * time.Duration(hours)).Unix(),
	}
	if err := db.GetDB().Create(&data).Error; err != nil {
		return nil, err
	}
	return &proto.UploadFileResp{
		Url: config2.AppConfigIns.App.DfsHost + "/" + dst,
	}, nil
}
