package orderquerybydayseq

/**
* 根据流水号查询订单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_query_by_day_seq_url = "/waimai/order/queryByDaySeq"

type PackageBagCartDetail struct {
    /**
    * 口袋id
    */
    CartId int32 `json:"cart_id"`
    /**
    * 原打包费
    */
    OriginalPackageBagFee float64 `json:"original_package_bag_fee"`
}
type OrderQueryByDaySeqRequest struct {
    /**
    *  门店下的当天订单流水号 流水号从1开始递增，标识当天第几个订单 
    */
    DaySeq int32 `json:"daySeq"`
    /**
    *  日期【yyyyMMdd】 
    */
    Date string `json:"date"`
}
type PackageBagFeeInfo struct {
    /**
    * 打包费模式, 0按照商品餐盒费收取，1按照订单总价收取，2按照订单口袋收取
    */
    PackageBagMode int32 `json:"package_bag_mode"`
    /**
    * 打包费金额
    */
    PackageBagFee float64 `json:"package_bag_fee"`
    /**
    * 口袋打包费明细
    */
    PackageBagCartDetail []PackageBagCartDetail `json:"package_bag_cart_detail"`
}
type OrderQueryByDaySeqResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data OrderQueryByDaySeqData `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderQueryByDaySeqData struct {
    BackupRecipientPhone string `json:"backupRecipientPhone"`
    CTime int64 `json:"cTime"`
    Caution string `json:"caution"`
    CityId int64 `json:"cityId"`
    DaySeq string `json:"daySeq"`
    DeliveryTime int64 `json:"deliveryTime"`
    Detail string `json:"detail"`
    DinnersNumber int64 `json:"dinnersNumber"`
    EPoiId string `json:"ePoiId"`
    Extras string `json:"extras"`
    HasInvoiced int64 `json:"hasInvoiced"`
    IncmpCode int64 `json:"incmpCode"`
    /**
    * 有降级的数据模块的集合
    */
    IncmpModules []int32 `json:"incmpModules"`
    InvMakeType int64 `json:"invMakeType"`
    InvoiceTitle string `json:"invoiceTitle"`
    IsFavorites bool `json:"isFavorites"`
    IsPoiFirstOrder bool `json:"isPoiFirstOrder"`
    IsPre int64 `json:"isPre"`
    IsThirdShipping int64 `json:"isThirdShipping"`
    Latitude float64 `json:"latitude"`
    LogisticsCancelTime int64 `json:"logisticsCancelTime"`
    LogisticsCode string `json:"logisticsCode"`
    LogisticsCompletedTime int64 `json:"logisticsCompletedTime"`
    LogisticsConfirmTime int64 `json:"logisticsConfirmTime"`
    LogisticsDispatcherMobile string `json:"logisticsDispatcherMobile"`
    LogisticsDispatcherName string `json:"logisticsDispatcherName"`
    LogisticsFetchTime int64 `json:"logisticsFetchTime"`
    LogisticsId int64 `json:"logisticsId"`
    LogisticsName string `json:"logisticsName"`
    LogisticsSendTime int64 `json:"logisticsSendTime"`
    LogisticsStatus int64 `json:"logisticsStatus"`
    Longitude float64 `json:"longitude"`
    OrderCompletedTime int64 `json:"orderCompletedTime"`
    OrderConfirmTime int64 `json:"orderConfirmTime"`
    OrderEntranceType int64 `json:"orderEntranceType"`
    OrderId int64 `json:"orderId"`
    OrderIdView int64 `json:"orderIdView"`
    OrderSendTime int64 `json:"orderSendTime"`
    OrderTagList string `json:"orderTagList"`
    OriginalPrice float64 `json:"originalPrice"`
    PayType int64 `json:"payType"`
    PickType int64 `json:"pickType"`
    PoiAddress string `json:"poiAddress"`
    PoiId int64 `json:"poiId"`
    PoiName string `json:"poiName"`
    PoiPhone string `json:"poiPhone"`
    PoiReceiveDetail string `json:"poiReceiveDetail"`
    RecipientAddress string `json:"recipientAddress"`
    RecipientAddressDesensitization string `json:"recipientAddressDesensitization"`
    RecipientName string `json:"recipientName"`
    RecipientPhone string `json:"recipientPhone"`
    ShipperPhone string `json:"shipperPhone"`
    ShippingFee float64 `json:"shippingFee"`
    Status int64 `json:"status"`
    TaxpayerId string `json:"taxpayerId"`
    Total float64 `json:"total"`
    UTime int64 `json:"uTime"`
    EstimateArrivalTime int64 `json:"estimateArrivalTime"`
    OrderCancelTime int64 `json:"orderCancelTime"`
    PhfOrderExtraData string `json:"phfOrderExtraData"`
    EtaModify string `json:"etaModify"`
    AddressOfReceivingAdministrative string `json:"addressOfReceivingAdministrative"`
    /**
    * 打包费模块
    */
    PackageBagFeeInfo PackageBagFeeInfo `json:"packageBagFeeInfo"`
    /**
    * 出餐考核时间
    */
    PoiMealAssessmentTime int64 `json:"poiMealAssessmentTime"`
    /**
    * 是否是商企通订单 1是商企通订单 2非商企通 0降级
    */
    SqtOrder int32 `json:"sqtOrder"`
    /**
    * 是否需要开发票 1需要开发票 2无需开发票 0降级
    */
    NeedSqtInvoice int32 `json:"needSqtInvoice"`
    /**
    * 赠品信息
    */
    GiftInfo string `json:"giftInfo"`
    /**
    * 商家应收款（分）
    */
    WmPoiReceiveCent int64 `json:"wmPoiReceiveCent"`
}



func (req *OrderQueryByDaySeqRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderQueryByDaySeqResponse, error) {
    resp, err := client.InvokeApi(order_query_by_day_seq_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderQueryByDaySeqResponse
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

func (response *OrderQueryByDaySeqResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

