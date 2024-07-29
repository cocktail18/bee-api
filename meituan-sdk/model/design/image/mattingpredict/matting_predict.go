package mattingpredict

/**
* 智能抠图服务 可指定对菜品或商品进行智能抠图，并返回指定物品结果图片
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const matting_predict_url = "/design/image/mattingPredict"

type MattingPredictResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data MattingPredictData `json:"data"`
    TraceId string `json:"traceId"`
}
type MattingPredictData struct {
    /**
    * 1:菜品抠图 2:商品抠图
    */
    Type int32 `json:"type"`
    /**
    * 图片二进制 Base64编码
    */
    Content string `json:"content"`
}
type MattingPredictRequest struct {
    /**
    * 图片二进制流，使用Base64编码 图片格式：JPG(JPEG)，PNG 图片像素尺寸：最小48*48像素，最大4096*4096像素 图片文件大小：最大 5MB
    */
    Content string `json:"content"`
    /**
    * 1:菜品抠图 2:商品抠图
    */
    Type int32 `json:"type"`
}



func (req *MattingPredictRequest) DoInvoke(client mtclient.MeituanClient) (*MattingPredictResponse, error) {
    resp, err := client.InvokeApi(matting_predict_url, 24, "", req)

    if err != nil {
        return nil, err
    }
    var response MattingPredictResponse
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

func (response *MattingPredictResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

