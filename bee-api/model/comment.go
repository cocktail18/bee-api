package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
)

type BeeComment struct {
	common.BaseModel
	Uid        int64              `gorm:"column:uid;type:bigint(20);comment:uid" json:"uid"`                             //业务ID，比如文章id
	RefId      int64              `gorm:"column:ref_id;type:bigint(20);comment:ref_id" json:"refId"`                     //业务ID，比如文章id
	Pid        int64              `gorm:"column:pid;type:bigint(20);comment:pid" json:"pid"`                             //引用的留言ID
	ShopId     int64              `gorm:"column:shop_id;type:bigint(20);comment:shop_id" json:"shopId"`                  //门店ID
	Type       enum.CommentType   `gorm:"column:type;type:int(11);comment:type" json:"type"`                             //0 网站留言 1 意见反馈 2 投诉建议 3 CMS文章评论 4 接口申请 5 预约/活动/报名 6 商城订单
	Pics       string             `gorm:"column:pics;type:varchar(2000);comment:pics" json:"pics"`                       // 图片数组
	File       string             `gorm:"column:file;type:varchar(255);comment:file" json:"file"`                        // 文件附件链接，可传图片、文件
	Mobile     string             `gorm:"column:mobile;type:varchar(100);comment:mobile" json:"mobile"`                  // 联系电话
	Content    string             `gorm:"column:content;type:varchar(2000);comment:content"  json:"content"`             // 留言/评论内容
	ExtJsonStr string             `gorm:"column:ext_json_str;type:varchar(2000);comment:ext_json_str" json:"extJsonStr"` //属性信息,JSON格式的字符串
	Status     enum.CommentStatus `gorm:"column:status;type:int(11);comment:状态" json:"status"`
	StatusStr  string             `gorm:"-" json:"statusStr"`
}

func (m *BeeComment) TableName() string {
	return "bee_comment"
}

func (m *BeeComment) FillData() {
	m.StatusStr = enum.CommentStatusMap[m.Status]
}
