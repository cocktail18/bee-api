package statusupdate

/**
* 品牌同步门店线上点业务开通状态变更
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const status_update_url = "/foodmop/shop/status/update"

type ResultData struct {
    /**
    * Key：vendorShopId（品牌门店Id） Value：OperateResulst（数据校验结果）
    */
    Data DataMap `json:"data"`
}
type OperateResulst struct {
    /**
    * 操作是否成功
    */
    Success bool `json:"success"`
    /**
    * 失败原因 在线状态不合法:vendorBizOpen未设置或者不在枚举值范围内； 门店类型不合法:没有传shopType标识，或者没有shopType场景的情况下shopType不为默认值 10； 门店请求存在重复: 请求中有重复的门店； 门店没有绑定品牌:请求的门店没有在美团绑定到对应品牌下；
    */
    FailMsg string `json:"failMsg"`
}
type VendorBizConfigRequest struct {
    /**
    *  厂商门店ID 
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    *  线上点业务开关，合同录入后，根据这个开关决定是否展示线上点入口 10: 开 20：关 
    */
    VendorBizOpen int32 `json:"vendorBizOpen"`
    /**
    *  分店类型，品牌方自定义（默认是10-无含义），美团可用于区分C端门店样式展示的区分（具体有需求同美团PM沟通） 
    */
    ShopType int32 `json:"shopType"`
}
type StatusUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type StatusUpdateRequest struct {
    Biz []VendorBizConfigRequest `json:"biz"`
}
type DataMap struct {
    OperateResulst OperateResulst `json:"OperateResulst"`
}



func (req *StatusUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*StatusUpdateResponse, error) {
    resp, err := client.InvokeApi(status_update_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response StatusUpdateResponse
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

func (response *StatusUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

