package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeePeiSongSearch struct{
    Fwf1Type  *int `json:"fwf1Type" form:"fwf1Type" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
