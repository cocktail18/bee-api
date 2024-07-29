package decorationdeleteimgeforsingle

/**
* 商家开放平台删除招牌
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const decoration_delete_imge_for_single_url = "/wmoper/ng/decorationop/deleteImgeForSingle"

type DecorationDeleteImgeForSingleResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type DecorationDeleteImgeForSingleRequest struct {
}



func (req *DecorationDeleteImgeForSingleRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DecorationDeleteImgeForSingleResponse, error) {
    resp, err := client.InvokeApi(decoration_delete_imge_for_single_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DecorationDeleteImgeForSingleResponse
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

func (response *DecorationDeleteImgeForSingleResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

