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

type SaveAddressReq struct {
	Id          int64    `json:"id" form:"id"`
	Address     *string  `json:"address" form:"address"`
	AreaStr     *string  `json:"areaStr" form:"areaStr"`
	CityId      *string  `json:"cityId" form:"cityId"`
	CityStr     *string  `json:"cityStr" form:"cityStr"`
	DistrictId  *string  `json:"districtId" form:"districtId"`
	IsDefault   *bool    `json:"isDefault" form:"isDefault"`
	Latitude    *float64 `json:"latitude" form:"latitude"`
	LinkMan     *string  `json:"linkMan" form:"linkMan"`
	Longitude   *float64 `json:"longitude" form:"longitude"`
	Mobile      *string  `json:"mobile" form:"mobile"`
	ProvinceId  *string  `json:"provinceId" form:"provinceId"`
	ProvinceStr *string  `json:"provinceStr" form:"provinceStr"`
	Status      *int64   `json:"status" form:"status"`
	Uid         *int64   `json:"uid" form:"uid"`
}
