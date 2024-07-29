package hourlycpcrtdatabyshop

/**
* cpc门店实时分时数据
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const hourly_cpc_rtdata_byshop_url = "/ad/report/getHourlyCpcRtDataByShop"

type HourlyCpcRtdataByshopResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []LaunchInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type HourlyCpcRtdataByshopRequest struct {
    /**
    * 子账号id
    */
    AccountId int32 `json:"accountId"`
    /**
    * 门店列表,单次最多支持50个
    */
    ShopIds []int64 `json:"shopIds"`
    /**
    * 点评门店id还是美团门店id，默认点评
    */
    ShopType int32 `json:"shopType"`
}
type LaunchInfo struct {
    /**
    * 天
    */
    Date string `json:"date"`
    /**
    * 小时
    */
    Hour int32 `json:"hour"`
    /**
    * 门店名字
    */
    ShopName string `json:"shop_name"`
    /**
    * 花费
    */
    Charge float64 `json:"charge"`
    /**
    * 曝光
    */
    Imp int32 `json:"imp"`
    /**
    * 点击
    */
    Click string `json:"click"`
}



func (req *HourlyCpcRtdataByshopRequest) DoInvoke(client mtclient.MeituanClient) (*HourlyCpcRtdataByshopResponse, error) {
    resp, err := client.InvokeApi(hourly_cpc_rtdata_byshop_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response HourlyCpcRtdataByshopResponse
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

func (response *HourlyCpcRtdataByshopResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

