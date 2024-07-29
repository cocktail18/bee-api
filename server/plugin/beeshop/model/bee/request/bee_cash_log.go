package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeCashLogSearch struct{
    Uid  *int `json:"uid" form:"uid" `
    request.PageInfo
}
