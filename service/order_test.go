package service

import (
	"gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderSrv_Create(t *testing.T) {
	config.InitConfig()
	ctx := GetTestContext()
	_, err := GetOrderSrv().Create(ctx, "127.0.0.1", &proto.CreateOrderReq{
		GoodsJsonStr: `[{"goodsId":1,"num":1,"skuId":1}]`,
	})
	assert.Equal(t, err, nil)
}
