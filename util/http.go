package util

import (
	"github.com/imroc/req/v3"
	jsoniter "github.com/json-iterator/go"
	"github.com/samber/lo"
	"time"
)

var defaultHttpClient = req.C().SetTimeout(time.Second * 3)

func GetDefaultHttpClient() *req.Client {
	return defaultHttpClient
}

// ToJsonWithoutErr 对象转 json 字符串，忽略异常
func ToJsonWithoutErr(data interface{}, defaultVal string) string {
	if lo.IsNil(data) {
		return defaultVal
	}
	jsonStr, err := ToJson(data)
	if err != nil {
		return defaultVal
	}
	return jsonStr
}

// ToJson 对象转 json 字符串
func ToJson(data interface{}) (string, error) {
	b, err := jsoniter.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
