package dada_sdk

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 定义状态常量
const (
	WAIT_ACCEPT         = 1    // 1
	WAIT_PICK           = 2    // 2
	ARRIVE_SHOP         = 100  // 100
	DELIVERING          = 3    // 3
	HAD_COMPLETE        = 4    // 4
	HAD_CANCEL          = 5    // 5
	RETURN_ING          = 9    // 9
	ORDER_CREATE_FAILED = 1000 // 1000
	RETURN_SURE         = 10   // 10
	RETURN_SHOP         = 11   // 11
)

// CallbackStatusDescriptions 状态描述映射
var CallbackStatusDescriptions = map[int]string{
	WAIT_ACCEPT:         "订单已创建成功-待接单",
	WAIT_PICK:           "已接单-待取货",
	ARRIVE_SHOP:         "已取货-待到店",
	DELIVERING:          "已取货-配送中",
	HAD_COMPLETE:        "已完成",
	HAD_CANCEL:          "已取消",
	RETURN_ING:          "妥投异常,返还中",
	ORDER_CREATE_FAILED: "系统异常,订单下发失败",
	RETURN_SURE:         "妥投异常,返还完成",
	RETURN_SHOP:         "返还到店",
}

// CallbackParam 表示回调参数
type CallbackParam struct {
	ClientId         string `json:"client_id"`          //	是	达达物流订单号，默认为空
	OrderId          string `json:"order_id"`           //	是	第三方订单ID，对应下单接口中的origin_id
	OrderStatus      int    `json:"order_status"`       //是	订单状态(待接单＝1,待取货＝2,骑士到店=100,配送中＝3,已完成＝4,已取消＝5, 已追加待接单=8,妥投异常之物品返回中=9, 妥投异常之物品返回完成=10, 售后取件单送达门店=6, 创建达达运单失败=1000）
	RepeatReasonType int    `json:"repeat_reason_type"` //否	重复回传状态原因(1-重新分配骑士，2-骑士转单)
	CancelReason     string `json:"cancel_reason"`      //	是	订单取消原因,其他状态下默认值为空字符串
	CancelFrom       int    `json:"cancel_from"`        //是	订单取消原因来源(1:达达配送员取消；2:商家主动取消；3:系统或客服取消；0:默认值)
	UpdateTime       int64  `json:"update_time"`        //	是	更新时间，时间戳除了创建达达运单失败=1000的精确毫秒，其他时间戳精确到秒
	Signature        string `json:"signature"`          //	是	对client_id, order_id, update_time的值进行字符串升序排列，再连接字符串，取md5值
	DmId             int    `json:"dm_id"`              //否	达达配送员id，接单以后会传
	DmName           string `json:"dm_name"`            //	否	配送员姓名，接单以后会传
	DmMobile         string `json:"dm_mobile"`          //	否	配送员手机号，接单以后会传
	FinishCode       string `json:"finish_code"`        //	否	收货码
}

// CallbackParamV2 表示回调参数
type CallbackParamV2 struct {
	ClientId         string `json:"clientId"`         //	是	达达物流订单号，默认为空
	OrderId          string `json:"orderId"`          //	是	第三方订单ID，对应下单接口中的origin_id
	OrderStatus      int    `json:"orderStatus"`      //是	订单状态(待接单＝1,待取货＝2,骑士到店=100,配送中＝3,已完成＝4,已取消＝5, 已追加待接单=8,妥投异常之物品返回中=9, 妥投异常之物品返回完成=10, 售后取件单送达门店=6, 创建达达运单失败=1000）
	RepeatReasonType int    `json:"repeatReasonType"` //否	重复回传状态原因(1-重新分配骑士，2-骑士转单)
	CancelReason     string `json:"cancelReason"`     //	是	订单取消原因,其他状态下默认值为空字符串
	CancelFrom       int    `json:"cancelFrom"`       //是	订单取消原因来源(1:达达配送员取消；2:商家主动取消；3:系统或客服取消；0:默认值)
	UpdateTime       int64  `json:"updateTime"`       //	是	更新时间，时间戳除了创建达达运单失败=1000的精确毫秒，其他时间戳精确到秒
	Signature        string `json:"signature"`        //	是	对client_id, order_id, update_time的值进行字符串升序排列，再连接字符串，取md5值
	DmId             int    `json:"dmId"`             //否	达达配送员id，接单以后会传
	DmName           string `json:"dmName"`           //	否	配送员姓名，接单以后会传
	DmMobile         string `json:"dmMobile"`         //	否	配送员手机号，接单以后会传
	FinishCode       string `json:"finishCode"`       //	否	收货码
}

func (param *CallbackParam) GetUpdateTime() time.Time {
	if param.OrderStatus == ORDER_CREATE_FAILED {
		return time.UnixMilli(param.UpdateTime)
	}
	return time.Unix(param.UpdateTime, 0)
}

// VerifySignature 验证签名是否正确
func VerifySignature(callbackParam CallbackParam) error {
	expectSign := generateSignature(callbackParam)
	actualSign := callbackParam.Signature

	if expectSign != actualSign {
		return fmt.Errorf("签名校验不通过")
	}
	return nil
}

// generateSignature 根据相关参数计算签名
func generateSignature(callbackParam CallbackParam) string {
	clientId := callbackParam.ClientId
	orderId := callbackParam.OrderId
	var updateTimeStr string
	if callbackParam.UpdateTime != 0 {
		updateTimeStr = strconv.FormatInt(callbackParam.UpdateTime, 10)
	}

	// 将签名相关字段加入列表
	list := []string{clientId, orderId, updateTimeStr}

	// 对字段进行升序排列
	sort.Strings(list)

	// 将排序后的参数拼接成一个字符串
	joinedStr := ""
	for _, s := range list {
		joinedStr += s
	}

	// 对拼接后的字符串进行MD5加密
	hash := md5.Sum([]byte(joinedStr))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}
