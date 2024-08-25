package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"sync"
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

func (s *CommentSrv) Add(c context.Context, req *proto.CommentAddReq) error {
	uid := kit.GetUid(c)
	return db.GetDB().Create(&model.BeeComment{
		BaseModel:  *kit.GetInsertBaseModel(c),
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
