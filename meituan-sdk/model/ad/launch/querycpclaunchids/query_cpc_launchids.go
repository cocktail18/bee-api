package querycpclaunchids

/**
* 获取推广id列表
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_cpc_launchids_url = "/ad/launch/queryCpcLaunchIds"

type LaunchObj struct {
    /**
    * 推广id
    */
    Launchid int32 `json:"launchid"`
    /**
    * 推广名称
    */
    LaunchName string `json:"launchName"`
    /**
    * 门店名称
    */
    ShopName string `json:"shopName"`
    /**
    * 投放状态
    */
    Status int32 `json:"status"`
    /**
    * 投放状态描述
    */
    StatusDesc string `json:"statusDesc"`
    /**
    * 门店信息
    */
    ShopId int64 `json:"shopId"`
}
type QueryCpcLaunchidsRequest struct {
    /**
    * 1有效，2过期，3暂停，4审核中，5审核驳回 6 删除 8 未投放（推广计划暂停） 9待推广，10 状态异常
    */
    Status []int32 `json:"status"`
    SonAdaccountId int32 `json:"sonAdaccountId"`
}
type QueryCpcLaunchidsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []LaunchObj `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryCpcLaunchidsRequest) DoInvoke(client mtclient.MeituanClient) (*QueryCpcLaunchidsResponse, error) {
    resp, err := client.InvokeApi(query_cpc_launchids_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response QueryCpcLaunchidsResponse
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

func (response *QueryCpcLaunchidsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

