package service

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type CommentSrv struct {
	BaseSrv
}

var commentSrvOnce sync.Once
var commentSrvInstance *CommentSrv

func GetCommentSrv() *CommentSrv {
	commentSrvOnce.Do(func() {
		commentSrvInstance = &CommentSrv{}
	})
	return commentSrvInstance
}

func (s *CommentSrv) Add(c *gin.Context, req *proto.CommentAddReq) error {
	userId := kit.GetUserId(c)
	uid := kit.GetUid(c)
	now := time.Now()
	return db.GetDB().Create(&model.BeeComment{
		BaseModel: common.BaseModel{
			UserId:     userId,
			IsDeleted:  false,
			DateAdd:    common.JsonTime(now),
			DateUpdate: common.JsonTime(now),
		},
		Uid:        uid,
		RefId:      req.RefId,
		Pid:        req.Pid,
		ShopId:     req.ShopId,
		Type:       req.Type,
		Pics:       req.Pics,
		File:       req.File,
		Mobile:     req.Mobile,
		Content:    req.Content,
		ExtJsonStr: req.ExtJsonStr,
		Status:     enum.CommentStatusNone,
	}).Error
}
