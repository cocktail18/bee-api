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
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"sync"
	yunlabasdk "yunlaba/gosdk"
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("找不到配送供应商配置，请先到后台添加")
		}
		return nil, err
	}
	return delivery.GetDeliveryImpl(item), nil
}

func (s *DeliverySrv) SupportPreOrder(ctx context.Context, t enum.DeliveryType) (bool, error) {
	impl, err := s.getByType(ctx, t)
	if err != nil {
		return false, err
	}
	return impl.SupportPreOrder(), nil
}

func (s *DeliverySrv) QueryDeliveryFee(ctx context.Context, t enum.DeliveryType, req *proto.QueryDeliverFeeReq) (*proto.QueryDeliverFeeResult, error) {
	impl, err := s.getByType(ctx, t)
	if err != nil {
		return nil, err
	}
	return impl.QueryDeliverFee(ctx, req)
}

func (s *DeliverySrv) AddOrderDirect(ctx context.Context, t enum.DeliveryType, req *proto.AddOrderDirectReq) (*proto.AddOrderDirectResp, error) {
	impl, err := s.getByType(ctx, t)
	if err != nil {
		return nil, err
	}
	return impl.AddOrderDirect(ctx, req)
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

func (s *DeliverySrv) BindYunlaba(ctx context.Context, shopId int64, appId string, state string) error {
	// 获取店铺信息
	shop, err := GetShopSrv().GetShopInfo(ctx, shopId, 0, 0)
	if err != nil {
		return err
	}
	_impl, err := s.getByType(ctx, enum.DeliveryTypeYunlaba)
	if err != nil {
		return err
	}
	impl := _impl.(*delivery.YunlabaDelivery)
	//info, err := impl.QueryShopInfo(ctx, &yunlabasdk.QueryShopReq{ShopId: cast.ToString(shopId)})
	//if err != nil {
	//	return err
	//}
	bindReq := &yunlabasdk.BindShopReq{
		Address:    shop.Address,
		Latitude:   shop.Latitude,
		Longitude:  shop.Longitude,
		Name:       shop.Name,
		Phone:      shop.LinkPhone,
		PictureUrl: "",
		ShopId:     cast.ToString(shop.Id),
		State:      state,
	}
	//if info == nil || info.ShopId == "" {
	return impl.BindShop(ctx, bindReq)
	//}

	//return impl.UpdateShopInfo(ctx, bindReq)
}
