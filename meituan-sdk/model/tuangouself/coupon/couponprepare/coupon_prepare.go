package couponprepare

/**
* 验券准备
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const coupon_prepare_url = "/tuangouself/coupon/prepare"

type CouponPrepareResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []CouponPrepareData `json:"data"`
    TraceId string `json:"traceId"`
}
type CouponPrepareData struct {
    /**
    * 券码
    */
    Code string `json:"code"`
    /**
    * 项目ID
    */
    DealId int64 `json:"dealId"`
    /**
    * 项目类型，0套餐，1代金券
    */
    DealType int32 `json:"dealType"`
    /**
    * deal对应的菜品信息
    */
    MenuInfo []DealMenu `json:"menuInfo"`
}
type Meal struct {
    /**
    * mealName
    */
    MealName string `json:"mealName"`
    /**
    * 菜品数量
    */
    MealNum string `json:"mealNum"`
}
type CouponPrepareRequest struct {
    /**
    *  券码集合 
    */
    Codes []string `json:"codes"`
}
type DealMenu struct {
    /**
    * 项目ID
    */
    DealId int64 `json:"dealId"`
    /**
    * 菜品分类
    */
    MenuTitle string `json:"menuTitle"`
    /**
    * 菜品列表
    */
    Meals []Meal `json:"meals"`
}



func (req *CouponPrepareRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CouponPrepareResponse, error) {
    resp, err := client.InvokeApi(coupon_prepare_url, 33, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CouponPrepareResponse
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

func (response *CouponPrepareResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

