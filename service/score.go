package service

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
	"time"
)

type ScoreSrv struct {
	BaseSrv
}

var scoreSrvOnce sync.Once
var scoreSrvInstance *ScoreSrv

func GetScoreSrv() *ScoreSrv {
	scoreSrvOnce.Do(func() {
		scoreSrvInstance = &ScoreSrv{}
	})
	return scoreSrvInstance
}

func (srv *ScoreSrv) Sign(c *gin.Context) error {
	uid := kit.GetUid(c)
	//获取上一次签到记录
	continutes := 1
	last := &model.BeeSignLog{}
	err := db.GetDB().Where("uid=?", uid).Order("id desc").First(last).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
	} else if err != nil {
		return err
	} else {
		addDate := carbon.CreateFromStdTime(time.Time(last.DateAdd))
		if addDate.ToDateString() == carbon.Now().ToDateString() {
			continutes = last.Continuous + 1
		}
	}

	log := &model.BeeSignLog{
		BaseModel:  *kit.GetInsertBaseModel(c),
		DateAdd:    common.JsonTime(time.Now()),
		Uid:        uid,
		Continuous: continutes,
	}
	//@todo 签到奖励
	return db.GetDB().Create(log).Error
}

func (srv *ScoreSrv) GetSignLogs(userId int64) ([]*model.BeeSignLog, error) {
	var list []*model.BeeSignLog
	now := carbon.Now()
	monthBegin := now.StartOfMonth().ToDateTimeString()
	monthEnd := now.EndOfMonth().ToDateTimeString()
	err := db.GetDB().Where("uid=? and date_add >= ? and date_add <= ?", userId, monthBegin, monthEnd).Order("date_add asc").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (srv *ScoreSrv) sameDay(t1, t2 time.Time) bool {
	return t1.Format("2006-01-02") == t2.Format("2006-01-02")
}
