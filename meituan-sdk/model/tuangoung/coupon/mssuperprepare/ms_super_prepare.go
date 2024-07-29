package mssuperprepare

/**
* 提供开放平台的预验券接口，聚合商品数据信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ms_super_prepare_url = "/tuangou/ng/coupon/msprepare"

type ReceiptPrepareMSDTO struct {
    /**
    * 项目开始时间
    */
    DealBeginTime string `json:"dealBeginTime"`
    /**
    * 项目ID
    */
    DealId int32 `json:"dealId"`
    /**
    * 券面值，即人们常说的市场价
    */
    DealValue float64 `json:"dealValue"`
    /**
    * 项目名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 最大可验证张数
    */
    Count int32 `json:"count"`
    /**
    * 量贩项目的单张券原价（普通券单张券原价与项目总价相同
    */
    SingleValue float64 `json:"singleValue"`
    /**
    * 团购券价格，即商家促销前的券购买价格，如首次购买有更多优惠这类需要在开店宝设置的促销
    */
    DealPrice float64 `json:"dealPrice"`
    /**
    * 套餐类券码对应套餐详细内容，代金券券码包含适用范围(如酒水除外)
    */
    DealMenu string `json:"dealMenu"`
    /**
    * 返回消息
    */
    Message string `json:"message"`
    /**
    * 最小消费张数
    */
    MinConsume int32 `json:"minConsume"`
    /**
    * 券码有效期
    */
    CouponEndTime string `json:"couponEndTime"`
    /**
    * 是否量贩：0：不是，1：是
    */
    Volume int32 `json:"volume"`
    /**
    * 操作状态,0表示成功
    */
    Result int32 `json:"result"`
    /**
    * 券购买价,即用户在购买团购券时所付的实际价格，单张券的用户实付，由于优惠分摊不保证均摊，当Count=1时此字段可用；当Count>1此字段不可用
    */
    CouponBuyPrice float64 `json:"couponBuyPrice"`
    /**
    * 券码
    */
    CouponCode string `json:"couponCode"`
    /**
    * 是否代金券,true代表代金券,false代表套餐券
    */
    Vourcher bool `json:"vourcher"`
    /**
    * 开店宝套餐名
    */
    RawTitle string `json:"rawTitle"`
    /**
    * 开店宝商品列表标题
    */
    PlatformTitle string `json:"platformTitle"`
    /**
    * 第三方商品ID(三方上单填入）
    */
    ThirdProductId string `json:"thirdProductId"`
    /**
    * C端场景中提单页、支付结果页、查看券码和订单页面的项目标题以及开店宝核销后展示的标题。C端场景包括美团App，点评App，美团小程序
    */
    ShortAttrTitle string `json:"shortAttrTitle"`
    /**
    * 返参内容与dealMenu一致，仅解析结构不同
    */
    DealMenuStd string `json:"dealMenuStd"`
}
type MsSuperPrepareRequest struct {
    /**
    *  券码 
    */
    Code string `json:"code"`
}
type MsSuperPrepareResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ReceiptPrepareMSDTO `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *MsSuperPrepareRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*MsSuperPrepareResponse, error) {
    resp, err := client.InvokeApi(ms_super_prepare_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response MsSuperPrepareResponse
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

func (response *MsSuperPrepareResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

