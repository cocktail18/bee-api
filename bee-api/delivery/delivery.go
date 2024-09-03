package delivery

import (
	"context"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
)

type Delivery interface {
	QueryOrder(ctx context.Context, orderId string) (*proto.QueryDeliveryResult, error)
	QueryDeliverFee(ctx context.Context, req *proto.QueryDeliverFeeReq) (*proto.QueryDeliverFeeResult, error)
	AddOrder(ctx context.Context, req *proto.AddDeliverOrderReq) error
	AddOrderDirect(ctx context.Context, req *proto.AddOrderDirectReq) (*proto.AddOrderDirectResp, error)
	CancelOrder(ctx context.Context, req *proto.CancelDeliverOrderReq) (*proto.CancelDeliverOrderRes, error)
	GetNotifyUrl(c context.Context, host string) string
	SupportPreOrder() bool
}

func GetDeliveryImpl(config *model.BeeDelivery) Delivery {
	switch config.Type {
	case enum.DeliveryTypeDada:
		return NewDadaDelivery(config)
	case enum.DeliveryTypeYunlaba:
		return NewYunlabaDelivery(config)
	}
	return nil
}
