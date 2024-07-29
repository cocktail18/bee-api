package poiupdatepromoteinfo

/**
* 更改门店公告信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const poi_updatepromoteinfo_url = "/wmoper/ng/poiop/poi/updatepromoteinfo"

type PoiUpdatepromoteinfoRequest struct {
    /**
    *  门店公告信息 
    */
    PromotionInfo string `json:"promotion_info"`
}
type PoiUpdatepromoteinfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *PoiUpdatepromoteinfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*PoiUpdatepromoteinfoResponse, error) {
    resp, err := client.InvokeApi(poi_updatepromoteinfo_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response PoiUpdatepromoteinfoResponse
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

func (response *PoiUpdatepromoteinfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

