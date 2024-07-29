package queryrecommendkeywords

/**
* 查询门店的推荐关键词
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_recommend_keywords_url = "/ad/launch/queryRecommendKeywords"

type QueryRecommendKeywordsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * { "code": "OP_SUCCESS", "msg": "成功", "traceId": "-6055251569466582277", "data": [ { "shopId": 92397439, "wordPackage": [//品类词包 { "competitive": "低",//竞争程度 "id": 390, "hotLevel": 1,//热度 "keyword": "韩国料理",//词包名（新建的时候传的） "feature": "",//特征 "itemType": 0, "wordList": [//词包包含的关键词 "朝鲜菜", "韩餐", "韩朝菜", "韩朝料理", "韩国菜", "韩国料理", "韩料", "韩食" ] }, { "competitive": "高", "id": 430, "hotLevel": 6, "keyword": "韩式烤肉", "feature": "同行买词,高热度", "itemType": 0, "wordList": [ "韩国烤肉", "韩式烤肉" ] } ], "topSearchWordPackage": [//热搜词 { "competitive": "低", "id": 0, "hotLevel": 6, "feature": "门店热搜", "itemType": 3,//如果要投热搜词，传该字段，内容固定为3 "wordList": [ "烤肉", "烧烤", "美食", "韩国料理", "韩餐", "附近美食", "附近的美食", "炸鸡", "韩式料理", "餐厅", "煎肉", "日料自助餐", "八爪鱼", "章鱼", "小牛" ] } ], "recommendWord": [//推荐自提词 { "competitive": "高", "hotLevel": 6, "feature": "同行买词,门店热搜", "word": "烤肉",//新建的时候传该字段 "wordType": 0 }, { "competitive": "高", "hotLevel": 6, "feature": "同行买词,门店热搜", "word": "烧烤", "wordType": 0 }, { "competitive": "高", "hotLevel": 3, "feature": "同行买词", "word": "餐饮", "wordType": 0 } ], "locationWordPackage": [//地址词 { "competitive": "低", "id": 0, "hotLevel": 4, "keyword": "商圈词",//新建的时候传该字段 "feature": "高性价比", "itemType": 0, "wordList": [ "中东大市场" ] }, { "competitive": "低", "id": 0, "hotLevel": 4, "keyword": "商场词", "feature": "高性价比", "itemType": 0, "wordList": [ "东环不夜城" ] } ] } ] }
    */
    Data []RecomendKeywordDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type RecomendKeywordDTO struct {
}
type QueryRecommendKeywordsRequest struct {
    /**
    *  需要查询的门店列表，多个门店所属类目必须一致 
    */
    ShopList int64 `json:"shopList"`
}



func (req *QueryRecommendKeywordsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryRecommendKeywordsResponse, error) {
    resp, err := client.InvokeApi(query_recommend_keywords_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryRecommendKeywordsResponse
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

func (response *QueryRecommendKeywordsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

