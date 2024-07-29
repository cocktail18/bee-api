package orderfoodslowappeal

/**
* 商家提交申诉
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_foodslow_appeal_url = "/waimai/ng/order/foodSlow/appeal"

type OrderFoodslowAppealRequest struct {
    /**
    *  订单ID(订单展示号或者真实号) 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  图片地址url,多个以英文逗号间隔 如果 appeal_type为1，图片不能为空 
    */
    AppealPictures string `json:"appealPictures"`
    /**
    *  商家申诉的类型 1.提交照片申诉 2.餐已经被骑手取走 
    */
    AppealType int32 `json:"appealType"`
}
type OrderFoodslowAppealResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *OrderFoodslowAppealRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderFoodslowAppealResponse, error) {
    resp, err := client.InvokeApi(order_foodslow_appeal_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderFoodslowAppealResponse
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

func (response *OrderFoodslowAppealResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

