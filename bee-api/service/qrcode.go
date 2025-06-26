package service

import (
	"encoding/base64"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/service/upload"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v3"
	"github.com/skip2/go-qrcode"
	"sync"
)

type QrCodeSrv struct {
	BaseSrv
}

var qrcodeSrvOnce sync.Once
var qrcodeSrvInstance *QrCodeSrv

func GetQrCodeSrv() *QrCodeSrv {
	qrcodeSrvOnce.Do(func() {
		qrcodeSrvInstance = &QrCodeSrv{}
	})
	return qrcodeSrvInstance
}

func (s *QrCodeSrv) GenByContent(c *gin.Context, content string, width int, expireHours int64) (string, error) {
	bs, err := qrcode.Encode(content, qrcode.Medium, width)
	if err != nil {
		return "", err
	}
	// 保存文件到指定路径
	// @todo 定时删除过期文件
	f, err := upload.BytesToFileHeader(bs)
	if err != nil {
		return "", err
	}
	filepath, filename, err := upload.NewOss().UploadFile(f)
	if err != nil {
		return "", err
	}
	resp, err := GetDfsSrv().SaveUploadedFile(c, c.Request.Host, f, kit.GetDomain(c), filename, filepath, expireHours)
	if err != nil {
		return "", err
	}
	return resp.Url, nil
}

func (s *QrCodeSrv) WxaUnLimit(c *gin.Context, scene string, page string, hyaline bool, autoColor bool, hours bool) (string, error) {
	wxApp, err := GetWxAppSrv(c)
	if err != nil {
		return "", err
	}
	bs, err := wxApp.GetWxAppQrCode(c, &weapp.UnlimitedQRCode{
		Scene:     scene,
		Page:      page,
		Width:     1280,
		AutoColor: autoColor,
		LineColor: weapp.Color{},
		IsHyaline: hyaline,
	})
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bs), nil
}
