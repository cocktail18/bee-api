package imagewatermarkremove

/**
* 通用图像去水印
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const image_watermark_remove_url = "/design/image/waterMarkRemove"

type ImageWatermarkRemoveRequest struct {
    /**
    * 图片二进制流，使用Base64编码 图片格式：JPG(JPEG)，PNG 图片像素尺寸：最小48*48像素，最大4096*4096像素 图片文件大小：最大 5MB
    */
    Content string `json:"content"`
    /**
    * 水印位置列表
    */
    Rects []Rects `json:"rects"`
}
type ImageWatermarkRemoveResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ImageWatermarkRemoveData `json:"data"`
    TraceId string `json:"traceId"`
}
type ImageWatermarkRemoveData struct {
    /**
    * 去水印后图片Base64编码二进制 
    */
    Content string `json:"content"`
}
type Rects struct {
    /**
    * x0, y0 是左上角坐标
    */
    X0 string `json:"x0"`
    /**
    * x0, y0 是左上角坐标
    */
    Y0 string `json:"y0"`
    /**
    * x1, y1 是右下角坐标
    */
    X1 string `json:"x1"`
    /**
    * x1, y1 是右下角坐标
    */
    Y1 string `json:"y1"`
}



func (req *ImageWatermarkRemoveRequest) DoInvoke(client mtclient.MeituanClient) (*ImageWatermarkRemoveResponse, error) {
    resp, err := client.InvokeApi(image_watermark_remove_url, 24, "", req)

    if err != nil {
        return nil, err
    }
    var response ImageWatermarkRemoveResponse
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

func (response *ImageWatermarkRemoveResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

