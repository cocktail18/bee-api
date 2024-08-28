package service

import (
	"context"
	"errors"
	"gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/delivery"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"sync"
)

type DeliverySrv struct {
	BaseSrv
}

var deliverSrvOnce sync.Once
var deliverSrvInstance *DeliverySrv

func GetDeliverySrv() *DeliverySrv {
	deliverSrvOnce.Do(func() {
		deliverSrvInstance = &DeliverySrv{}
	})
	return deliverSrvInstance
}

func (s *DeliverySrv) getImpl(ctx context.Context, str string) (delivery.Delivery, error) {
	deliveryType := enum.DeliveryStrMap[str]
	if deliveryType == 0 {
		return nil, errors.New("暂不支持的配送商")
	}
	return s.getByType(ctx, deliveryType)
}

func (s *DeliverySrv) getByType(ctx context.Context, deliveryType enum.DeliveryType) (delivery.Delivery, error) {
	var item = &model.BeeDelivery{}
	if err := db.GetDB().Where("user_id = ? and `type` = ? and is_deleted = 0", kit.GetUserId(ctx), deliveryType).Order("id desc").Take(item).Error; err != nil {
		return nil, err
	}
	return delivery.GetDeliveryImpl(item), nil
}

func (s *DeliverySrv) QueryDeliveryFee(ctx context.Context, t enum.DeliveryType, req *proto.QueryDeliverFeeReq) (*proto.QueryDeliverFeeResult, error) {
	impl, err := s.getByType(ctx, t)
	if err != nil {
		return nil, err
	}
	return impl.QueryDeliverFee(ctx, req)
}

func (s *DeliverySrv) AddOrder(ctx context.Context, t enum.DeliveryType, preOrderNo string) error {
	impl, err := s.getByType(ctx, t)
	if err != nil {
		return err
	}
	return impl.AddOrder(ctx, &proto.AddDeliverOrderReq{
		DeliveryNo:  preOrderNo,
		EnableReset: false,
	})
}

func (s *DeliverySrv) CancelOrder(ctx context.Context, t enum.DeliveryType, req *proto.CancelDeliverOrderReq) (*proto.CancelDeliverOrderRes, error) {
	impl, err := s.getByType(ctx, t)
	if err != nil {
		return nil, err
	}
	return impl.CancelOrder(ctx, req)
}

func (s *DeliverySrv) GetNotifyUrl(ctx context.Context, t enum.DeliveryType) (string, error) {
	impl, err := s.getByType(ctx, t)
	if err != nil {
		return "", err
	}
	return impl.GetNotifyUrl(ctx, config.GetHost()), nil
}

func (s *DeliverySrv) QueryDetail(ctx context.Context, t enum.DeliveryType, peisongOrderNo string) (*proto.QueryDeliveryResult, error) {
	impl, err := s.getByType(ctx, t)
	if err != nil {
		return nil, err
	}
	return impl.QueryOrder(ctx, peisongOrderNo)
}
