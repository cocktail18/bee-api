package tuangoureceiptreverseconsume

/**
* 撤销验券
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tuangou_receipt_reverseconsume_url = "/ddzh/tuangou/receipt/reverseconsume"

type ReceiptConsumeResult struct {
    /**
    * 订单号
    */
    OrderId string `json:"orderId"`
    /**
    * 券码
    */
    ReceiptCode string `json:"receiptCode"`
    /**
    * 套餐ID
    */
    DealId int64 `json:"dealId"`
    /**
    * 团购ID
    */
    DealgroupId int64 `json:"dealgroupId"`
    /**
    * 商品名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 商品售卖价格
    */
    DealPrice float64 `json:"dealPrice"`
    /**
    * 商品市场价
    */
    DealMarketPrice float64 `json:"dealMarketPrice"`
    /**
    * 用户手机号
    */
    Mobile string `json:"mobile"`
}
type TuangouReceiptReverseconsumeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type TuangouReceiptReverseconsumeRequest struct {
    /**
    *  商家在自研系统或第三方服务商系统内登录的帐号，仅用于记录验券者的信息，该字段不参与任何验券校验逻辑 
    */
    AppShopAccount string `json:"appShopAccount"`
    /**
    *  商家在自研系统或第三方服务商系统内登陆的用户名，仅用于记录验券者的信息，该字段不参与任何验券校验逻辑 
    */
    AppShopAccountName string `json:"appShopAccountName"`
    /**
    *  团购券码，注： 撤销核销只能撤销当前核销且核销未超过10分钟的券 
    */
    ReceiptCode string `json:"receiptCode"`
    /**
    *  套餐ID，注意：对应dealId，非dealgroupId 
    */
    DealId string `json:"dealId"`
}
type Data struct {
    Result []ReceiptConsumeResult `json:"result"`
}



func (req *TuangouReceiptReverseconsumeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TuangouReceiptReverseconsumeResponse, error) {
    resp, err := client.InvokeApi(tuangou_receipt_reverseconsume_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TuangouReceiptReverseconsumeResponse
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

func (response *TuangouReceiptReverseconsumeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

