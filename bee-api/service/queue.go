package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"sync"
)

type QueueSrv struct {
	BaseSrv
}

var queueSrvOnce sync.Once
var queueSrvInstance *QueueSrv

func GetQueueSrv() *QueueSrv {
	queueSrvOnce.Do(func() {
		queueSrvInstance = &QueueSrv{}
	})
	return queueSrvInstance
}

func (srv *QueueSrv) GetMyQueueInfo(c context.Context, queueId int64, status string) ([]*proto.MyQueue, error) {
	res := make([]*proto.MyQueue, 0)
	var myQueueList = make([]*model.BeeUserQueue, 0)
	if err := db.GetDB().Where(map[string]interface{}{
		"is_deleted": 0,
	}).Find(&myQueueList).Error; err != nil {
		return nil, err
	}

	if len(myQueueList) == 0 {
		return nil, nil
	}

	var queueInfoList = make([]*model.BeeQueue, 0)
	if err := db.GetDB().Where("is_deleted = 0 and id in ?", lo.Map(myQueueList, func(item *model.BeeUserQueue, index int) int64 {
		return item.Id
	})).Find(&queueInfoList).Error; err != nil {
		return nil, err
	}
	queueInfoById := lo.KeyBy(queueInfoList, func(item *model.BeeQueue) int64 {
		return item.Id
	})
	for _, item := range myQueueList {
		queueInfo := queueInfoById[item.QueueId]
		if queueInfo == nil {
			continue
		}

		if item.Status != enum.BeeUserQueueStatusDone {
			if item.Number < queueInfo.CurNumber {
				item.Status = enum.BeeUserQueueStatusNormal
			} else if item.Number == queueInfo.CurNumber {
				item.Status = enum.BeeUserQueueStatusDoing
			} else if item.Number > queueInfo.CurNumber {
				item.Status = enum.BeeUserQueueStatusOver
			}
		}
		myQueueLog := &proto.MyQueueLog{
			Number:    item.Number,
			Status:    item.Status,
			StatusStr: "",
			Name:      queueInfo.Name,
		}
		myQueueLog.FillData()
		res = append(res, &proto.MyQueue{
			QueuingLog: myQueueLog,
			QueuingUpType: &proto.MyQueuingUpType{
				CurNumber: queueInfo.CurNumber,
				Minitus:   queueInfo.Minutes,
			},
		})
	}

	return res, nil
}

func (srv *QueueSrv) GetQueueTypes(c context.Context, status string) ([]*model.BeeQueue, error) {
	var queueInfoList = make([]*model.BeeQueue, 0)
	dbIns := db.GetDB().Where("is_deleted = ?", 0)
	if status != "" {
		dbIns = dbIns.Where("status = ?", cast.ToInt32(status))
	}
	err := dbIns.Find(&queueInfoList).Error
	if err != nil {
		return nil, err
	}
	return queueInfoList, err
}

func (srv *QueueSrv) Get(c context.Context, queueId int64, mobile string) error {
	var curQueueInfo = &model.BeeUserQueue{}
	if err := db.GetDB().Where("queue_id = ? and is_deleted = 0 ", queueId).Take(&curQueueInfo).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else {
		return enum.NewBussErr(nil, 30005, "您已取过号了")
	}
	var queueInfo = &model.BeeQueue{}
	err := db.GetDB().Where(map[string]interface{}{
		"is_deleted": 0,
		"id":         queueId,
	}).Take(&queueInfo).Error
	if err != nil {
		return err
	}
	if queueInfo.MaxNumber <= queueInfo.CurNumber {
		return enum.NewBussErr(nil, 30006, "排号已满，下次再来")
	}
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		//@todo 并发处理
		if err := tx.Model(&model.BeeQueue{}).Where(map[string]interface{}{
			"is_deleted": 0,
			"id":         queueId,
		}).Update("number_get", gorm.Expr("number_get + ?", 1)).Error; err != nil {
			return err
		}
		if err := tx.Create(&model.BeeUserQueue{
			BaseModel: *kit.GetInsertBaseModel(c),
			Uid:       kit.GetUid(c),
			QueueId:   queueId,
			Number:    queueInfo.NumberGet + 1,
			Status:    enum.BeeUserQueueStatusNormal,
		}).Error; err != nil {
			return err
		}
		return nil
	})
}