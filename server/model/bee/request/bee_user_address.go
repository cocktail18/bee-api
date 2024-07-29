package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeUserAddressSearch struct{
    Status  *int `json:"status" form:"status" `
    Uid  *int `json:"uid" form:"uid" `
    request.PageInfo
}
