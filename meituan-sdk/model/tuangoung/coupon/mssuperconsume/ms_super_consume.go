package mssuperconsume

/**
* 提供开放平台的验券接口，聚合商品数据信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ms_super_consume_url = "/tuangou/ng/coupon/msconsume"

type ReceiptConsumeMSDTO struct {
    /**
    * 验证券码数组
    */
    CouponCodes []string `json:"couponCodes"`
    /**
    * 项目名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 项目ID
    */
    DealId int32 `json:"dealId"`
    /**
    * 券面值，即人们常说的市场价
    */
    DealValue float32 `json:"dealValue"`
    /**
    * 开店宝套餐名
    */
    RawTitle string `json:"rawTitle"`
    /**
    * 开店宝商品列表标题
    */
    PlatformTitle string `json:"platformTitle"`
    /**
    * C端场景中提单页、支付结果页、查看券码和订单页面的项目标题以及开店宝核销后展示的标题。C端场景包括美团App，点评App，美团小程序
    */
    ShortAttrTitle string `json:"shortAttrTitle"`
    /**
    * 第三方商品ID
    */
    ThirdProductId string `json:"thirdProductId"`
    /**
    * 商家实收 = 售价salePrice - 商促bizCost
    */
    ReceiptInfoDTOList []ReceiptInfoDTO `json:"receiptInfoDTOList"`
}
type ReceiptInfoDTO struct {
}
type MsSuperConsumeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ReceiptConsumeMSDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type MsSuperConsumeRequest struct {
    /**
    *  券码 
    */
    Code string `json:"code"`
    /**
    *  核销张数（支持该券码关联的订单下其他券码一起核销，核销本张券码传1） 
    */
    Num int32 `json:"num"`
}



func (req *MsSuperConsumeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*MsSuperConsumeResponse, error) {
    resp, err := client.InvokeApi(ms_super_consume_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response MsSuperConsumeResponse
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

func (response *MsSuperConsumeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

