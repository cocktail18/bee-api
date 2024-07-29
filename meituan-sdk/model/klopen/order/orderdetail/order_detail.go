package orderdetail

/**
* 查询订单详细信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_detail_url = "/kl/open/order/detail"

type SkuTicketDetailDto struct {
    TicketName string `json:"ticketName"`
    TicketUrlList []string `json:"ticketUrlList"`
}
type OrderDetailResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderDetailRequest struct {
    /**
    *  快驴订单编号 
    */
    OrderId string `json:"orderId"`
    /**
    *  品牌标识，除华住外其他服务商必传 
    */
    BrandIdentify string `json:"brandIdentify"`
}
type Data struct {
    OrderId string `json:"orderId"`
    Status string `json:"status"`
    CreateDate string `json:"createDate"`
    ReceiverName string `json:"receiverName"`
    ReceiverPhone string `json:"receiverPhone"`
    AddressDetail string `json:"addressDetail"`
    ExpectTime string `json:"expectTime"`
    Remark string `json:"remark"`
    ProTotalAmount string `json:"proTotalAmount"`
    FreightAmount string `json:"freightAmount"`
    DiscountAmount string `json:"discountAmount"`
    ReceivableAmount string `json:"receivableAmount"`
    TotalAmount string `json:"totalAmount"`
    CancelReason string `json:"cancelReason"`
    CancelRemark string `json:"cancelRemark"`
    CancelUser string `json:"cancelUser"`
    PayType string `json:"payType"`
    VerStatus string `json:"verStatus"`
    OrderProducts OrderDetailSkuItemDto `json:"orderProducts"`
    DriverName string `json:"driverName"`
    DriverPhone string `json:"driverPhone"`
    PlateNumber string `json:"plateNumber"`
}
type OrderDetailSkuItemDto struct {
    OrderId string `json:"orderId"`
    SkuCode int32 `json:"skuCode"`
    SkuName string `json:"skuName"`
    SpecificationsDes string `json:"specificationsDes"`
    Unit string `json:"unit"`
    Price string `json:"price"`
    Quantity int32 `json:"quantity"`
    DeliveryQuantity int32 `json:"deliveryQuantity"`
    DifferenceAmount string `json:"differenceAmount"`
    DifferenceWeight string `json:"differenceWeight"`
    SignWeight string `json:"signWeight"`
    EstimateWeight string `json:"estimateWeight"`
    SkuTicketDetailDtoList []SkuTicketDetailDto `json:"skuTicketDetailDtoList"`
}



func (req *OrderDetailRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderDetailResponse, error) {
    resp, err := client.InvokeApi(order_detail_url, 27, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderDetailResponse
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

func (response *OrderDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

