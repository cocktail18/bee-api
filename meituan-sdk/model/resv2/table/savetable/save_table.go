package savetable

/**
* 保存桌位
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const save_table_url = "/resv2/table/save"

type SaveTableResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 返回的数据
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type SaveTableRequest struct {
    Data []Data `json:"data"`
}
type Data struct {
    /**
    *  桌位id 
    */
    TableId string `json:"tableId"`
    /**
    *  桌位名称 
    */
    TableName string `json:"tableName"`
    /**
    *  桌台类型 0:大厅 1:包间 
    */
    TableType int32 `json:"tableType"`
    /**
    *  最高可坐 
    */
    MaxPeople int32 `json:"maxPeople"`
    /**
    *  最低可坐 默认为1 
    */
    MinPeople int32 `json:"minPeople"`
    /**
    *  楼层桌位描述 
    */
    TableLocationDesc string `json:"tableLocationDesc"`
    /**
    *  桌台特点 0:大厅卡座 1:隔断 2:会场 3:普通包间 4:宴会大厅 5:大厅散台 6:多功能厅 7:豪华包间 8:多桌包间 
    */
    TableLocationFeature int32 `json:"tableLocationFeature"`
    /**
    *  包间特点 0:靠窗 1:会客区 2:独立卫生间 3:小型厨房 4:卡拉ok 5:可棋牌 6:海景房 7:液晶电视 
    */
    RoomFeature int32 `json:"roomFeature"`
    /**
    *  服务费 
    */
    ServiceFee int32 `json:"serviceFee"`
    /**
    *  包间费 
    */
    RoomFee int32 `json:"roomFee"`
    /**
    *  是否线上直销(满足条件的订单可以系统自动分配到这个桌位，不需要餐厅人员手动分配。即美大推单到ERP厂商后，能自动接单返回的桌位) 0: 否 1: 是 
    */
    IsOnlineSale int32 `json:"isOnlineSale"`
}



func (req *SaveTableRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SaveTableResponse, error) {
    resp, err := client.InvokeApi(save_table_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SaveTableResponse
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

func (response *SaveTableResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

