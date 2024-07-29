package querybypage

/**
* 品牌查询全量门店Id列表（非必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_by_page_url = "/foodmop/shop/queryByPage"

type QueryByPageRequest struct {
    /**
    *  查询起始位置（批量限制20） 
    */
    Offset int32 `json:"offset"`
    /**
    *  默认500 查询单页大小（单次查询最大500） 
    */
    Limit int32 `json:"limit"`
}
type QueryByPageResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data VendorShopIdListPageDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type VendorShopIdListPageDTO struct {
    /**
    * 该品牌下配置过线上点门店总数
    */
    Total int32 `json:"total"`
    /**
    * 查询起始位置
    */
    Offset int32 `json:"offset"`
    /**
    * 查询单页大小
    */
    Limit int32 `json:"limit"`
    /**
    * 厂商门店ID列表（批量最大限制500）
    */
    VendorShopIdList []string `json:"vendorShopIdList"`
}



func (req *QueryByPageRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryByPageResponse, error) {
    resp, err := client.InvokeApi(query_by_page_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryByPageResponse
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

func (response *QueryByPageResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

