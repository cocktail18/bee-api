package queryposter

/**
* 商家开放平台海报查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_poster_url = "/waimai/ng/decoration/queryPoster"

type QueryPosterRequest struct {
    Biz string `json:"biz"`
}
type QueryPosterResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryPosterData `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryPosterData struct {
    /**
    * 页面
    */
    PageNum int32 `json:"pageNum"`
    /**
    * 每页展示数
    */
    PageSize int32 `json:"pageSize"`
    /**
    * 总数
    */
    Count int32 `json:"count"`
    /**
    * 海报详情
    */
    WmAppPostVos []WmAppPostVos `json:"wmAppPostVos"`
}
type WmAppPostFoodVos struct {
    /**
    * 菜品id
    */
    AppFoodCode int64 `json:"app_food_code"`
    /**
    * 菜品名称
    */
    FoodName string `json:"food_name"`
    /**
    * 标签名
    */
    TagName string `json:"tag_name"`
    /**
    * 是否有图片
    */
    HasPic bool `json:"has_pic"`
}
type WmAppPostVos struct {
    /**
    * 海报id
    */
    PosterId int64 `json:"posterId"`
    /**
    * 图片地址
    */
    SourcePicUrl string `json:"sourcePicUrl"`
    /**
    * 商品信息
    */
    WmAppPostFoodVos []WmAppPostFoodVos `json:"wmAppPostFoodVos"`
    /**
    * 展示顺序
    */
    Sequence int32 `json:"sequence"`
    /**
    * 审核状态 0：免审 1：审核中 2： 审核通过 3：审核驳回
    */
    VerifyStatus int32 `json:"verifyStatus"`
    /**
    * 驳回原因
    */
    RejectReason string `json:"rejectReason"`
    /**
    * 是否在c端展示
    */
    Valid int32 `json:"valid"`
    /**
    * 海报生效开始日期,默认-1(格式：都是秒)：
    */
    DisplayStartDay int32 `json:"display_start_day"`
    /**
    * 海报生效结束日期,默认-1(格式：都是秒)：
    */
    DisplayEndDay int32 `json:"display_end_day"`
    /**
    * 海报生效开始时间,默认-1(格式：0至246060-1)
    */
    DisplayStime int32 `json:"display_stime"`
    /**
    * 海报生效结束时间,默认-1(格式：0至246060-1)
    */
    DisplayEtime int32 `json:"display_etime"`
    /**
    * 展示周期 海报投放周几,传入1，2，以此类推，默认[1,2,3,4,5,6,7]
    */
    DisplayWeeks string `json:"display_weeks"`
}



func (req *QueryPosterRequest) DoInvoke(client mtclient.MeituanClient) (*QueryPosterResponse, error) {
    resp, err := client.InvokeApi(query_poster_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response QueryPosterResponse
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

func (response *QueryPosterResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

