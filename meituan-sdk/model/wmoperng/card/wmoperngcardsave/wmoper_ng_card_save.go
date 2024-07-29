package wmoperngcardsave

/**
* 保存安心卡
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_ng_card_save_url = "/wmoper/ng/card/save"

type WmoperNgCardSaveResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperNgCardSaveRequest struct {
    /**
    *  安心卡内容详情：当syncIM=1 或者 syncOpenReceipt=1 或者 syncPost=1 时，安心卡内容详情的人员信息不能为空 
    */
    CardDetailList []CardDetailVO `json:"cardDetailList"`
    /**
    *  温馨话术。举例，我们向您承诺：1、外卖食品安全皆可追踪，请您安心食用。2、疫情期间，建议您选择“无接触配送”服务。3、让我们一起同舟共济，共渡疫情。 
    */
    Tips string `json:"tips"`
    /**
    *  是否同步展示在“店铺首页-门店海报”第一位。0不同步，1同步。商家选中此项时，安心卡保存成功后，会在门店海报中显示安心卡信息。当商家有特型海报，默认将安心卡保存成特型海报的第一位 
    */
    SyncPost int32 `json:"syncPost"`
    /**
    *  安心卡类型。0普通版，1升级版。当type=1 升级版时，upgradeDetailList为必填，不可为空 
    */
    Type int32 `json:"type"`
    /**
    *  是否通过在线联系自动同步给用户。0不同步，1同步。向商家展示时默认值必须为0 。建议商家选中此项时，提示商家只有在美团外卖商家端开通在线联系后才可勾选此项，否则安心卡信息会创建失败 
    */
    SyncIM int32 `json:"syncIM"`
    /**
    *  安心卡升级版的安心保障模块详情。当安心卡类型为普通版（type=0），此字段为[]。当安心卡类型为升级版（type=1），此字段为必填。 
    */
    UpgradeDetailList []UpgradeDetailVO `json:"upgradeDetailList"`
    /**
    *  是否同步打印小票。0不同步，1同步。服务商后续可根据此字段，判断是否需要将安心卡信息打印在小票中 
    */
    SyncOpenReceipt int32 `json:"syncOpenReceipt"`
    TestingItemList []TestingItemParam `json:"testingItemList"`
}
type TestingItemParam struct {
    /**
    *  安心保障上传图片url 
    */
    Url string `json:"url"`
    /**
    *  核酸检测日0点时间戳 
    */
    TestingTime int32 `json:"testingTime"`
    /**
    *  描述 
    */
    Desc string `json:"desc"`
}
type Data struct {
}
type UpgradeDetailVO struct {
    /**
    *  1全员体温检测正常；2全员佩戴口罩；3全员定时洗手；4店内每日消毒。当type=1时，desc必为 全员体温检测正常；当type=2时，desc必为 全员佩戴口罩；当type=3时，desc必为 全员定时洗手；当type=4时，desc必为 店内每日消毒；其他取值type和desc 不限定；type取值不可重复 
    */
    Type int32 `json:"type"`
    /**
    *  主标题 
    */
    Desc string `json:"desc"`
    /**
    *  图片地址，参见 7.6.12 上传安心卡图片 
    */
    Url string `json:"url"`
    /**
    *  副标题 
    */
    SubTitle string `json:"subTitle"`
}
type CardDetailVO struct {
    /**
    *  岗位 
    */
    Job string `json:"job"`
    /**
    *  姓名 
    */
    Name string `json:"name"`
    /**
    *  体温，格式为 xx.x，只能一位小数 
    */
    Temperature string `json:"temperature"`
}



func (req *WmoperNgCardSaveRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperNgCardSaveResponse, error) {
    resp, err := client.InvokeApi(wmoper_ng_card_save_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperNgCardSaveResponse
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

func (response *WmoperNgCardSaveResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

