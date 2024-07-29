package getbatchorderrefundinfo

/**
* 批量查询退款订单信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_batch_order_refund_info_url = "/waimai/ng/order/getBatchOrderRefundInfo"

type GetBatchOrderRefundInfoRequest struct {
    /**
    *  订单id列表 
    */
    OrderIds []string `json:"orderIds"`
}
type OrderRefundInfo struct {
    /**
    * 订单id
    */
    OrderId int64 `json:"order_id"`
    /**
    * 1-部分退款；2-全部退款
    */
    RefundType string `json:"refund_type"`
    /**
    * 退款申请时间
    */
    RefundApplyTime int64 `json:"refund_apply_time"`
    /**
    * 退款原因
    */
    Reason string `json:"reason"`
    /**
    * 0：未处理；1：商家驳回退款请求；2、商家同意退款；3、客服驳回退款请求；4、客服帮商家同意退款；5、超过3小时自动同意；6、系统自动确认；7：用户取消退款申请；8：用户取消退款申诉
    */
    ResType int32 `json:"res_type"`
    /**
    * 是否申诉退款：0-否；1-是
    */
    IsAppeal int32 `json:"is_appeal"`
    /**
    * 退款图片：Json数组；用户在申请退款时上传的退款图片；上限为9张图片；
    */
    Pictures string `json:"pictures"`
    /**
    * 退款金额，refund_type=2（全部退款）无该金额
    */
    Money string `json:"money"`
    /**
    * 订单业务打标枚举，即业务身份
    */
    OrderTagList string `json:"order_tag_list"`
    /**
    * 售后发起人角色： * 1-用户发起售后 * 2-客服帮助用户发起售后 * 3-重复支付自动发起退款 * 6-用户发起申诉 * 7-商家发起售后 * 0-其他
    */
    ApplyRole int32 `json:"apply_role"`
    /**
    * 退款申请单唯一标识（申诉入参 after_sale_id等价）
    */
    RefundId int64 `json:"refund_id"`
    /**
    * 商家是否可以申诉，1-可以、0 -不可以
    */
    PoiCanAppeal int32 `json:"poi_can_appeal"`
    /**
    * 用户申请退款类型，1-用户申请 、2-用户第一次申诉 、3-用户第二次申诉
    */
    UserApplyType int32 `json:"user_apply_type"`
    /**
    * 退款商品信息（仅部分退款有该字段）
    */
    Food []WmOrderPartRefundFoodInfo `json:"food"`
}
type RespData struct {
    /**
    * 订单id
    */
    OrderId int64 `json:"order_id"`
    /**
    * 全部或部分退款信息列表
    */
    OrderRefundInfo []OrderRefundInfo `json:"OrderRefundInfo"`
}
type GetBatchOrderRefundInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data RespData `json:"data"`
    TraceId string `json:"traceId"`
}
type WmOrderPartRefundFoodInfo struct {
    /**
    * 订单记录的商品行id，对于相同sku的商品可用于标识，单订单下唯一
    */
    DetailId int64 `json:"detail_id"`
    /**
    * APP方菜品id，最大长度128，不同门店可以重复，同一门店内不能重复
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    * 退款菜品名称
    */
    FoodName string `json:"food_name"`
    /**
    * sku码
    */
    SkuId string `json:"sku_id"`
    /**
    * 单位
    */
    Spec string `json:"spec"`
    /**
    * 商品价格
    */
    FoodPrice float64 `json:"food_price"`
    /**
    * 商品数量
    */
    Count int32 `json:"count"`
    /**
    * 打包盒数量
    */
    BoxNum float64 `json:"box_num"`
    /**
    * 打包盒价格
    */
    BoxPrice float64 `json:"box_price"`
    /**
    * 菜品原价，单位元
    */
    OriginFoodPrice float64 `json:"origin_food_price"`
    /**
    * 退款价格，单位元
    */
    RefundPrice float64 `json:"refund_price"`
    /**
    * 订单数据状态标记。当订单中部分字段的数据因内部交互异常或网络等原因延迟生成（超时），导致开发者当前获取的订单数据不完整，此时平台对订单数据缺失情况进行标记。如不完整，建议尝试重新查询。注意，平台仅对部分模块的数据完整性进行监察标记（参考incmp_modules字段）。参考值： -1：有数据降级 0：无数据降级
    */
    IncmpCode int32 `json:"incmp_code"`
    /**
    * 0：订单商品详情 1：订单优惠信息 2：商品优惠详情 3：订单用户会员信息 4：订单维度的商家对账信息 5：订单维度的商家对账信息(元) 6：订单收货人地址 7：订单配送方式 8：开放平台用户id 9：部分退款商品信息 10：退货退款物流信息 11：部分订单基本信息(包括订单优惠信息、订单商品详情、门店信息等) 12：sku信息 13：spu信息 14：商品信息(可能是sku或spu等商品相关信息获取时降级)
    */
    IncmpModules int32 `json:"incmp_modules"`
}



func (req *GetBatchOrderRefundInfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetBatchOrderRefundInfoResponse, error) {
    resp, err := client.InvokeApi(get_batch_order_refund_info_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetBatchOrderRefundInfoResponse
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

func (response *GetBatchOrderRefundInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

