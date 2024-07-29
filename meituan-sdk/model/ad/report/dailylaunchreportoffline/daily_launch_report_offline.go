package dailylaunchreportoffline

/**
* 推广分日报告
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const daily_launch_report_offline_url = "/ad/report/getDailyDataByLaunchOffline"

type DailyLaunchReportOfflineRequest struct {
    /**
    * 推广id列表,单次最多支持50个
    */
    LaunchIds []int64 `json:"launchIds"`
    /**
    * 平台（0全平台，1点评，2美团）
    */
    Platform int32 `json:"platform"`
    /**
    * 查询开始时间
    */
    BeginTime string `json:"beginTime"`
    /**
    * 查询截止时间
    */
    EndTime string `json:"endTime"`
}
type DailyLaunchReportOfflineResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 对象列表
    */
    Data []LaunchInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type LaunchInfo struct {
    /**
    * 天数
    */
    Date string `json:"date"`
    /**
    * 推广id
    */
    LaunchId int32 `json:"launch_id"`
    /**
    * 推广名称
    */
    LaunchName string `json:"launch_name"`
    /**
    * 门店名称
    */
    ShopName string `json:"shop_name"`
    /**
    * 花费
    */
    Charge float64 `json:"charge"`
    /**
    * 曝光
    */
    Imp int32 `json:"imp"`
    /**
    * 点击
    */
    Click int32 `json:"click"`
    /**
    * 商户浏览量
    */
    Pv int32 `json:"pv"`
    /**
    * 感兴趣
    */
    InterestAction int32 `json:"interest_action"`
    /**
    * 图片点击
    */
    ViewPhoto int32 `json:"view_photo"`
    /**
    * 评论点击
    */
    ViewReview int32 `json:"view_review"`
    /**
    * 优惠促销点击
    */
    ViewYh int32 `json:"view_yh"`
    /**
    * 团购点击
    */
    ViewTg int32 `json:"view_tg"`
    /**
    * 地址点击
    */
    ViewAddress int32 `json:"view_address"`
    /**
    * 预付商品点击
    */
    ViewYf int32 `json:"view_yf"`
    /**
    * 分享
    */
    Share int32 `json:"share"`
    /**
    * 收藏
    */
    AddFavor int32 `json:"add_favor"`
    /**
    * 预约及意向
    */
    YyyxAction int32 `json:"yyyx_action"`
    /**
    * 电话点击
    */
    Call int32 `json:"call"`
    /**
    * 在线咨询点击
    */
    Chat int32 `json:"chat"`
    /**
    * 预约量
    */
    YyOrder int32 `json:"yy_order"`
    /**
    * 优惠促销领取
    */
    OrderPromo int32 `json:"order_promo"`
    /**
    * 团购订单量
    */
    OrderTg int32 `json:"order_tg"`
    /**
    * 智能支付订单量
    */
    OfflineCnt int32 `json:"offline_cnt"`
    /**
    * 预付商品订单量
    */
    YfCnt int32 `json:"yf_cnt"`
    /**
    * 闪惠买单量
    */
    OrderSh int32 `json:"order_sh"`
    /**
    * 订单量
    */
    PayCnt int32 `json:"pay_cnt"`
    /**
    * 优惠预订订单量
    */
    ReserveCnt int32 `json:"reserve_cnt"`
    /**
    * 技师点击
    */
    ViewStylist int32 `json:"view_stylist"`
    /**
    * 商家会员卡点击量
    */
    ClickCard int32 `json:"click_card"`
    /**
    * 签到
    */
    AddCheckin int32 `json:"add_checkin"`
    /**
    * 商家会员卡购买量
    */
    CardCnt int32 `json:"card_cnt"`
    /**
    * 店铺信息点击
    */
    ViewInfo int32 `json:"view_info"`
    /**
    * 款式点击
    */
    ViewStyle int32 `json:"view_style"`
    /**
    * 意向客流
    */
    YxAction int32 `json:"yx_action"`
    /**
    * 款式一口价订单量
    */
    OrderKuanshi int32 `json:"order_kuanshi"`
    /**
    * 扫码支付订单
    */
    ScanpayCnt int32 `json:"scanpay_cnt"`
}



func (req *DailyLaunchReportOfflineRequest) DoInvoke(client mtclient.MeituanClient) (*DailyLaunchReportOfflineResponse, error) {
    resp, err := client.InvokeApi(daily_launch_report_offline_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response DailyLaunchReportOfflineResponse
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

func (response *DailyLaunchReportOfflineResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

