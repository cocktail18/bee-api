package wmoperuploadcardimage

/**
* 上传安心卡图片
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_upload_card_image_url = "/wmoper/card/uploadCardImage"

type WmoperUploadCardImageResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data WmoperUploadCardImageData `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperUploadCardImageRequest struct {
    /**
    *  图片名称。图片名称必须以.jpg或.png或.jpeg结尾 
    */
    ImageName string `json:"imageName"`
    /**
    *  Base64格式图片内容。 
    */
    File string `json:"file"`
}
type WmoperUploadCardImageData struct {
    /**
    * 图片URL
    */
    PicUrl string `json:"pic_url"`
}



func (req *WmoperUploadCardImageRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperUploadCardImageResponse, error) {
    resp, err := client.InvokeApi(wmoper_upload_card_image_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperUploadCardImageResponse
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

func (response *WmoperUploadCardImageResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

