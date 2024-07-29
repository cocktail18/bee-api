package trigger

/**
* 摩西机器人会话接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const trigger_url = "/moses/dialog/trigger"

type NluResult struct {
}
type TDialogAct struct {
    /**
    * 意图类型。intentClarify：意图澄清； faqTask：问答类型的意图； normalTask：正常Task
    */
    Act string `json:"act"`
}
type ReportData struct {
    /**
    * 意图ID列表
    */
    IntentIdList string `json:"intentIdList"`
    /**
    * 当前轮对话之后是否还处于Task执行中
    */
    IsInTask string `json:"isInTask"`
}
type TriggerData struct {
    /**
    * 版本
    */
    Version string `json:"version"`
    /**
    * 请求query
    */
    Query string `json:"query"`
    /**
    * 请求Id
    */
    RequestId string `json:"requestId"`
    /**
    * 包括意图信息，返回话术信息等，是一个列表
    */
    DmResult []DmResult `json:"dmResult"`
    /**
    * 上报数据节点
    */
    ReportData ReportData `json:"reportData"`
    /**
    * 扩展字段
    */
    ExtData ExtData2 `json:"extData"`
    /**
    * NLU识别结果
    */
    NluResult NluResult `json:"nluResult"`
}
type TriggerResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data TriggerData `json:"data"`
    TraceId string `json:"traceId"`
}
type ExtData2 struct {
    /**
    * 问答类型意图答案的图片列表
    */
    AnswerPicList string `json:"answerPicList"`
}
type ExtData struct {
    /**
    * 选项卡的答案
    */
    MosesTabAnswer string `json:"mosesTabAnswer"`
}
type TriggerRequest struct {
    /**
    *  版本，当前固定2.0.0 
    */
    Version string `json:"version"`
    /**
    *  请求文本内容 
    */
    Query string `json:"query"`
    /**
    *  消息id，接入方自己生成传入，保证 每次不一样，全局唯一 
    */
    MessageId string `json:"messageId"`
    /**
    *  会话id，根据机器人相关配置进行传入 
    */
    SessionId string `json:"sessionId"`
    /**
    *  消息类型，客服场景固定为TEXT 
    */
    MessageType string `json:"messageType"`
    /**
    *  当前固定为MOSES 
    */
    Source string `json:"source"`
    /**
    *  应用信息 
    */
    AppInfo AppInfo `json:"appInfo"`
    /**
    *  用户信息，json字符串形式 
    */
    UserInfo string `json:"userInfo"`
    /**
    *  json字符串形式，内部传入随路的 额外字段。 是否需要传入、需要传入哪些是与机器人配置相关，请与相关配置人员沟通 
    */
    DeviceContext string `json:"deviceContext"`
}
type DmResult struct {
    /**
    * 意图识别信息
    */
    TDialogAct TDialogAct `json:"tDialogAct"`
    /**
    * 回复类型 。text：普通文本； list：列表
    */
    SolutionType string `json:"solutionType"`
    /**
    * 回复结果。如果为空，表示执行失败了。
    */
    Solution string `json:"solution"`
    /**
    * spokenText和solution的值一般是一样的
    */
    SpokenText string `json:"spokenText"`
    /**
    * 答案的类型 。默认的类型为空；mosesTab：表示选项卡类型，需要用户选择，此时具体的选项信息在extData.mosesTabAnswer字段的json中
    */
    AnswerType string `json:"answerType"`
    /**
    * 扩展信息，目前选项卡的答案在该字段的mosesTabAnswer属性中
    */
    ExtData ExtData `json:"extData"`
}
type AppInfo struct {
    /**
    *  所配置机器人的key 
    */
    RobotKey string `json:"robotKey"`
}



func (req *TriggerRequest) DoInvoke(client mtclient.MeituanClient) (*TriggerResponse, error) {
    resp, err := client.InvokeApi(trigger_url, 28, "", req)

    if err != nil {
        return nil, err
    }
    var response TriggerResponse
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

func (response *TriggerResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

