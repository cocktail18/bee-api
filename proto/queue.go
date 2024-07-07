package proto

import "gitee.com/stuinfer/bee-api/enum"

type MyQueue struct {
	QueuingLog    *MyQueueLog      `json:"queuingLog"`
	QueuingUpType *MyQueuingUpType `json:"queuingUpType"`
}

type MyQueueLog struct {
	Number    int64 `json:"number"`
	Status    enum.BeeUserQueueStatus
	StatusStr string `json:"statusStr"`
	Name      string `json:"name"`
}

type MyQueuingUpType struct {
	CurNumber int64 `json:"curNumber"`
	Minitus   int64 `json:"minitus"`
}
