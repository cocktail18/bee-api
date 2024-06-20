package enum

type BalanceType int

const (
	BalanceTypeBalance           BalanceType = 1
	BalanceTypeFreeze            BalanceType = 2
	BalanceTypeFxCommisionPaying BalanceType = 3
	BalanceTypeGrowth            BalanceType = 4
	BalanceTypeScore             BalanceType = 5
)
