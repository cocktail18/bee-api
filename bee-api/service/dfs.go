package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/spf13/cast"
	"mime/multipart"
	"path/filepath"
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

func (s DfsSrv) SaveUploadedFile(c context.Context, host string, file *multipart.FileHeader, domain, filename, dst string, hours int64) (*proto.UploadFileResp, error) {
	if domain == "" {
		domain = kit.GetDomain(c)
	}
	data := model.BeeUploadFile{
		BaseModel: *kit.GetInsertBaseModel(c),
		Domain:    domain,
		Filename:  filename,
		Dst:       dst,
		Size:      file.Size,
		ExpireAt:  time.Now().Add(time.Hour * time.Duration(hours)).Unix(),
	}
	if err := db.GetDB().Create(&data).Error; err != nil {
		return nil, err
	}
	return &proto.UploadFileResp{
		Url:          s.FillFileUrl(host, dst),
		Name:         dst,
		Id:           data.Id,
		Type:         filepath.Ext(filename),
		OriginalName: file.Filename,
		Size:         cast.ToString(file.Size),
	}, nil
}

func (s DfsSrv) FillFileUrl(host, path string) string {
	if strings.HasPrefix(path, "http") {
		return path
	}
	return strings.TrimRight(config.GetDfsHost(host), "/") + "/" + strings.TrimLeft(path, "/")
}
