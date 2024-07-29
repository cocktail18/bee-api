package dishupdatecat

/**
* 新增／更新菜品分类
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_update_cat_url = "/waimai/dish/updateCat"

type DishUpdateCatRequest struct {
    /**
    *  原始的菜品分类名称，更新时为必填 
    */
    OldCatName string `json:"oldCatName"`
    /**
    *  菜品排序【数字越小，排名越靠前,不同分类顺序可以相同 】新建时必须 
    */
    Sequence int32 `json:"sequence"`
    /**
    *  菜品分类名称 
    */
    CatName string `json:"catName"`
    /**
    *  分类信息。-1：不修改分类信息，0：普通分类；1：必选分类；2：可单独结算分类 
    */
    CategoryMode int32 `json:"categoryMode"`
    /**
    *  分组描述，默认空,最多255字，空代表不做任何修改。-1代表清空 
    */
    CategoryDescription string `json:"categoryDescription"`
    /**
    *  置顶开关。-1：不修改该信息；0：关闭； 1：开启 
    */
    TopFlag int32 `json:"topFlag"`
    /**
    *  置顶时间，默认空。-1代表清空。设置分组置顶时段。每个分组可从每周的周一至周日中选择时间，每天最多支持5个时段，时段之间不能有重合，时间跨度不能超过24h。 
    */
    TimeZone string `json:"timeZone"`
}
type DishUpdateCatResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *DishUpdateCatRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishUpdateCatResponse, error) {
    resp, err := client.InvokeApi(dish_update_cat_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishUpdateCatResponse
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

func (response *DishUpdateCatResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

