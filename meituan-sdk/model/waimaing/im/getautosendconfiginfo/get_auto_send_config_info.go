package getautosendconfiginfo

/**
* 查询门店“智能回复机器人”的开关状态
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_auto_send_config_info_url = "/waimai/ng/im/autosend/getAutoSendConfigInfo"

type GetAutoSendConfigInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []AutoSendConfigInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type GetAutoSendConfigInfoRequest struct {
}
type AutoSendConfigInfo struct {
    /**
    * 自动回复类型 （这些类型后续可能还会扩展，接口同步自动更新 ） 1- 售后自动回复 7 - 忙碌状态自动回复 9 - 顾客催单场景自动回复 10 - 骑手取餐后补充备注场景自动回复 11 - 顾客补充备注（美配）场景自动回复 12 - 顾客二次强调备注 13 - 顾客要求退款 14 - 顾客要求修改地址 15 - 顾客咨询优惠信息 16 - 顾客咨询发票信息 17 - 顾客闲聊 18 - 邀请评价顾客评价自动回复 19 - 顾客咨询营业时间 20 - 顾客咨询店铺位置 21 - 顾客投诉食物中毒、异物 22 - 顾客反馈餐品少送 24 - 无人机咨询 25 - 顾客咨询不同时间能否配送 26 - 顾客咨询能否赠送赠品 27 - 顾客反馈食品安全问题
    */
    Type int32 `json:"type"`
    /**
    * 开关状态 0关闭 1开启
    */
    Status int32 `json:"status"`
    /**
    * 场景名称 如 售后自动回复 同type中的备注
    */
    SceneName string `json:"sceneName"`
    /**
    * 回复内容
    */
    Content string `json:"content"`
}



func (req *GetAutoSendConfigInfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetAutoSendConfigInfoResponse, error) {
    resp, err := client.InvokeApi(get_auto_send_config_info_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetAutoSendConfigInfoResponse
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

func (response *GetAutoSendConfigInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

