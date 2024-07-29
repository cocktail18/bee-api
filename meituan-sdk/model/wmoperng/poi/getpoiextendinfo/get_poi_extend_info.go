package getpoiextendinfo

/**
* 查询门店二维码
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_poi_extend_info_url = "/wmoper/ng/poi/getPoiExtendInfo"

type GetPoiExtendInfoRequest struct {
}
type GetPoiExtendInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GetPoiExtendInfoData `json:"data"`
    TraceId string `json:"traceId"`
}
type GetPoiExtendInfoData struct {
    /**
    * 门店点菜页的H5链接
    */
    H5LinkUrl string `json:"h5LinkUrl"`
    /**
    * 用来生产二维码的门店点菜页链接（如要二维码需要业务自己生成）
    */
    QrCodeUrl string `json:"qrCodeUrl"`
    /**
    *  门店小程序二维码图片url
    */
    MiniProgramUrl string `json:"miniProgramUrl"`
}



func (req *GetPoiExtendInfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetPoiExtendInfoResponse, error) {
    resp, err := client.InvokeApi(get_poi_extend_info_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetPoiExtendInfoResponse
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

func (response *GetPoiExtendInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

