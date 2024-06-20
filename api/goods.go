package api

import (
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type GoodsApi struct {
	BaseApi
}

func (api GoodsApi) List(c *gin.Context) {
	categoryId := cast.ToInt64(c.PostForm("categoryId"))
	page := cast.ToInt(c.PostForm("page"))
	pageSize := cast.ToInt(c.PostForm("pageSize"))
	list, err := service.GetGoodsSrv().GetGoodsList(categoryId, page, pageSize)
	if err != nil {
		api.Fail(c, enum.ResCodeFail, err.Error())
		return
	}
	api.Success(c, list)
}

func (api GoodsApi) Detail(c *gin.Context) {
	id := cast.ToInt64(c.Query("id"))
	detail, err := service.GetGoodsSrv().GetGoodsDetail(c, id, "")
	if err != nil {
		api.Fail(c, enum.ResCodeFail, err.Error())
		return
	}
	api.Success(c, detail)
}

func (api GoodsApi) GoodsAddition(c *gin.Context) {
	goodsId := cast.ToInt64(c.Query("goodsId"))
	res, err := service.GetGoodsSrv().GetGoodsAddition(c, goodsId)
	api.Res(c, res, err)
}

func (api GoodsApi) Price(c *gin.Context) {
	goodsId := cast.ToInt64(c.PostForm("goodsId"))
	propertyChildIds := c.PostForm("propertyChildIds")
	resp, err := service.GetGoodsSrv().GetPrice(c, goodsId, propertyChildIds)
	api.Res(c, resp, err)
}

// TimesSchedule 获取商品时段定价日历数据
func (api GoodsApi) TimesSchedule(c *gin.Context) {
	//@todo
	api.Res(c, nil, enum.ErrEmpty)
}

func (api BaseApi) CategoryAll(c *gin.Context) {
	list, err := service.GetGoodsSrv().GetCategoryAll()
	if err != nil {
		api.Fail(c, enum.ResCodeFail, err.Error())
		return
	}
	api.Success(c, list)
}
