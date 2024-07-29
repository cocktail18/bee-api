package wmoperimageupload

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

const wmoper_image_upload_url = "/wmoper/image/upload"

type WmoperImageUploadResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 菜品code
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperImageUploadRequest struct {
    /**
    *  图片名称。图片名称必须以.jpg或.png或.jpeg结尾 
    */
    ImageName string `json:"imageName"`
    /**
    *  base64字符串的文件内容。 
    */
    File string `json:"file"`
}



func (req *WmoperImageUploadRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperImageUploadResponse, error) {
    resp, err := client.InvokeApi(wmoper_image_upload_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperImageUploadResponse
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

func (response *WmoperImageUploadResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

