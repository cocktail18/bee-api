package supplysavetable

/**
* 保存/更新桌型桌位
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const supply_save_table_url = "/resv2/table/supply/saveOrUpdateTable"

type VendorTableSyncDTO struct {
    /**
    *  桌型：0-大厅 1-包间 3-自定义 
    */
    TableType int32 `json:"tableType"`
    /**
    *  tableType=3(特殊桌型)时，此字段必填，用作C端桌型名称展示 
    */
    TableTypeName string `json:"tableTypeName"`
    TableGroupList []VendorTableGroupDTO `json:"tableGroupList"`
}
type SupplySaveTableRequest struct {
    VendorTableSyncDTOList []VendorTableSyncDTO `json:"vendorTableSyncDTOList"`
    /**
    *  请求唯一id 
    */
    Uuid string `json:"uuid"`
}
type SupplySaveTableResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
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



func (req *SupplySaveTableRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SupplySaveTableResponse, error) {
    resp, err := client.InvokeApi(supply_save_table_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SupplySaveTableResponse
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

func (response *SupplySaveTableResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

