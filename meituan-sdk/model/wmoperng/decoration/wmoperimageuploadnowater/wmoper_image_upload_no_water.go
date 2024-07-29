package wmoperimageuploadnowater

/**
* 上传门店装修图片
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_image_upload_no_water_url = "/wmoper/image/uploadNoWater"

type WmoperImageUploadNoWaterData struct {
    /**
    * 图片地址
    */
    PicUrl string `json:"pic_url"`
    /**
    * 图片code
    */
    PicCode string `json:"pic_code"`
}
type WmoperImageUploadNoWaterResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data WmoperImageUploadNoWaterData `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperImageUploadNoWaterRequest struct {
    /**
    *  图片名称 文件名只能是字母或数字,且必须以.jpg结尾 
    */
    ImageName string `json:"imageName"`
    /**
    *  Base64格式的图片字符串。 
    */
    File string `json:"file"`
    /**
    *  图片类型。1海报，2招牌，3安心卡 
    */
    PicType int32 `json:"picType"`
}



func (req *WmoperImageUploadNoWaterRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperImageUploadNoWaterResponse, error) {
    resp, err := client.InvokeApi(wmoper_image_upload_no_water_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperImageUploadNoWaterResponse
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

func (response *WmoperImageUploadNoWaterResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

