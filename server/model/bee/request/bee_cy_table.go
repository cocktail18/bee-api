package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeCyTableSearch struct{
    ShopId  *int `json:"shopId" form:"shopId" `
    TableNum  string `json:"tableNum" form:"tableNum" `
    Remark  string `json:"remark" form:"remark" `
    Uid  *int `json:"uid" form:"uid" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
