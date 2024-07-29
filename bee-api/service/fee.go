package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
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

func (fee *FeeSrv) ListPeiSong(c context.Context) ([]*model.BeePeiSong, error) {
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

func (fee *FeeSrv) GetPeiSongByDistance(c context.Context, distance float64) (*model.BeePeiSong, error) {
	list := &model.BeePeiSong{}
	err := db.GetDB().Where("user_id = ? and is_deleted = 0 and distance < ?", kit.GetUserId(c), distance).Order("distance desc").First(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (fee *FeeSrv) GetPeiSongById(c context.Context, id int64) (*model.BeePeiSong, error) {
	list := &model.BeePeiSong{}
	err := db.GetDB().Where("user_id = ? and is_deleted = 0 and id = ?", kit.GetUserId(c), id).Take(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
