package mtclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"meituan/gosdk/constants"
	"meituan/gosdk/log"
	"meituan/gosdk/util"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const timestampKey = "timestamp"
const businessIdKey = "businessId"
const developerIdKey = "developerId"
const charset = "UTF-8"
const charsetKey = "charset"
const versionKey = "version"
const version = "2"
const bizKey = "biz"
const appAuthTokenKey = "appAuthToken"
const signMapKey = "sign"

type HttpClient struct {
	*http.Client
}

type DefaultMeituanClient struct {
	DeveloperId int64
	SignKey     string
	HttpClient  *http.Client
}

func NewDefaultClient(developerId int64, signKey string) MeituanClient {
	httpClient := &http.Client{
		Timeout: time.Second * 8,
	}
	return &DefaultMeituanClient{DeveloperId: developerId, SignKey: signKey, HttpClient: httpClient}
}

func NewDefaultClientWithTimeout(developerId int64, signKey string, timeoutSec int64) MeituanClient {
	httpClient := &http.Client{
		Timeout: time.Second * time.Duration(timeoutSec),
	}
	return &DefaultMeituanClient{DeveloperId: developerId, SignKey: signKey, HttpClient: httpClient}
}

// InvokeApi 方法用于自定义路径和参数调用美团api，
// apiPath用于指定API的路径（不包含域名），例如/waimai/order/queryById
// data字段用于指定业务参数，可使用map[string]string，也可使用包含业务字段的结构体
func (client *DefaultMeituanClient) InvokeApi(apiPath string, businessId int, appAuthToken string, data interface{}) ([]byte, error) {
	paramMap, err := buildRequestParam(client.SignKey, businessId, appAuthToken, client.DeveloperId, data)
	if err != nil {
		log.Warn(fmt.Sprintf("calcSign encounted error:%s", err))
	}

	encodedData := encodeParamMap(paramMap)
	log.Debug(fmt.Sprintf("encoded request data is:%s", encodedData))

	req, err := http.NewRequest("POST", constants.MeituanOpenDomain+apiPath, strings.NewReader(encodedData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Sdk-Info", "go-sdk")
	req.Header.Add("DeveloperId", paramMap[developerIdKey])

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func (client *DefaultMeituanClient) GetToken(businessId int, code string) ([]byte, error) {
	paramMap := make(map[string]string)
	paramMap[timestampKey] = strconv.FormatInt(time.Now().Unix(), 10)
	paramMap[businessIdKey] = strconv.Itoa(businessId)
	paramMap[developerIdKey] = strconv.FormatInt(client.DeveloperId, 10)
	paramMap[charsetKey] = charset
	paramMap["code"] = code
	paramMap["grantType"] = "authorization_code"
	sign, err := util.CalculateSign(paramMap, client.SignKey)
	if err != nil {
		return nil, err
	}
	paramMap[signMapKey] = sign

	encodedData := encodeParamMap(paramMap)
	log.Debug(fmt.Sprintf("encoded request data is:%s", encodedData))

	req, err := http.NewRequest("POST", constants.MeituanOpenDomain+"/oauth/token", strings.NewReader(encodedData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Sdk-Info", "go-sdk")
	req.Header.Add("DeveloperId", paramMap[developerIdKey])

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func (client *DefaultMeituanClient) RefreshToken(businessId int, refreshToken string) ([]byte, error) {
	paramMap := make(map[string]string)
	paramMap[timestampKey] = strconv.FormatInt(time.Now().Unix(), 10)
	paramMap[businessIdKey] = strconv.Itoa(businessId)
	paramMap[developerIdKey] = strconv.FormatInt(client.DeveloperId, 10)
	paramMap[charsetKey] = charset
	paramMap["refreshToken"] = refreshToken
	paramMap["grantType"] = "refresh_token"
	paramMap["scope"] = "all"
	sign, err := util.CalculateSign(paramMap, client.SignKey)
	if err != nil {
		return nil, err
	}
	paramMap[signMapKey] = sign

	encodedData := encodeParamMap(paramMap)
	log.Debug(fmt.Sprintf("encoded request data is:%s", encodedData))

	req, err := http.NewRequest("POST", constants.MeituanOpenDomain+"/oauth/refresh", strings.NewReader(encodedData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Sdk-Info", "go-sdk")
	req.Header.Add("DeveloperId", paramMap[developerIdKey])

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func buildRequestParam(signKey string, businessId int, appAuthToken string, developerId int64, data interface{}) (map[string]string, error) {
	paramMap := make(map[string]string)

	paramMap[timestampKey] = strconv.FormatInt(time.Now().Unix(), 10)
	paramMap[businessIdKey] = strconv.Itoa(businessId)
	paramMap[developerIdKey] = strconv.FormatInt(developerId, 10)
	paramMap[charsetKey] = charset
	paramMap[versionKey] = version

	if data != nil {
		paramMap[bizKey] = toJson(data)
	}

	if appAuthToken != "" {
		paramMap[appAuthTokenKey] = appAuthToken
	}

	sign, err := util.CalculateSign(paramMap, signKey)
	if err != nil {
		return nil, err
	}
	paramMap[signMapKey] = sign

	return paramMap, nil
}

func encodeParamMap(paramMap map[string]string) string {
	urlParamValues := url.Values{}
	for key, value := range paramMap {
		urlParamValues.Add(key, value)
	}

	return urlParamValues.Encode()
}

func toJson(data interface{}) string {
	bData, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(bData)
}
