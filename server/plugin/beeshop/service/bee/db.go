package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

func GetBeeDB() *gorm.DB {
	db := global.GetGlobalDBByDBName("bee")
	if db != nil {
		return db
	}
	return global.GVA_DB
}
