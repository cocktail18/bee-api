package queryshopmenu

/**
* 查询门店菜谱
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_shop_menu_url = "/foodmop/sku/queryShopMenu"

type ResultData struct {
    /**
    * 门店Id
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    * 门店前台类目列表
    */
    VendorSellerCategoryList []VendorSellerCategoryDTO `json:"vendorSellerCategoryList"`
}
type SellerCategoryBasicDTO struct {
    /**
    * 前台类目 id，即类目编码
    */
    VendorSellerCategoryId string `json:"vendorSellerCategoryId"`
    /**
    * 类目名称
    */
    CategoryName string `json:"categoryName"`
    /**
    * 类目排序 rank > 0
    */
    Rank int32 `json:"rank"`
    /**
    * 类目层级 1：一级类目 2：二级类目
    */
    Level int32 `json:"level"`
    /**
    * 类目类型 10：普通类目 30：餐具类目 品牌默认 10 
    */
    Type int32 `json:"type"`
    /**
    * 永久置顶开关 false：关闭 true：开启
    */
    Top bool `json:"top"`
    /**
    * 分时置顶时间
    */
    TopTime []TimeDTO `json:"topTime"`
    /**
    * 分时展示时间
    */
    ShowTime []TimeDTO `json:"showTime"`
    /**
    * 类目描述
    */
    Description string `json:"description"`
    /**
    * 配料标签，该字段未实际投入使用
    */
    TagList []int32 `json:"tagList"`
    /**
    * 父类目 id
    */
    ParentVendorSellerCategoryId string `json:"parentVendorSellerCategoryId"`
    /**
    * 类目是否在页面隐藏
    */
    Hide bool `json:"hide"`
}
type TimeDTO struct {
    /**
    * 一周的第几天，eg： 1：星期一 2：星期二 ...以此类推
    */
    DayOfWeek int32 `json:"dayOfWeek"`
    /**
    * 时间段
    */
    Range []TimeRangeDTO `json:"range"`
}
type TimeRangeDTO struct {
    /**
    * 开始时间 不能为空，只能为 xx:xx 格式
    */
    Start string `json:"start"`
    /**
    * 结束时间 不能为空，只能为 xx:xx 格式
    */
    End string `json:"end"`
}
type QueryShopMenuRequest struct {
    /**
    *  厂商门店id 
    */
    VendorShopId string `json:"vendorShopId"`
}
type VendorShopProductSimpleDTO struct {
    /**
    * 商品 id
    */
    VendorProductId string `json:"vendorProductId"`
    /**
    * 商品门店上下架状态 0 ：下架 1：上架
    */
    ShelfStatus int32 `json:"shelfStatus"`
    /**
    * 商品形式： 1：单品 2：套餐
    */
    VendorProductForm int32 `json:"vendorProductForm"`
}
type VendorSellerCategoryDTO struct {
    /**
    * 前台类目基本信息
    */
    SellerCategoryBasic SellerCategoryBasicDTO `json:"sellerCategoryBasic"`
    /**
    * 商品 SPU 列表，必存在于商品池中
    */
    VendorShopProductSimpleList []VendorShopProductSimpleDTO `json:"vendorShopProductSimpleList"`
}
type QueryShopMenuResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryShopMenuRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryShopMenuResponse, error) {
    resp, err := client.InvokeApi(query_shop_menu_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryShopMenuResponse
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

func (response *QueryShopMenuResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

