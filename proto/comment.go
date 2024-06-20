package proto

import "gitee.com/stuinfer/bee-api/enum"

type CommentAddReq struct {
	RefId      int64            `json:"refId"`      //业务ID，比如文章id
	Pid        int64            `json:"pid"`        //引用的留言ID
	ShopId     int64            `json:"shopId"`     //门店ID
	Type       enum.CommentType `json:"type"`       //0 网站留言 1 意见反馈 2 投诉建议 3 CMS文章评论 4 接口申请 5 预约/活动/报名 6 商城订单
	Pics       string           `json:"pics"`       // 图片数组
	File       string           `json:"file"`       // 文件附件链接，可传图片、文件
	Mobile     string           `json:"mobile"`     // 联系电话
	Content    string           `json:"content"`    // 留言/评论内容
	ExtJsonStr string           `json:"extJsonStr"` //属性信息,JSON格式的字符串

	Ip string
}
