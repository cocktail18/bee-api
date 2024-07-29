package orderquerybyid

/**
* 根据订单Id查询订单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_query_by_id_url = "/waimai/order/queryById"

type PackageBagCartDetail struct {
    /**
    * 口袋id
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
type OrderQueryByIdRequest struct {
    /**
    *  订单号 
    */
    OrderId int64 `json:"orderId"`
}
type EtaModifyVO struct {
    /**
    * 发起申请时间
    */
    UetaCtime int32 `json:"ueta_ctime"`
    /**
    * 用户申请修改eta时间
    */
    UetaTime int32 `json:"ueta_time"`
    /**
    * 截止时间
    */
    UetaExpire int32 `json:"ueta_expire"`
    /**
    * 审核状态ueta_status 1.申请，2.通过，3.商家拒绝，4.超时未处理，5.订单完成或取消自动拒绝 6.订单处于退款中自动拒绝 7.商家点击平台原因拒绝（骑手已接单、骑手已到店、骑手已取餐、骑手已送达）8.商家点击平台原因拒绝（跑腿自动锁单状态）9.商家点击平台原因拒绝（专人直送）)
    */
    UetaStatus int32 `json:"ueta_status"`
}
type OrderQueryByIdResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data OrderQueryByIdData `json:"data"`
    TraceId string `json:"traceId"`
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
type OrderQueryByIdData struct {
    /**
    * 订单创建时间 整性数据，单位是秒
    */
    CTime int64 `json:"cTime"`
    /**
    * caution字段为用户下单时的备注 ，具体表现：有备注且开通隐私号功能，格式内容为'备注内容+收餐人隐私号+脱敏的真实号' ；有备注未开通隐私号功能，格式内容为“备注内容”；无备注但开通隐私号功能，格式内容为'收餐人隐私号13141047707_6179,手机号186****9921' ；无备注未开通隐私号功能，不显示；用户选择到店自取，备注内容显示”到店自取”。
    */
    Caution string `json:"caution"`
    /**
    * 城市Id
    */
    CityId int64 `json:"cityId"`
    /**
    * 门店当天的订单流水号 每天流水号从1开始
    */
    DaySeq int32 `json:"daySeq"`
    /**
    * 用户预计送达时间，“立即送达”时为0 整性数据，单位是秒, 如取餐类型为到店自取，则该时间为用户自选的到店自取时间。
    */
    DeliveryTime int64 `json:"deliveryTime"`
    /**
    * 订单菜品详情 参照 detail 指的是订单中都有哪些菜品
    */
    Detail string `json:"detail"`
    /**
    * 就餐人数 用餐人数（0：用户没有选择用餐人数；1-10：用户选择的用餐人数；-10：10人以上用餐；88：用户需要餐具；99：用户不需要餐具），该信息默认不推送，如有需求可在开发者中心订阅
    */
    DinnersNumber int32 `json:"dinnersNumber"`
    /**
    * ERP方门店id 指的是三方系统中的门店Id
    */
    EPoiId string `json:"ePoiId"`
    /**
    * 订单扩展信息 参考 extras 指的是订单所参加的优惠等信息
    */
    Extras string `json:"extras"`
    /**
    * 是否需要发票 1-需要发票；0-不需要
    */
    HasInvoiced int32 `json:"hasInvoiced"`
    /**
    * 订单数据状态标记。当订单中部分字段的数据因内部交互异常或网络等原因延迟生成（超时），导致开发者当前获取的订单数据不完整，此时平台对订单数据缺失情况进行标记。如不完整，建议尝试重新查询。注意，平台仅对部分模块的数据完整性进行监察标记（参考incmp_modules字段）。参考值： -1：有数据降级 0：无数据降级
    */
    IncmpCode int32 `json:"incmpCode"`
    /**
    * 有降级的数据模块的集合
    */
    IncmpModules []int32 `json:"incmpModules"`
    InvMakeType int64 `json:"invMakeType"`
    /**
    * 发票抬头 如果用户选择需要发票，此字段是用户填写的发票抬头
    */
    InvoiceTitle string `json:"invoiceTitle"`
    /**
    * 用户是否收藏此门店 该信息默认不展示，如有需求可联系开放平台工作人员开通
    */
    IsFavorites bool `json:"isFavorites"`
    /**
    * 用户是否第一次在此门店点餐 该信息默认不展示，如有需求可联系开放平台工作人员开通
    */
    IsPoiFirstOrder bool `json:"isPoiFirstOrder"`
    IsPre int64 `json:"isPre"`
    /**
    * 是否第三方配送 0：否；1：是 目前基本上不支持第三方配送
    */
    IsThirdShipping int32 `json:"isThirdShipping"`
    /**
    * 实际送餐地址纬度 美团使用的是高德坐标系
    */
    Latitude float64 `json:"latitude"`
    /**
    * 配送类型码 参考附录配送类型码
    */
    LogisticsCode string `json:"logisticsCode"`
    /**
    * 实际送餐地址经度 美团使用的是高德坐标系
    */
    Longitude float64 `json:"longitude"`
    /**
    * 标识订单是否为团餐定制菜单 1：团餐定制菜单。0或空：外卖菜单
    */
    OrderEntranceType int32 `json:"orderEntranceType"`
    /**
    * 订单Id
    */
    OrderId int64 `json:"orderId"`
    /**
    * 订单展示Id 指的是C端用户在外卖App上看到的订单号
    */
    OrderIdView int64 `json:"orderIdView"`
    /**
    * 用户下单时间 整性数据，单位是秒
    */
    OrderSendTime int64 `json:"orderSendTime"`
    /**
    * 订单业务打标枚举，即业务身份 json字符串格式的数组，例：'[16]'
    */
    OrderTagList string `json:"orderTagList"`
    /**
    * 订单原价
    */
    OriginalPrice float32 `json:"originalPrice"`
    /**
    * 订单支付类型 （1：货到付款；2：在线支付）
    */
    PayType int32 `json:"payType"`
    /**
    * 门店地址
    */
    PoiAddress string `json:"poiAddress"`
    /**
    * 门店Id 指的是外卖中的门店Id
    */
    PoiId int64 `json:"poiId"`
    /**
    * 门店名称
    */
    PoiName string `json:"poiName"`
    /**
    * 门店服务电话
    */
    PoiPhone string `json:"poiPhone"`
    /**
    * 商家对账信息 参考poiReceiveDetail
    */
    PoiReceiveDetail string `json:"poiReceiveDetail"`
    /**
    * 脱敏收货人地址。订单完成3小时后展示“为保护顾客隐私，地址已隐藏”。实际的地址 @#后是经纬度反查的地址，是用户订餐时定位的地址
    */
    RecipientAddress string `json:"recipientAddress"`
    /**
    * 脱敏收货人地址。订单完成3小时后展示“为保护顾客隐私，地址已隐藏”
    */
    RecipientAddressDesensitization string `json:"recipientAddressDesensitization"`
    /**
    * 收货人名称
    */
    RecipientName string `json:"recipientName"`
    /**
    * 收货人电话 C端启用号码保护时该字段展示隐私号，如15696424612_7472。C端未启用号码保护或隐私号降级时，该字段为真实手机号，如12345678901
    */
    RecipientPhone string `json:"recipientPhone"`
    /**
    * 骑手电话 跟logisticsDispatcherMobile值相同，以此字段为准
    */
    ShipperPhone string `json:"shipperPhone"`
    /**
    * 配送费用
    */
    ShippingFee float32 `json:"shippingFee"`
    /**
    * 订单状态 查看备注信息
    */
    Status int32 `json:"status"`
    /**
    * 发票税号
    */
    TaxpayerId string `json:"taxpayerId"`
    /**
    * 总价 用户实际支付金额
    */
    Total float32 `json:"total"`
    /**
    * 订单更新时间
    */
    UTime int64 `json:"uTime"`
    /**
    * 备份隐私号 具体为json格式的字符串数组 1、当recipientPhone字段为隐私号时，该字段作为其备份，目前为1个号码，将来为可能为多个，用英文逗号隔开“,”。 2、有至少2个可用的隐私号时，该字段返回备用的隐私号；如果只有1个隐私号可用或所有隐私号均不可用时，该字段返回内容为空。 3、recipientPhone和backupRecipientPhone的隐私号间没有优劣之分，任何一个隐私号故障时，请尝试用其他隐私号联系用户。当所有隐私号都故障且美团触发降级时，重新获取的recipientPhone字段才会展示真实号。
    */
    BackupRecipientPhone string `json:"backupRecipientPhone"`
    LogisticsCancelTime int64 `json:"logisticsCancelTime"`
    /**
    * 配送完成时间
    */
    LogisticsCompletedTime int64 `json:"logisticsCompletedTime"`
    /**
    * 配送单确认时间，骑手接单时间
    */
    LogisticsConfirmTime int64 `json:"logisticsConfirmTime"`
    /**
    * 骑手电话
    */
    LogisticsDispatcherMobile string `json:"logisticsDispatcherMobile"`
    /**
    * 骑手姓名
    */
    LogisticsDispatcherName string `json:"logisticsDispatcherName"`
    /**
    * 骑手取单时间
    */
    LogisticsFetchTime int64 `json:"logisticsFetchTime"`
    /**
    * 配送方ID
    */
    LogisticsId int32 `json:"logisticsId"`
    /**
    * 配送方名称
    */
    LogisticsName string `json:"logisticsName"`
    /**
    * 配送单下单时间
    */
    LogisticsSendTime int64 `json:"logisticsSendTime"`
    /**
    * 配送订单状态code 参考《补充相关字段说明》
    */
    LogisticsStatus int32 `json:"logisticsStatus"`
    /**
    * 商家确认时间
    */
    OrderConfirmTime int64 `json:"orderConfirmTime"`
    /**
    * 取餐类型 0：普通取餐；1：到店取餐 该信息默认不展示，如有需求可联系开放平台工作人员开通
    */
    PickType int32 `json:"pickType"`
    /**
    * 订单完成时间
    */
    OrderCompletedTime int64 `json:"orderCompletedTime"`
    EtaModify EtaModifyVO `json:"etaModify"`
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
    /**
    * 预计送达时间
    */
    EstimateArrivalTime int64 `json:"estimateArrivalTime"`
}



func (req *OrderQueryByIdRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderQueryByIdResponse, error) {
    resp, err := client.InvokeApi(order_query_by_id_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderQueryByIdResponse
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

func (response *OrderQueryByIdResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

