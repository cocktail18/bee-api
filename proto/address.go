package proto

import "gitee.com/stuinfer/bee-api/model"

type UserAddressResp struct {
	model.BeeUserAddress
	RegionDistrict *model.BeeRegion `json:"regionDistrict"`
	RegionCity     *model.BeeRegion `json:"regionCity"`
	RegionProvince *model.BeeRegion `json:"regionProvince"`
}

type UserAddressDetailResp struct {
	Info    *UserAddressResp       `json:"info"`
	ExtJson map[string]interface{} `json:"extJson"`
}
