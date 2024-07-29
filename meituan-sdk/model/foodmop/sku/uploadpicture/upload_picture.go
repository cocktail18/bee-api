package uploadpicture

/**
* 图片上传（必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const upload_picture_url = "/foodmop/sku/uploadPicture"

type UploadPictureResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type UploadPictureRequest struct {
    /**
    *  图片url 仅支持 https 的图片 url 链接且不能包含中文 
    */
    PicUrl string `json:"picUrl"`
}



func (req *UploadPictureRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*UploadPictureResponse, error) {
    resp, err := client.InvokeApi(upload_picture_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response UploadPictureResponse
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

func (response *UploadPictureResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

