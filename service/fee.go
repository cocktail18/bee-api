package service

import (
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/gin-gonic/gin"
	"sync"
)

type FeeSrv struct {
	BaseSrv
}

var feeSrvOnce sync.Once
var feeSrvInstance *FeeSrv

func GetFeeSrv() *FeeSrv {
	feeSrvOnce.Do(func() {
		feeSrvInstance = &FeeSrv{}
	})
	return feeSrvInstance
}

func (fee *FeeSrv) ListPeiSong(c *gin.Context) ([]*model.BeePeiSong, error) {
	list := make([]*model.BeePeiSong, 0)
	err := db.GetDB().Where("user_id = ? and is_deleted = 0", kit.GetUserId(c)).Find(&list).Error
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, enum.ErrEmpty
	}
	return list, nil
}
