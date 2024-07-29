package enum

type BeeQueueStatus int32
type BeeUserQueueStatus int32 // 0 排队中 1 受理中 2 已处理 3 已过号

var BeeQueueStatusMap = map[BeeQueueStatus]string{
	BeeQueueStatusNormal: "正常",
	BeeQueueStatusClose:  "关闭",
	BeeQueueStatusFull:   "满号",
}
var BeeUserQueueStatusMap = map[BeeUserQueueStatus]string{
	BeeUserQueueStatusNormal: "排队中",
	BeeUserQueueStatusDoing:  "受理中",
	BeeUserQueueStatusDone:   "已处理",
	BeeUserQueueStatusOver:   "已过号",
}

const (
	BeeQueueStatusNormal BeeQueueStatus = 0
	BeeQueueStatusClose  BeeQueueStatus = 1
	BeeQueueStatusFull   BeeQueueStatus = 2

	BeeUserQueueStatusNormal BeeUserQueueStatus = 0
	BeeUserQueueStatusDoing  BeeUserQueueStatus = 1
	BeeUserQueueStatusDone   BeeUserQueueStatus = 2
	BeeUserQueueStatusOver   BeeUserQueueStatus = 3
)
