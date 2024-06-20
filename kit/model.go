package kit

import (
	"gitee.com/stuinfer/bee-api/common"
	"github.com/gin-gonic/gin"
	"time"
)

func GetInsertBaseModel(c *gin.Context) *common.BaseModel {
	return &common.BaseModel{
		UserId:     GetUserId(c),
		IsDeleted:  false,
		DateAdd:    common.JsonTime(time.Now()),
		DateUpdate: common.JsonTime(time.Now()),
	}
}

func GetInsertBaseModelWithUserId(userId int64) *common.BaseModel {
	return &common.BaseModel{
		UserId:     userId,
		IsDeleted:  false,
		DateAdd:    common.JsonTime(time.Now()),
		DateUpdate: common.JsonTime(time.Now()),
	}
}
