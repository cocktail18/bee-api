package proto

import "gitee.com/stuinfer/bee-api/model"

type ShopInfoResp struct {
	Info    *model.BeeShopInfo     `json:"info"`
	ExtJson map[string]interface{} `json:"extJson"`
}

type ListShopReq struct {
	PageReqCommon
	NameLike  string  `json:"nameLike" form:"nameLike"`
	Latitude  float64 `json:"curlatitude" form:"curlatitude"`
	Longitude float64 `json:"curlongitude" form:"curlongitude"`
}

type ListShopResp []*model.BeeShopInfo
