package querylaunchpremiumlist

/**
* 查询关键词列表
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_launch_premiumlist_url = "/ad/launch/queryLaunchPremiumList"

type QueryLaunchPremiumlistRequest struct {
    /**
    *  关键词综合态 ONLINE ( 1 , "推广中" , "推广中" , 1 ) , SUBMITTED ( 2 , "待审核" , "待审核" , 2 ) , REJECTED ( 3 , "被驳回" , "被驳回" , 3 ) , LAUNCH_TIMEOUT ( 4 , "推广已结束" , "推广已结束" , 13 ) , LAUNCH_OFFLINE ( 5 , "推广已暂停" , "推广已暂停" , 7 ) , LAUNCH_SUBMITTED ( 6 , "推广待审核" , "推广待审核" , 9 ) , LAUNCH_REJECTED ( 7 , "推广被驳回" , "推广被驳回" , 4 ) , DELETED ( 8 , "已删除" , "已删除" , 15 ) , LAUNCH_EDITING ( 9 , "推广草稿中" , "推广草稿中" , 12 ) , LAUNCH_NOT_STARTED ( 10 , "待推广" , "待推广" , 5 ) , LAUNCH_EXCEPTION ( 11 , "推广状态异常" , "推广状态异常" , 10 ) , LAUNCH_INVALID ( 12 , "已下线" , "已下线" , 16 ) , 
    */
    Status []Interger `json:"status"`
    /**
    *  分页页码，默认为1 
    */
    PageNum int32 `json:"pageNum"`
    /**
    *  分页数 
    */
    PageSize int32 `json:"pageSize"`
}
type Interger struct {
}
type LaunchKeywordPremiumDTO struct {
}
type QueryLaunchPremiumlistResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * { "launchId": 36172770, "shopID": 1884389535, "shopName": "好雨知时节休闲娱乐", "launchName": "门店20231207BLwN2B", "launchPremiumId": 3604522,//关键词溢价ID "launchPremiumStatus": 1, "launchPremiumStatusDesc": "推广中", "launchStatus": 1, "launchStatusDesc": "推广中", "mtBidPrice": 11,//美团出价，单位分 "dpBidPrice": 11,//点评出价，单位分 "premiumContent": "{\"keywords\":\"养生,商圈词\",\"displayKeywords\":\"养生,商圈词\",\"itemTypes\":\"3\",\"displayPremiumContent\":\"词包｜门店热搜词、养生、商圈词\"}", "premiumName": "关键词出价100" }
    */
    Data []LaunchKeywordPremiumDTO `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryLaunchPremiumlistRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryLaunchPremiumlistResponse, error) {
    resp, err := client.InvokeApi(query_launch_premiumlist_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryLaunchPremiumlistResponse
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

func (response *QueryLaunchPremiumlistResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

