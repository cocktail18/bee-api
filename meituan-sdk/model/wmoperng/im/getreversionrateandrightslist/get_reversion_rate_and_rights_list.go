package getreversionrateandrightslist

/**
* 回复率及权益查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_reversion_rate_and_rights_list_url = "/wmoper/ng/im/getReversionRateAndRightsList"

type GetReversionRateAndRightsListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type RightsList struct {
    Code int64 `json:"code"`
    Name string `json:"name"`
}
type Data struct {
    BenefitStatus int64 `json:"benefitStatus"`
    ReversionRate string `json:"reversionRate"`
    TargetReversionRate string `json:"targetReversionRate"`
    UrgentBenefitStatus int64 `json:"urgentBenefitStatus"`
    UrgentReversionRate string `json:"urgentReversionRate"`
    UrgentTargetReversionRate string `json:"urgentTargetReversionRate"`
    RightsList []RightsList `json:"rightsList"`
    TempRightsList []RightsList `json:"tempRightsList"`
}
type GetReversionRateAndRightsListRequest struct {
}



func (req *GetReversionRateAndRightsListRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetReversionRateAndRightsListResponse, error) {
    resp, err := client.InvokeApi(get_reversion_rate_and_rights_list_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetReversionRateAndRightsListResponse
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

func (response *GetReversionRateAndRightsListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

