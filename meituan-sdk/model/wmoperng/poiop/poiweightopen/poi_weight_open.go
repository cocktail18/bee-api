package poiweightopen

/**
* 门店开启加权
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const poi_weight_open_url = "/wmoper/ng/poiop/poi/weight/open"

type PoiWeightOpenRequest struct {
}
type PoiWeightOpenResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data PoiWeightOpenData `json:"data"`
    TraceId string `json:"traceId"`
}
type PoiWeightOpenData struct {
    /**
    * 开启是否成功
    */
    Status bool `json:"status"`
    /**
    * 开启失败文案
    */
    Remark string `json:"remark"`
}



func (req *PoiWeightOpenRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*PoiWeightOpenResponse, error) {
    resp, err := client.InvokeApi(poi_weight_open_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response PoiWeightOpenResponse
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

func (response *PoiWeightOpenResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

