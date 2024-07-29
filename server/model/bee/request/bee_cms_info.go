package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeCmsInfoSearch struct{
    Title  string `json:"title" form:"title" `
    Type  *int `json:"type" form:"type" `
    request.PageInfo
}
