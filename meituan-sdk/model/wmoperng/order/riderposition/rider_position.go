package riderposition

/**
* 自配订单同步配送信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const rider_position_url = "/wmoper/ng/order/riderPosition"

type RiderPositionResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type RiderPositionRequest struct {
    /**
    *  订单号 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  第三方配送商物流单号，字段不为空时长度必须小于32位 
    */
    ThirdCarrierId string `json:"thirdCarrierId"`
    /**
    *  code 说明 0 配送单发往配送 10 配送单已确认 15 骑手已到店 20 骑手已取餐 40 骑手已送达 100 配送单已取消 
    */
    LogisticsStatus int32 `json:"logisticsStatus"`
    /**
    *  配送人名称 
    */
    CourierName string `json:"courierName"`
    /**
    *  配送员联系电话。 格式为纯数字，仅支持11位手机号、固定电话、400电话、用“#”“_”“,”拼接分机号的电话 
    */
    CourierPhone string `json:"courierPhone"`
    /**
    *  第三方配送方式 配送code 配送方 10001 顺丰 10002 达达 10003 闪送 10004 蜂鸟 10005 UU跑腿 10006 快跑者 10007 极客快送 10008 点我达 10009 同达 10010 生活半径 10011 邻趣 10012 趣送 10013 快服务 10014 菜鸟新配盟 10015 商家自建配送 10016 风先生 10017 其他 10018 美团配送 
    */
    ThirdLogisticsId int32 `json:"thirdLogisticsId"`
    /**
    *  骑手当前的纬度 
    */
    Latitude string `json:"latitude"`
    /**
    *  骑手当前的经度 
    */
    Longitude string `json:"longitude"`
    /**
    *  骑手信息采集时间戳 
    */
    BackFlowTime int32 `json:"backFlowTime"`
    /**
    *  配送员电话类型。0是真实号，1是隐私号 
    */
    CourierPhoneType int32 `json:"courierPhoneType"`
}



func (req *RiderPositionRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*RiderPositionResponse, error) {
    resp, err := client.InvokeApi(rider_position_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response RiderPositionResponse
    err = json.Unmarshal(resp, &response)
    if err != nil {
        return nil, err
    }

    if response.IsSuccess() {
        return &response, nil
    } else {
        return &response, &api_error.ApiError{Code: response.Code, Msg: response.Msg, TraceId: response.TraceId}
    }
}

func (response *RiderPositionResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

