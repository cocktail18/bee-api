package deleteimgeforsingle

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

const delete_imge_for_single_url = "/waimai/ng/decoration/deleteImgeForSingle"

type DeleteImgeForSingleResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type DeleteImgeForSingleRequest struct {
}



func (req *DeleteImgeForSingleRequest) DoInvoke(client mtclient.MeituanClient) (*DeleteImgeForSingleResponse, error) {
    resp, err := client.InvokeApi(delete_imge_for_single_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response DeleteImgeForSingleResponse
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

func (response *DeleteImgeForSingleResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

