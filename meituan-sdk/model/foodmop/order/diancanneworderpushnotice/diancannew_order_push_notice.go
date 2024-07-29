package diancanneworderpushnotice

/**
* 订单确认接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const diancannew_order_push_notice_url = "/foodmop/order/diancannew/order/pushNotice"

type DiancannewOrderPushNoticeRequest struct {
    /**
    *  美团订单号 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  请求时间戳 
    */
    Timestamp int64 `json:"timestamp"`
    /**
    *  商户订单号，接单时必传 
    */
    VendorOrderId string `json:"vendorOrderId"`
    /**
    *  取餐码，仅以二维码形式展示给用户，非必传 
    */
    TakeMark string `json:"takeMark"`
    /**
    *  取餐口令 如品牌自己生成取餐口令，则接单时必传（需审核） 可以二维码/字符串形式展示，以二维码展示时，takeMark优先级高于takeCode（两者最多仅有一个会转换为二维码） 
    */
    TakeCode string `json:"takeCode"`
    /**
    *  确认订单结果： SUCCESS(0, "接单") FAIL(1, "拒单") 
    */
    OrderConfirmStatus int32 `json:"orderConfirmStatus"`
    /**
    *  （必传）用户预计取餐时间； （举例）假设用户10:00下单，如果10分钟（该时间依赖门店出单效率与积压订单数量）后可取餐，则需传递的预计取餐时间为10:10； （重点提示）用户依赖取餐时间到店，如若传错需要重新进行技术对接 
    */
    TakeTime int64 `json:"takeTime"`
    /**
    *  预约单会给预计开始制作时间，时间戳（预约单必传） 
    */
    MakeTime int64 `json:"makeTime"`
    /**
    *  处理时间 
    */
    HandleTime int64 `json:"handleTime"`
    /**
    *  发票子渠道，发票需要用到（需要确认开票方式） 
    */
    OcOrderType int32 `json:"ocOrderType"`
    /**
    *  发表主渠道，发票需要用到（需要确认开票方式） 
    */
    ServiceType int32 `json:"serviceType"`
    /**
    *  操作者Id 
    */
    OperatorId string `json:"operatorId"`
    /**
    *  操作者名称 
    */
    OperatorName string `json:"operatorName"`
    /**
    *  拒单原因，拒单时必填： DUPLICATE_ORDER(1, "重复订单") SHOP_CLOSE(2, "店铺已打烊") PROMOTION_ERROR(3, "营销信息错误") SHOP_TIME_NOT_SUPPORT(4, "该时段不支持（预约单）") SELL_OUT(5, "菜品售完") ORDER_ERROR(6, "其他原因") 
    */
    VendorConfirmReject int32 `json:"vendorConfirmReject"`
    /**
    *  拒单描述，拒单时必填 
    */
    RejectMessage string `json:"rejectMessage"`
    /**
    *  业务类型 10-MOP 
    */
    BizType int32 `json:"bizType"`
}
type DiancannewOrderPushNoticeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *DiancannewOrderPushNoticeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DiancannewOrderPushNoticeResponse, error) {
    resp, err := client.InvokeApi(diancannew_order_push_notice_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DiancannewOrderPushNoticeResponse
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

func (response *DiancannewOrderPushNoticeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

