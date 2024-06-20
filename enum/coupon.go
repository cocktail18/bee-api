package enum

type CouponStatus int8
type CouponMoneyType int32
type CouponDateEndType int32

const (
	CouponStatusNormal CouponStatus = 0
	CouponStatusUsing  CouponStatus = 1
	CouponStatusUsed   CouponStatus = 2

	CouponMoneyTypeFixed CouponMoneyType = 0 //固定
	CouponMoneyTypeRatio CouponMoneyType = 1 //比例

	CouponDateEndTypeFixed CouponDateEndType = 0 //指定日期过期
	CouponDateEndTypeDelay CouponDateEndType = 1 //几天后过期
)

var CouponStatusMap = map[CouponStatus]string{
	CouponStatusNormal: "正常",
	CouponStatusUsing:  "使用中",
	CouponStatusUsed:   "已使用",
}
