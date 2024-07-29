package batchupdatespubyid

/**
* 批量通过美团spuid修改商品部分信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_update_spuby_id_url = "/waimai/ng/dish/batchUpdateSpuById"

type BatchUpdateSpubyIdRequest struct {
    FoodData []WmProductSpu `json:"food_data"`
}
type WmProductSpu struct {
    /**
    *  商品真实id 
    */
    OriginSpuId int64 `json:"origin_spu_id"`
    /**
    *  商品描述 
    */
    SpuDesc string `json:"spu_desc"`
    /**
    *  商品图片（图片id/URL） 
    */
    SpuPicture string `json:"spu_picture"`
    /**
    *  商品图片多图（图片id/URL），英文逗号隔开，最多5张。 注意： 1、如果spu_picture和spu_pictures都有值，则以spu_picture为第一张，spu_pictures从第二张开始生效； 2、如果仅spu_picture有值，则商品只有一张图片； 3、如果仅spu_pictures有值，则商品图片为spu_pictures生效； 4、如果2个字段都没有值，则商品无图。 
    */
    SpuPictures string `json:"spu_pictures"`
    /**
    *  图文详情URL 英文逗号分隔不同图片链接 不同图片没有no排序，C端按照逗号分隔后的顺序展示 
    */
    SpuLongPictures string `json:"spu_longPictures"`
}
type BatchUpdateSpubyIdResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
}



func (req *BatchUpdateSpubyIdRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*BatchUpdateSpubyIdResponse, error) {
    resp, err := client.InvokeApi(batch_update_spuby_id_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response BatchUpdateSpubyIdResponse
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

func (response *BatchUpdateSpubyIdResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

