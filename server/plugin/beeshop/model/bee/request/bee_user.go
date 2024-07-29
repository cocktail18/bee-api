package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeUserSearch struct{
    Id  *int `json:"id" form:"id" `
    ShowUid  *int `json:"showUid" form:"showUid" `
    CardNumber  string `json:"cardNumber" form:"cardNumber" `
    IsVirtual  *bool `json:"isVirtual" form:"isVirtual" `
    Status  *int `json:"status" form:"status" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
