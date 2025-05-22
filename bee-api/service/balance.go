package service

import (
	"context"
	"errors"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"sync"
	"time"
)

type BalanceSrv struct {
	BaseSrv
}

var balanceSrvOnce sync.Once
var balanceSrvInstance *BalanceSrv

func GetBalanceSrv() *BalanceSrv {
	balanceSrvOnce.Do(func() {
		balanceSrvInstance = &BalanceSrv{}
	})
	return balanceSrvInstance
}

func (srv *BalanceSrv) InitAmount(ctx context.Context, uid int64) (*model.BeeUserAmount, error) {
	amount := model.BeeUserAmount{
		BaseModel:         *kit.GetInsertBaseModel(ctx),
		Uid:               uid,
		Balance:           decimal.Zero,
		Freeze:            decimal.Zero,
		FxCommisionPaying: decimal.Zero,
		Growth:            decimal.Zero,
		Score:             decimal.Zero,
		ScoreLastRound:    decimal.Zero,
		TotalPayAmount:    decimal.Zero,
		TotalPayNumber:    decimal.Zero,
		TotalScore:        decimal.Zero,
		TotalWithdraw:     decimal.Zero,
		TotalConsumed:     decimal.Zero,
		Pwd:               "",
		Salt:              "",
	}
	return &amount, db.GetDB().Create(&amount).Error
}

func (srv *BalanceSrv) GetAmount(ctx context.Context, uid int64) (*model.BeeUserAmount, error) {
	var amount model.BeeUserAmount
	err := db.GetDB().Where("uid = ?", uid).Take(&amount).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return srv.InitAmount(ctx, uid)
	}
	return &amount, err
}

func (srv *BalanceSrv) balanceType2field(t enum.BalanceType) string {
	switch t {
	case enum.BalanceTypeBalance:
		return "balance"
	case enum.BalanceTypeFreeze:
		return "freeze"
	case enum.BalanceTypeFxCommisionPaying:
		return "fx_commision_paying"
	case enum.BalanceTypeGrowth:
		return "growth"
	case enum.BalanceTypeScore:
		return "score"
	}
	panic("未定义的财产类型")
}

func (srv *BalanceSrv) OperAmountByTx(c context.Context, tx *gorm.DB, uid int64, amountType enum.BalanceType, num decimal.Decimal, orderId string, mark string, extraTx ...func(tx *gorm.DB) error) (*model.BeeUserAmount, error) {
	shouldCommit := false
	if tx == nil {
		tx = db.GetDB().Begin()
		shouldCommit = true
	}
	err := func() error {
		field := srv.balanceType2field(amountType)
		if num.IsZero() {
			// 没发生改变
		} else if num.IsNegative() {
			err := tx.Model(&model.BeeUserAmount{}).Where("uid = ? and "+field+" >= ?", uid, num.String()).
				Updates(map[string]interface{}{field: gorm.Expr(field+" + ?", num), "date_update": time.Now()}).Error
			if err != nil {
				return err
			}
		} else {
			err := tx.Model(&model.BeeUserAmount{}).Where("uid = ?", uid).
				Updates(map[string]interface{}{field: gorm.Expr(field+" + ?", num), "date_update": time.Now()}).Error
			if err != nil {
				return err
			}
		}
		for _, fun := range extraTx {
			if err := fun(tx); err != nil {
				return err
			}
		}
		return tx.Save(&model.BeeUserBalanceLog{
			BaseModel:   *kit.GetInsertBaseModel(c),
			OrderId:     orderId,
			BalanceType: amountType,
			Num:         num,
			Uid:         uid,
			Mark:        mark,
		}).Error
	}()
	if err != nil {
		if shouldCommit {
			tx.Rollback()
		}
		return nil, err
	}
	if shouldCommit {
		if err := tx.Commit().Error; err != nil {
			return nil, err
		}
	}
	return srv.GetAmount(c, uid)
}

func (srv *BalanceSrv) OperAmount(c context.Context, uid int64, amountType enum.BalanceType, num decimal.Decimal, orderId string, mark string, extraTx ...func(tx *gorm.DB) error) (*model.BeeUserAmount, error) {
	return srv.OperAmountByTx(c, nil, uid, amountType, num, orderId, mark, extraTx...)
}
