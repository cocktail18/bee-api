package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/golang-module/carbon/v2"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
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

func (srv *ScoreSrv) Sign(c context.Context) error {
	uid := kit.GetUid(c)
	//获取上一次签到记录
	continues := 1
	last := &model.BeeSignLog{}
	err := db.GetDB().Where("uid=?", uid).Order("id desc").First(last).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
	} else if err != nil {
		return err
	} else {
		addDate := carbon.CreateFromStdTime(time.Time(last.DateAdd))
		if addDate.ToDateString() == carbon.Now().ToDateString() {
			continues = last.Continues + 1
		}
	}
	score := int64(1) // @todo 配置积分规则
	if continues%7 == 0 {
		score = score + int64(continues/7)
	}
	log := &model.BeeSignLog{
		BaseModel: *kit.GetInsertBaseModel(c),
		Uid:       uid,
		Continues: continues,
		Score:     score,
	}

	_, err = GetBalanceSrv().OperAmount(c, uid, enum.BalanceTypeScore, decimal.NewFromInt(score), "sign_"+util.GetRandInt64(), "签到领积分", func(tx *gorm.DB) error {
		if err := tx.Create(log).Error; err != nil {
			return err
		}
		return tx.Create(&model.BeeScoreLog{
			BaseModel: *kit.GetInsertBaseModel(c),
			Uid:       kit.GetUid(c),
			Score:     decimal.NewFromInt(score),
			Remark:    "签到领积分",
		}).Error
	})
	return err
}

func (srv *ScoreSrv) GetSignLogs(c context.Context) ([]*model.BeeSignLog, error) {
	var list []*model.BeeSignLog
	now := carbon.Now()
	monthBegin := now.StartOfMonth().ToDateTimeString()
	monthEnd := now.EndOfMonth().ToDateTimeString()
	err := db.GetDB().Where("uid=? and date_add >= ? and date_add <= ?", kit.GetUid(c), monthBegin, monthEnd).Order("date_add asc").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (srv *ScoreSrv) sameDay(t1, t2 time.Time) bool {
	return t1.Format("2006-01-02") == t2.Format("2006-01-02")
}

func (srv *ScoreSrv) GetScoreLogs(c context.Context, page int64, size int64) ([]*model.BeeScoreLog, error) {
	var list []*model.BeeScoreLog
	err := db.GetDB().Where("uid=?", kit.GetUid(c)).Order("id desc").Offset(int((page - 1) * size)).Limit(int(size)).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (srv *ScoreSrv) WxaGroup(c context.Context) error {
	score := decimal.NewFromInt(10)
	_, err := GetBalanceSrv().OperAmount(c, kit.GetUid(c), enum.BalanceTypeScore, score, "wxagroup_"+util.GetRandInt64(), "微信分享领积分", func(tx *gorm.DB) error {
		return tx.Create(&model.BeeScoreLog{
			BaseModel: *kit.GetInsertBaseModel(c),
			Uid:       kit.GetUid(c),
			Score:     score,
			Remark:    "微信分享领积分",
		}).Error
	})
	return err
}
