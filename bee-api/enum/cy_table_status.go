package enum

type CyTableStatus int8

var CyTableStatusStrMap = map[CyTableStatus]string{
	CyTableStatusConfirm: "确认中",
	CyTableStatusDoing:   "已下厨",
	CyTableStatusDone:    "已上菜",
}

const (
	CyTableStatusConfirm CyTableStatus = 0 //确认中
	CyTableStatusDoing   CyTableStatus = 1 //已下厨
	CyTableStatusDone    CyTableStatus = 2 //已上菜
)
