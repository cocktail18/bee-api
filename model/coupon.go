package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/shopspring/decimal"
	"time"
)

// 用户优惠券列表
type BeeUserCoupon struct {
	common.BaseModel
	Uid           int64                `gorm:"column:uid;type:bigint(20);comment:用户id" json:"uid"`
	DateStart     common.JsonTime      `gorm:"column:date_start;type:datetime(3);comment:可以使用时间" json:"dateStart"`
	ExpiryMillis  int64                `gorm:"column:expiry_millis;type:bigint(11);comment:过期时间（毫秒）" json:"expiryMillis"`
	Money         decimal.Decimal      `gorm:"column:money;type:decimal(20,2);comment:金额" json:"money"`
	MoneyHreshold decimal.Decimal      `gorm:"column:money_hreshold;type:bigint(11);comment:满xx可用" json:"moneyHreshold"`
	MoneyMin      decimal.Decimal      `gorm:"column:money_min;type:bigint(11);comment:优惠券面额范围" json:"moneyMin"`
	MoneyMax      decimal.Decimal      `gorm:"column:money_max;type:bigint(11);comment:优惠券面额范围" json:"moneyMax"`
	MoneyType     enum.CouponMoneyType `gorm:"column:money_type;type:bigint(11);comment:优惠券类型" json:"moneyType"`
	Name          string               `gorm:"column:name;type:varchar(100);comment:优惠券名字" json:"name"`
	OnlyFreight   bool                 `gorm:"column:only_freight;type:bool;comment:仅抵扣运费" json:"onlyFreight"`
	Pid           int64                `gorm:"column:pid;type:bigint(11);comment:优惠券id" json:"pid"`
	ShopId        int64                `gorm:"column:shop_id;type:bigint(11);comment:店铺id" json:"shopId"`
	Source        string               `gorm:"column:source;type:varchar(100);comment:管理员批量发放" json:"source"`
	Status        enum.CouponStatus    `gorm:"column:status;type:int(11);comment:使用状态，0未使用，1使用中，2已使用" json:"status"`
}

func (m *BeeUserCoupon) TableName() string {
	return "bee_user_coupon"
}

func (m *BeeUserCoupon) CanUse(amount decimal.Decimal) bool {
	now := time.Now()
	dateStart := time.Time(m.DateStart)
	if amount.LessThanOrEqual(decimal.Zero) {
		return false
	}
	if !(now.After(dateStart) && time.UnixMilli(m.ExpiryMillis).After(now)) {
		return false
	}
	if amount.LessThan(m.MoneyHreshold) && !m.MoneyHreshold.Equal(decimal.Zero) {
		return false
	}
	if m.Status != enum.CouponStatusNormal {
		return false
	}
	return true
}

func (m *BeeUserCoupon) IsExpire() bool {
	now := time.Now()
	return time.UnixMilli(m.ExpiryMillis).Before(now)
}

// 优惠券
type BeeCoupon struct {
	common.BaseModel
	BatchSendUid         int64                  `gorm:"column:batch_send_uid;type:bigint(11);comment:批量" json:"batchSendUid"`
	DateEndDays          int                    `gorm:"column:date_end_days;type:int(11);comment:第n天失效" json:"dateEndDays"`
	DateEnd              common.JsonTime        `gorm:"column:date_end;type:datetime(3);comment:失效日期;default:null" json:"dateEnd"`
	DateEndType          enum.CouponDateEndType `gorm:"column:date_end_type;type:bigint(11);comment:失效类型，1n天后失效，0指定" json:"dateEndType"`
	DateStart            common.JsonTime        `gorm:"column:date_start;type:datetime(3);comment:生效日期;default:null" json:"dateStart"`
	DateStartDays        int64                  `gorm:"column:date_start_days;type:bigint(11);comment:第n天生效" json:"dateStartDays"`
	DateStartType        int                    `gorm:"column:date_start_type;type:bigint(11);comment:生效时间，1立即，2次日，0指定时间" json:"dateStartType"`
	MoneyHreshold        decimal.Decimal        `gorm:"column:money_hreshold;type:decimal(20,2);comment:满xx可用" json:"moneyHreshold"`
	MoneyMax             decimal.Decimal        `gorm:"column:money_max;type:decimal(20,2);comment:优惠券最大金额" json:"moneyMax"`
	MoneyMin             decimal.Decimal        `gorm:"column:money_min;type:decimal(20,2);comment:优惠券最小金额" json:"moneyMin"`
	MoneyType            enum.CouponMoneyType   `gorm:"column:money_type;type:bigint(11);comment:类型，1满xx，2折扣" json:"moneyType"`
	Name                 string                 `gorm:"column:name;type:varchar(100);comment:优惠券名字" json:"name"`
	NeedAmount           decimal.Decimal        `gorm:"column:need_amount;type:decimal(20,2);comment:需要支付" json:"needAmount"`
	NeedScore            decimal.Decimal        `gorm:"column:need_score;type:decimal(20,2);comment:需要积分" json:"needScore"`
	NeedSignedContinuous int64                  `gorm:"column:need_signed_continuous;type:bigint(11);comment:连续签到多少天可得" json:"needSignedContinuous"`
	NumberLeft           int64                  `gorm:"column:number_left;type:bigint(11);comment:剩余数量" json:"numberLeft"`
	NumberPersonMax      int64                  `gorm:"column:number_person_max;type:bigint(11);comment:每人最多领取多少" json:"numberPersonMax"`
	NumberTotle          int64                  `gorm:"column:number_totle;type:bigint(11);comment:总数量" json:"numberTotle"`
	NumberUsed           int64                  `gorm:"column:number_used;type:bigint(11);comment:使用数量" json:"numberUsed"`
	OnlyFreight          bool                   `gorm:"column:only_freight;type:tinyint(1);comment:仅抵扣运费" json:"onlyFreight"`
	SendBirthday         bool                   `gorm:"column:send_birthday;type:tinyint(1);comment:生日赠送" json:"sendBirthday"`
	SendInviteM          bool                   `gorm:"column:send_invite_m;type:tinyint(1);comment:邀请新人注册赠送" json:"sendInviteM"`
	SendInviteS          bool                   `gorm:"column:send_invite_s;type:tinyint(1);comment:被邀请赠送" json:"sendInviteS"`
	SendRegister         bool                   `gorm:"column:send_register;type:tinyint(1);comment:注册赠送" json:"sendRegister"`
	ShopId               int64                  `gorm:"column:shop_id;type:bigint(11);comment:店铺id" json:"shopId"`
	ShowInFront          bool                   `gorm:"column:show_in_front;type:tinyint(1);comment:展示" json:"showInFront"`
	Status               int64                  `gorm:"column:status;type:bigint(11);comment:状态，0是正常" json:"status"`
	Type                 string                 `gorm:"column:type;type:varchar(100);comment:优惠券类型描述" json:"type"`
}

func (m *BeeCoupon) TableName() string {
	return "bee_coupon"
}

type BeeUserCouponLog struct {
	common.BaseModel
	CouponId int64              `gorm:"column:coupon_id;type:bigint(11);comment:优惠券id" json:"couponId"`
	Uid      int64              `gorm:"column:uid;type:bigint(11);comment:用户id" json:"uid"`
	Uniq     string             `gorm:"column:uniq;type:varchar(100);comment:唯一值去重" json:"uniq"`
	Remark   string             `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark"`
	Type     enum.CouponLogType `gorm:"column:type;type:bigint(11);comment:类型，1发放，2使用" json:"type"`
}

func (m *BeeUserCouponLog) TableName() string {
	return "bee_user_coupon_log"
}
