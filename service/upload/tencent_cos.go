package upload

import (
	"context"
	"errors"
	"fmt"
	config2 "gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/logger"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

type TencentCOS struct{}

// UploadFile upload file to COS
func (*TencentCOS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client := NewClient()
	f, openError := file.Open()
	if openError != nil {
		logger.GetLogger().Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), config2.AppConfigIns.Upload.TencentCOS.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(err)
	}
	return config2.AppConfigIns.Upload.TencentCOS.BaseURL + "/" + config2.AppConfigIns.Upload.TencentCOS.PathPrefix + "/" + fileKey, fileKey, nil
}

// DeleteFile delete file form COS
func (*TencentCOS) DeleteFile(key string) error {
	client := NewClient()
	name := config2.AppConfigIns.Upload.TencentCOS.PathPrefix + "/" + key
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		logger.GetLogger().Error("function bucketManager.Delete() failed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}
	return nil
}

// NewClient init COS client
func NewClient() *cos.Client {
	urlStr, _ := url.Parse("https://" + config2.AppConfigIns.Upload.TencentCOS.Bucket + ".cos." + config2.AppConfigIns.Upload.TencentCOS.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config2.AppConfigIns.Upload.TencentCOS.SecretID,
			SecretKey: config2.AppConfigIns.Upload.TencentCOS.SecretKey,
		},
	})
	return client
}
