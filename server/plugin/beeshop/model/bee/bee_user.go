// 自动生成模板BeeUser
package bee

import (
	"time"
)

// beeUser表 结构体  BeeUser
type BeeUser struct {
	Id                    *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                                                  //id字段
	UserId                *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`                                          //商店用户ID
	IsDeleted             *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`                                            //已删除
	DateAdd               *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                                                 //创建时间
	DateUpdate            *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`                                        //更新时间
	DateDelete            *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`                                        //删除时间
	ShowUid               *int       `json:"showUid" form:"showUid" gorm:"column:show_uid;comment:用户show_id;size:19;"`                                    //用户show_id
	AvatarUrl             string     `json:"avatarUrl" form:"avatarUrl" gorm:"column:avatar_url;comment:头像;size:100;"`                                    //头像
	CardNumber            string     `json:"cardNumber" form:"cardNumber" gorm:"column:card_number;comment:会员卡号;size:100;"`                               //会员卡号
	City                  string     `json:"city" form:"city" gorm:"column:city;comment:城市;size:100;"`                                                    //城市
	DateLogin             *time.Time `json:"dateLogin" form:"dateLogin" gorm:"column:date_login;comment:登录时间;"`                                           //登录时间
	Gender                *int       `json:"gender" form:"gender" gorm:"column:gender;comment:性别;size:19;"`                                               //性别
	IpAdd                 string     `json:"ipAdd" form:"ipAdd" gorm:"column:ip_add;comment:注册ip;size:100;"`                                              //注册ip
	IpLogin               string     `json:"ipLogin" form:"ipLogin" gorm:"column:ip_login;comment:登录ip;size:100;"`                                        //登录ip
	IsFaceCheck           *bool      `json:"isFaceCheck" form:"isFaceCheck" gorm:"column:is_face_check;comment:人脸识别;"`                                    //人脸识别
	IsIdcardCheck         *bool      `json:"isIdcardCheck" form:"isIdcardCheck" gorm:"column:is_idcard_check;comment:身份证识别;"`                             //身份证识别
	IsManager             *bool      `json:"isManager" form:"isManager" gorm:"column:is_manager;comment:管理员;"`                                            //管理员
	IsSeller              *bool      `json:"isSeller" form:"isSeller" gorm:"column:is_seller;comment:销售人员;"`                                              //销售人员
	IsSendRegisterCoupons *bool      `json:"isSendRegisterCoupons" form:"isSendRegisterCoupons" gorm:"column:is_send_register_coupons;comment:已发放注册优惠券;"` //已发放注册优惠券
	IsShopManager         *bool      `json:"isShopManager" form:"isShopManager" gorm:"column:is_shop_manager;comment:店长;"`                                //店长
	IsTeamLeader          *bool      `json:"isTeamLeader" form:"isTeamLeader" gorm:"column:is_team_leader;comment:团长;"`                                   //团长
	IsTeamMember          *bool      `json:"isTeamMember" form:"isTeamMember" gorm:"column:is_team_member;comment:团员;"`                                   //团员
	IsUserAttendant       *bool      `json:"isUserAttendant" form:"isUserAttendant" gorm:"column:is_user_attendant;comment:关注用户;"`                        //关注用户
	IsVirtual             *bool      `json:"isVirtual" form:"isVirtual" gorm:"column:is_virtual;comment:是否虚拟人;"`                                          //是否虚拟人
	Nick                  string     `json:"nick" form:"nick" gorm:"column:nick;comment:名字;size:100;"`                                                    //名字
	Province              string     `json:"province" form:"province" gorm:"column:province;comment:所在省;size:100;"`                                       //所在省
	VipLevel              *int       `json:"vipLevel" form:"vipLevel" gorm:"column:vip_level;comment:vip等级;size:19;"`                                     //vip等级
	Source                *int       `json:"source" form:"source" gorm:"column:source;comment:注册来源;size:19;"`                                             //注册来源
	Status                *int       `json:"status" form:"status" gorm:"column:status;comment:状态;size:19;"`                                               //状态
}

// TableName beeUser表 BeeUser自定义表名 bee_user
func (BeeUser) TableName() string {
	return "bee_user"
}
