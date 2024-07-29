package actdiscountlist

/**
* 批量查询折扣商品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const act_discount_list_url = "/waimai/ng/act/discount/list"

type ActDiscountListRequest struct {
    /**
    *  分页偏移 
    */
    Offset int32 `json:"offset"`
    /**
    *  分页查询，每页查询的数量不大于200 
    */
    Limit int32 `json:"limit"`
}
type ActDiscountListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []DiscountInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type DiscountInfo struct {
    /**
    * 折扣价格，单位元
    */
    ActPrice float64 `json:"act_price"`
    /**
    * APP方商品id
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    * 当日活动库存，-1表示无限库存
    */
    DayLimit int32 `json:"day_limit"`
    /**
    * 活动开始时间，单位秒
    */
    StartTime int32 `json:"start_time"`
    /**
    * 活动结束时间，单位秒
    */
    EndTime int32 `json:"end_time"`
    /**
    * 每单限购
    */
    OrderLimit int32 `json:"order_limit"`
    /**
    * 生效时间段，多个用英文逗号分隔，最多支持3个时间段
    */
    Period string `json:"period"`
    /**
    * 参与活动的用户类型，0为不限制，1为新用户，2为老用户
    */
    UserType int32 `json:"user_type"`
    /**
    * 生效活动周期，1-7表示周一至周日，多个星期用英文逗号分隔
    */
    WeeksTime string `json:"weeks_time"`
    /**
    * 商品名称
    */
    Name string `json:"name"`
    /**
    * 商品原价，单位元
    */
    OriginPrice float64 `json:"origin_price"`
    /**
    * 当日剩余可购买数
    */
    Stock int32 `json:"stock"`
    /**
    * 活动状态，2表示待生效，1表示生效，0表示过期
    */
    Status int32 `json:"status"`
}



func (req *ActDiscountListRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ActDiscountListResponse, error) {
    resp, err := client.InvokeApi(act_discount_list_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ActDiscountListResponse
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

func (response *ActDiscountListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

