package enum

type CyTableStatus int8

const (
	CyTableStatusConfirm CyTableStatus = 0 //确认中
	CyTableStatusDoing   CyTableStatus = 1 //已下厨
	CyTableStatusDone    CyTableStatus = 2 //已上菜
)
