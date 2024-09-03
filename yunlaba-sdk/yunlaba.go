package yunlaba_sdk

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/spf13/cast"
	"log"
	"strings"
	"time"
)

// 文档地址： https://console-docs.apipost.cn/preview/100a3773567008af/77e6c9c9a13d21be?target_id=001
// 测试环境域名：https://pao.ylb.io具体接口名
// 线上环境域名：https://xxxx/具体接口名
// 共用测试appId：9999 测试secret：YH0QTIHfnyh2
const (
	host      = "" //@todo
	debugHost = "https://pao.ylb.io/shop/message/open"
)

type Cmd string
type LogisticsStatus int //logisticsStatus 枚举： 1：骑手收餐入箱 2：骑手配送中 3：骑手已送达 4：商家出餐 5：骑手到店取餐

var (
	CmdPushErrCallback   Cmd = "push.ErrorCallback"
	CmdDeliveryStateSync Cmd = "order.DeliveryStateSync"

	LogisticsStatusWaitingForDelivery LogisticsStatus = 1
	LogisticsStatusDelivering         LogisticsStatus = 2
	LogisticsStatusDelivered          LogisticsStatus = 3
	LogisticsStatusShopReady          LogisticsStatus = 4
	LogisticsStatusPickedUp           LogisticsStatus = 5
)

type ResponseBody struct {
	Code   int             `json:"code"`
	ErrMsg string          `json:"errMsg"`
	Data   json.RawMessage `json:"data,omitempty"`
}

func (b ResponseBody) Error() string {
	return b.ErrMsg
}

type Response struct {
	Sign      string        `json:"sign"`
	Cmd       string        `json:"cmd"`
	Source    int           `json:"source"`
	Body      *ResponseBody `json:"body"`
	Version   string        `json:"version"`
	Timestamp int           `json:"timestamp"`
}

type YunlabaSdk struct {
	AppKey    string `json:"app_key"` //对应appid
	AppSecret string `json:"app_secret"`
	Host      string `json:"host"`
	Debug     bool   `json:"debug"` //是否debug模式
	client    *req.Client
}

type QueryShopReq struct {
	ShopId string `json:"shopId"`
}

type QueryShopRes struct {
	ShopId     string  `json:"shopId"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Phone      string  `json:"phone"`
	PictureUrl string  `json:"pictureUrl"`
}

type BindShopReq struct {
	Address    string  `json:"address,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
	Name       string  `json:"name,omitempty"`
	Phone      string  `json:"phone,omitempty"`
	PictureUrl string  `json:"pictureUrl,omitempty"`
	ShopId     string  `json:"shopId,omitempty"`
	State      string  `json:"state,omitempty"`
}

type CreateOrderReq struct {
	CreatedTime  string `json:"created_time,omitempty"`
	DaySeq       int64  `json:"day_seq,omitempty"`
	DeliveryTime int    `json:"delivery_tim,omitempty"`
	Foods        []struct {
		BoxNum   int    `json:"box_num,omitempty"`
		BoxPrice int    `json:"box_price,omitempty"`
		Name     string `json:"name,omitempty"`
		Price    string `json:"price,omitempty"`
		Quantity int    `json:"quantity,omitempty"`
	} `json:"foods,omitempty"`
	NotifyUrl          string `json:"notify_url,omitempty"`
	OrderId            string `json:"order_id,omitempty"`
	OriginalPrice      string `json:"original_price,omitempty"`
	PaidPrice          string `json:"paid_price,omitempty"`
	Quantity           int    `json:"quantity,omitempty"`
	RecipientAddress   string `json:"recipient_address,omitempty"`
	RecipientLatitude  string `json:"recipient_latitude,omitempty"`
	RecipientLongitude string `json:"recipient_longitude,omitempty"`
	RecipientName      string `json:"recipient_name,omitempty"`
	RecipientPhone     string `json:"recipient_phone,omitempty"`
	Remarks            string `json:"remarks,omitempty"`
	ShopAddress        string `json:"shop_address,omitempty"`
	ShopId             string `json:"shop_id,omitempty"` // 门店唯一标识ID
	ShopName           string `json:"shop_name,omitempty"`
	ShopPhone          string `json:"shop_phone,omitempty"`
	Source             string `json:"source,omitempty"` // 订单来源
	UpdatedTime        string `json:"updated_time,omitempty"`
}

type DeliveryStateSync struct {
	OrderId         string          `json:"orderId"`
	LogisticsStatus LogisticsStatus `json:"logisticsStatus"`
	CourierName     string          `json:"courierName"`
	CourierPhone    string          `json:"courierPhone"`
	Latitude        float64         `json:"latitude"`
	Longitude       float64         `json:"longitude"`
}

type ErrorCallback struct {
	TraceId  string `json:"traceId"`
	PushType string `json:"pushType"`
	ErrorMsg string `json:"errorMsg"`
}

type CancelOrderRequest struct {
	Source       string `json:"source"`
	ShopId       string `json:"shop_id"`
	OrderId      string `json:"order_id"`
	CancelReason string `json:"cancel_reason"`
}

type Request struct {
	Sign      string `json:"sign"`
	Cmd       string `json:"cmd"`
	Source    int    `json:"source"`
	Body      string `json:"body"`
	Version   string `json:"version"`
	Timestamp int64  `json:"timestamp"`
}

type NotifyData struct {
	Sign      string          `json:"sign"`
	Cmd       string          `json:"cmd"`
	Source    int             `json:"source"`
	Body      json.RawMessage `json:"body"`
	Version   string          `json:"version"`
	Timestamp int64           `json:"timestamp"`
}

// generateParams 构建请求参数
func (sdk *YunlabaSdk) generateParams(cmd string, body any) (*Request, error) {
	bodyBs, _ := json.Marshal(body)
	_req := &Request{
		Cmd:       cmd,
		Source:    cast.ToInt(sdk.AppKey),
		Body:      string(bodyBs),
		Version:   "1.0",
		Timestamp: time.Now().Unix(),
	}

	_req.Sign = sdk.signature(_req)
	return _req, nil
}

// signature 计算签名
func (sdk *YunlabaSdk) signature(req *Request) string {
	// Append appSecret to the string
	finalSignStr := fmt.Sprintf("body=%s&cmd=%s&source=%s&timestamp=%d&version=%s%s", req.Body, req.Cmd, sdk.AppKey, req.Timestamp, req.Version, sdk.AppSecret)
	// MD5 hash
	if sdk.Debug {
		log.Println("finalSignStr", finalSignStr)
	}
	hash := md5.Sum([]byte(finalSignStr))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

func (sdk *YunlabaSdk) getHost() string {
	if sdk.Debug {
		return debugHost
	}
	return sdk.Host
}

func (sdk *YunlabaSdk) post(ctx context.Context, cmd string, traceId string, body any, res any) error {
	reqBody, err := sdk.generateParams(cmd, body)
	if err != nil {
		return err
	}
	var result = &Response{}
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	if traceId != "" {
		headers["traceId"] = traceId
	}
	rs, err := sdk.client.R().SetContext(ctx).SetHeaders(headers).SetBody(reqBody).SetSuccessResult(result).Post(fmt.Sprintf("%s", sdk.getHost()))
	if err != nil {
		return err
	}
	if rs.IsErrorState() {
		return rs.Err
	}
	if result.Body.Code != 0 {
		return result.Body
	}
	if res != nil {
		_ = json.Unmarshal(result.Body.Data, res)
	}
	return nil
}

func (sdk *YunlabaSdk) GetResponse(cmd string, body *ResponseBody) (*Request, error) {
	return sdk.generateParams(cmd, body)
}

func NewYunlabaSdk(appKey, appSecret, host string, debug bool) *YunlabaSdk {
	sdk := &YunlabaSdk{
		AppKey:    appKey,
		AppSecret: appSecret,
		Host:      host,
		Debug:     debug,
		client:    req.C(),
	}
	if sdk.Debug {
		sdk.client.EnableDebugLog()
		sdk.client.DevMode()
	}
	return sdk
}
