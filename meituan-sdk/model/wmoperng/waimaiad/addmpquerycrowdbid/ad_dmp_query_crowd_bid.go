package addmpquerycrowdbid

/**
* 获取三方人群包出价
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ad_dmp_query_crowd_bid_url = "/wmoper/ng/waimaiad/dmp/queryCrowdBid"

type AdDmpQueryCrowdBidResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ADOrientInfoResp `json:"data"`
    TraceId string `json:"traceId"`
}
type AdDmpQueryCrowdBidRequest struct {
    /**
    *  外卖广告的产品类型，1表示点金推广 
    */
    Channel int32 `json:"channel"`
}
type ADOrientInfoResp struct {
    Result Result `json:"result"`
    /**
    * 出价信息
    */
    OrientInfoList []ADOrientInfo `json:"orientInfoList"`
}
type ADOrientInfo struct {
    /**
    * 人群包id
    */
    CrowdPackId int64 `json:"crowdPackId"`
    /**
    * 出价
    */
    Price int32 `json:"price"`
    /**
    * 人群包标签
    */
    TopicCrowdCN string `json:"topicCrowdCN"`
    /**
    * 人群包名称
    */
    TopicCrowdName string `json:"topicCrowdName"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *AdDmpQueryCrowdBidRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*AdDmpQueryCrowdBidResponse, error) {
    resp, err := client.InvokeApi(ad_dmp_query_crowd_bid_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response AdDmpQueryCrowdBidResponse
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

func (response *AdDmpQueryCrowdBidResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

