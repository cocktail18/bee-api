package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"strings"
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

func (s DfsSrv) SaveUploadedFile(c context.Context, host string, domain string, filename, dst string, hours int64) (*proto.UploadFileResp, error) {
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
		Url: s.FillFileUrl(host, dst),
	}, nil
}

func (s DfsSrv) FillFileUrl(host, path string) string {
	if strings.HasPrefix(path, "http") {
		return path
	}
	return strings.TrimRight(config.GetDfsHost(host), "/") + "/" + strings.TrimLeft(path, "/")
}
