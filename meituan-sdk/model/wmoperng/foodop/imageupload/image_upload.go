package imageupload

/**
* 上传菜品图片
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const image_upload_url = "/wmoper/ng/foodop/image/upload"

type ImageUploadRequest struct {
    /**
    *  图片名称 
    */
    ImgName string `json:"img_name"`
    /**
    *  图片BASE64字符串 
    */
    Img string `json:"img"`
}
type ImageUploadResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 图片ID
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *ImageUploadRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ImageUploadResponse, error) {
    resp, err := client.InvokeApi(image_upload_url, 16, appAuthToken, req)

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

