package supplystockupdate

/**
* 库存同步
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const supply_stock_update_url = "/resv2/stock/supply/updateStock"

type VendorStockUpdateDTO struct {
    /**
    *  库存数 
    */
    StockCount int32 `json:"stockCount"`
    /**
    *  库存类型 1-有限库存 2-无限库存 
    */
    StockType int32 `json:"stockType"`
    SkuDTO VendorSkuDTO `json:"skuDTO"`
}
type SupplyStockUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type SupplyStockUpdateRequest struct {
    VendorStockUpdateDTO VendorStockUpdateDTO `json:"vendorStockUpdateDTO"`
    Uuid string `json:"uuid"`
}
type VendorSkuDTO struct {
    /**
    *  桌型类型 1-大厅 2-包间 3-特殊桌型 
    */
    TableType int32 `json:"tableType"`
    /**
    *  桌型名称 tableType=3(特殊桌型)时，此字段必填 
    */
    TableName string `json:"tableName"`
    /**
    *  该桌位最少可坐的人数 
    */
    MinPeople int32 `json:"minPeople"`
    /**
    *  该桌位最多可坐的人数 
    */
    MaxPeople int32 `json:"maxPeople"`
    /**
    *  美团桌型id， 推荐该字段必传，或者 「tableType、tableName、 minPeople 、 maxPeople 」四个字段必传 
    */
    TableTypeId int64 `json:"tableTypeId"`
    /**
    *  美团 桌位id, 推荐该字段必传，或者 「tableType、tableName、 minPeople 、 maxPeople 」四个字段必传 
    */
    TableGroupId int64 `json:"tableGroupId"`
    /**
    *  餐段 
    */
    Period int32 `json:"period"`
    /**
    *  日期 秒级时间戳 
    */
    Day int32 `json:"day"`
}



func (req *SupplyStockUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SupplyStockUpdateResponse, error) {
    resp, err := client.InvokeApi(supply_stock_update_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SupplyStockUpdateResponse
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

func (response *SupplyStockUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

