package api

import (
	"encoding/json"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type CyTableAPi struct {
	BaseApi
}

func (p CyTableAPi) AddOrder(c *gin.Context) {
	goodsJsonStr := c.PostForm("goodsJsonStr")
	var goods []*proto.BeeOrderGoods
	if err := json.Unmarshal([]byte(goodsJsonStr), &goods); err != nil {
		p.Res(c, nil, err)
		return
	}
	err := service.GetCyTableSrv().AddOrder(c, c.ClientIP(), goods)
	p.Res(c, "success", err)
}

// Token 分配虚拟用户跟token
func (p CyTableAPi) Token(c *gin.Context) {
	//● tableId 桌码的id，不是名称也不是code，是id
	tableId := cast.ToInt64(c.PostForm("id"))
	//● key 桌子密钥
	key := c.PostForm("k")
	if tableId == 0 || key == "" {
		p.Res(c, "", enum.ErrParamError)
		return
	}
	resp, err := service.GetCyTableSrv().Token(c, tableId, key)
	p.Res(c, resp, err)
}
