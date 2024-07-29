package batchupdatelaunchtimeslot

/**
* 批量更新投放每周时间段
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_update_launch_timeslot_url = "/ad/launch/batchUpdateLaunchTimeSlot"

type BatchOptResult struct {
}
type BatchUpdateLaunchTimeslotResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * private List successObjects; private Map failObjects;
    */
    Data BatchOptResult `json:"data"`
    TraceId string `json:"traceId"`
}
type BatchUpdateLaunchTimeslotRequest struct {
    /**
    *  timeSlot：周一至周末的时间片7*24小时=168位，0表示关闭，1表示开启 
    */
    LaunchTimeSlotList string `json:"launchTimeSlotList"`
}



func (req *BatchUpdateLaunchTimeslotRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*BatchUpdateLaunchTimeslotResponse, error) {
    resp, err := client.InvokeApi(batch_update_launch_timeslot_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response BatchUpdateLaunchTimeslotResponse
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

func (response *BatchUpdateLaunchTimeslotResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

