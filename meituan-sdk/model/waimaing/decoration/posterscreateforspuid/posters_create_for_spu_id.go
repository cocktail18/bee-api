package posterscreateforspuid

/**
* 门店装修-新建海报接口(通过spuid)
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const posters_create_for_spu_id_url = "/waimai/ng/decoration/postersCreateForSpuId"

type PostersCreateForSpuIdResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data PostersCreateForSpuIdData `json:"data"`
    TraceId string `json:"traceId"`
}
type PostersCreateForSpuIdData struct {
    /**
    * 海报id
    */
    PostId int64 `json:"post_id"`
}
type PostersCreateForSpuIdRequest struct {
    /**
    *  海报绑定商品列表,逗号表达式,如'1,2,3' 注：商品必须通过开放平台创建 
    */
    SpuIds string `json:"spuIds"`
    /**
    *  图片code 
    */
    PicCode string `json:"picCode"`
    /**
    *  海报生效开始日期,默认-1(格式：都是秒) 
    */
    DisplayStartDay int32 `json:"displayStartDay"`
    /**
    *  海报生效结束日期,默认-1(格式：都是秒) 
    */
    DisplayEndDay int32 `json:"displayEndDay"`
    /**
    *  展示周期 海报投放周几,传入1，2，以此类推，默认[1,2,3,4,5,6,7] 
    */
    DisplayWeeks string `json:"displayWeeks"`
    /**
    *  海报生效开始时间,默认-1(格式：0至24*60*60-1) 
    */
    DisplayStime string `json:"displayStime"`
    /**
    *  海报生效结束时间,默认-1(格式：0至24*60*60-1) 
    */
    DisplayEtime string `json:"displayEtime"`
    /**
    *  是否审核通过后立即使用 
    */
    ValidStatus int32 `json:"validStatus"`
    /**
    *  海报展示顺序,默认0 
    */
    Sequence int32 `json:"sequence"`
    /**
    *  海报id（新建的时候postid是0 ,只有修改的时候是不为空） 
    */
    PostId int32 `json:"postId"`
}



func (req *PostersCreateForSpuIdRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*PostersCreateForSpuIdResponse, error) {
    resp, err := client.InvokeApi(posters_create_for_spu_id_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response PostersCreateForSpuIdResponse
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

func (response *PostersCreateForSpuIdResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

