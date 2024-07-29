package proto

import (
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/model"
)

type GetUserDetailResp struct {
	Base *model.BeeUser `json:"base"`
}

type GetUserAmountResp struct {
	model.BeeUserAmount
}

type CashLogReq struct {
	PageReqCommon

	Behavior *enum.BeeCashLogBehavior //	0 收入； 1 支出		false
	OrderId  string                   //业务订单号		false
	Type     *enum.BeeCashLogType     //交易类型 0 充值 11 提现申请 12 提现失败 1 提现成功 2 支付订单 3 退款 4支付预约报名费
	// 5 知识付费 6分销返佣 7 分享商品奖励； 8 优惠买单； 100 押金；101 押金退还； 110 购买会员； 120 货款收入；
	// 130 分润收入； 140 积分兑换		false
	Types string //	交易类型，同时获取多种类型请用英文逗号分隔	query	false
}

type CashLogResp struct {
	PageRespCommon

	Result []*model.BeeCashLog `json:"result"`
}

type PayLogsReq struct {
	PageReqCommon

	Recharge *bool             //参数如果传true，那么只会返回所有用户的充值记录（用来支付订单的支付记录不会返回），不传该参数返回所有记录
	Status   enum.PayLogStatus // 0 未支付；1 已支付；2 失败
}

type PayLogsResp []*model.BeePayLog

type BindWxMobileReq struct {
	Code          string `json:"code,omitempty" form:"code"`
	EncryptedData string `json:"encryptedData,omitempty" form:"encryptedData"` //	用户信息加密数据
	Iv            string `json:"iv,omitempty" form:"iv"`
	Pwd           string `json:"pwd,omitempty" form:"pwd"`
}
