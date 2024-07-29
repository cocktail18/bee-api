package getcustomlabels

/**
* 获取自定义标签
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_custom_labels_url = "/wmoper/ng/valueadded/getCustomLabels"

type GetCustomLabelsRequest struct {
    /**
    *  版本号。 1.0老数据，2.0新数据 
    */
    Version string `json:"version"`
}
type LabelValue struct {
    /**
    * 特征值选项中文名，显示用
    */
    Label string `json:"label"`
    /**
    * 描述
    */
    Desc string `json:"desc"`
    /**
    * 取值，提交时需要提交
    */
    Value string `json:"value"`
}
type LableResponse struct {
    /**
    * 特征中文名，显示用
    */
    Name string `json:"name"`
    /**
    * 特征英文名，提交时需要提交特征英文名
    */
    Field string `json:"field"`
    /**
    * 特征类型：1 单值，2 多值 ,3 文本。如果是1，则提交对应特征值时只能提交一个，如果是2，则可以提交多个值，如果是3，则可以自己填写值。
    */
    Type string `json:"type"`
    /**
    * 是否必填，true必填，false非必填
    */
    Required string `json:"required"`
    /**
    * 当特征类型为1和2的时候，提供的供选择的特征值json数组，具体说明如下
    */
    LabelValue LabelValue `json:"labelValue"`
    /**
    * 这个特征的描述说明，包括它的提填写规则说明。
    */
    Desc string `json:"desc"`
}
type GetCustomLabelsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 自定义标签列表，每个列表元素是一个特征模版json,json里的属性对应不同的特征。示例如下：
    */
    Data []LableResponse `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *GetCustomLabelsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetCustomLabelsResponse, error) {
    resp, err := client.InvokeApi(get_custom_labels_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetCustomLabelsResponse
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

func (response *GetCustomLabelsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

