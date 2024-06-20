package service

import (
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/gin-gonic/gin"
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

func (srv *BalanceSrv) OperAmount(c *gin.Context, userId int64, amountType enum.BalanceType, num decimal.Decimal, orderId string, mark string) (*model.BeeUserAmount, error) {
	var amount model.BeeUserAmount

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		field := srv.balanceType2field(amountType)
		if num.IsZero() {
			// 没发生改变
		} else if num.IsNegative() {
			err := tx.Where("user_id = ? and "+field+" >= ?", userId, num.String()).
				Updates(map[string]interface{}{field: gorm.Expr(field+" + ?", num), "updated_at": time.Now()}).Error
			if err != nil {
				return err
			}
		} else {
			err := tx.Where("user_id = ?", userId).
				Updates(map[string]interface{}{field: gorm.Expr(field+" + ?", num), "updated_at": time.Now()}).Error
			if err != nil {
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
	})
	return &amount, err
}
