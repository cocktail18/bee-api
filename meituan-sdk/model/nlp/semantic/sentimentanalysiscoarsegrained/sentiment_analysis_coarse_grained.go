package sentimentanalysiscoarsegrained

/**
* 粗粒度情感分析
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const sentiment_analysis_coarse_grained_url = "/nlp/semantic/api/sentiment_analysis/coarse_grained"

type SentimentAnalysisCoarseGrainedResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * sentimentResultMap:情感分析的结果，statusCode：返回状态，errorMessage：状态错误时返回的信息
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type SentimentAnalysisCoarseGrainedRequest struct {
    /**
    *  输入的文本，通过情感分析模型得到文本对应的情感倾向（输入文本长度最好不超过512字） 
    */
    Text string `json:"text"`
}



func (req *SentimentAnalysisCoarseGrainedRequest) DoInvoke(client mtclient.MeituanClient) (*SentimentAnalysisCoarseGrainedResponse, error) {
    resp, err := client.InvokeApi(sentiment_analysis_coarse_grained_url, 42, "", req)

    if err != nil {
        return nil, err
    }
    var response SentimentAnalysisCoarseGrainedResponse
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

func (response *SentimentAnalysisCoarseGrainedResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

