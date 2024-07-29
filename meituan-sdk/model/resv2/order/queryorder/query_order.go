package queryorder

/**
* 查询订单信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_order_url = "/resv2/order/queryById"

type QueryOrderResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    /**
    * 预订订单号
    */
    OrderSerializedId string `json:"orderSerializedId"`
    /**
    * 状态 0:待创建订单 10:新建订单 20:撤单 30:已拒单 40:已接 50:已到店 60:未到店 70:已取消 80:失败 90:申述到店 100:申述未到店 110:失效 120:关闭 
    */
    Status int32 `json:"status"`
    /**
    * 就餐时间 unix时间戳 单位秒
    */
    BookingTime int32 `json:"bookingTime"`
    /**
    * 就餐人数
    */
    Number int32 `json:"number"`
    /**
    * 姓名
    */
    Name string `json:"name"`
    /**
    * 性别 10 - 女士 20 - 先生
    */
    Gender int32 `json:"gender"`
    /**
    * 手机号 status为10、20、30、80时为空 商家接单后提供加密手机号 解密参考https://developer.meituan.com/docs/biz/biz_resv2_897172d9-aae6-46b3-8369-1ca24069b897
    */
    Phone string `json:"phone"`
    /**
    * 附加要求
    */
    Requirements Requirements `json:"requirements"`
    /**
    * 礼遇信息
    */
    Recepts []ReceptParam `json:"recepts"`
    Deposit DepositParam `json:"deposit"`
}
type QueryOrderRequest struct {
    /**
    *  订单id 
    */
    OrderSerializedId string `json:"orderSerializedId"`
}
type DepositParam struct {
    /**
    * 订金支付金额
    */
    PayAmount string `json:"payAmount"`
    /**
    * 订金状态 1-待支付 2-支付成功 3-支付失败 4-退款中 5-退款成功 6-退款失败 7-商家收取 8-打款成功 9-超时未支付 10-取消支付
    */
    PayStatus int32 `json:"payStatus"`
    /**
    * 订金退款金额
    */
    RefundAmount string `json:"refundAmount"`
    /**
    * 订金抵扣金额
    */
    DeductAmount string `json:"deductAmount"`
    /**
    * 订金结算给商家金额
    */
    SettleAmount string `json:"settleAmount"`
}
type ReceptParam struct {
    /**
    * 礼遇名称
    */
    Name string `json:"name"`
    /**
    * 礼遇描述
    */
    Desc string `json:"desc"`
}
type Requirements struct {
    /**
    * 桌台类型 0-大厅 1-包间 3-自定义
    */
    TableType int32 `json:"tableType"`
    /**
    * 桌台类型名称
    */
    TableTypeName string `json:"tableTypeName"`
    /**
    * 备注
    */
    Remark string `json:"remark"`
}



func (req *QueryOrderRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryOrderResponse, error) {
    resp, err := client.InvokeApi(query_order_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryOrderResponse
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

func (response *QueryOrderResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

