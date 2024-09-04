package service

import (
	"context"
	"encoding/json"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/delivery"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/util"
	"go.uber.org/zap"
	"sync"
	yunlabasdk "yunlaba/gosdk"
)

type YunlabaSrv struct {
	BaseSrv
}

var yunlabaSrvOnce sync.Once
var yunlabaSrvInstance *YunlabaSrv

func GetYunlabaSrv() *YunlabaSrv {
	yunlabaSrvOnce.Do(func() {
		yunlabaSrvInstance = &YunlabaSrv{}
	})
	return yunlabaSrvInstance
}

func (s *YunlabaSrv) transferLogisticsStatus(status yunlabasdk.LogisticsStatus) enum.OrderPaisongStatus {
	switch status {
	case yunlabasdk.LogisticsStatusWaitingForDelivery:
		return enum.OrderPaisongStatusPickup
	case yunlabasdk.LogisticsStatusDelivering:
		return enum.OrderPaisongStatusDelivering
	case yunlabasdk.LogisticsStatusDelivered:
		return enum.OrderPaisongStatusFinish
	case yunlabasdk.LogisticsStatusShopReady:
		return enum.OrderPaisongStatusWaiting
	case yunlabasdk.LogisticsStatusPickedUp:
		return enum.OrderPaisongStatusPickup
	default:
		return enum.OrderPaisongStatusNone
	}
}

func (s *YunlabaSrv) Notify(c context.Context, param *yunlabasdk.NotifyData) (*yunlabasdk.Request, error) {
	//@todo 检验真实性
	logger.GetLogger().Info("云喇叭回调", zap.Any("param", param))

	_impl, err := GetDeliverySrv().getByType(c, enum.DeliveryTypeYunlaba)
	if err != nil {
		return nil, err
	}
	impl := _impl.(*delivery.YunlabaDelivery)
	switch yunlabasdk.Cmd(param.Cmd) {
	case yunlabasdk.CmdDeliveryStateSync:
		var data yunlabasdk.DeliveryStateSync
		if err := json.Unmarshal([]byte(param.Body), &data); err != nil {
			logger.GetLogger().Info("云喇叭回调解析失败", zap.Any("param", param), zap.Error(err))
			return nil, err
		}

		peisongOrderInfo, err := GetOrderSrv().GetPeisongOrderInfoByPeisongOrderNo(c, data.OrderId)
		if err != nil {
			return nil, err
		}

		if err := db.GetDB().Create(&model.BeeOrderPeisongLog{
			BaseModel:      *kit.GetInsertBaseModel(c),
			OrderId:        peisongOrderInfo.OrderId,
			PeisongOrderNo: peisongOrderInfo.PeisongOrderNo,
			Log:            util.ToJsonWithoutErr(param, ""),
			Remark:         "",
		}).Error; err != nil {
			logger.GetLogger().Error("创建配送订单日志失败", zap.Error(err))
			return nil, err
		}
		logger.GetLogger().Info("当前配送订单状态", zap.Any("peisongOrderInfo", peisongOrderInfo), zap.Any("status", peisongOrderInfo.Status))
		if peisongOrderInfo.Status == enum.OrderPaisongStatusFinish {
			logger.GetLogger().Warn("配送订单已完成，忽略回调", zap.String("peisongOrderNo", peisongOrderInfo.PeisongOrderNo))
			return impl.GetResponse("resp."+param.Cmd, &yunlabasdk.ResponseBody{
				Code:   0,
				ErrMsg: "success",
			})
		}
		// 以yunlaba为准，所以不需要映射
		peisongOrderInfo.Status = s.transferLogisticsStatus(data.LogisticsStatus)
		peisongOrderInfo.ThirdStatus = util.ToJsonWithoutErr(param, "")
		if err := GetOrderSrv().UpdatePeisongOrderInfoStatus(c, peisongOrderInfo); err != nil {
			return nil, err
		}
		return impl.GetResponse("resp."+param.Cmd, &yunlabasdk.ResponseBody{
			Code:   0,
			ErrMsg: "success",
		})
	case yunlabasdk.CmdPushErrCallback:
		var data yunlabasdk.ErrorCallback
		if err := json.Unmarshal([]byte(param.Body), &data); err != nil {
			logger.GetLogger().Info("云喇叭回调解析失败", zap.Any("param", param), zap.Error(err))
			return nil, err
		}
		peisongOrderId := data.TraceId
		peisongOrderInfo, err := GetOrderSrv().GetPeisongOrderInfoByPeisongOrderId(c, peisongOrderId)
		if err != nil {
			return nil, err
		}
		if err := db.GetDB().Create(&model.BeeOrderPeisongLog{
			BaseModel:      *kit.GetInsertBaseModel(c),
			OrderId:        peisongOrderInfo.OrderId,
			PeisongOrderNo: peisongOrderInfo.PeisongOrderNo,
			Log:            util.ToJsonWithoutErr(param, ""),
			Remark:         "type:" + data.PushType + ",fail:" + data.ErrorMsg,
		}).Error; err != nil {
			logger.GetLogger().Error("创建配送订单日志失败", zap.Error(err))
			return nil, err
		}
		logger.GetLogger().Info("当前配送订单状态", zap.Any("peisongOrderInfo", peisongOrderInfo), zap.Any("status", peisongOrderInfo.Status))
		if peisongOrderInfo.Status == enum.OrderPaisongStatusFinish {
			logger.GetLogger().Warn("配送订单已完成，忽略回调", zap.String("peisongOrderNo", peisongOrderInfo.PeisongOrderNo))
			return impl.GetResponse("resp."+param.Cmd, &yunlabasdk.ResponseBody{
				Code:   0,
				ErrMsg: "success",
			})
		}
		// 以yunlaba为准，所以不需要映射
		peisongOrderInfo.Status = enum.OrderPaisongStatusDadaFail
		peisongOrderInfo.ThirdStatus = util.ToJsonWithoutErr(param, "")
		if err := GetOrderSrv().UpdatePeisongOrderInfoStatus(c, peisongOrderInfo); err != nil {
			return nil, err
		}
		return impl.GetResponse("resp."+param.Cmd, &yunlabasdk.ResponseBody{
			Code:   0,
			ErrMsg: "success",
		})
	default:
		return impl.GetResponse("resp."+param.Cmd, &yunlabasdk.ResponseBody{
			Code:   0,
			ErrMsg: "success",
		})
	}
}
