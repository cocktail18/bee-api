package mealcodesave

/**
* 取餐码信息下发
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const meal_code_save_url = "/peisong/order/mealCode/saveMealCodeByPkgId"

type MealCodeSaveRequest struct {
    /**
    * 美团配送内部订单id，最长不超过32个字符
    */
    MtPeisongId string `json:"mt_peisong_id"`
    /**
    *  取餐码业务详情信息，内容为JSON格式，示例如下： { "type": 0, "cabinetNum": "xxxxxx", "cabinetDoor": "xxxxx,xxxxx", "cabinetCode": "xxxxxx" } type：取餐类型，int类型，包含0（存餐及更新）、1（撤餐），必填且必须为0或1； cabinetNum：柜号，String类型，最长不超过100字符，非必填，同步撤餐状态时可为空，若存在多个时可以逗号隔开； cabinetDoor：门号（取餐柜的格口号），String类型，最长不超过100字符，必填，同步撤餐状态时可为空，若存在多个时可以逗号隔开； cabinetCode：取餐码，String类型，最长不超过12个字符，必填，同步撤餐状态时可为空，一个配送订单对应一个取餐码； 取餐码信息会同步到骑手侧展示并将取餐码转换为二维码展示。 
    */
    Infos string `json:"infos"`
}
type MealCodeSaveResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *MealCodeSaveRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*MealCodeSaveResponse, error) {
    resp, err := client.InvokeApi(meal_code_save_url, 19, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response MealCodeSaveResponse
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

func (response *MealCodeSaveResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

