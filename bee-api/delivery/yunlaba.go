package delivery

import (
	"context"
	"errors"
	"gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"time"
	yunlabasdk "yunlaba/gosdk"
)

type YunlabaDelivery struct {
	sdk *yunlabasdk.YunlabaSdk
}

var _ Delivery = (*YunlabaDelivery)(nil)

func NewYunlabaDelivery(cfg *model.BeeDelivery) *YunlabaDelivery {
	return &YunlabaDelivery{
		sdk: yunlabasdk.NewYunlabaSdk(cfg.AppKey, cfg.AppSecret, cfg.SourceId, cfg.Debug),
	}
}

func (d *YunlabaDelivery) SupportPreOrder() bool {
	return false
}

func (d *YunlabaDelivery) BindShop(ctx context.Context, req *yunlabasdk.BindShopReq) error {
	return d.sdk.BindShop(ctx, req)
}

func (d *YunlabaDelivery) QueryShopInfo(ctx context.Context, req *yunlabasdk.QueryShopReq) (*yunlabasdk.QueryShopRes, error) {
	var res = &yunlabasdk.QueryShopRes{}
	if err := d.sdk.QueryShopInfo(ctx, req, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (d *YunlabaDelivery) UpdateShopInfo(ctx context.Context, req *yunlabasdk.BindShopReq) error {
	return d.sdk.UpdateShopInfo(ctx, req)
}

func (d *YunlabaDelivery) QueryOrder(ctx context.Context, orderId string) (*proto.QueryDeliveryResult, error) {
	return nil, errors.New("云喇叭不支持查询订单信息")
}

func (d *YunlabaDelivery) QueryDeliverFee(ctx context.Context, req *proto.QueryDeliverFeeReq) (*proto.QueryDeliverFeeResult, error) {
	return &proto.QueryDeliverFeeResult{
		DeliveryNo: util.GetUUID(), //trace id
	}, nil
}

func (d *YunlabaDelivery) AddOrder(ctx context.Context, req *proto.AddDeliverOrderReq) error {
	return errors.New("不支持预下单")
}

func (d *YunlabaDelivery) AddOrderDirect(ctx context.Context, req *proto.AddOrderDirectReq) (*proto.AddOrderDirectResp, error) {
	now := time.Now()
	err := d.sdk.CreateOrder(ctx, req.DeliveryNo, &yunlabasdk.CreateOrderReq{
		CreatedTime:        cast.ToString(now.Unix()),
		DaySeq:             req.DaySeq,
		DeliveryTime:       0,
		Foods:              nil,
		NotifyUrl:          d.GetNotifyUrl(ctx, config.GetHost()),
		OrderId:            req.OriginId, //订单号order_id参数请使用由纯数字组成的字符串，确保风格统一
		OriginalPrice:      req.CargoPrice.String(),
		PaidPrice:          req.CargoPrice.String(),
		Quantity:           0,
		RecipientAddress:   req.ReceiverAddress,
		RecipientLatitude:  cast.ToString(req.ReceiverLat),
		RecipientLongitude: cast.ToString(req.ReceiverLng),
		RecipientName:      req.ReceiverName,
		RecipientPhone:     req.ReceiverPhone,
		Remarks:            req.Info,
		ShopAddress:        req.ShopAddress,
		ShopId:             req.ShopNo,
		ShopName:           req.ShopName,
		ShopPhone:          req.ShopPhone,
		Source:             "",
		UpdatedTime:        "",
	})
	if err != nil {
		return nil, err
	}
	return &proto.AddOrderDirectResp{
		QueryDeliverFeeResult: proto.QueryDeliverFeeResult{
			DeliveryNo: req.DeliveryNo,
		},
	}, nil
}

func (d *YunlabaDelivery) CancelOrder(ctx context.Context, req *proto.CancelDeliverOrderReq) (*proto.CancelDeliverOrderRes, error) {
	err := d.sdk.CancelOrder(ctx, &yunlabasdk.CancelOrderRequest{
		OrderId:      req.OrderId,
		CancelReason: req.CancelReason,
		Source:       util.IF(req.Source == "", "自营", req.Source),
		ShopId:       req.ShopId,
	})
	if err != nil {
		return nil, err
	}
	return &proto.CancelDeliverOrderRes{}, nil
}

func (d *YunlabaDelivery) GetNotifyUrl(c context.Context, host string) string {
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
	return host + "/" + kit.GetDomain(c) + "/notify/yunlaba"
}

func (d *YunlabaDelivery) GetResponse(cmd string, body *yunlabasdk.ResponseBody) (*yunlabasdk.Request, error) {
	return d.sdk.GetResponse(cmd, body)
}
