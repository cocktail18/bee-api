package model

import (
	"encoding/json"
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/util"
)

type BeeCyTable struct {
	common.BaseModel
	ShopId   int64  `gorm:"column:shop_id;type:bigint(20)" json:"shopId"`
	TableNum string `gorm:"column:table_num;type:varchar(255)" json:"tableNum"` // 桌号
	Remark   string `gorm:"column:remark;type:varchar(255)" json:"remark"`      // 备注
	Uid      int64  `gorm:"column:uid;type:bigint(20)" json:"uid"`              // 虚拟用户id
}

func (m *BeeCyTable) TableName() string {
	return "bee_cy_table"
}

var key = []byte("a95RNzFJ1nbi") // 16, 24 or 32 bytes long key for AES-128, AES-192 or AES-256 respectively

func EncryptCyTableData(info *BeeCyTable) string {
	bs := util.ToJsonWithoutErr(info, "{}")
	return util.AesEncrypt(key, []byte(bs))
}

func DecryptCyTableData(data string, info *BeeCyTable) error {
	bs, err := util.Decrypt(key, []byte(data))
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, info)
}
