package enum

type BalanceType int

const (
	BalanceTypeBalance           BalanceType = 1
	BalanceTypeFreeze            BalanceType = 2
	BalanceTypeFxCommisionPaying BalanceType = 3
	BalanceTypeGrowth            BalanceType = 4
	BalanceTypeScore             BalanceType = 5
)

var BalanceTypeMap = map[BalanceType]string{
	BalanceTypeBalance:           "余额",
	BalanceTypeFreeze:            "冻结",
	BalanceTypeFxCommisionPaying: "分销佣金",
	BalanceTypeGrowth:            "成长值",
	BalanceTypeScore:             "积分",
}
