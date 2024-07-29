package updatepremiumshelfstatus

/**
* 更新配料属性上下架状态（必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const update_premium_shelf_status_url = "/foodmop/sku/shelf/premium/update"

type UpdatePremiumShelfStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * key：premiumCode value：同步结果，成功返回 “成功”，失败则返回相应的异常信息 类型为Map 
    */
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type ResultData struct {
}
type UpdatePremiumShelfStatusRequest struct {
    /**
    *  更新范围 0：全国 1：门店 
    */
    Scope int32 `json:"scope"`
    /**
    *  门店 id 列表 premiumScope == 1 时必传vendorShopIdList.get(0)是 门店配料对应门店 premiumScope == 0 scope == 1 时必传 最多 15 个，所有 vendorShopId 都在美团侧存在映射，若存在无映射 vendorShopId，失败并在错误信息中返回无映射 vendorShopId 
    */
    VendorShopIdList []string `json:"vendorShopIdList"`
    /**
    *  配料 code 列表 最多 50 个 
    */
    PremiumCodeList []string `json:"premiumCodeList"`
    /**
    *  上下架状态 0：下架 1：上架 
    */
    ShelfStatus int32 `json:"shelfStatus"`
    /**
    *  0：品牌配料 1：门店配料 默认品牌配料 
    */
    PremiumScope int32 `json:"premiumScope"`
}



func (req *UpdatePremiumShelfStatusRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*UpdatePremiumShelfStatusResponse, error) {
    resp, err := client.InvokeApi(update_premium_shelf_status_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response UpdatePremiumShelfStatusResponse
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

func (response *UpdatePremiumShelfStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

