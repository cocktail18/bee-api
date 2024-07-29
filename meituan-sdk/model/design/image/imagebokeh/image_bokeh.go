package imagebokeh

/**
* 图片虚景滤镜特效 对菜品等物体图片，添加虚景滤镜，对菜品等主要物体以外的背景进行虚化处理，突出菜品等主要物体本身。
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const image_bokeh_url = "/design/image/bokeh"

type ImageBokehResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ImageBokehData `json:"data"`
    TraceId string `json:"traceId"`
}
type ImageBokehData struct {
    /**
    * 图片二进制 Base64编码
    */
    Content string `json:"content"`
}
type ImageBokehRequest struct {
    /**
    * 图片二进制流，使用Base64编码 图片格式：JPG(JPEG)，PNG 图片像素尺寸：最小48*48像素，最大4096*4096像素 图片文件大小：最大 5MB
    */
    Content string `json:"content"`
}



func (req *ImageBokehRequest) DoInvoke(client mtclient.MeituanClient) (*ImageBokehResponse, error) {
    resp, err := client.InvokeApi(image_bokeh_url, 24, "", req)

    if err != nil {
        return nil, err
    }
    var response ImageBokehResponse
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

func (response *ImageBokehResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

