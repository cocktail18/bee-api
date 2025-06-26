package upload

import (
	"bytes"
	"fmt"
	config2 "gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/util"
	"mime/multipart"
	"os"
)

// OSS 对象存储接口
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// NewOss OSS的实例化方法
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
func NewOss() OSS {
	switch config2.AppConfigIns.Upload.OssType {
	case "local":
		return &Local{}
	case "qiniu":
		return &Qiniu{}
	case "tencent-cos":
		return &TencentCOS{}
	case "aliyun-oss":
		return &AliyunOSS{}
	case "huawei-obs":
		return HuaWeiObs
	case "aws-s3":
		return &AwsS3{}
	default:
		return &Local{}
	}
}

func BytesToFileHeader(data []byte) (*multipart.FileHeader, error) {
	filename := util.GetUUIDStr()
	// 创建一个multipart.Writer，用来创建一个multipart表单文件
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// 创建文件字段并将字节数据写入
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}

	// 将字节数据写入part
	_, err = part.Write(data)
	if err != nil {
		return nil, err
	}

	// 关闭writer以完成multipart数据的创建
	writer.Close()

	// 创建一个multipart.FileHeader
	fileHeader := &multipart.FileHeader{
		Filename: filename,
		Header:   make(map[string][]string),
		Size:     int64(len(data)),
	}

	// 为了能够使用FileHeader，您需要将文件数据写入一个临时文件
	tmpFile, err := os.CreateTemp("", "file-")
	if err != nil {
		return nil, err
	}
	defer tmpFile.Close()

	// 将数据写入临时文件
	_, err = tmpFile.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}

	// 将文件的临时路径赋值给FileHeader的 `Content-Disposition` 字段
	fileHeader.Header.Set("Content-Disposition", fmt.Sprintf("form-data; name=\"file\"; filename=\"%s\"", filename))

	// 将tmpFile对象作为文件源
	fileHeader.Size = int64(len(buf.Bytes()))
	fileHeader.Header.Set("Content-Type", "application/octet-stream")

	return fileHeader, nil
}
