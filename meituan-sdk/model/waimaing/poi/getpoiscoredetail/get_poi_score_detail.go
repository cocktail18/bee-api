package getpoiscoredetail

/**
* 查询店铺分数据
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_poi_score_detail_url = "/waimai/ng/poi/getPoiScoreDetail"

type GetPoiScoreDetailResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GetPoiScoreDetailData `json:"data"`
    TraceId string `json:"traceId"`
}
type ExcellentList struct {
    /**
    * 文案name
    */
    ContentName string `json:"contentName"`
    /**
    * 当前指标描述
    */
    YourShopText string `json:"yourShopText"`
    /**
    * 门店当前指标
    */
    YourShop string `json:"yourShop"`
    /**
    * 分数
    */
    Score string `json:"score"`
    ExtraFields []ExtraFields `json:"extraFields"`
}
type RepairList struct {
    /**
    * 文案name
    */
    ContentName string `json:"contentName"`
    /**
    * 当前指标描述
    */
    YourShopText string `json:"yourShopText"`
    /**
    * 门店当前指标
    */
    YourShop string `json:"yourShop"`
    /**
    * 分数
    */
    Score string `json:"score"`
    ExtraFields []ExtraFields `json:"extraFields"`
}
type GetPoiScoreDetailRequest struct {
}
type GetPoiScoreDetailData struct {
    /**
    * 店铺总分数
    */
    TotalScore int32 `json:"totalScore"`
    /**
    * 待提升项
    */
    RepairList []RepairList `json:"repairList"`
    /**
    * 满分项
    */
    ExcellentList []ExcellentList `json:"excellentList"`
}
type ExtraFields struct {
    /**
    * 占比
    */
    HoverTextWeight string `json:"hoverTextWeight"`
    /**
    * hover文本描述
    */
    HoverText string `json:"hoverText"`
}



func (req *GetPoiScoreDetailRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetPoiScoreDetailResponse, error) {
    resp, err := client.InvokeApi(get_poi_score_detail_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetPoiScoreDetailResponse
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

func (response *GetPoiScoreDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

