package common

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"strings"
	"time"
)

type BaseModel struct {
	Id         int64        `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	UserId     int64        `gorm:"column:user_id;type:bigint(20);<-:create" json:"userId"` //主账号uid
	IsDeleted  bool         `gorm:"column:is_deleted;type:tinyint(1)" json:"isDeleted"`
	DateAdd    JsonTime     `gorm:"column:date_add;type:datetime(3);default:null;<-:create" json:"dateAdd"`
	DateUpdate JsonTime     `gorm:"column:date_update;type:datetime(3);default:null" json:"dateUpdate"`
	DateDelete sql.NullTime `gorm:"column:date_delete;type:datetime(3)" json:"dateDelete"`
}

// 定义个类型别名
type JsonTime time.Time

func (m *JsonTime) ToString() string {
	return time.Time(*m).Format(carbon.DateTimeLayout)
}

// 实现它的json序列化方法
func (m *JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(*m).Format(carbon.DateTimeLayout)) // Format内即是你想转换的格式
	return []byte(stamp), nil
}

// 实现json.Unmarshaler接口
func (m *JsonTime) UnmarshalJSON(data []byte) error {
	dataStr := string(data)
	item := carbon.Parse(strings.Trim(dataStr, "\""), carbon.Local)
	*m = JsonTime(item.StdTime())
	return nil
}

func (m JsonTime) Value() (driver.Value, error) {
	if time.Time(m).IsZero() {
		return nil, nil
	}
	return time.Time(m), nil
}

func (m *JsonTime) Scan(value interface{}) error {
	var result sql.NullTime
	if err := result.Scan(value); err != nil {
		return err
	}
	*m = JsonTime(result.Time)
	return nil
}
