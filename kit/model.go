package kit

import (
	"context"
	"gitee.com/stuinfer/bee-api/common"
	"time"
)

func GetInsertBaseModel(c context.Context) *common.BaseModel {
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
