package querycoupon

/**
* 查询粉丝群发券信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_coupon_url = "/waimai/ng/im/queryCoupon"

type QueryCouponResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryCoupon `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryCoupon struct {
    /**
    * 总页数
    */
    TotalPage int32 `json:"totalPage"`
    /**
    * 总条数
    */
    Total int32 `json:"total"`
    /**
    * 券配置ID
    */
    CouponConfigId string `json:"couponConfigId"`
    /**
    * 券名称
    */
    CouponName string `json:"couponName"`
    /**
    * 优惠券金额
    */
    Price int32 `json:"price"`
    /**
    * 使用门槛
    */
    LimitPrice int32 `json:"limitPrice"`
    /**
    * INIT(0,"处理中"), STARTED(1,"进行中"), FINISH(2,"已结束"), STOP(3,"已停止"),
    */
    Status int32 `json:"status"`
    /**
    * 活动创建时间
    */
    CreateTime string `json:"createTime"`
    /**
    * 活动结束时间
    */
    EndTime string `json:"endTime"`
    /**
    * 券领取后有效期，单位秒
    */
    ValidTime int64 `json:"validTime"`
    /**
    * 剩余的券数量 等于总数-已领取的
    */
    Remain int32 `json:"remain"`
    /**
    * 商品折扣，0到98，98代表0.98折
    */
    Discount int32 `json:"discount"`
    /**
    * 商品名称
    */
    SkuName string `json:"skuName"`
    /**
    * 券类型，1：门店满减券，2:商品券
    */
    Category int32 `json:"category"`
}
type QueryCouponRequest struct {
    /**
    *  页码 从 1 开始 
    */
    Page int32 `json:"page"`
    /**
    *  页数据条数 
    */
    PageSize int32 `json:"pageSize"`
}



func (req *QueryCouponRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryCouponResponse, error) {
    resp, err := client.InvokeApi(query_coupon_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryCouponResponse
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

func (response *QueryCouponResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

