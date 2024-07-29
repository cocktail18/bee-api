package proto

import "github.com/shopspring/decimal"

type GetWxPayInfoRes struct {
	TimeStamp  string `json:"timeStamp"`
	OutTradeId string `json:"outTradeId"`
	Package    string `json:"package"`
	PaySign    string `json:"paySign"`
	Appid      string `json:"appid"`
	Sign       string `json:"sign"`
	SignType   string `json:"signType"`
	PrepayId   string `json:"prepayId"`
	NonceStr   string `json:"nonceStr"`
}

type PayBillRes struct {
	RealMoney decimal.Decimal `json:"realMoney"`
	Amount    decimal.Decimal `json:"amount"`
	Balance   decimal.Decimal `json:"balance"`
	Behavior  int             `json:"behavior"`
	Freeze    decimal.Decimal `json:"freeze"`
	Id        int64           `json:"id"`
	Remark    string          `json:"remark"`
	Type      int             `json:"type"`
	TypeStr   string          `json:"typeStr"`
	Uid       int64           `json:"uid"`
	UserId    int64           `json:"userId"`
}
