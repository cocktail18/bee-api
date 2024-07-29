package productdealgroupoffline

/**
* 下架团单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const product_dealgroup_offline_url = "/ddzhkh/shangpin/dealgroup/offline"

type ProductDealgroupOfflineResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Response `json:"data"`
    TraceId string `json:"traceId"`
}
type Response struct {
    /**
    * 团单id
    */
    DealGroupId int32 `json:"dealGroupId"`
}
type ProductDealgroupOfflineRequest struct {
    /**
    *  其他原因 
    */
    Remark string `json:"remark"`
    /**
    *  来源，体检 : physicalExamination 
    */
    Source string `json:"source"`
    /**
    *  团单id 
    */
    DealGroupId int32 `json:"dealGroupId"`
    /**
    *  下架原因，90104-其他, 90133-特定时间范围内不接待团购用户（如春节、店庆）, 90134-店内不再经营该项目, 90135-转让/装修/停业/转行, 90136-与其他团单重复, 90137-被销售私自上线, 90138:销量不好 
    */
    ReasonType int32 `json:"reasonType"`
}



func (req *ProductDealgroupOfflineRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ProductDealgroupOfflineResponse, error) {
    resp, err := client.InvokeApi(product_dealgroup_offline_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ProductDealgroupOfflineResponse
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

func (response *ProductDealgroupOfflineResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

