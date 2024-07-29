package queryimsubscribestatusbyepoi

/**
* 查询门店外卖非接单M消息订阅状态
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_im_subscribe_status_by_epoi_url = "/wmoper/ng/im/queryImSubscribeStatusByEpoi"

type Data struct {
    /**
    * 查询结果编码，0成功，其他失败
    */
    Code int32 `json:"code"`
    /**
    * 查询结果描述
    */
    Message string `json:"message"`
    /**
    * 门店外卖IM消息订阅状态，0未订阅，1订阅中
    */
    Status int32 `json:"status"`
}
type QueryImSubscribeStatusByEpoiResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryImSubscribeStatusByEpoiRequest struct {
}



func (req *QueryImSubscribeStatusByEpoiRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryImSubscribeStatusByEpoiResponse, error) {
    resp, err := client.InvokeApi(query_im_subscribe_status_by_epoi_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryImSubscribeStatusByEpoiResponse
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

func (response *QueryImSubscribeStatusByEpoiResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

