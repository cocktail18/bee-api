package yunlaba_sdk

import "context"

func (sdk *YunlabaSdk) CreateOrder(ctx context.Context, traceId string, req *CreateOrderReq) error {
	return sdk.post(ctx, "order.create", traceId, req, nil)
}

func (sdk *YunlabaSdk) CancelOrder(ctx context.Context, req *CancelOrderRequest) error {
	return sdk.post(ctx, "order.cancel.push", "", req, nil)
}

func (sdk *YunlabaSdk) QueryShopInfo(ctx context.Context, req *QueryShopReq, res *QueryShopRes) error {
	return sdk.post(ctx, "resp.shop.get", "", req, res)
}

func (sdk *YunlabaSdk) BindShop(ctx context.Context, req *BindShopReq) error {
	return sdk.post(ctx, "shop.bind.msg", "", req, nil)
}

func (sdk *YunlabaSdk) UpdateShopInfo(ctx context.Context, req *BindShopReq) error {
	return sdk.post(ctx, "shop.change.push", "", req, nil)
}
