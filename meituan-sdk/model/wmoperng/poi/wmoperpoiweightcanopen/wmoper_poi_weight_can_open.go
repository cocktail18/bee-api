package wmoperpoiweightcanopen

/**
* 门店是否可开启加权
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_poi_weight_can_open_url = "/wmoper/ng/poi/weight/canOpen"

type WmoperPoiWeightCanOpenResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data WmoperPoiWeightCanOpenData `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperPoiWeightCanOpenRequest struct {
}
type WmoperPoiWeightCanOpenData struct {
    /**
    * 是否可开启加权
    */
    CanWeight bool `json:"canWeight"`
    /**
    * 开启状态(1待开启、2开启中、3开启结束、4过期)
    */
    OpenWeightStatus int32 `json:"openWeightStatus"`
    /**
    * 过期日期(未开启加权，有首次营业才有值)(yyyy-mm-dd HH:mm:ss)
    */
    OverDueDate string `json:"overDueDate"`
    /**
    * 秒级unix timestamp，开启加权时间(开启过加权的才有值)
    */
    WeightStartTime int32 `json:"weightStartTime"`
    /**
    * 单位秒，加权剩余时间（开启过加权，未加权结束的才有值）
    */
    SurPlusTime int32 `json:"surPlusTime"`
    /**
    * 未通过硬性门槛
    */
    UnPassHardThresholdResults []UnPassHardThresholdResults `json:"unPassHardThresholdResults"`
    /**
    * 未通过软性门槛
    */
    UnPassSoftThresholdResults []UnPassSoftThresholdResults `json:"unPassSoftThresholdResults"`
}
type UnPassSoftThresholdResults struct {
    /**
    * 未通过的软性门槛code（-1是否配置了门店活动）
    */
    Code int32 `json:"code"`
    /**
    * 未通过的软性门槛描述文案（对于code只有是/否两种取值范围的，如-1，没有描述文案）
    */
    Remark string `json:"remark"`
}
type UnPassHardThresholdResults struct {
    /**
    * 未通过的硬性门槛code（1商品信息 2是否支持在线支付 3是否营业 4营业时长 5是否有完整资质）
    */
    Code int32 `json:"code"`
    /**
    * 未通过的硬性门槛描述文案（对于code只有是/否两种取值范围的，如2、3、5，没有描述文案）
    */
    Remark string `json:"remark"`
}



func (req *WmoperPoiWeightCanOpenRequest) DoInvoke(client mtclient.MeituanClient) (*WmoperPoiWeightCanOpenResponse, error) {
    resp, err := client.InvokeApi(wmoper_poi_weight_can_open_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response WmoperPoiWeightCanOpenResponse
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

func (response *WmoperPoiWeightCanOpenResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

