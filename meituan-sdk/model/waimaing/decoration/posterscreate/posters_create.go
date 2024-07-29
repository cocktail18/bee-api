package posterscreate

/**
* 商家开放平台创建海报
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const posters_create_url = "/waimai/ng/decoration/postersCreate"

type PostersCreateRequest struct {
    /**
    *  海报绑定商品列表,逗号表达式,如'1,2,3' ， 注： 商品必须通过开放平台创建，商品来源可参考门店装修-商品查询 
    */
    AppFoodCodes string `json:"appFoodCodes"`
    /**
    *  图片code，服务商可通过 上传无水印照片 获取 
    */
    PicCode string `json:"picCode"`
    /**
    *  海报生效开始日期,默认-1(格式：都是秒) 
    */
    DisplayStartDay int64 `json:"displayStartDay"`
    /**
    *  海报生效结束日期,默认-1(格式：都是秒) 
    */
    DisplayEndDay int64 `json:"displayEndDay"`
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
    *  是否审核通过后立即使用 1:是;0:否 
    */
    ValidStatus string `json:"validStatus"`
    /**
    *  海报展示顺序,默认0 
    */
    Sequence int32 `json:"sequence"`
    /**
    *  海报id（新建的时候postid是0 ,只有修改的时候是不为空 
    */
    PostId int64 `json:"postId"`
}
type PostersCreateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data PostersCreateData `json:"data"`
    TraceId string `json:"traceId"`
}
type PostersCreateData struct {
    /**
    * 海报id
    */
    PostId int64 `json:"post_id"`
}



func (req *PostersCreateRequest) DoInvoke(client mtclient.MeituanClient) (*PostersCreateResponse, error) {
    resp, err := client.InvokeApi(posters_create_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response PostersCreateResponse
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

func (response *PostersCreateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

