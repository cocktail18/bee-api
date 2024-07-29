package watermark

/**
* 通用水印/文字涂鸦识别接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const water_mark_url = "/design/image/waterMark"

type WaterMarkRequest struct {
    /**
    * 图片二进制流，使用Base64编码 图片格式：JPG(JPEG)，PNG 图片像素尺寸：最小48*48像素，最大4096*4096像素 图片文件大小：最大 5MB
    */
    Content string `json:"content"`
}
type WaterMarkResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 水印标签列表
    */
    Data []WaterMark `json:"data"`
    TraceId string `json:"traceId"`
}
type WaterMark struct {
    /**
    * 水印坐标
    */
    X0 int32 `json:"x0"`
    /**
    * 水印坐标
    */
    Y0 int32 `json:"y0"`
    /**
    * 水印坐标
    */
    X1 int32 `json:"x1"`
    /**
    * 水印坐标
    */
    Y1 int32 `json:"y1"`
    /**
    * 标签
    */
    Label string `json:"label"`
    /**
    * 置信度得分
    */
    Score string `json:"score"`
}



func (req *WaterMarkRequest) DoInvoke(client mtclient.MeituanClient) (*WaterMarkResponse, error) {
    resp, err := client.InvokeApi(water_mark_url, 24, "", req)

    if err != nil {
        return nil, err
    }
    var response WaterMarkResponse
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

func (response *WaterMarkResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

