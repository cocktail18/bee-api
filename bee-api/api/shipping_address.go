package api

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type ShippingAddressApi struct {
	BaseApi
}

func (api ShippingAddressApi) List(c *gin.Context) {
	resp, err := service.GetAddressSrv().ListAddress(c)
	api.Res(c, resp, err)
}

func (api ShippingAddressApi) Add(c *gin.Context) {
	linkMan := c.PostForm("linkMan") // xxx
	address := c.PostForm("address") // 广东省广州市天河区天府路1号广州市天河区人民政府(天府路西)
	mobile := c.PostForm("mobile")
	isDefault := cast.ToBool(c.PostForm("isDefault"))
	latitude := cast.ToFloat64(c.PostForm("latitude"))
	longitude := cast.ToFloat64(c.PostForm("longitude"))
	provinceId := c.PostForm("provinceId")
	cityId := c.PostForm("cityId")
	districtId := c.PostForm("districtId")

	userAddress := &model.BeeUserAddress{
		Address:     address,
		AreaStr:     "",
		CityId:      cityId,
		CityStr:     "",
		DistrictId:  districtId,
		IsDefault:   isDefault,
		Latitude:    latitude,
		LinkMan:     linkMan,
		Longitude:   longitude,
		Mobile:      mobile,
		ProvinceId:  provinceId,
		ProvinceStr: "",
		Status:      0,
	}
	resp, err := service.GetAddressSrv().AddAddress(c, userAddress)
	api.Res(c, resp, err)
}

func (api ShippingAddressApi) Detail(c *gin.Context) {
	id := cast.ToInt64(c.Query("id"))
	resp, err := service.GetAddressSrv().GetAddress(c, id)
	api.Res(c, resp, err)
}

func (api ShippingAddressApi) Update(c *gin.Context) {
	linkMan := c.PostForm("linkMan") // xxx
	address := c.PostForm("address") // 广东省广州市天河区天府路1号广州市天河区人民政府(天府路西)
	mobile := c.PostForm("mobile")
	isDefault := cast.ToBool(c.PostForm("isDefault"))
	latitude := cast.ToFloat64(c.PostForm("latitude"))
	longitude := cast.ToFloat64(c.PostForm("longitude"))
	provinceId := c.PostForm("provinceId")
	cityId := c.PostForm("cityId")
	districtId := c.PostForm("districtId")
	id := cast.ToInt64(c.PostForm("id"))

	userAddress := &model.BeeUserAddress{
		BaseModel: common.BaseModel{
			Id:     id,
			UserId: kit.GetUserId(c),
		},
		Address:     address,
		AreaStr:     "",
		CityId:      cityId,
		CityStr:     "",
		DistrictId:  districtId,
		IsDefault:   isDefault,
		Latitude:    latitude,
		LinkMan:     linkMan,
		Longitude:   longitude,
		Mobile:      mobile,
		ProvinceId:  provinceId,
		ProvinceStr: "",
		Status:      0,
		Uid:         kit.GetUid(c),
	}
	_, err := service.GetAddressSrv().SaveAddress(c, userAddress)
	api.Res(c, "success", err)
}

func (api ShippingAddressApi) Delete(c *gin.Context) {
	id := cast.ToInt64(c.PostForm("id"))
	err := service.GetAddressSrv().DeleteAddress(kit.GetUid(c), id)
	api.Res(c, "success", err)
}

func (api ShippingAddressApi) Default(c *gin.Context) {
	res, err := service.GetAddressSrv().Default(c)
	api.Res(c, res, err)
}
