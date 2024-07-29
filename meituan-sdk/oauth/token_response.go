package oauth

type TokenResponse struct {
	Code    int
	Msg     string
	Data    TokenData
	TraceId string
}

type TokenData struct {
	AccessToken  string
	ExpireIn     int
	RefreshToken string
	Scope        string
	OpBizCode    string
}

func (response *TokenResponse) IsSuccess() bool {
	return response.Code == 0
}
