package batchcreatecpclaunchs

/**
* 新建推广
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_create_cpc_launchs_url = "/ad/launch/batchCreateCpcLaunchs"

type FailShop struct {
}
type BatchCreateCpcLaunchsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data BatchCreateCpcLaunchsData `json:"data"`
    TraceId string `json:"traceId"`
}
type BatchCreateCpcLaunchsData struct {
    /**
    * 成功的门店列表
    */
    SuccessShops []int64 `json:"successShops"`
    /**
    * 创建失败的门店map，key为原因，value为门店列表
    */
    FailShops FailShop `json:"failShops"`
}
type BatchCreateCpcLaunchsRequest struct {
    /**
    * 门店列表，一个门店对应一个推广,单次最多支持50个
    */
    ShopList []int64 `json:"shopList"`
    /**
    * 1点评，2美团
    */
    ShopType int32 `json:"shopType"`
    /**
    * 日预算
    */
    BudgetBasic int32 `json:"budgetBasic"`
    /**
    * 高峰日上浮比例
    */
    BudgetFloatRatio int32 `json:"budgetFloatRatio"`
    /**
    * 是否支持匀速投放 0 不支持 1 支持 默认支持
    */
    DeliverySpeed int32 `json:"deliverySpeed"`
    /**
    * 点评侧出价（单位分）
    */
    DpPrice int32 `json:"dpPrice"`
    /**
    * 美团侧出价(单位分)
    */
    MtPrice int32 `json:"mtPrice"`
    /**
    * 7*24小时定向选择 默认7*24
    */
    TimeList []int64 `json:"timeList"`
    /**
    *  1 默认，2是终使用门店头图，3开启系统优选 
    */
    SmartPic int32 `json:"smartPic"`
    SonAdaccountId int32 `json:"sonAdaccountId"`
    /**
    *  shopId：创建投放用到的门店 keywordPremiums：关键词定向jsonArray premiumName：溢价关键词定向名 keywords：词包列表，用于新增品类词包，地址词 words:用于新增推荐词 itemTypes：只在选择门店热搜词的情况下传固定值3 dpBidPrice：点评出价（单位分） mtBidPrice：美团出价（单位分） 
    */
    ShopKeywordPremiumListMap string `json:"shopKeywordPremiumListMap"`
    NeedEcpc int32 `json:"needEcpc"`
}



func (req *BatchCreateCpcLaunchsRequest) DoInvoke(client mtclient.MeituanClient) (*BatchCreateCpcLaunchsResponse, error) {
    resp, err := client.InvokeApi(batch_create_cpc_launchs_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response BatchCreateCpcLaunchsResponse
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

func (response *BatchCreateCpcLaunchsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

