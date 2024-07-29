package foodcatupdate

/**
* 创建或更新菜品分类
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_cat_update_url = "/wmoper/ng/foodop/foodCat/update"

type FoodCatUpdateRequest struct {
    /**
    *  原始的菜品分类名称，更新时必须 
    */
    CategoryNameOrigin string `json:"category_name_origin"`
    /**
    *  新建的菜品分类名称 
    */
    CategoryName string `json:"category_name"`
    /**
    *  菜品分类排序，数字越小，排名越靠前，新建时必须 
    */
    Sequence int32 `json:"sequence"`
    /**
    *  分组描述，默认空。最多255字，空代表不做任何修改。-1代表清空 
    */
    CategoryDescription string `json:"category_description"`
    /**
    *  分组信息。-1：不修改分类信息，0：普通分类；1：必选分类；2：可单独结算分类 
    */
    CategoryMode int32 `json:"category_mode"`
    /**
    *  置顶开关。1-不修改该信息；0 关闭； 1 开启 
    */
    TopFlag int32 `json:"top_flag"`
    /**
    *  置顶时间，默认空。-1代表清空。 设置分组置顶时段，每个分组可从每周的周一至周日中选择时间，每天最多支持5个时段，时段之间不能有重合，时间跨度不能超过24h。 
    */
    TimeZone string `json:"time_zone"`
}
type FoodCatUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *FoodCatUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodCatUpdateResponse, error) {
    resp, err := client.InvokeApi(food_cat_update_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodCatUpdateResponse
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

func (response *FoodCatUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

