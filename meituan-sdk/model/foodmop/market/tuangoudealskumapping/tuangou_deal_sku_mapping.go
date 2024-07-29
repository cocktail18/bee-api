package tuangoudealskumapping

/**
* 团购券菜品关系同步（使用团券兑换功能-必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tuangou_deal_sku_mapping_url = "/foodmop/market/tuangou/deal/sku/mapping"

type Combo struct {
    /**
    *  商家套餐id，下单时会透传该字段 
    */
    VendorComboId string `json:"vendorComboId"`
    /**
    *  套餐分组列表，单次请求最大数数量限制10 
    */
    ComboGroupList []ComboGroup `json:"comboGroupList"`
    /**
    *  套餐名称 
    */
    ComboName string `json:"comboName"`
}
type TuangouDealSkuMappingResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}
type ComboItem struct {
    /**
    *  菜品排序 
    */
    Rank int32 `json:"rank"`
    /**
    *  商家skuId 
    */
    VendorSkuId string `json:"vendorSkuId"`
    /**
    *  商家spuId 
    */
    VendorSpuId string `json:"vendorSpuId"`
}
type TuangouDealSkuMappingRequest struct {
    /**
    *  团购券id 
    */
    DealId int32 `json:"dealId"`
    /**
    *  团购券映射类型，单品or套餐 10：团购映射单品 20：团购套餐 特殊说明： 团购券只兑换单sku，映射类型为单品 团购券兑换属于同一spu下多个sku中的一个，映射类型为套餐 团购券兑换单分组、多分组套餐，映射类型为套餐 
    */
    MappingType int32 `json:"mappingType"`
    /**
    *  厂商的skuid，mappingType为10-单品时非空 
    */
    VendorSkuId string `json:"vendorSkuId"`
    /**
    *  业务类型： 10：点餐 
    */
    Bizcode int32 `json:"bizcode"`
    /**
    *  映射关系作用的维度 10：门店 20：品牌 
    */
    Dimension int32 `json:"dimension"`
    /**
    *  商家门店id，dimension为10时必填 
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    *  套餐分组列表，mappingType为20-套餐时非空 
    */
    Combo Combo `json:"combo"`
}
type ComboGroup struct {
    /**
    *  套餐排序 
    */
    Rank int32 `json:"rank"`
    /**
    *  最大选择的商品数 
    */
    MaxNum int32 `json:"maxNum"`
    /**
    *  最小选择的商品数，需要等于最大选择的商品数 
    */
    MinNum int32 `json:"minNum"`
    /**
    *  套餐分组名称,不超过16个字符 
    */
    GroupName string `json:"groupName"`
    /**
    *  套餐分组下的菜品，单次请求最大数数量限制300 
    */
    ComboItemList []ComboItem `json:"comboItemList"`
}



func (req *TuangouDealSkuMappingRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TuangouDealSkuMappingResponse, error) {
    resp, err := client.InvokeApi(tuangou_deal_sku_mapping_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TuangouDealSkuMappingResponse
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

func (response *TuangouDealSkuMappingResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

