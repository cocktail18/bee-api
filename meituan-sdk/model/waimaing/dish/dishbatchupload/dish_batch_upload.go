package dishbatchupload

/**
* 批量上传／更新菜品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_batch_upload_url = "/waimai/dish/batchUpload"

type DishBatchUploadResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type DishBatchUploadRequest struct {
    /**
    * ERP门店id 最大长度100
    */
    EPoiId string `json:"ePoiId"`
    /**
    *  菜品列表json字符串， 参考dish说明 
    */
    Dishes string `json:"dishes"`
    /**
    * 为true时会将没有传进来的sku删除
    */
    SkuOverwrite bool `json:"skuOverwrite"`
}



func (req *DishBatchUploadRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishBatchUploadResponse, error) {
    resp, err := client.InvokeApi(dish_batch_upload_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishBatchUploadResponse
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

func (response *DishBatchUploadResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

