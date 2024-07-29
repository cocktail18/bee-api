package decrypt

/**
* 针对加密字段解密
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const decrypt_url = "/wmoper/ng/waimaiad/common/decrypt"

type DecryptResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    Result Result `json:"result"`
    Value string `json:"value"`
}
type DecryptRequest struct {
    /**
    * 加密后的字符串，英文逗号分隔，最多传入10个密文
    */
    EncryptedValue string `json:"encryptedValue"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *DecryptRequest) DoInvoke(client mtclient.MeituanClient) (*DecryptResponse, error) {
    resp, err := client.InvokeApi(decrypt_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response DecryptResponse
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

func (response *DecryptResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

