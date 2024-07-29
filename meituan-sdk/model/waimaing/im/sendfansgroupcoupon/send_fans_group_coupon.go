package sendfansgroupcoupon

/**
* 在粉丝群内主动建券并发券
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const send_fans_group_coupon_url = "/waimai/ng/im/sendFansGroupCoupon"

type SendFansGroupCouponResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data SendCouponMsg `json:"data"`
    TraceId string `json:"traceId"`
}
type SendFansGroupCouponRequest struct {
    /**
    *  群ID 偏移 
    */
    GroupId int64 `json:"groupId"`
    /**
    *  券领取后有效期，单位秒 
    */
    ValidTime int32 `json:"validTime"`
    /**
    *  门槛 满减券必传 
    */
    LimitPrice int32 `json:"limitPrice"`
    /**
    *  满减券优惠金额， 满减金额需要在使用门槛的3-50% 满减券必传 
    */
    Price int32 `json:"price"`
    /**
    *  三方商品ID 商品券必传 
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    *  商品折扣，0到98，98代表0.98折 商品券必传 
    */
    Discount int32 `json:"discount"`
    /**
    *  券数量，1到9999 
    */
    Count int32 `json:"count"`
    /**
    *  券类型，1：门店满减券，2:商品券 
    */
    ActType int32 `json:"actType"`
    /**
    *  活动有效期，秒 
    */
    ActivityValidTime int32 `json:"activityValidTime"`
    /**
    *  商品ID 商品券（和app_food_code至少传一个 
    */
    SpuId int64 `json:"spuId"`
}
type SendCouponMsg struct {
    /**
    * 券配置ID
    */
    CouponConfigId string `json:"couponConfigId"`
}



func (req *SendFansGroupCouponRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SendFansGroupCouponResponse, error) {
    resp, err := client.InvokeApi(send_fans_group_coupon_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SendFansGroupCouponResponse
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

func (response *SendFansGroupCouponResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

