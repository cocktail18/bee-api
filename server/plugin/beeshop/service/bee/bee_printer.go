package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeePrinterService struct{}

// CreateBeePrinter 创建beePrinter表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePrinterService *BeePrinterService) CreateBeePrinter(beePrinter *bee.BeePrinter) (err error) {
	beePrinter.DateAdd = utils.NowPtr()
	beePrinter.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beePrinter).Error
	return err
}

// DeleteBeePrinter 删除beePrinter表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePrinterService *BeePrinterService) DeleteBeePrinter(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeePrinter{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeePrinterByIds 批量删除beePrinter表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePrinterService *BeePrinterService) DeleteBeePrinterByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeePrinter{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeePrinter 更新beePrinter表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePrinterService *BeePrinterService) UpdateBeePrinter(beePrinter bee.BeePrinter, shopUserId int) (err error) {
	beePrinter.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeePrinter{}).Where("id = ? and user_id = ?", beePrinter.Id, shopUserId).Updates(&beePrinter).Error
	return err
}

// GetBeePrinter 根据id获取beePrinter表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePrinterService *BeePrinterService) GetBeePrinter(id string, shopUserId int) (beePrinter bee.BeePrinter, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beePrinter).Error
	return
}

// GetBeePrinterInfoList 分页获取beePrinter表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePrinterService *BeePrinterService) GetBeePrinterInfoList(info beeReq.BeePrinterSearch, shopUserId int) (list []bee.BeePrinter, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeePrinter{})
	db = db.Where("user_id = ?", shopUserId)
	var beePrinters []bee.BeePrinter
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beePrinters).Error
	return beePrinters, total, err
}
