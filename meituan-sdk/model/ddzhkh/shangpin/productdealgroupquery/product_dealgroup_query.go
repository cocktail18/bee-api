package productdealgroupquery

/**
* 查询团单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const product_dealgroup_query_url = "/ddzhkh/shangpin/dealgroup/query"

type ProductDealgroupQueryRequest struct {
    /**
    *  团单id 
    */
    DealGroupId int32 `json:"dealGroupId"`
    /**
    *  来源，体检 : physicalExamination，亲子：parentChildFun，商场: shoppingMall，医疗医美：medicalBeauty，数码家电：digitalHouseholdAppliance，爱车：car 
    */
    Source string `json:"source"`
}
type ProductDealgroupQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    /**
    *  团单id
    */
    DealGroupId int32 `json:"dealGroupId"`
    /**
    * 团单名
    */
    ProductName string `json:"productName"`
    /**
    * 团购售价
    */
    RetailPrice float64 `json:"retailPrice"`
    /**
    * 团购原价
    */
    OriginalPrice float64 `json:"originalPrice"`
    /**
    * 售卖结束时间，yyyy-MM-dd
    */
    SaleEndDate string `json:"saleEndDate"`
    /**
    * 审核状态，1-未提交，2-审批中，3-审核通过，4-变更未提交，5-变更审批中，6-待商户确认，8-审核不通过
    */
    ApprovalStatus int32 `json:"approvalStatus"`
    /**
    * 发布状态，0-删除，1-未发布，2-已发布
    */
    PublishStatus int32 `json:"publishStatus"`
    /**
    * 上架状态，0-下架，1-上架
    */
    DisplayStatus int32 `json:"displayStatus"`
    /**
    * 库存
    */
    StockAmount int32 `json:"stockAmount"`
    /**
    * 驳回原因
    */
    RejectReason string `json:"rejectReason"`
    OpPoiIds []string `json:"opPoiIds"`
}



func (req *ProductDealgroupQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ProductDealgroupQueryResponse, error) {
    resp, err := client.InvokeApi(product_dealgroup_query_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ProductDealgroupQueryResponse
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

func (response *ProductDealgroupQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

