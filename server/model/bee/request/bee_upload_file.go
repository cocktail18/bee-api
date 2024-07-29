package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BeeUploadFileSearch struct{
    StartDateAdd  *time.Time  `json:"startDateAdd" form:"startDateAdd"`
    EndDateAdd  *time.Time  `json:"endDateAdd" form:"endDateAdd"`
    StartDateUpdate  *time.Time  `json:"startDateUpdate" form:"startDateUpdate"`
    EndDateUpdate  *time.Time  `json:"endDateUpdate" form:"endDateUpdate"`
    Filename  string `json:"filename" form:"filename" `
    StartExpireAt  *int  `json:"startExpireAt" form:"startExpireAt"`
    EndExpireAt  *int  `json:"endExpireAt" form:"endExpireAt"`
    request.PageInfo
}
