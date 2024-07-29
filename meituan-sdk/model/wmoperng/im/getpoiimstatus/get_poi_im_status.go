package getpoiimstatus

/**
* 查询门店IM状态
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_poi_im_status_url = "/wmoper/ng/im/getPoiIMStatus"

type GetPoiImStatusData struct {
    /**
    * ERP 应用标识
    */
    AppId int64 `json:"app_id"`
    /**
    * ERP 商户门店id
    */
    EpoiId string `json:"epoiId"`
    /**
    * 0:关闭, 1:开启
    */
    ImStatus int32 `json:"im_status"`
}
type GetPoiImStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GetPoiImStatusData `json:"data"`
    TraceId string `json:"traceId"`
}
type GetPoiImStatusRequest struct {
}



func (req *GetPoiImStatusRequest) DoInvoke(client mtclient.MeituanClient) (*GetPoiImStatusResponse, error) {
    resp, err := client.InvokeApi(get_poi_im_status_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response GetPoiImStatusResponse
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

func (response *GetPoiImStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

