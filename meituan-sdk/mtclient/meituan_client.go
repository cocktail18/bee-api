package mtclient

type MeituanClient interface {
	InvokeApi(apiPath string, businessId int, appAuthToken string, data interface{}) ([]byte, error)

	GetToken(businessId int, code string) ([]byte, error)

	RefreshToken(businessId int, refreshToken string) ([]byte, error)
}
