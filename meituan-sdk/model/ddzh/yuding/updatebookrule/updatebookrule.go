package updatebookrule

/**
* 三方门店推送预订规则至平台
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const updatebookrule_url = "/ddzh/yuding/updatebookrule"

type UpdatebookruleResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type UpdatebookruleRequest struct {
    /**
    *  是否可退，0-可退，1-不可退 
    */
    Refundable int32 `json:"refundable"`
    /**
    *  提前可订时间，分钟数，如：90，表示90分钟 
    */
    AddOrderBefore int32 `json:"addOrderBefore"`
    /**
    *  提前可退时间，在refundable为不可退时，该字段可以不传，传了也是无效的 分钟数，如：90，表示90分钟 
    */
    RefundBefore int32 `json:"refundBefore"`
    /**
    *  通知方式，2-IVR，8-PC弹窗，16-APP通知。例如 [2,8] 一般三方仅需要传入2就可以了，或者不传 
    */
    NotifyTypeList []int32 `json:"notifyTypeList"`
    /**
    *  通知电话，必须是手机号码，座机不行 
    */
    NotifyPhone string `json:"notifyPhone"`
    /**
    *  不可预订日期, 如[1565766000000,1565767000000] 
    */
    DisableDateList []int64 `json:"disableDateList"`
    /**
    *  可预订开始时段，如：“01:00” 
    */
    ReservePeriodBegin string `json:"reservePeriodBegin"`
    /**
    *  可预订结束时段，如：“23:00” 
    */
    ReservePeriodEnd string `json:"reservePeriodEnd"`
    /**
    *  手动接单开始时段，如：“01:00” 
    */
    AcceptPeriodBegin string `json:"acceptPeriodBegin"`
    /**
    *  手动接单结束时段，如：“23:00” 
    */
    AcceptPeriodEnd string `json:"acceptPeriodEnd"`
    /**
    *  特殊价格日期 
    */
    SpecialPriceDateList []int64 `json:"specialPriceDateList"`
    /**
    *  三方扩展字段json串，如需要开票信息，则直接传{"needBill":1} 
    */
    ThirdExtraInfo string `json:"thirdExtraInfo"`
    /**
    *  是否支持非规则退款 
    */
    IrregularRefund bool `json:"irregularRefund"`
    /**
    *  最晚到店时间 
    */
    LatestArrivalTime string `json:"latestArrivalTime"`
    /**
    *  展示天数 
    */
    BookingDayNumber string `json:"bookingDayNumber"`
    /**
    *  可预订楼层 
    */
    BookableLayer string `json:"bookableLayer"`
    /**
    *  舞台以及区域底图 
    */
    StageAreaPictures string `json:"stageAreaPictures"`
    /**
    *  画布信息 
    */
    CanvasSize string `json:"canvasSize"`
    /**
    *  是否有座位图，是否开通预定小程序，0-否，1-是 
    */
    HasSeatPicture int32 `json:"hasSeatPicture"`
    /**
    *  最晚可退时间 
    */
    LatestRefundTime string `json:"latestRefundTime"`
    /**
    *  最晚延迟时间点，2023-09-08 20:00:00 
    */
    LatestPeriodRulePoint string `json:"latestPeriodRulePoint"`
    /**
    *  延迟分钟数 
    */
    DelayRuleMinute string `json:"delayRuleMinute"`
}



func (req *UpdatebookruleRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*UpdatebookruleResponse, error) {
    resp, err := client.InvokeApi(updatebookrule_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response UpdatebookruleResponse
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

func (response *UpdatebookruleResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

