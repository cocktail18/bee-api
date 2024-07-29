package tuangoudealqueryshopdeal

/**
* 获取团购信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tuangou_deal_queryshopdeal_url = "/ddzh/tuangou/deal/queryshopdeal"

type TuangouDealQueryshopdealResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    Result []RshopDeal `json:"result"`
}
type RshopDeal struct {
    /**
    * 套餐ID
    */
    DealId int64 `json:"dealId"`
    /**
    * 团购ID
    */
    DealGroupId int64 `json:"dealGroupId"`
    /**
    * 团购开始售卖时间
    */
    BeginDate string `json:"beginDate"`
    /**
    * 团购结束售卖时间
    */
    EndDate string `json:"endDate"`
    /**
    * 套餐名称
    */
    Title string `json:"title"`
    /**
    * 套餐价格
    */
    Price float64 `json:"price"`
    /**
    * 套餐原价
    */
    MarketPrice float64 `json:"marketPrice"`
    /**
    * 团购券开始服务时间
    */
    ReceiptBeginDate string `json:"receiptBeginDate"`
    /**
    * 团购券结束服务时间
    */
    ReceiptEndDate string `json:"receiptEndDate"`
    /**
    * 售卖状态
    */
    SaleStatus int32 `json:"saleStatus"`
    /**
    * 团购状态
    */
    DealGroupStatus int32 `json:"dealGroupStatus"`
    SaleChannelName string `json:"saleChannelName"`
    /**
    * 团购类型
    */
    DealType string `json:"dealType"`
    ExtraMap string `json:"extraMap"`
}
type TuangouDealQueryshopdealRequest struct {
    /**
    *  页码，默认值为1 
    */
    Offset int32 `json:"offset"`
    /**
    *  页大小，默认值为100，超过100，需要分页 
    */
    Limit int32 `json:"limit"`
    /**
    *  售卖平台（1：点评 2：美团 默认为1） 
    */
    Source int32 `json:"source"`
}



func (req *TuangouDealQueryshopdealRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TuangouDealQueryshopdealResponse, error) {
    resp, err := client.InvokeApi(tuangou_deal_queryshopdeal_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TuangouDealQueryshopdealResponse
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

func (response *TuangouDealQueryshopdealResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

