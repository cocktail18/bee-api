package proto

type AuthorizeResp struct {
	Openid     string `json:"openid"`
	Token      string `json:"token"`
	Uid        int64  `json:"uid"`
	UserId     int64  `json:"userId"`
	SessionKey string `json:"sessionKey"`
	UserLevel  int64  `json:"userLevel"`
}
