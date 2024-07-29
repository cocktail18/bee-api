package wmoperngqueryorderdetail

/**
* 查询订单详情(展示费率相关字段)
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoperng_query_order_detail_url = "/wmoper/ng/order/queryDetail"

type PackageBagCartDetail struct {
    /**
    * 口袋Id
    */
    CartId int32 `json:"cart_id"`
    /**
    * 原打包费
    */
    OriginalPackageBagFee float64 `json:"original_package_bag_fee"`
    /**
    * 活动打包费
    */
    ActivityPackageBagFee float64 `json:"activity_package_bag_fee"`
}
type EtaModifyVO struct {
    /**
    * 发起申请时间
    */
    UetaCtime int32 `json:"uetaCtime"`
    /**
    * 用户申请修改eta时间
    */
    UetaTime int32 `json:"uetaTime"`
    /**
    * 截止时间
    */
    UetaExpire int32 `json:"uetaExpire"`
    /**
    * 审核状态ueta_status
    */
    UetaStatus int32 `json:"uetaStatus"`
}
type WmoperngQueryOrderDetailResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data WmoperngQueryOrderDetailData `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperngQueryOrderDetailRequest struct {
    /**
    *  订单id 
    */
    OrderId int64 `json:"orderId"`
}
type PackageBagFeeInfo struct {
    /**
    * 打包费模式
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
type Activity struct {
    /**
    * 优惠金额中美团承担的部分
    */
    MtCharge float64 `json:"mtCharge"`
    /**
    * 说明
    */
    Remark string `json:"remark"`
    /**
    * 优惠金额中商家承担的部分
    */
    PoiCharge float64 `json:"poiCharge"`
    /**
    * 活动ID
    */
    ActDetailId int64 `json:"actDetailId"`
    /**
    * 活动类型
    */
    Type int32 `json:"type"`
    /**
    * 该活动用户实际享受优惠金额
    */
    ReduceFee float64 `json:"reduceFee"`
    /**
    * 顾客预先支付金额，顾客已经提前支付的神抢手券
    */
    ConsumerPrepay float32 `json:"consumer_prepay"`
}
type Json struct {
    /**
    * 商品分类，是否为套餐，1套餐商品
    */
    AttrType int32 `json:"attrType"`
    /**
    * 分类名称
    */
    Cate string `json:"cate"`
    /**
    * 套餐明细
    */
    PackageDeatil string `json:"packageDeatil"`
}
type WmoperngQueryOrderDetailData struct {
    /**
    * 订单金额
    */
    Total float64 `json:"total"`
    /**
    * 菜品信息 List类型
    */
    Detail []Detail `json:"detail"`
    /**
    * 备注。去掉隐私号，去掉脱敏的真实号，去掉预订人手机号。
    */
    Caution string `json:"caution"`
    /**
    * 订单取消时间
    */
    OrderCancelTime int64 `json:"orderCancelTime"`
    /**
    * 商户收到时间
    */
    OrderReceiveTime int64 `json:"orderReceiveTime"`
    /**
    * 商家对账信息，此字段为JSON字符串
    */
    PoiReceiveDetail string `json:"poiReceiveDetail"`
    /**
    * 配送单取消时间
    */
    LogisticsCancelTime int64 `json:"logisticsCancelTime"`
    /**
    * 配送单确认时间
    */
    LogisticsConfirmTime int64 `json:"logisticsConfirmTime"`
    /**
    * 收货人电话。字段默认内容“***********”。
    */
    RecipientPhone string `json:"recipientPhone"`
    /**
    * 配送方式
    */
    LogisticsCode string `json:"logisticsCode"`
    DetailSize int64 `json:"detailSize"`
    /**
    * 实际送餐地址经度。字段默认内容“0.0”。
    */
    Longitude float64 `json:"longitude"`
    ActivitySize int64 `json:"activitySize"`
    /**
    * 商户确认时间
    */
    OrderConfirmTime int64 `json:"orderConfirmTime"`
    /**
    * 活动信息 List类型
    */
    Activity []Activity `json:"activity"`
    /**
    * 订单展示Id
    */
    OrderId int64 `json:"orderId"`
    /**
    * 即时单：字段值为配送eta系统提供的预计送达时间 到店自取单：字段值为用户到店取餐时间 预订单：字段值为用户下单时填写的预计送达时间
    */
    EstimateArrivalTime int64 `json:"estimateArrivalTime"`
    /**
    * 订单状态 2 新订单，4 已接单，8 订单完成，9 订单取消
    */
    Status int32 `json:"status"`
    /**
    * 收货地址
    */
    RecipientAddress string `json:"recipientAddress"`
    /**
    * 流水单号
    */
    DaySeq int32 `json:"daySeq"`
    /**
    * 用户下单时间
    */
    OrderSendTime int64 `json:"orderSendTime"`
    /**
    * 配送完成时间
    */
    LogisticsCompletedTime int64 `json:"logisticsCompletedTime"`
    /**
    * 配送单下单时间
    */
    LogisticsSendTime int64 `json:"logisticsSendTime"`
    /**
    * 门店配送费
    */
    ShippingFee float64 `json:"shippingFee"`
    /**
    * 订单完成时间
    */
    OrderCompletedTime int64 `json:"orderCompletedTime"`
    /**
    * 配送员电话
    */
    ShipperPhone string `json:"shipperPhone"`
    /**
    * 用户预计送达时间
    */
    DeliveryTime int64 `json:"deliveryTime"`
    /**
    * 商品原价
    */
    OriginalPrice float64 `json:"originalPrice"`
    /**
    * 收货人名字
    */
    RecipientName string `json:"recipientName"`
    /**
    * 骑手取单时间
    */
    LogisticsFetchTime int64 `json:"logisticsFetchTime"`
    /**
    * 支付类型，1表货到付款，2表在线支付
    */
    PayType int32 `json:"payType"`
    /**
    * 透传给第三方的偏移后的user_id
    */
    UserIdView string `json:"userIdView"`
    /**
    * 实际送餐地址纬度。字段默认内容“0.0”。
    */
    Latitude float64 `json:"latitude"`
    /**
    * 外卖门店ID
    */
    WmPoiId int64 `json:"wmPoiId"`
    /**
    * 取餐类型（0：普通取餐；1：到店取餐）
    */
    PickType int32 `json:"pickType"`
    EtaModify EtaModifyVO `json:"etaModify"`
    /**
    * 收餐地址行政区域
    */
    AddressOfReceivingAdministrative string `json:"address_of_receiving_administrative"`
    /**
    * 打包费模块
    */
    PackageBagFeeInfo PackageBagFeeInfo `json:"package_bag_fee_info"`
    /**
    * 是否是商企通订单 1是商企通订单 2非商企通 0降级
    */
    SqtOrder int32 `json:"sqtOrder"`
    /**
    * 是否需要开发票 1需要开发票 2无需开发票 0降级
    */
    NeedSqtInvoice int32 `json:"needSqtInvoice"`
}
type Detail struct {
    /**
    * 菜品规格名称
    */
    Spec string `json:"spec"`
    MtSpuId int64 `json:"mtSpuId"`
    /**
    * 单位
    */
    Unit string `json:"unit"`
    /**
    * 商品单价
    */
    Price float64 `json:"price"`
    /**
    * 菜品属性
    */
    FoodProperty string `json:"foodProperty"`
    /**
    * 菜品code
    */
    AppFoodCode string `json:"appFoodCode"`
    TagName string `json:"tagName"`
    /**
    * sku编码
    */
    SkuId string `json:"skuId"`
    /**
    * 餐盒数量
    */
    BoxNum float64 `json:"boxNum"`
    /**
    * 商品数量
    */
    Quantity float64 `json:"quantity"`
    MtTagId int64 `json:"mtTagId"`
    MtSkuId int64 `json:"mtSkuId"`
    /**
    * 餐盒价格
    */
    BoxPrice float64 `json:"boxPrice"`
    /**
    * 菜品名称
    */
    FoodName string `json:"foodName"`
    /**
    * 订单记录的商品行id，对于相同sku商品可用于标识区分，单订单下唯一
    */
    DetailId int64 `json:"detailId"`
    /**
    * 详情参见7.3.1的sub_detail_list
    */
    SubDetailList []SubDetailList `json:"subDetailList"`
    /**
    * 套餐商品详情(如果订单含有新版套餐，且品牌不支持新版套餐子商品subDetailList，则仍使用本字段兼容返回套餐详情)(示例为新版套餐返回，旧版套餐不包含规格、数量信息)
    */
    DetailExtra Json `json:"detailExtra"`
}
type SubDetailList struct {
}



func (req *WmoperngQueryOrderDetailRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperngQueryOrderDetailResponse, error) {
    resp, err := client.InvokeApi(wmoperng_query_order_detail_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperngQueryOrderDetailResponse
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

func (response *WmoperngQueryOrderDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

