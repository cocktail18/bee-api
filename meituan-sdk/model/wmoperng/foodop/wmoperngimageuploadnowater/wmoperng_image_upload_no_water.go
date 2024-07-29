package wmoperngimageuploadnowater

/**
* 上传无水印图片
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoperng_image_upload_no_water_url = "/wmoper/ng/foodop/image/uploadNoWater"

type WmoperngImageUploadNoWaterRequest struct {
    /**
    *  图片名称 
    */
    ImgName string `json:"img_name"`
    /**
    *  图片BASE64字符串 
    */
    Img string `json:"img"`
    /**
    *  图片类型：1.海报，2.招牌，3.安心卡 
    */
    PicType int32 `json:"pic_type"`
}
type WmoperngImageUploadNoWaterData struct {
    /**
    * 图片code
    */
    PicCode string `json:"pic_code"`
    /**
    * 图片url
    */
    PicUrl string `json:"pic_url"`
}
type WmoperngImageUploadNoWaterResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data WmoperngImageUploadNoWaterData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *WmoperngImageUploadNoWaterRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperngImageUploadNoWaterResponse, error) {
    resp, err := client.InvokeApi(wmoperng_image_upload_no_water_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperngImageUploadNoWaterResponse
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

func (response *WmoperngImageUploadNoWaterResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

