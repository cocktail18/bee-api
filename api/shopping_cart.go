package api

import (
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

type ShoppingCartAPi struct {
	BaseApi
}

func (api ShoppingCartAPi) Info(c *gin.Context) {
	uid := kit.GetUid(c)
	resp, err := service.GetShoppingCartSrv().GetShoppingCart(c, uid)
	api.Res(c, resp, err)
}

func (api ShoppingCartAPi) Add(c *gin.Context) {
	uid := kit.GetUid(c)
	goodsId := cast.ToInt64(c.PostForm("goodsId"))
	number := cast.ToInt64(c.PostForm("number"))
	sku := c.PostForm("sku")
	//[{"optionId":142479,"optionValueId":1205135},{"optionId":142480,"optionValueId":1205137},{"optionId":142481,"optionValueId":1205139}]
	addition := c.PostForm("addition")
	logrus.Infof("添加购物车，%v %v %v %v %v", uid, goodsId, number, sku, addition)

	resp, err := service.GetShoppingCartSrv().Add(c, goodsId, number, sku, addition)
	api.Res(c, resp, err)
}

func (api ShoppingCartAPi) ModifyNumber(c *gin.Context) {
	uid := kit.GetUid(c)
	key := c.PostForm("key")
	number := cast.ToInt64(c.PostForm("number"))
	resp, err := service.GetShoppingCartSrv().ModifyNumber(c, uid, key, number)
	api.Res(c, resp, err)
}

func (api ShoppingCartAPi) Remove(c *gin.Context) {
	key := c.PostForm("key")
	resp, err := service.GetShoppingCartSrv().Remove(c, key)
	api.Res(c, resp, err)
}

func (api ShoppingCartAPi) Empty(c *gin.Context) {
	resp, err := service.GetShoppingCartSrv().Empty(c)
	api.Res(c, resp, err)
}
