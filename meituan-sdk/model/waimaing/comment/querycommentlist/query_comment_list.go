package querycommentlist

/**
* 查询门店评价信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_comment_list_url = "/waimai/ng/comment/queryCommentList"

type QueryCommentListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []CommentInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryCommentListRequest struct {
    Biz string `json:"biz"`
}
type CommentOrderDetail struct {
    FoodName string `json:"food_name"`
    Count int64 `json:"count"`
}
type CommentInfo struct {
    Result string `json:"result"`
    /**
    * 评价ID
    */
    CommentId int64 `json:"comment_id"`
    /**
    * 用户首次评价的内容
    */
    CommentContent string `json:"comment_content"`
    /**
    * 口味评分
    */
    OrderCommentScore int32 `json:"order_comment_score"`
    /**
    * 商家评分(星级评分)
    */
    FoodCommentScore int32 `json:"food_comment_score"`
    /**
    * 配送评价分数
    */
    DeliveryCommentScore int32 `json:"delivery_comment_score"`
    /**
    * 追评内容
    */
    AddComment string `json:"add_comment"`
    /**
    * 用户追评时间
    */
    AddCommentTime string `json:"add_comment_time"`
    /**
    * 包装评价分数
    */
    PackingScore int32 `json:"packing_score"`
    /**
    * 用户首次评价的时间
    */
    CommentTime string `json:"comment_time"`
    /**
    * 评价标签
    */
    CommentLables string `json:"comment_lables"`
    /**
    * 配送时间。单位为分钟
    */
    ShipTime int32 `json:"ship_time"`
    /**
    * 用户提交的评价图片，多个图片url以英文逗号隔开。
    */
    CommentPictures string `json:"comment_pictures"`
    /**
    * 用户点赞的菜品
    */
    PraiseFoodList string `json:"praise_food_list"`
    /**
    * 用户点踩的菜品
    */
    CriticFoodList string `json:"critic_food_list"`
    /**
    * 商家回复该评价的状态。0未回复，1已回复
    */
    ReplyStatus int32 `json:"reply_status"`
    CommentOrderDetail []CommentOrderDetail `json:"comment_order_detail"`
    /**
    * 商家回复内容
    */
    EReplyContent string `json:"e_reply_content"`
    /**
    * 商家回复时间(精确到天)
    */
    EReplyTime string `json:"e_reply_time"`
    CommentType int64 `json:"comment_type"`
    Valid int64 `json:"valid"`
    /**
    * 配送超时时间。当配送超时时长<10分钟时，展示“超时10分钟内送达”。当配送超时时长≥10分钟时，展示“超时10分钟以上送达”。当配送未超时，不展示。原ship_time字段已废弃。
    */
    OverDeliveryTimeDesc string `json:"over_delivery_time_desc"`
}



func (req *QueryCommentListRequest) DoInvoke(client mtclient.MeituanClient) (*QueryCommentListResponse, error) {
    resp, err := client.InvokeApi(query_comment_list_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response QueryCommentListResponse
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

func (response *QueryCommentListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

