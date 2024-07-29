package supplytablequery

/**
* 查询桌型桌位
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const supply_table_query_url = "/resv2/table/supply/queryTable"

type SupplyTableQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []VendorTableDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type VendorTableDTO struct {
    /**
    * 桌型
    */
    TableType int32 `json:"tableType"`
    /**
    * 桌型名称
    */
    TableTypeName string `json:"tableTypeName"`
    /**
    * 桌型id
    */
    TableTypeId int64 `json:"tableTypeId"`
    /**
    * 桌位信息
    */
    TableGroupList []VendorTableGroupDTO `json:"tableGroupList"`
}
type VendorTableGroupDTO struct {
    /**
    * 该桌位最多可坐的人数
    */
    MaxPeople int32 `json:"maxPeople"`
    /**
    * 该桌位最少可坐的人数
    */
    MixPeople int32 `json:"mixPeople"`
    /**
    * 桌位id
    */
    TableGroupId int64 `json:"tableGroupId"`
}
type SupplyTableQueryRequest struct {
}



func (req *SupplyTableQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SupplyTableQueryResponse, error) {
    resp, err := client.InvokeApi(supply_table_query_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SupplyTableQueryResponse
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

func (response *SupplyTableQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

