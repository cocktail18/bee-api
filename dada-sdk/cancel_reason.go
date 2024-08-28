package dada_sdk

type CancelReason struct {
	Id     int64
	Reason string
}

var CancelReasons = []CancelReason{
	{Id: 1, Reason: "没有配送员接单"},
	{Id: 2, Reason: "配送员没来取货"},
	{Id: 3, Reason: "配送员态度太差"},
	{Id: 4, Reason: "顾客取消订单"},
	{Id: 5, Reason: "订单填写错误"},
	{Id: 34, Reason: "配送员让我取消此单"},
	{Id: 35, Reason: "配送员不愿上门取货"},
	{Id: 36, Reason: "我不需要配送了"},
	{Id: 37, Reason: "配送员以各种理由表示无法完成订单"},
	{Id: 10000, Reason: "其他"},
}
