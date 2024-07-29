package querystorehavebuild

/**
* 查询门店是否有建群资格
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_store_have_build_url = "/wmoper/ng/im/queryStoreHaveBuild"

type QueryStoreHaveBuildResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryStoreHaveBuild `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryStoreHaveBuild struct {
    /**
    * 是否有建群能力 0 - 没有 1 - 有
    */
    Status bool `json:"status"`
}
type QueryStoreHaveBuildRequest struct {
}



func (req *QueryStoreHaveBuildRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryStoreHaveBuildResponse, error) {
    resp, err := client.InvokeApi(query_store_have_build_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryStoreHaveBuildResponse
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

func (response *QueryStoreHaveBuildResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

