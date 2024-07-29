package tuangoureceiptconsume

/**
* 验券
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tuangou_receipt_consume_url = "/ddzh/tuangou/receipt/consume"

type ReceiptConsumeResult struct {
    /**
    * 订单ID
    */
    OrderId string `json:"orderId"`
    /**
    * 验证券码
    */
    ReceiptCode string `json:"receiptCode"`
    /**
    * 套餐ID（若验证的券所对应的商品为团购时，该字段必返回）
    */
    DealId int64 `json:"dealId"`
    /**
    * 团购ID，团购ID与套餐ID是一对多的关系（若验证的券所对应的商品为团购时，该字段必返回）
    */
    DealgroupId int64 `json:"dealgroupId"`
    /**
    * 商品名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 商品类型 1、泛商品如丽人派样活动商品等（若验证的券所对应的商品非团购时，该字段必返回）
    */
    ProductType int64 `json:"productType"`
    /**
    * 商品ID（若验证的券所对应的商品非团购时，该字段必返回，product_item_id含义参考商品管理API）
    */
    ProductItemId int64 `json:"productItemId"`
    /**
    * 商品售卖价格
    */
    DealPrice float64 `json:"dealPrice"`
    /**
    * 商品市场价
    */
    DealMarketPrice float64 `json:"dealMarketPrice"`
    /**
    * 用户手机号，形如：185****1212
    */
    Mobile string `json:"mobile"`
    /**
    * 业务类型 0:普通团购 203:拼团 205:次卡 217:通兑标品
    */
    BizType int32 `json:"bizType"`
    /**
    * 该券码所在的订单的支付明细，如果一笔订单包含两个券码a、b，在核销a、b券码时返回信息一致，都是该订单的支付明细 (如果订单有多张券可以通过订单券码分摊金额查询接口查询分摊信息)
    */
    PaymentDetail []RpaymentDetail `json:"paymentDetail"`
    /**
    * 券过期时间
    */
    ReceiptEndDate string `json:"receiptEndDate"`
    /**
    * 商品快照
    */
    ProductSnapshot string `json:"productSnapshot"`
    /**
    * 流水号
    */
    FlowId string `json:"flowId"`
    /**
    * 是否为团购次卡
    */
    TgTimesCardFlag bool `json:"tgTimesCardFlag"`
    /**
    * 团购次卡可用总次数
    */
    PurchaseToConsumeRatio string `json:"purchaseToConsumeRatio"`
    /**
    * 团购次卡单次卡价格
    */
    TimesCardUnitPrice string `json:"timesCardUnitPrice"`
    /**
    * 附加项目信息列表
    */
    AdditionItemInfos []RadditionItemInfo `json:"additionItemInfos"`
    /**
    * 补贴金额（平台分摊）
    */
    PlatformAmount string `json:"platformAmount"`
    /**
    * 补贴金额（商户分摊）
    */
    MerchantAmount string `json:"merchantAmount"`
}
type TuangouReceiptConsumeRequest struct {
    /**
    *  三方设备类型（未上线），枚举如下： 兑币机 娃娃机 
    */
    ThirdDeviceType string `json:"thirdDeviceType"`
    /**
    *  核销设备ID（未上线） 
    */
    ThirdDeviceId string `json:"thirdDeviceId"`
    /**
    *  三方设备所在城市，地址格式：市-区， 例如“北京市-密云区”（未上线），城市范围： 城市范围 
    */
    ThirdDeviceCity string `json:"thirdDeviceCity"`
    /**
    *  验券数量, 不可多于100个 
    */
    Count int32 `json:"count"`
    /**
    *  团购券码，必须未验证 
    */
    ReceiptCode string `json:"receiptCode"`
    /**
    *  三方位置类型（未上线），枚举如下： 幼儿园 小学 初中 高中&amp;职业技术学校 大学 培训机构 酒店 4S店 步行街 机场 火车站 汽车站 地铁站 工厂 社区 办公楼 医院 政府机构 旅游景点 KTV 游艺厅 电影院 商场 网吧 酒吧 会所 咖啡厅 餐饮店 奶茶店 夜总会 书报亭 其他 
    */
    ThirdLocationType string `json:"thirdLocationType"`
    /**
    *  核销场地ID（未上线） 
    */
    ThirdVenueId string `json:"thirdVenueId"`
    /**
    *  商家在自研系统或第三方服务商系统内登陆的用户名，仅用于记录验券者的信息，该字段不参与任何验券校验逻辑 
    */
    AppShopAccountName string `json:"appShopAccountName"`
    /**
    *  请求ID，用于标识幂等性 
    */
    RequestId string `json:"requestId"`
    /**
    *  商家在自研系统或第三方服务商系统内登录的帐号，仅用于记录验券者的信息，该字段不参与任何验券校验逻辑 
    */
    AppShopAccount string `json:"appShopAccount"`
    /**
    *  核销设备经纬度，格式"经度,纬度"，示例："116.424966,39.930207"（未上线） 
    */
    ThirdDeviceLocation string `json:"thirdDeviceLocation"`
}
type RpaymentDetail struct {
    /**
    * 支付详情ID
    */
    PaymentDetailId string `json:"paymentDetailId"`
    /**
    * 支付金额
    */
    Amount string `json:"amount"`
    /**
    * 金额类型 类型说明： 2：抵用券 5：积分 6：立减 8：商户抵用券 10：C端美团支付 12：优惠代码 15：美团立减 17：商家立减 18：美团商家立减 21：次卡 22：打折卡 23：B端美团支付 24：全渠道会员券 25：pos支付 26：线下认款平台 28：商家折上折 29：美团分销支付
    */
    AmountType int32 `json:"amountType"`
}
type Data struct {
    Result []ReceiptConsumeResult `json:"result"`
}
type RadditionItemInfo struct {
    /**
    * 服务名
    */
    ServiceTitle string `json:"serviceTitle"`
    /**
    * 调价规则当前金额(当前金额，单位：元)
    */
    AdjustedAmount string `json:"adjustedAmount"`
    /**
    * 加项商品ID
    */
    ProductId string `json:"productId"`
    /**
    * 三方加项商品ID
    */
    AddItemThirdPartyId string `json:"addItemThirdPartyId"`
}
type TuangouReceiptConsumeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *TuangouReceiptConsumeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TuangouReceiptConsumeResponse, error) {
    resp, err := client.InvokeApi(tuangou_receipt_consume_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TuangouReceiptConsumeResponse
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

func (response *TuangouReceiptConsumeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

