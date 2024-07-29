package newimagefooddet

/**
* 菜品主体检测
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const new_image_food_det_url = "/design/image/newFoodDet"

type NewImageFoodDetRequest struct {
    /**
    *  图片二进制流，使用Base64编码 图片格式：JPG(JPEG)，PNG 图片像素尺寸：最小48*48像素，最大4096*4096像素 图片文件大小：最大 10MB 
    */
    Content string `json:"content"`
}
type FoodDetectResult struct {
    Detect string `json:"detect"`
}
type NewImageFoodDetResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data FoodDetectResult `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *NewImageFoodDetRequest) DoInvoke(client mtclient.MeituanClient) (*NewImageFoodDetResponse, error) {
    resp, err := client.InvokeApi(new_image_food_det_url, 24, "", req)

    if err != nil {
        return nil, err
    }
    var response NewImageFoodDetResponse
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

func (response *NewImageFoodDetResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

