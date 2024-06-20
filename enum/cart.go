package enum

type CartStatus int

const (
	CartStatusNormal CartStatus = 1
)

var (
	CartStatusStrMap = map[CartStatus]string{
		CartStatusNormal: "正常",
	}
)
