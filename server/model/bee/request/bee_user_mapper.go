package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeUserMapperSearch struct{
    Uid  *int `json:"uid" form:"uid" `
    Source  *int `json:"source" form:"source" `
    OpenId  string `json:"openId" form:"openId" `
    request.PageInfo
}
