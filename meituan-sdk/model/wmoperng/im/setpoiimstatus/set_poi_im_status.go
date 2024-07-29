package setpoiimstatus

/**
* 设置门店IM状态
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const set_poi_im_status_url = "/wmoper/ng/im/setPoiIMStatus"

type SetPoiImStatusRequest struct {
    /**
    *  IM状态 0 关闭 1 开启 
    */
    ImStatus int32 `json:"imStatus"`
}
type SetPoiImStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *SetPoiImStatusRequest) DoInvoke(client mtclient.MeituanClient) (*SetPoiImStatusResponse, error) {
    resp, err := client.InvokeApi(set_poi_im_status_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response SetPoiImStatusResponse
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

func (response *SetPoiImStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

