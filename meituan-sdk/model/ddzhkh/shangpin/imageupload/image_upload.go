package imageupload

/**
* 图片上传
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const image_upload_url = "/ddzhkh/shangpin/image/upload"

type ImageUploadRequest struct {
    /**
    *  图片二进制流，使用Base64编码，最大不超过5M 
    */
    ImageContent string `json:"imageContent"`
}
type ImageUploadResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ImageInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type ImageInfo struct {
    /**
    * 请求成功图片的key
    */
    ImageKey string `json:"imageKey"`
}



func (req *ImageUploadRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ImageUploadResponse, error) {
    resp, err := client.InvokeApi(image_upload_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ImageUploadResponse
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

func (response *ImageUploadResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

