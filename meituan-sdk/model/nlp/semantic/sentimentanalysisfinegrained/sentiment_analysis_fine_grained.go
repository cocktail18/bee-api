package sentimentanalysisfinegrained

/**
* 细粒度情感分析
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const sentiment_analysis_fine_grained_url = "/nlp/semantic/api/sentiment_analysis/fine_grained"

type SentimentAnalysisFineGrainedRequest struct {
    /**
    *  输入的文本，通过细粒度情感分析之后，可以得到各个维度的情感。（输入的文本长度不超过512字) 
    */
    Text string `json:"text"`
}
type SentimentAnalysisFineGrainedResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * results是细粒度情感分析的结果，statusCode是返回状态，errorMessage未错误时返回的信息
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *SentimentAnalysisFineGrainedRequest) DoInvoke(client mtclient.MeituanClient) (*SentimentAnalysisFineGrainedResponse, error) {
    resp, err := client.InvokeApi(sentiment_analysis_fine_grained_url, 42, "", req)

    if err != nil {
        return nil, err
    }
    var response SentimentAnalysisFineGrainedResponse
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

func (response *SentimentAnalysisFineGrainedResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

