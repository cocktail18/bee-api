package imagefilter

/**
* 图片添加滤镜效果，对图片整体进行定制化色彩调整。
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const image_filter_url = "/design/image/filter"

type ImageFilterRequest struct {
    /**
    * 图片二进制流，使用Base64编码 图片格式：JPG(JPEG)，PNG 图片像素尺寸：最小48*48像素，最大4096*4096像素 图片文件大小：最大 5MB
    */
    Content string `json:"content"`
}
type ImageFilterResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ImageFilterData `json:"data"`
    TraceId string `json:"traceId"`
}
type ImageFilterData struct {
    /**
    * 图片二进制 Base64编码
    */
    Content string `json:"content"`
}



func (req *ImageFilterRequest) DoInvoke(client mtclient.MeituanClient) (*ImageFilterResponse, error) {
    resp, err := client.InvokeApi(image_filter_url, 24, "", req)

    if err != nil {
        return nil, err
    }
    var response ImageFilterResponse
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

func (response *ImageFilterResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

