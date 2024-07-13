package service

import (
	"context"
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

func (srv *BalanceSrv) GetAmount(userId int64) (*model.BeeUserAmount, error) {
	var amount model.BeeUserAmount
	err := db.GetDB().Where("uid = ?", userId).Take(&amount).Error
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

func (srv *BalanceSrv) OperAmountByTx(c context.Context, tx *gorm.DB, userId int64, amountType enum.BalanceType, num decimal.Decimal, orderId string, mark string, extraTx ...func(tx *gorm.DB) error) (*model.BeeUserAmount, error) {
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
			err := tx.Model(&model.BeeUserAmount{}).Where("user_id = ? and "+field+" >= ?", userId, num.String()).
				Updates(map[string]interface{}{field: gorm.Expr(field+" + ?", num), "date_update": time.Now()}).Error
			if err != nil {
				return err
			}
		} else {
			err := tx.Model(&model.BeeUserAmount{}).Where("user_id = ?", userId).
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
			Uid:         userId,
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
		tx.Commit()
	}
	return srv.GetAmount(userId)
}

func (srv *BalanceSrv) OperAmount(c context.Context, userId int64, amountType enum.BalanceType, num decimal.Decimal, orderId string, mark string, extraTx ...func(tx *gorm.DB) error) (*model.BeeUserAmount, error) {
	return srv.OperAmountByTx(c, nil, userId, amountType, num, orderId, mark, extraTx...)
}
