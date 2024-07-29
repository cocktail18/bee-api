package service

import (
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
)

type NoticeSrv struct {
	BaseSrv
}

var noticeSrvOnce sync.Once
var noticeSrvInstance *NoticeSrv

func GetNoticeSrv() *NoticeSrv {
	noticeSrvOnce.Do(func() {
		noticeSrvInstance = &NoticeSrv{}
	})
	return noticeSrvInstance
}

func (srv *NoticeSrv) GetLastOne() (*model.BeeNotice, error) {
	var noticeInfo model.BeeNotice
	err := db.GetDB().Where("is_show=?", 1).Order("date_update desc").Take(&noticeInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &noticeInfo, err
}

func (srv *NoticeSrv) GetDetail(id int64) (*model.BeeNotice, error) {
	var noticeInfo model.BeeNotice
	err := db.GetDB().Where("id=?", id).Take(&noticeInfo).Error
	return &noticeInfo, err
}
