package enum

type CommentType int32
type CommentStatus int32

var CommentTypeMap = map[CommentType]string{
	CommentTypeOrder: "商城订单",
}

var CommentStatusMap = map[CommentStatus]string{
	CommentStatusNone: "待审核",
}

const (
	CommentTypeOrder CommentType = 6

	CommentStatusNone CommentStatus = 0 //待审核
)
