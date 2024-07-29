package recognizemenu

/**
* 菜单OCR识别 识别并结构化输出菜单图片中的菜名、价格和单位。
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const recognize_menu_url = "/design/ocr/recognizeMenu"

type RecognizeMenuRequest struct {
    /**
    *  图片二进制流，使用Base64编码 图片格式：JPG(JPEG)，PNG 图片像素尺寸：最小48*48像素，最大4096*4096像素 图片文件大小：最大 5MB 
    */
    Content string `json:"content"`
}
type RecognizeMenuResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data RecognizeMenuData `json:"data"`
    TraceId string `json:"traceId"`
}
type RecognizeMenuData struct {
    /**
    * 菜名数目
    */
    NDishNum int32 `json:"nDishNum"`
    /**
    * 菜单识别结果
    */
    Menus string `json:"menus"`
}



func (req *RecognizeMenuRequest) DoInvoke(client mtclient.MeituanClient) (*RecognizeMenuResponse, error) {
    resp, err := client.InvokeApi(recognize_menu_url, 24, "", req)

    if err != nil {
        return nil, err
    }
    var response RecognizeMenuResponse
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

func (response *RecognizeMenuResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

