package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/shopspring/decimal"
)

// BeeUser
type BeeUser struct {
	common.BaseModel
	ShowUid                    int64              `gorm:"column:show_uid;type:bigint(20);comment:用户show_id" json:"show_uid"`
	AvatarUrl                  string             `gorm:"column:avatar_url;type:varchar(100);comment:头像" json:"avatar_url"`
	BirthdayProcessSuccessYear int64              `gorm:"column:birthday_process_success_year;type:bigint(11);comment:生日任务成功年份" json:"birthday_process_success_year"`
	CardNumber                 string             `gorm:"column:card_number;type:varchar(100);comment:会员卡号" json:"card_number"`
	City                       string             `gorm:"column:city;type:varchar(100);comment:城市" json:"city"`
	DateLogin                  common.JsonTime    `gorm:"column:date_login;type:datetime;comment:登录时间" json:"date_login"`
	Gender                     int                `gorm:"column:gender;type:bigint(11);comment:性别" json:"gender"`
	IpAdd                      string             `gorm:"column:ip_add;type:varchar(100);comment:注册ip" json:"ip_add"`
	IpLogin                    string             `gorm:"column:ip_login;type:varchar(100);comment:登录ip" json:"ip_login"`
	IsFaceCheck                bool               `gorm:"column:is_face_check;type:tinyint(1);comment:人脸识别" json:"is_face_check"`
	IsIdcardCheck              bool               `gorm:"column:is_idcard_check;type:tinyint(1);comment:身份证识别" json:"is_idcard_check"`
	IsManager                  bool               `gorm:"column:is_manager;type:tinyint(1);comment:管理员" json:"is_manager"`
	IsSeller                   bool               `gorm:"column:is_seller;type:tinyint(1);comment:销售人员" json:"is_seller"`
	IsSendRegisterCoupons      bool               `gorm:"column:is_send_register_coupons;type:tinyint(1);comment:已发放注册优惠券" json:"is_send_register_coupons"`
	IsShopManager              bool               `gorm:"column:is_shop_manager;type:tinyint(1);comment:店长" json:"is_shop_manager"`
	IsTeamLeader               bool               `gorm:"column:is_team_leader;type:tinyint(1);comment:团长" json:"is_team_leader"`
	IsTeamMember               bool               `gorm:"column:is_team_member;type:tinyint(1);comment:团员" json:"is_team_member"`
	IsUserAttendant            bool               `gorm:"column:is_user_attendant;type:tinyint(1);comment:关注用户" json:"is_user_attendant"`
	IsVirtual                  bool               `gorm:"column:is_virtual;type:tinyint(1);comment:是否虚拟人" json:"is_virtual"`
	Nick                       string             `gorm:"column:nick;type:varchar(100);comment:名字" json:"nick"`
	Province                   string             `gorm:"column:province;type:varchar(100);comment:所在省" json:"province"`
	VipLevel                   int64              `gorm:"column:vip_level;type:bigint(20);comment:vip等级" json:"vip_level"`
	Source                     enum.BeeUserSource `gorm:"column:source;type:bigint(11);comment:注册来源" json:"source"`
	Status                     enum.BeeUserStatus `gorm:"column:status;type:bigint(11);comment:状态" json:"status"`
}

func (m *BeeUser) TableName() string {
	return "bee_user"
}

func (m *BeeUser) GetUid() int64 {
	return m.Id
}

func (m *BeeUser) GetUserId() int64 {
	return m.UserId
}

type BeeUserMapper struct {
	common.BaseModel

	Uid     int64  `gorm:"column:uid;type:bigint(20)" json:"uid"`
	Source  int64  `gorm:"column:source;type:bigint(20)" json:"source"`
	OpenId  string `gorm:"column:open_id;unique;type:varchar(200)" json:"openId"`
	UnionId string `gorm:"column:union_id;type:varchar(200)" json:"unionId"`
}

func (m *BeeUserMapper) TableName() string {
	return "bee_user_mapper"
}

type BeeCashLog struct {
	common.BaseModel
	Amount   decimal.Decimal         `gorm:"column:amount;type:decimal(10,2);comment:金额" json:"amount"`
	Balance  decimal.Decimal         `gorm:"column:balance;type:decimal(10,2);comment:剩余可用余额" json:"balance"`
	Behavior enum.BeeCashLogBehavior `gorm:"column:behavior;type:int(10);comment:0 收入 1 支出" json:"behavior"`
	Freeze   decimal.Decimal         `gorm:"column:freeze;type:decimal(10,2);comment:剩余冻结金额" json:"freeze"`
	Remark   string                  `gorm:"column:remark;type:varchar(100);comment:remark" json:"remark"`
	Type     enum.BeeCashLogType     `gorm:"column:type;type:int(10);comment:类型" json:"type"`
	TypeStr  string                  `gorm:"-" json:"typeStr"`
	Uid      int                     `gorm:"column:uid;type:bigint(20)" json:"uid"`
}

func (b *BeeCashLog) TableName() string {
	return "bee_cash_log"
}

func (b *BeeCashLog) FillData() {
	b.TypeStr = enum.BeeCashLogTypeMap[b.Type]
}

type BeeUserDynamicCode struct {
	common.BaseModel
	Uid      int64  `gorm:"column:uid;type:bigint(20)" json:"uid"`
	Code     string `gorm:"column:code;type:varchar(100)" json:"code"`
	ExpireAt int64  `gorm:"column:expire_at;type:bigint(20)" json:"expireAt"`
}