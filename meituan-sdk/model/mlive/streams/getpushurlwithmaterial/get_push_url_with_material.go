package getpushurlwithmaterial

/**
* 用直播物料获取推流地址
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_push_url_with_material_url = "/mlive/streams/getPushUrlWithLiveMaterial"

type ProductsSub struct {
    /**
    *  商品讲解预设的问题&amp;回答 generateType=2时，可不传；其它情况，必传 
    */
    Questions []QuestionsSub `json:"questions"`
    /**
    *  商品名称，必传 
    */
    ProductName string `json:"productName"`
    /**
    *  商品讲解台词内容，必传 
    */
    Content string `json:"content"`
}
type StreamParam struct {
    /**
    *  直播模组，0-正式直播，1-测试直播（美团直播助手创建测试直播，仅分享用户可见） 必传，建议传0 
    */
    LiveModel int64 `json:"liveModel"`
    /**
    *  非必传，用户设备id 
    */
    DeviceId string `json:"deviceId"`
}
type GetPushUrlWithMaterialResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []LiveStreamDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type QuestionsSub struct {
    /**
    *  商品问题 questions非空时，必传 
    */
    Question string `json:"question"`
    /**
    *  商品回答 questions非空时，必传 
    */
    Answer string `json:"answer"`
}
type LiveMateria struct {
    /**
    *  必传，生成类型，枚举如下： 1-模板，表示直播评论回复、口播等都是预先设置的，不能与用户交互； 2-开放式，表示都是实时生成的 3-表示枚举1、2的情况都有 
    */
    GenerateType int64 `json:"generateType"`
    /**
    *  非必传，数字人声音id （废弃字段，强烈建议改为使用voiceIds字段） 
    */
    VoiceId string `json:"voiceId"`
    /**
    *  必传，与观众互动的方式，枚举如下： 1-语音交互，仅口播回复 2-评论交互，仅评论回复 3-表示枚举1、2的情况都有 
    */
    InteractType int64 `json:"interactType"`
    /**
    *  必传，交互类型，表示是否与用户有互动，枚举如下： 1-交互式，有交互 2-非交互，无交互 
    */
    LiveType int64 `json:"liveType"`
    /**
    *  必传，预览图url，合成的数字人预览封面（包括形象、背景等配置），一般size=1 非内网域名线下提前周知 
    */
    PreviewLayouts []string `json:"previewLayouts"`
    /**
    *  必传，预览视频url，合成的数字人1分钟预览视频，一般size=1 要求mp4格式 非内网域名线下提前周知 
    */
    PreviewVideos []string `json:"previewVideos"`
    /**
    *  必传，商品列表 
    */
    Products []ProductsSub `json:"products"`
    /**
    *  非必传，数字人形象id（废弃字段，强烈建议改为使用characterIds字段） 
    */
    CharacterId string `json:"characterId"`
    /**
    *  必传，数字人形象id列表，单场直播使用的1~n个数字人形象 
    */
    CharacterIds []string `json:"characterIds"`
    /**
    *  必传，数字人声音id列表，单场直播使用的1~n个数字人声音 
    */
    VoiceIds []string `json:"voiceIds"`
}
type LiveStreamDTO struct {
    /**
    * 推流地址的过期时间
    */
    Expires int64 `json:"expires"`
    /**
    * 推流地址
    */
    PushUrl string `json:"pushUrl"`
}
type GetPushUrlWithMaterialRequest struct {
    /**
    *  必传，直播推流相关参数 
    */
    StreamParam StreamParam `json:"streamParam"`
    /**
    *  直播物料。 同一个liveid首次必传，重复获取推流地址时，非必传 
    */
    LiveMateria LiveMateria `json:"liveMateria"`
    /**
    *  必传，直播id 
    */
    LiveId int64 `json:"liveId"`
}



func (req *GetPushUrlWithMaterialRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetPushUrlWithMaterialResponse, error) {
    resp, err := client.InvokeApi(get_push_url_with_material_url, 50, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetPushUrlWithMaterialResponse
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

func (response *GetPushUrlWithMaterialResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

