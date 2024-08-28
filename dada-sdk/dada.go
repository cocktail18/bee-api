package dada_sdk

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/spf13/cast"
	"sort"
	"strings"
	"time"
)

// 文档地址： http://newopen.imdada.cn/#/development/file/code
// 测试环境域名：https://newopen.qa.imdada.cn/具体接口名
// 线上环境域名：https://newopen.imdada.cn/具体接口名
const (
	host      = "https://newopen.imdada.cn"
	debugHost = "https://newopen.qa.imdada.cn"
)

type BussError struct {
	Msg  string
	Code int
	Data any
}

func (b BussError) Error() string {
	return b.Msg
}

type DadaSdk struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	SourceId  string `json:"source_id"`
	Debug     bool   `json:"debug"` //是否debug模式
	client    *req.Client
}

// generateParams 构建请求参数
func (sdk *DadaSdk) generateParams(body interface{}) (map[string]interface{}, error) {
	appSecret := sdk.AppSecret

	data := map[string]interface{}{
		"source_id": sdk.SourceId,
		"app_key":   sdk.AppKey,
		"timestamp": time.Now().Unix(),
		"format":    "json",
		"v":         "1.0",
		"body":      body,
	}

	// Convert the body to JSON string
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	data["body"] = string(bodyJson)

	data["signature"] = sdk.signature(data, appSecret)

	return data, nil
}

// signature 计算签名
func (sdk *DadaSdk) signature(data map[string]interface{}, appSecret string) string {
	// Sort the keys
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build the sign string
	signStr := ""
	for _, k := range keys {
		signStr += fmt.Sprintf("%s%s", k, cast.ToString(data[k]))
	}
	// Append appSecret to the string
	finalSignStr := appSecret + signStr + appSecret
	// MD5 hash
	hash := md5.Sum([]byte(finalSignStr))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

func (sdk *DadaSdk) getHost() string {
	if sdk.Debug {
		return debugHost
	}
	return host
}

func (sdk *DadaSdk) post(ctx context.Context, api string, body interface{}, res any) error {
	reqBody, err := sdk.generateParams(body)
	if err != nil {
		return err
	}
	var result = &Response{}
	rs, err := sdk.client.R().SetContext(ctx).SetHeaders(map[string]string{
		"Content-Type": "application/json",
	}).SetBodyJsonMarshal(reqBody).SetSuccessResult(result).Post(fmt.Sprintf("%s%s", sdk.getHost(), api))
	if err != nil {
		return err
	}
	if rs.IsErrorState() {
		return rs.Err
	}
	if result.Status != "success" {
		return &BussError{
			Msg:  result.Msg,
			Code: result.Code,
			Data: result,
		}
	}
	if res == nil {
		return nil
	}
	return json.Unmarshal(result.Result, res)
}

func NewDadaSdk(appKey, appSecret, sourceId string, debug bool) *DadaSdk {
	sdk := &DadaSdk{
		AppKey:    appKey,
		AppSecret: appSecret,
		SourceId:  sourceId,
		Debug:     debug,
		client:    req.C(),
	}
	if sdk.Debug {
		sdk.client.EnableDebugLog()
		sdk.client.DevMode()
	}
	return sdk
}
