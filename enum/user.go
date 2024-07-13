package enum

type BeeUserStatus int32
type BeeUserSource int32

const (
	BeeUserSourceWx BeeUserSource = 1

	BeeUserStatusNormal BeeUserStatus = 0
)

var BeeUserSourceMap = map[BeeUserSource]string{
	BeeUserSourceWx: "微信",
}
