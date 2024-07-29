package gettakefoodcodelist

/**
* 获取取餐码标签
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_take_food_code_list_url = "/waimai/ng/order/getTakeFoodCodeList"

type GetTakeFoodCodeListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GetTakeFoodCodeListData `json:"data"`
    TraceId string `json:"traceId"`
}
type GetTakeFoodCodeListRequest struct {
    /**
    *  订单id 
    */
    OrderId int64 `json:"orderId"`
}
type GetTakeFoodCodeListData struct {
    /**
    * 取餐码标签列表
    */
    TakeFoodCodeList []TakeFoodCodeList `json:"takeFoodCodeList"`
}
type TakeFoodCodeList struct {
    /**
    * 取餐码
    */
    TakeFoodCode string `json:"takeFoodCode"`
    /**
    * 识别码
    */
    IdentCode string `json:"identCode"`
    /**
    * 收餐人姓名
    */
    RecipientName string `json:"recipientName"`
    /**
    * 收餐电话
    */
    RecipientPhone string `json:"recipientPhone"`
    /**
    * 菜品名称
    */
    FoodName string `json:"foodName"`
    /**
    * 商家名称
    */
    PoiName string `json:"poiName"`
    /**
    * 收餐地址
    */
    RecipientAddress string `json:"recipientAddress"`
    /**
    * 送达时间
    */
    EstimatedDeliveryTime string `json:"estimatedDeliveryTime"`
}



func (req *GetTakeFoodCodeListRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetTakeFoodCodeListResponse, error) {
    resp, err := client.InvokeApi(get_take_food_code_list_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetTakeFoodCodeListResponse
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

func (response *GetTakeFoodCodeListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

