package enum

type BeeQueueStatus int32
type BeeUserQueueStatus int32 // 0 排队中 1 受理中 2 已处理 3 已过号

const (
	BeeQueueStatusNormal BeeQueueStatus = 0
	BeeQueueStatusClose  BeeQueueStatus = 1
	BeeQueueStatusFull   BeeQueueStatus = 2

	BeeUserQueueStatusNormal BeeUserQueueStatus = 0
	BeeUserQueueStatusDoing  BeeUserQueueStatus = 1
	BeeUserQueueStatusDone   BeeUserQueueStatus = 2
	BeeUserQueueStatusOver   BeeUserQueueStatus = 3
)
