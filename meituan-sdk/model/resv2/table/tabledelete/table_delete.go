package tabledelete

/**
* 删除桌位
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const table_delete_url = "/resv2/table/delete"

type TableDeleteResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 返回的数据
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type TableDeleteRequest struct {
    /**
    *  桌位的tableId 
    */
    TableId string `json:"tableId"`
}



func (req *TableDeleteRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TableDeleteResponse, error) {
    resp, err := client.InvokeApi(table_delete_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TableDeleteResponse
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

func (response *TableDeleteResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

