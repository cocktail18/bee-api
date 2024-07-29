package querycard

/**
* 查询安心卡
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_card_url = "/waimai/ng/card/query"

type QueryCardData struct {
    AppId int64 `json:"app_id"`
    /**
    * 安心卡id
    */
    CardId int32 `json:"cardId"`
    /**
    * ERP门店id
    */
    EpoiId string `json:"epoiId"`
    /**
    * 外卖门店名
    */
    WmPoiName string `json:"wmPoiName"`
    /**
    * 商家是否已签署协议。0未签署，1已签署。如果商家未签署协议， 需先让商家在美团外卖商家端中签署协议后才可使用安心卡功能。签署路径：进入“美团外卖商家端-门店运营-外卖安心卡”后，同意协议即可
    */
    ProtocolStatus int32 `json:"protocolStatus"`
    /**
    * 安心卡落地页url
    */
    LandPageUrl string `json:"landPageUrl"`
    /**
    * 是否同步打印小票。0不同步，1同步。当syncOpenReceipt\u003d1， 需将cardDetailList随小票打印
    */
    SyncOpenReceipt int32 `json:"syncOpenReceipt"`
    /**
    * 审核状态。0正常，1被配置下线，2审核失败。当auditStatus\u003d1或2时， 商家安心卡将会下线，相应的门店海报、在线联系自动同步给用户、随小票打印、C端安心卡标签也会失效。
    */
    AuditStatus int32 `json:"auditStatus"`
    /**
    * 温馨话术。文案举例：“我们向您承诺：1、外卖食品安全皆可追踪，请您安心食用。2、疫情期间，建议您选择“无接触配送”服务。3、让我们一起同舟共济，共渡疫情。”
    */
    Tips string `json:"tips"`
    /**
    * 商家生成的安心卡上填写的时间，单位秒
    */
    CardTime int32 `json:"cardTime"`
    /**
    * 是否同步展示在“店铺首页-门店海报”第一位。0不同步，1同步
    */
    SyncPost int32 `json:"syncPost"`
    /**
    * 安心卡类型。0普通版，1升级版
    */
    Type int32 `json:"type"`
    /**
    * 是否通过在线联系自动同步给用户。0不同步，1同步
    */
    SyncIM int32 `json:"syncIM"`
    /**
    * 商家生成安心卡的时间
    */
    Ctime int64 `json:"ctime"`
    /**
    * 安心卡升级版的安心保障模块详情。当syncOpenReceipt\u003d1，即把安心卡信息同步打印小票时， 无须将此信息随小票打印
    */
    UpgradeDetailList []UpgradeDetailList `json:"upgradeDetailList"`
    /**
    * 安心卡内容详情。当syncOpenReceipt\u003d1，即把安心卡信息同步打印小票时， 必须将此信息随小票打印
    */
    CardDetailList []CardDetailList `json:"cardDetailList"`
    /**
    * 核酸信息。此字段会根据安心卡的具体策略，按区域开放给门店使用。
    */
    TestingItemList []TestingItemList `json:"testingItemList"`
    /**
    * 审核文案。若审核失败或被封禁会传该文案
    */
    AuditTip string `json:"auditTip"`
}
type CardDetailList struct {
    /**
    * 岗位
    */
    Job string `json:"job"`
    /**
    * 姓名
    */
    Name string `json:"name"`
    /**
    * 体温。格式为 xx.x，只能一位小数
    */
    Temperature string `json:"temperature"`
}
type QueryCardRequest struct {
}
type TestingItemList struct {
    /**
    * 核酸信息图片url
    */
    Url string `json:"url"`
}
type QueryCardResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryCardData `json:"data"`
    TraceId string `json:"traceId"`
}
type UpgradeDetailList struct {
    /**
    * 1全员体温检测正常；2全员佩戴口罩；3全员定时洗手；4店内每日消毒
    */
    Type int32 `json:"type"`
    /**
    * 图片地址
    */
    Url string `json:"url"`
    /**
    * 主标题
    */
    Desc string `json:"desc"`
    /**
    * 副标题
    */
    SubTitle string `json:"subTitle"`
}



func (req *QueryCardRequest) DoInvoke(client mtclient.MeituanClient) (*QueryCardResponse, error) {
    resp, err := client.InvokeApi(query_card_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response QueryCardResponse
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

func (response *QueryCardResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

