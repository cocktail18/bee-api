package wmoperfoodbatchquerylist

/**
* 批量查询外卖菜品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_food_batch_query_list_url = "/wmoper/ng/food/batchQuery"

type WmoperFoodBatchQueryListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []WmoperFoodBatchQueryListData `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperFoodBatchQueryListData struct {
    EpoiId string `json:"epoiId"`
    AppFoodCode string `json:"app_food_code"`
    Name string `json:"name"`
    Description string `json:"description"`
    Price float64 `json:"price"`
    MinOrderCount int64 `json:"min_order_count"`
    Unit string `json:"unit"`
    BoxNum float64 `json:"box_num"`
    BoxPrice float64 `json:"box_price"`
    CategoryName string `json:"category_name"`
    IsSoldOut int64 `json:"is_sold_out"`
    Picture string `json:"picture"`
    Sequence int64 `json:"sequence"`
    Skus string `json:"skus"`
    Ctime int64 `json:"ctime"`
    Utime int64 `json:"utime"`
    SpuAttr string `json:"spuAttr"`
    MtSpuId int64 `json:"mt_spu_id"`
    MtTagId int64 `json:"mt_tag_id"`
    TagName string `json:"tag_name"`
    OriginSpuId int64 `json:"origin_spu_id"`
    Pictures string `json:"pictures"`
    HighLight string `json:"high_light"`
    Speciality int64 `json:"speciality"`
    IsNotSingle int64 `json:"is_not_single"`
    MonthSaled int64 `json:"monthSaled"`
    /**
    * 图文详情URL
    */
    LongPictures string `json:"longPictures"`
}
type WmoperFoodBatchQueryListRequest struct {
    /**
    *  商品code列表，最多20个，英文逗号隔开【,】 
    */
    AppFoodCodes string `json:"app_food_codes"`
}



func (req *WmoperFoodBatchQueryListRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperFoodBatchQueryListResponse, error) {
    resp, err := client.InvokeApi(wmoper_food_batch_query_list_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperFoodBatchQueryListResponse
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

func (response *WmoperFoodBatchQueryListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

