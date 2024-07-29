package createorupdate

/**
* 创建/更新会员卡模板数据
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const create_or_update_url = "/jmcard/cards/create-or-update"

type CardTemplate struct {
    /**
    *  卡模板key，标记一套卡模板，必须确保多个卡模板key之间的唯一性，否则报错 
    */
    TemplateKey string `json:"templateKey"`
    /**
    *  卡模板基础信息 ​ 
    */
    CardTemplateBaseInfo CardTemplateBaseInfo `json:"cardTemplateBaseInfo"`
    /**
    *  卡模板会员权益 
    */
    MembershipCardBenefitInfo MembershipCardBenefitInfo `json:"membershipCardBenefitInfo"`
    /**
    *  卡模板栏位信息 
    */
    MembershipCardDetailInfo MembershipCardDetailInfo `json:"membershipCardDetailInfo"`
}
type CardBaseInfo struct {
    /**
    *  会员卡领卡按钮文案，4个字以内。 ​ 
    */
    Button string `json:"button"`
    /**
    *  会员注册信息设置 ​ 
    */
    RegisterInfos []RegisterStructure `json:"registerInfos"`
    /**
    *  卡有效期，可选：半年（HAFE_YEAR），一年（ONE_YEAR），两年（TWO_YEAR），长期有效（LONGTIME） 
    */
    ValidityPeriodType string `json:"validityPeriodType"`
}
type MembershipCardBenefitInfo struct {
    /**
    *  默认会员权益 ​ ​ 
    */
    DefaultBenefit DefaultBenefitStructure `json:"defaultBenefit"`
    /**
    *  新增会员权益 
    */
    NewBenefits []NewBenefitStructure `json:"newBenefits"`
}
type MembershipCardDetailInfo struct {
    /**
    *  默认中心栏位配置，中心栏位+列表栏位数量小于等于3 ​ 
    */
    DefaultBaseDetailList []DefaultMainEntrance `json:"defaultBaseDetailList"`
    /**
    *  必须为空list 
    */
    NewBaseDetailList []NewDetailStruct `json:"newBaseDetailList"`
    /**
    *  默认列表栏位， 默认列表栏位可选类型： "MEMBER_RIGHTS"：会员权益 "CAN_USE_POI"：适用门店 
    */
    DefaultListDetailList []string `json:"defaultListDetailList"`
    /**
    *  必须为空list 
    */
    NewListDetailList []NewDetailStruct `json:"newListDetailList"`
}
type DefaultMainEntrance struct {
    /**
    *  中心栏位类型， 可选： "POINT":积分 "BALANCE":余额 "COUPON":优惠券 
    */
    Type string `json:"type"`
}
type CardTemplateBaseInfo struct {
    /**
    *  会员卡模板展示名称，尽量短，超长截断展示，最长限制32个字符 
    */
    CardName string `json:"cardName"`
    /**
    *  logo展示url， 图片格式支持jpeg和png，尺寸大小： 100x100px 限制256个字符 
    */
    Logo string `json:"logo"`
    /**
    *  小图, 用户领卡后，在卡包展示会员卡信息，尺寸大小： 1440x276px 图片格式支持jpeg和png 限制128个字符 ​ 
    */
    SmallPic string `json:"smallPic"`
    /**
    *  中图。 领卡引导页使用（门店详情页等），尺寸大小： 1440x462px 图片格式支持jpeg和png 限制128个字符 
    */
    MiddlePic string `json:"middlePic"`
    /**
    *  大图， 领卡页使用，尺寸大小： 1440x810px 图片格式支持jpeg和png 限制128个字符 ​ 
    */
    BigPic string `json:"bigPic"`
    /**
    *  背景色。 控制领卡按钮颜色等 仅支持格式如：#ff6633 ​ 
    */
    BgColor string `json:"bgColor"`
    /**
    *  营销文案。限制32个字符以内 
    */
    PromotionText string `json:"promotionText"`
}
type CreateOrUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Resp `json:"data"`
    TraceId string `json:"traceId"`
}
type DefaultBenefitStructure struct {
    /**
    *  联系方式 
    */
    Contact MerchantContact `json:"contact"`
    /**
    *  会员卡详细介绍，限制512个字符 
    */
    DetailInfo string `json:"detailInfo"`
    /**
    *  会员卡使用须知，限制512个字符 
    */
    Instruction string `json:"instruction"`
}
type Resp struct {
    /**
    * 业务状态码
    */
    Status string `json:"status"`
    /**
    * 业务发生异常时的描述信息
    */
    ErrMsg string `json:"errMsg"`
}
type NewBenefitStructure struct {
    /**
    *  新增权益标题，限制16个字符 
    */
    Title string `json:"title"`
    /**
    *  新增权益内容，限制512个字符 
    */
    Content string `json:"content"`
}
type CreateOrUpdateRequest struct {
    /**
    *  会员卡基础信息 
    */
    CardBaseInfo CardBaseInfo `json:"cardBaseInfo"`
    /**
    *  卡模板信息，可传多个 
    */
    CardTemplateList []CardTemplate `json:"cardTemplateList"`
    /**
    *  默认卡模板key。每个卡模板都应该有一个key，需要指定一个卡模板为默认卡模板；该key必须能定位到一张卡模板，否则报错 
    */
    DefaultTemplateKey string `json:"defaultTemplateKey"`
}
type RegisterStructure struct {
    /**
    *  类型， 可选：姓名（NAME）、性别(SEX)、电话(PHONE)、生日(BIRTHDAY)、城市(CITY) 
    */
    Type string `json:"type"`
}
type NewDetailStruct struct {
}
type MerchantContact struct {
    /**
    *  类型，手机（MOBILE），座机（TEL） 
    */
    Type string `json:"type"`
    /**
    *  联系方式值， 手机号或者座机号 请填写真实的电话号码 
    */
    Value string `json:"value"`
}



func (req *CreateOrUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CreateOrUpdateResponse, error) {
    resp, err := client.InvokeApi(create_or_update_url, 15, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CreateOrUpdateResponse
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

func (response *CreateOrUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

