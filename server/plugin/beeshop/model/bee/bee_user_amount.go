// 自动生成模板BeeUserAmount
package bee

import (
	"time"
)

// 用户资产 结构体  BeeUserAmount
type BeeUserAmount struct {
	Id                *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                                       //id字段
	UserId            *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`                               //商店用户ID
	IsDeleted         *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`                                 //已删除
	DateAdd           *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                                      //创建时间
	DateUpdate        *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`                             //更新时间
	DateDelete        *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`                             //删除时间
	Uid               *int       `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:19;"`                                           //用户id
	Balance           *float64   `json:"balance" form:"balance" gorm:"column:balance;comment:余额;size:22;"`                                 //余额
	Freeze            *float64   `json:"freeze" form:"freeze" gorm:"column:freeze;comment:冻结金额;size:20;"`                                  //冻结金额
	FxCommisionPaying *float64   `json:"fxCommisionPaying" form:"fxCommisionPaying" gorm:"column:fx_commision_paying;comment:佣金;size:20;"` //佣金
	Growth            *float64   `json:"growth" form:"growth" gorm:"column:growth;comment:成长值;size:20;"`                                   //成长值
	Score             *float64   `json:"score" form:"score" gorm:"column:score;comment:积分;size:20;"`                                       //积分
	ScoreLastRound    *float64   `json:"scoreLastRound" form:"scoreLastRound" gorm:"column:score_last_round;comment:上一个周期的积分;size:20;"`    //上一个周期的积分
	TotalPayAmount    *float64   `json:"totalPayAmount" form:"totalPayAmount" gorm:"column:total_pay_amount;comment:支付总额;size:20;"`        //支付总额
	TotalPayNumber    *float64   `json:"totalPayNumber" form:"totalPayNumber" gorm:"column:total_pay_number;comment:支付订单数;size:20;"`       //支付订单数
	TotalScore        *float64   `json:"totalScore" form:"totalScore" gorm:"column:total_score;comment:总积分;size:20;"`                      //总积分
	TotalWithdraw     *float64   `json:"totalWithdraw" form:"totalWithdraw" gorm:"column:total_withdraw;comment:撤回次数;size:20;"`            //撤回次数
	TotalConsumed     *float64   `json:"totalConsumed" form:"totalConsumed" gorm:"column:total_consumed;comment:累计消费金额;size:20;"`          //累计消费金额
}

// TableName 用户资产 BeeUserAmount自定义表名 bee_user_amount
func (BeeUserAmount) TableName() string {
	return "bee_user_amount"
}
