package h5url

/**
* 获取骑手位置H5
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const h5url_url = "/peisong/order/rider/location/h5url"

type H5urlResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data H5urlData `json:"data"`
    TraceId string `json:"traceId"`
}
type H5urlData struct {
    /**
    * 骑手位置 H5 URL。跑腿帮送服务包返回的H5链接中，订单完成8小时后则不再展示骑手位置信息，其他服务包返回的H5链接中，订单完成24小时后则不再展示骑手位置信息；美配会根据访问流量进行时段控制。
    */
    RiderLocationUrl string `json:"riderLocationUrl"`
}
type H5urlRequest struct {
    /**
    *  美团配送内部订单id，最长不超过32个字符 
    */
    MtPeisongId string `json:"mt_peisong_id"`
}



func (req *H5urlRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*H5urlResponse, error) {
    resp, err := client.InvokeApi(h5url_url, 19, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response H5urlResponse
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

func (response *H5urlResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

