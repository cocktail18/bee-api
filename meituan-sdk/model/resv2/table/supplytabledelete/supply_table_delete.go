package supplytabledelete

/**
* 删除桌型/桌位
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const supply_table_delete_url = "/resv2/table/supply/deleteTable"

type SupplyTableDeleteRequest struct {
    VendorTableDeleteDTO VendorTableDeleteDTO `json:"vendorTableDeleteDTO"`
    /**
    *  请求唯一标识符 
    */
    Uuid string `json:"uuid"`
}
type SupplyTableDeleteResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type VendorTableDeleteDTO struct {
    /**
    *  桌型 
    */
    TableType int32 `json:"tableType"`
    /**
    *  桌型名称 
    */
    TableTypeName string `json:"tableTypeName"`
    /**
    *  删除对象 1-删除桌型 2-删除桌位 
    */
    OperateType int32 `json:"operateType"`
    TableGroupList []VendorTableGroupDTO `json:"tableGroupList"`
}
type VendorTableGroupDTO struct {
    /**
    *  该桌位最多可坐的人数 
    */
    MaxPeople int32 `json:"maxPeople"`
    /**
    *  该桌位最少可坐的人数 
    */
    MinPeople int32 `json:"minPeople"`
}



func (req *SupplyTableDeleteRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SupplyTableDeleteResponse, error) {
    resp, err := client.InvokeApi(supply_table_delete_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SupplyTableDeleteResponse
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

func (response *SupplyTableDeleteResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

