package enum

type BeeShopStatus int32

var BeeShopStatusMap = map[BeeShopStatus]string{
	BeeShopStatusNormal: "正常",
}

const (
	BeeShopStatusNormal BeeShopStatus = 1 //正常
)
