package oauth

import (
	"encoding/json"
	"fmt"
	"meituan/gosdk/api_error"
	"meituan/gosdk/log"
	"meituan/gosdk/mtclient"
	"strconv"
)

type GetTokenRequest struct {
	BusinessId int    //api对应的业务id，请参考API文档的说明赋值
	Code       string //OAuth授权码，通过授权回调链接从美团获取
}

type RefreshTokenRequest struct {
	BusinessId   int    //api对应的业务id，请参考API文档的说明赋值
	RefreshToken string //获取token时同时获得的refreshToken
}

// 创建一个新的GetTokenRequest
func NewGetTokenRequest(businessId int, code string) GetTokenRequest {
	return GetTokenRequest{BusinessId: businessId, Code: code}
}

// 创建一个新的RefreshTokenRequest
func NewRefreshTokenRequest(businessId int, refreshToken string) RefreshTokenRequest {
	return RefreshTokenRequest{BusinessId: businessId, RefreshToken: refreshToken}
}

// GetToken 此方法用于通过授权码获取token
func (req GetTokenRequest) GetToken(client mtclient.MeituanClient) (*TokenResponse, error) {
	resp, err := client.GetToken(req.BusinessId, req.Code)
	log.Debug(fmt.Sprintf("getToken response=%s, err=%s", resp, err))
	if err != nil {
		return nil, err
	}
	return parseTokenResponse(resp)
}

// TokenRefresh 此方法用于用于刷新token
func (req RefreshTokenRequest) TokenRefresh(client mtclient.MeituanClient) (*TokenResponse, error) {
	resp, err := client.RefreshToken(req.BusinessId, req.RefreshToken)
	log.Debug(fmt.Sprintf("RefreshToken response=%s, err=%s", resp, err))
	if err != nil {
		return nil, err
	}
	return parseTokenResponse(resp)
}

// parseTokenResponse 解析响应并返回TokenResponse或错误。
func parseTokenResponse(resp []byte) (*TokenResponse, error) {
	var response TokenResponse
	err := json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	// 假设response.IsSuccess()检查的是response.Code == 0
	if response.IsSuccess() {
		return &response, nil
	}

	// 如果ApiError的Code字段是字符串类型，则需要转换
	return &response, &api_error.ApiError{
		Code:    strconv.Itoa(response.Code),
		Msg:     response.Msg,
		TraceId: response.TraceId,
	}
}
