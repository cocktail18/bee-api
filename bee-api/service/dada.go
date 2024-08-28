package service

import (
	"context"
	dadasdk "dada/gosdk"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"sync"
)

type DadaSrv struct {
	BaseSrv
}

var dadaSrvOnce sync.Once
var dadaSrvInstance *DadaSrv

func GetDadaSrv() *DadaSrv {
	dadaSrvOnce.Do(func() {
		dadaSrvInstance = &DadaSrv{}
	})
	return dadaSrvInstance
}

func (s *DadaSrv) Notify(c context.Context, param *dadasdk.CallbackParam) error {
	peisongOrderNo := param.OrderId
	peisongOrderInfo, err := GetOrderSrv().GetPeisongOrderInfoByPeisongOrderNo(c, peisongOrderNo)
	if err != nil {
		return err
	}
	if err = db.GetDB().Create(&model.BeeOrderPeisongLog{
		BaseModel:      *kit.GetInsertBaseModel(c),
		OrderId:        peisongOrderInfo.OrderId,
		PeisongOrderNo: peisongOrderNo,
		Log:            util.ToJsonWithoutErr(param, ""),
		Remark:         s.GetRemark(param),
	}).Error; err != nil {
		logger.GetLogger().Error("创建配送订单日志失败", zap.Error(err))
		return err
	}
	logger.GetLogger().Info("当前配送订单状态", zap.Any("peisongOrderInfo", peisongOrderInfo), zap.Any("peisongOrderInfo", peisongOrderInfo))
	if peisongOrderInfo.Status == enum.OrderPaisongStatusFinish {
		logger.GetLogger().Warn("配送订单已完成，忽略回调", zap.String("peisongOrderNo", peisongOrderInfo.PeisongOrderNo))
		return nil
	}
	// 以dada为准，所以不需要映射
	peisongOrderInfo.Status = enum.OrderPaisongStatus(param.OrderStatus)
	peisongOrderInfo.ThirdStatus = util.ToJsonWithoutErr(param, "")
	return GetOrderSrv().UpdatePeisongOrderInfoStatus(c, peisongOrderInfo)
}

func (s *DadaSrv) GetRemark(param *dadasdk.CallbackParam) string {
	if info, ok := dadasdk.CallbackStatusDescriptions[param.OrderStatus]; ok {
		if param.OrderStatus == dadasdk.HAD_CANCEL {
			return info + "，原因：" + param.CancelReason
		} else {
			return info
		}
	} else {
		return "未知状态" + cast.ToString(param.OrderStatus)
	}
}
