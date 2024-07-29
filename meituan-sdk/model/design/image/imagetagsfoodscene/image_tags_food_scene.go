package imagetagsfoodscene

/**
* 图像标签识别
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const image_tags_food_scene_url = "/design/image/tagsFoodScene"

type ImageTagsFoodSceneRequest struct {
    /**
    * 图片二进制流，使用Base64编码 图片格式：JPG(JPEG)，PNG 图片像素尺寸：最小48*48像素，最大4096*4096像素 图片文件大小：最大 5MB
    */
    Content string `json:"content"`
}
type ImageTagsFoodSceneData struct {
    /**
    * 菜品标签
    */
    Tags []Tags `json:"tags"`
}
type ImageTagsFoodSceneResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ImageTagsFoodSceneData `json:"data"`
    TraceId string `json:"traceId"`
}
type Tags struct {
    /**
    * 标签名称 [宣传图，美食，Logo，菜单，小票，其他，内部环境，外部环境]
    */
    Name string `json:"name"`
    /**
    * 置信度 范围0-1 0为非菜品图片，1为菜品图片，0~1区间表示为菜品的可信度。
    */
    Prob float64 `json:"prob"`
}



func (req *ImageTagsFoodSceneRequest) DoInvoke(client mtclient.MeituanClient) (*ImageTagsFoodSceneResponse, error) {
    resp, err := client.InvokeApi(image_tags_food_scene_url, 24, "", req)

    if err != nil {
        return nil, err
    }
    var response ImageTagsFoodSceneResponse
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

func (response *ImageTagsFoodSceneResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

