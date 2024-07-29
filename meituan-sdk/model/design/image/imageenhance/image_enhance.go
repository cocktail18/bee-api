package imageenhance

/**
* 图片智能美化
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const image_enhance_url = "/design/image/enhance"

type ImageEnhanceResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ImageEnhanceData `json:"data"`
    TraceId string `json:"traceId"`
}
type ImageEnhanceRequest struct {
    /**
    * 图片二进制流，使用Base64编码 图片格式：JPG(JPEG)，PNG 图片像素尺寸：最小48*48像素，最大4096*4096像素 图片文件大小：最大 5MB
    */
    Content string `json:"content"`
}
type ImageEnhanceData struct {
    /**
    * 图片二进制 Base64编码
    */
    Content string `json:"content"`
}



func (req *ImageEnhanceRequest) DoInvoke(client mtclient.MeituanClient) (*ImageEnhanceResponse, error) {
    resp, err := client.InvokeApi(image_enhance_url, 24, "", req)

    if err != nil {
        return nil, err
    }
    var response ImageEnhanceResponse
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

func (response *ImageEnhanceResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

