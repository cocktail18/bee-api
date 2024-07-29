package updatespushelfstatus

/**
* 更新商品上下架状态（必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const update_spu_shelf_status_url = "/foodmop/sku/updateSpuShelfStatus"

type ResultData struct {
}
type UpdateSpuShelfStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * key：vendorSpuId value：同步结果，成功返回 “成功”，失败则返回相应的异常信息
    */
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type UpdateSpuShelfStatusRequest struct {
    /**
    *  更新范围 0：全国 1：门店 
    */
    Scope int32 `json:"scope"`
    /**
    *  门店 id 列表 spuScope == 1 时必传vendorShopIdList.get(0)是 门店商品对应门店 spuScope == 0 scope == 1 时必传 最多 30 个，所有 vendorShopId 都在美团侧存在映射，若存在无映射 vendorShopId，失败并在错误信息中返回无映射 vendorShopId 
    */
    VendorShopIdList []string `json:"vendorShopIdList"`
    /**
    *  上下架状态 0：下架 1：上架 
    */
    ShelfStatus int32 `json:"shelfStatus"`
    /**
    *  商品 spuId 列表 每次最多 50 个 
    */
    VendorSpuIdList []string `json:"vendorSpuIdList"`
    /**
    *  商品类型 0：品牌商品 1：门店商品 默认品牌纬度 
    */
    SpuScope int32 `json:"spuScope"`
}



func (req *UpdateSpuShelfStatusRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*UpdateSpuShelfStatusResponse, error) {
    resp, err := client.InvokeApi(update_spu_shelf_status_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response UpdateSpuShelfStatusResponse
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

func (response *UpdateSpuShelfStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

