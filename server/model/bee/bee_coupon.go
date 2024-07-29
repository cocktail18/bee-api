// 自动生成模板BeeCoupon
package bee

import (
	"time"
)

// 优惠券 结构体  BeeCoupon
type BeeCoupon struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    BatchSendUid  *int `json:"batchSendUid" form:"batchSendUid" gorm:"column:batch_send_uid;comment:批量;size:19;"`  //批量 
    DateEndDays  *int `json:"dateEndDays" form:"dateEndDays" gorm:"column:date_end_days;comment:第n天失效;size:19;"`  //第n天失效 
    DateEnd  *time.Time `json:"dateEnd" form:"dateEnd" gorm:"column:date_end;comment:失效日期;"`  //失效日期 
    DateEndType  *int `json:"dateEndType" form:"dateEndType" gorm:"column:date_end_type;comment:失效类型，1n天后失效，0指定;size:19;"`  //失效类型，1n天后失效，0指定 
    DateStart  *time.Time `json:"dateStart" form:"dateStart" gorm:"column:date_start;comment:生效日期;"`  //生效日期 
    DateStartDays  *int `json:"dateStartDays" form:"dateStartDays" gorm:"column:date_start_days;comment:第n天生效;size:19;"`  //第n天生效 
    DateStartType  *int `json:"dateStartType" form:"dateStartType" gorm:"column:date_start_type;comment:生效时间，1立即，2次日，0指定时间;size:19;"`  //生效时间，1立即，2次日，0指定时间 
    MoneyHreshold  *float64 `json:"moneyHreshold" form:"moneyHreshold" gorm:"column:money_hreshold;comment:满xx可用;size:20;"`  //满xx可用 
    MoneyMax  *float64 `json:"moneyMax" form:"moneyMax" gorm:"column:money_max;comment:优惠券最大金额;size:20;"`  //优惠券最大金额 
    MoneyMin  *float64 `json:"moneyMin" form:"moneyMin" gorm:"column:money_min;comment:优惠券最小金额;size:20;"`  //优惠券最小金额 
    MoneyType  *int `json:"moneyType" form:"moneyType" gorm:"column:money_type;comment:类型，1满xx，2折扣;size:19;"`  //类型，1满xx，2折扣 
    Name  string `json:"name" form:"name" gorm:"column:name;comment:优惠券名字;size:100;"`  //优惠券名字 
    NeedAmount  *float64 `json:"needAmount" form:"needAmount" gorm:"column:need_amount;comment:需要支付;size:20;"`  //需要支付 
    NeedScore  *float64 `json:"needScore" form:"needScore" gorm:"column:need_score;comment:需要积分;size:20;"`  //需要积分 
    NeedSignedContinuous  *int `json:"needSignedContinuous" form:"needSignedContinuous" gorm:"column:need_signed_continuous;comment:连续签到多少天可得;size:19;"`  //连续签到多少天可得 
    NumberLeft  *int `json:"numberLeft" form:"numberLeft" gorm:"column:number_left;comment:剩余数量;size:19;"`  //剩余数量 
    NumberPersonMax  *int `json:"numberPersonMax" form:"numberPersonMax" gorm:"column:number_person_max;comment:每人最多领取多少;size:19;"`  //每人最多领取多少 
    NumberTotle  *int `json:"numberTotle" form:"numberTotle" gorm:"column:number_totle;comment:总数量;size:19;"`  //总数量 
    NumberUsed  *int `json:"numberUsed" form:"numberUsed" gorm:"column:number_used;comment:使用数量;size:19;"`  //使用数量 
    OnlyFreight  *bool `json:"onlyFreight" form:"onlyFreight" gorm:"column:only_freight;comment:仅抵扣运费;"`  //仅抵扣运费 
    SendBirthday  *bool `json:"sendBirthday" form:"sendBirthday" gorm:"column:send_birthday;comment:生日赠送;"`  //生日赠送 
    SendInviteM  *bool `json:"sendInviteM" form:"sendInviteM" gorm:"column:send_invite_m;comment:邀请新人注册赠送;"`  //邀请新人注册赠送 
    SendInviteS  *bool `json:"sendInviteS" form:"sendInviteS" gorm:"column:send_invite_s;comment:被邀请赠送;"`  //被邀请赠送 
    SendRegister  *bool `json:"sendRegister" form:"sendRegister" gorm:"column:send_register;comment:注册赠送;"`  //注册赠送 
    ShopId  *int `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:店铺id;size:19;"`  //店铺id 
    ShowInFront  *bool `json:"showInFront" form:"showInFront" gorm:"column:show_in_front;comment:展示;"`  //展示 
    Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态，0是正常;size:19;"`  //状态，0是正常 
    Type  string `json:"type" form:"type" gorm:"column:type;comment:优惠券类型描述;size:100;"`  //优惠券类型描述 
}


// TableName 优惠券 BeeCoupon自定义表名 bee_coupon
func (BeeCoupon) TableName() string {
    return "bee_coupon"
}

