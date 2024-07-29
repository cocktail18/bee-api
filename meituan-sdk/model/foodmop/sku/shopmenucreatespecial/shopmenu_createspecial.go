package shopmenucreatespecial

/**
* 创建或更新门店特殊类目（非必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const shopmenu_createspecial_url = "/foodmop/sku/shopmenu/createSpecial"

type SellerCategoryBasicDTO struct {
    /**
    *  前台类目 id，即类目编码 
    */
    VendorSellerCategoryId string `json:"vendorSellerCategoryId"`
    /**
    *  类目名称 
    */
    CategoryName string `json:"categoryName"`
    /**
    *  类目排序 rank &gt; 0 
    */
    Rank int32 `json:"rank"`
    /**
    *  类目层级 1：一级类目 2：二级类目 
    */
    Level int32 `json:"level"`
    /**
    *  类目类型 10：普通类目 30：餐具类目 品牌默认填 10 即可 不为空 餐具类目只能放餐具商品 餐具类目只能有一个 餐具类目只能为一级类目 
    */
    Type int32 `json:"type"`
    /**
    *  永久置顶开关 false：关闭 true：开启 top 是永久置顶，topTime是分时置顶，两个可以单独配置 top 优先级高，top = true 时即永久置顶，top = false 是根据 topTime 判定类目是否置顶 
    */
    Top bool `json:"top"`
    /**
    *  分时置顶时间 
    */
    TopTime []TimeDTO `json:"topTime"`
    /**
    *  分时展示时间 
    */
    ShowTime []TimeDTO `json:"showTime"`
    /**
    *  类目描述 
    */
    Description string `json:"description"`
    /**
    *  类目标签 RECOMMEND(-95, "本周推荐"), SPECIAL_OFFER(-94, "本店特供") 星巴克品牌定制，其他品牌不填 
    */
    TagList []int32 `json:"tagList"`
    /**
    *  父类目 id level == 2 时必填 parentVendorSellerCategoryId != vendorSellerCategoryId，parentVendorSellerCategoryId 必须存在于一级类目 id 中，一级类目没有父类目 id 
    */
    ParentVendorSellerCategoryId string `json:"parentVendorSellerCategoryId"`
    Hide bool `json:"hide"`
}
type TimeDTO struct {
    /**
    *  一周的第几天，eg： 1：星期一 2：星期二 ...以此类推 不为空，只能为 1~7 
    */
    DayOfWeek int32 `json:"dayOfWeek"`
    /**
    *  时间段 
    */
    Range []TimeRangeDTO `json:"range"`
}
type TimeRangeDTO struct {
    /**
    *  开始时间 不能为空，只能为 xx:xx 格式 
    */
    Start string `json:"start"`
    /**
    *  结束时间 不能为空，只能为 xx:xx 格式 
    */
    End string `json:"end"`
}
type VendorShopProductSimpleDTO struct {
    /**
    *  商品id 
    */
    VendorProductId string `json:"vendorProductId"`
    /**
    *  商品门店上下架状态 0 ：下架 1：上架 如果updateShelfStatus=true 则不能为空，只能为 0 or 1 
    */
    ShelfStatus int32 `json:"shelfStatus"`
    /**
    *  商品形式： 1：单品 2：套餐 
    */
    VendorProductForm int32 `json:"vendorProductForm"`
}
type ShopmenuCreatespecialRequest struct {
    /**
    *  门店菜谱 
    */
    VendorShopMenu VendorShopMenuDTO `json:"vendorShopMenu"`
    /**
    *  同步模式 1：全量覆盖模式 2：增量模式 只针对特殊类目，特殊类目不影响正常类目 
    */
    Append int32 `json:"append"`
}
type VendorShopMenuDTO struct {
    /**
    *  门店Id 
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    *  门店前台类目列表 最多 50 个 
    */
    VendorSellerCategoryList []VendorSellerCategoryDTO `json:"vendorSellerCategoryList"`
}
type VendorSellerCategoryDTO struct {
    /**
    *  前台类目基本信息 
    */
    SellerCategoryBasic SellerCategoryBasicDTO `json:"sellerCategoryBasic"`
    /**
    *  商品 SPU 列表，必存在于商品池中 SellerCategoryBasicDTO . level == 2 时必填 最多 50 个 
    */
    VendorShopProductSimpleList []VendorShopProductSimpleDTO `json:"vendorShopProductSimpleList"`
}
type ShopmenuCreatespecialResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 成功：true，失败：false
    */
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *ShopmenuCreatespecialRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ShopmenuCreatespecialResponse, error) {
    resp, err := client.InvokeApi(shopmenu_createspecial_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ShopmenuCreatespecialResponse
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

func (response *ShopmenuCreatespecialResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

