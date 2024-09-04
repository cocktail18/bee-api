package delivery

import (
	"context"
	dadasdk "dada/gosdk"
	"errors"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/gin-gonic/gin"
)

type DadaDelivery struct {
	sdk *dadasdk.DadaSdk
}

var _ Delivery = (*DadaDelivery)(nil)

func NewDadaDelivery(cfg *model.BeeDelivery) *DadaDelivery {
	return &DadaDelivery{
		sdk: dadasdk.NewDadaSdk(cfg.AppKey, cfg.AppSecret, cfg.SourceId, cfg.Debug),
	}
}

func (d *DadaDelivery) AddOrderDirect(ctx context.Context, req *proto.AddOrderDirectReq) (*proto.AddOrderDirectResp, error) {
	return nil, errors.New("not support")
}

func (d *DadaDelivery) SupportPreOrder() bool {
	return true
}

func (d *DadaDelivery) QueryOrder(ctx context.Context, orderId string) (*proto.QueryDeliveryResult, error) {
	res, err := d.sdk.QueryOrder(ctx, &dadasdk.QueryOrderReq{
		OrderId: orderId,
	})
	if err != nil {
		return nil, err
	}
	return &proto.QueryDeliveryResult{
		OrderId:          res.OrderId,
		StatusCode:       res.StatusCode,
		StatusMsg:        res.StatusMsg,
		TransporterName:  res.TransporterName,
		TransporterPhone: res.TransporterPhone,
		TransporterLng:   res.TransporterLng,
		TransporterLat:   res.TransporterLat,
		DeliveryFee:      res.DeliveryFee,
		Tips:             res.Tips,
		CouponFee:        res.CouponFee,
		InsuranceFee:     res.InsuranceFee,
		ActualFee:        res.ActualFee,
		Distance:         res.Distance,
		CreateTime:       res.CreateTime,
		AcceptTime:       res.AcceptTime,
		FetchTime:        res.FetchTime,
		FinishTime:       res.FinishTime,
		CancelTime:       res.CancelTime,
		OrderFinishCode:  res.OrderFinishCode,
		DeductFee:        res.DeductFee,
	}, nil
}

func (d *DadaDelivery) QueryDeliverFee(ctx context.Context, req *proto.QueryDeliverFeeReq) (*proto.QueryDeliverFeeResult, error) {
	//TODO implement me
	res, err := d.sdk.QueryDeliverFee(ctx, &dadasdk.QueryDeliverFeeReq{
		ShopNo:          req.ShopNo,
		OriginId:        req.OriginId,
		CargoPrice:      req.CargoPrice,
		IsPrepay:        req.IsPrepay,
		ReceiverName:    req.ReceiverName,
		ReceiverAddress: req.ReceiverAddress,
		Callback:        req.Callback,
		CargoWeight:     req.CargoWeight,
		ReceiverLat:     req.ReceiverLat,
		ReceiverLng:     req.ReceiverLng,
		ReceiverPhone:   req.ReceiverPhone,
		ReceiverTel:     req.ReceiverTel,
		Tips:            req.Tips,
		Info:            req.Info,
		CargoType:       req.CargoType,
		CargoNum:        req.CargoNum,
	})
	if err != nil {
		return nil, err
	}
	return &proto.QueryDeliverFeeResult{
		Distance:     res.Distance,
		DeliveryNo:   res.DeliveryNo,
		Fee:          res.Fee,
		DeliverFee:   res.DeliverFee,
		CouponFee:    res.CouponFee,
		Tips:         res.Tips,
		InsuranceFee: res.InsuranceFee,
	}, nil
}

func (d *DadaDelivery) AddOrder(ctx context.Context, req *proto.AddDeliverOrderReq) error {
	return d.sdk.AddAfterQuery(ctx, &dadasdk.AddAfterQueryReq{
		DeliveryNo:  req.DeliveryNo,
		EnableReset: req.EnableReset,
		Info:        req.Info,
	})
}

func (d *DadaDelivery) CancelOrder(ctx context.Context, req *proto.CancelDeliverOrderReq) (*proto.CancelDeliverOrderRes, error) {
	res, err := d.sdk.CancelOrder(ctx, &dadasdk.CancelOrderReq{
		OrderId:        req.OrderId,
		CancelReasonId: req.CancelReasonId,
		CancelReason:   req.CancelReason,
	})
	if err != nil {
		return nil, err
	}
	return &proto.CancelDeliverOrderRes{
		DeductFee: res.DeductFee,
	}, nil
}

func (d *DadaDelivery) GetNotifyUrl(c context.Context, host string) string {
	if host == "" {
		ginContext, ok := c.(*gin.Context)
		if ok {
			scheme := ginContext.Request.URL.Scheme
			// Default to HTTP if Scheme is empty
			if scheme == "" {
				if ginContext.Request.TLS != nil {
					scheme = "https"
				} else {
					scheme = "http"
				}
			}
			host = scheme + "://" + ginContext.Request.Host
		}
	}
	return host + "/" + kit.GetDomain(c) + "/notify/dada"
}
