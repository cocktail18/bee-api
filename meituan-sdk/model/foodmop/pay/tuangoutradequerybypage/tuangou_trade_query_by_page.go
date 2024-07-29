package tuangoutradequerybypage

/**
* 对账分页查询账单信息（非必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tuangou_trade_query_by_page_url = "/foodmop/pay/tuangou/trade/queryByPage"

type TradeDetail struct {
    /**
    * 美团券码，bizType=10团购业务 买单订单号,bizType=20买单业务 点餐订单号，biztype=40点餐业务
    */
    CouponCode string `json:"couponCode"`
    /**
    * 美团团单ID，bizType=10 买单项目ID，bizType=20 不传，bizType=40
    */
    DealId int32 `json:"dealId"`
    /**
    * 美团团单名称,bizType in (10,20)有值 不传，bizType=40点餐
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 原价，单位元，保留两位小数 bizType=40点餐，为0.00
    */
    OriginalPrice float64 `json:"originalPrice"`
    /**
    * 券售价/买单消费金额，单位元，保留两位小数 bizType=40点餐，为0.00
    */
    SalePrice float64 `json:"salePrice"`
    /**
    * 服务费，单位元，保留两位小数 bizType=10团购，服务费=售价salePrice - 进价buyPrice； bizType=20买单，美团扣除的商家佣金 bizType=40点餐，美团扣除的商家佣金，actType为10消费时为正数，20撤销消费为负数
    */
    ServiceFee float64 `json:"serviceFee"`
    /**
    * 商家优惠，单位元，保留两位小数 bizType=40点餐，商户优惠，actType为10消费时为正数，20撤销消费为负数
    */
    BizCost float64 `json:"bizCost"`
    /**
    * 结算金额，单位元，保留两位小数 bizType=10团购，结算金额=进价 buyPrice- 商家优惠bizCost； bizType=20买单，结算金额即指商家实际到账的钱 bizType=40点餐，结算金额即指商家实际到账的钱，actType为10消费时为正数，20撤销消费为负数
    */
    SettlementAmount float64 `json:"settlementAmount"`
    /**
    * 进价，单位元，保留两位小数 bizType=40点餐，为0.00
    */
    Buyprice float64 `json:"buyprice"`
    /**
    * 退款金额，单位元，保留两位小数 bizType=40点餐，退款金额，actType为10消费时为0.00，20撤销消费为0.00
    */
    BizRefund float64 `json:"bizRefund"`
    /**
    * 流水类型 10 : 消费 20 : 撤销消费 30 : 已消费退 40 : 购买 50 : 未消费退
    */
    ActType int32 `json:"actType"`
    /**
    * 交易流水发生时间，毫秒时间戳，需要转换成日期格式，需使用北京时间，即UTC+8时区 具体释义视actType（流水类型）字段取值而定 例：若actType=10，即消费流水，则actTime为消费时间；若actType=20，即撤销消费流水，则actTime为撤销消费时间；其余取值情况以此类推
    */
    ActTime string `json:"actTime"`
    /**
    * 收益时间，毫秒时间戳，需要转换成日期格式，需使用北京时间，即UTC+8时区
    */
    ClearTime string `json:"clearTime"`
    /**
    * 美团门店Id,Long类型
    */
    PoiIdL int64 `json:"poiIdL"`
    /**
    * 收银门店Id即商家门店Id
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    * 收银订单号 bizType=10且券码用于点餐，必传，即商家点餐订单号（当商家点餐拒单或未接单，则该值为空） bizType=10且券码未用于点餐，非必传，即商家上传的商家订单号 当BizType=40即点餐业务时，该字段必传，表示商家点餐订单号
    */
    VendorOrderId string `json:"vendorOrderId"`
    /**
    * 扩展信息
    */
    ExtendInfo ExtendInfo `json:"extendInfo"`
}
type ResultData struct {
    /**
    * 流水数据，查询成功时返回
    */
    TradeDetailList []TradeDetail `json:"tradeDetailList"`
}
type TuangouTradeQueryByPageResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type ExtendInfo struct {
    /**
    * 服务费率，保留4位小数 bizType=40点餐必传，单位1，保留4位小数。如0.0060表示千分之六的费率
    */
    ServiceFeeRate float64 `json:"serviceFeeRate"`
    /**
    * 团购券抵扣金额，单位元，保留两位小数 bizType=40点餐必传，即美团团购券（含代金券、套餐券）在该笔点餐订单中总计抵扣掉的金额，actType为10消费时为正数，20撤销消费为负数
    */
    CouponExchangeAmount float64 `json:"couponExchangeAmount"`
    /**
    * 美团订单号 bizType=10且券码用于点餐，必传，即美团系统内点餐订单号
    */
    OrderId int64 `json:"orderId"`
}
type TuangouTradeQueryByPageRequest struct {
    /**
    *  收益日期，即触发收益结算的日期 格式为"yyyy-MM-dd" 注： 结算信息完备的正常情况下，收益日期即核销日期，举例：团购券在5月1日核销，正常情况下5月1日触发收益计算，商家在5月2日收到打款，此时“收益日期”字段为5月1日 若结算信息不全，会导致延期结算，收益日期为商家收款日期-1，举例：团购券在5月1日核销，6月1日才触发收益计算，商家在6月2日收到打款，此时“收益日期”字段为6月1日 
    */
    ClearDate string `json:"clearDate"`
    /**
    *  查询条数（最大不超过1000） 举例： 查询900条明细，需调用1次，limit=900； 查询1200条明细，需调用2次，第1次 limit=1000，第2次limit=200； 查询2300条明细；需调用3次，第1次 limit=1000，第2次limit=1000，第3次 limit=300； 依此类推 
    */
    Limit int32 `json:"limit"`
    /**
    *  查询起始位置(从0开始，根据实际数据量依次增大) 举例： 查询900条明细，需调用1次，offset=0； 查询1200条明细，需调用2次，第1次 offset=0，第2次offset=1000； 查询2300条明细；需调用3次，第1次 offset=0，第2次offset=1000，第3次 offset=2000； 依此类推 
    */
    Offset int32 `json:"offset"`
    /**
    *  业务类型 10： 团 购 20： 买 单 40：点餐 
    */
    BizType int32 `json:"bizType"`
}



func (req *TuangouTradeQueryByPageRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TuangouTradeQueryByPageResponse, error) {
    resp, err := client.InvokeApi(tuangou_trade_query_by_page_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TuangouTradeQueryByPageResponse
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

func (response *TuangouTradeQueryByPageResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

