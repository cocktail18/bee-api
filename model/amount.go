package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/shopspring/decimal"
)

// BeeUserAmount 用户资产
type BeeUserAmount struct {
	common.BaseModel
	Uid               int64           `gorm:"column:uid;type:bigint(20);comment:用户id" json:"uid"`
	Balance           decimal.Decimal `gorm:"column:balance;type:double;comment:余额" json:"balance"`
	Freeze            decimal.Decimal `gorm:"column:freeze;type:decimal(20,2);comment:冻结金额" json:"freeze"`
	FxCommisionPaying decimal.Decimal `gorm:"column:fx_commision_paying;type:decimal(20,2);comment:佣金" json:"fx_commision_paying"`
	Growth            decimal.Decimal `gorm:"column:growth;type:decimal(20,2);comment:成长值" json:"growth"`
	Score             decimal.Decimal `gorm:"column:score;type:decimal(20,2);comment:积分" json:"score"`
	ScoreLastRound    decimal.Decimal `gorm:"column:score_last_round;type:decimal(20,2);comment:上一个周期的积分" json:"score_last_round"`
	TotalPayAmount    decimal.Decimal `gorm:"column:total_pay_amount;type:decimal(20,2);comment:支付总额" json:"total_pay_amount"`
	TotalPayNumber    decimal.Decimal `gorm:"column:total_pay_number;type:decimal(20,2);comment:支付订单数" json:"total_pay_number"`
	TotalScore        decimal.Decimal `gorm:"column:total_score;type:decimal(20,2);comment:总积分" json:"total_score"`
	TotalWithdraw     decimal.Decimal `gorm:"column:total_withdraw;type:decimal(20,2);comment:撤回次数" json:"total_withdraw"`
	TotalConsumed     decimal.Decimal `gorm:"column:total_consumed;type:decimal(20,2);comment:累计消费金额" json:"total_consumed"`
	Pwd               string          `gorm:"column:pwd;type:varchar(100);comment:密码" json:"-"`
	Salt              string          `gorm:"column:salt;type:varchar(100);comment:salt" json:"-"`
}

func (m *BeeUserAmount) TableName() string {
	return "bee_user_amount"
}

func (m *BeeUserAmount) GetPwdEncode(pwd string) string {
	return util.Md5(m.Salt + pwd + m.Salt)
}
