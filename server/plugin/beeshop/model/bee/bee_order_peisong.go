// 自动生成模板BeeOrderPeisong
package bee

import (
	"github.com/shopspring/decimal"
	"time"
)

// beeOrderPeisong表 结构体  BeeOrderPeisong
type BeeOrderPeisong struct {
	Id             *int            `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                                   //id字段
	UserId         *int            `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`                           //商店用户ID
	IsDeleted      *bool           `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`                             //已删除
	DateAdd        *time.Time      `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                                  //创建时间
	DateUpdate     *time.Time      `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`                         //更新时间
	DateDelete     *time.Time      `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`                         //删除时间
	OrderId        *int            `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单id;size:19;"`                          //订单id
	PeisongOrderId string          `json:"peisongOrderId" form:"peisongOrderId" gorm:"column:peisong_order_id;comment:配送订单id;size:100;"` //配送订单id
	Status         *int            `json:"status" form:"status" gorm:"column:status;comment:状态;size:19;"`                                //状态
	PeisongOrderNo string          `json:"peisongOrderNo" form:"peisongOrderNo" gorm:"column:peisong_order_no;comment:配送订单号;size:100;"`  //配送订单号
	Type           *int            `json:"type" form:"type" gorm:"column:type;comment:配送类型;size:19;"`                                    //配送类型
	Money          decimal.Decimal `json:"money" form:"money" gorm:"column:money;comment:配送费;size:10;"`                                  //配送费
	RealMoney      decimal.Decimal `json:"realMoney" form:"realMoney" gorm:"column:real_money;comment:实际支付金额;size:10;"`                  //实际支付金额
	DeductFee      decimal.Decimal `json:"deductFee" gorm:"column:deduct_fee;type:decimal(10,2);comment:取消配送扣费"`
	IsCancel       *bool           `json:"isCancel" form:"isCancel" gorm:"column:is_cancel;type:tinyint(1);comment:是否取消"`
	ErrMsg         string          `json:"errMsg" form:"errMsg" gorm:"column:err_msg;comment:配送错误信息;size:1000;"`                     //配送错误信息
	ThirdStatus    string          `json:"thirdStatus" form:"thirdStatus" gorm:"column:third_status;comment:第三方状态;size:3000;"`       //第三方状态
	RetryTimes     *int            `json:"retryTimes" form:"retryTimes" gorm:"column:retry_times;comment:重试次数;size:19;"`             //重试次数
	LastRetryUnix  *int            `json:"lastRetryUnix" form:"lastRetryUnix" gorm:"column:last_retry_unix;comment:上次重试时间;size:19;"` //上次重试时间
	ReqData        string          `json:"reqData" form:"reqData" gorm:"column:req_data;comment:请求数据;"`                              //请求数据
}

// TableName beeOrderPeisong表 BeeOrderPeisong自定义表名 bee_order_peisong
func (BeeOrderPeisong) TableName() string {
	return "bee_order_peisong"
}
