package bee

import (
	"context"
	"fmt"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
	"github.com/medivhzhan/weapp/v3"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

type BeeCyTableService struct{}

// CreateBeeCyTable 创建桌号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCyTableService *BeeCyTableService) CreateBeeCyTable(beeCyTable *bee.BeeCyTable) (err error) {
	beeCyTable.Uid = nil
	beeCyTable.Key = lo.RandomString(10, lo.LettersCharset)
	beeCyTable.DateAdd = utils.NowPtr()
	beeCyTable.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeCyTable).Error
	return err
}

// DeleteBeeCyTable 删除桌号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCyTableService *BeeCyTableService) DeleteBeeCyTable(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeCyTable{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeCyTableByIds 批量删除桌号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCyTableService *BeeCyTableService) DeleteBeeCyTableByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeCyTable{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeCyTable 更新桌号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCyTableService *BeeCyTableService) UpdateBeeCyTable(beeCyTable bee.BeeCyTable, shopUserId int) (err error) {
	beeCyTable.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeCyTable{}).Where("id = ? and user_id = ?", beeCyTable.Id, shopUserId).Updates(&beeCyTable).Error
	return err
}

// GetBeeCyTable 根据id获取桌号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCyTableService *BeeCyTableService) GetBeeCyTable(id string, shopUserId int) (beeCyTable bee.BeeCyTable, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeCyTable).Error
	return
}

// GetBeeCyTableInfoList 分页获取桌号信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCyTableService *BeeCyTableService) GetBeeCyTableInfoList(info beeReq.BeeCyTableSearch, shopUserId int) (list []bee.BeeCyTable, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeCyTable{})
	db = db.Where("user_id = ?", shopUserId)
	var beeCyTables []bee.BeeCyTable
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ShopId != nil {
		db = db.Where("shop_id = ?", info.ShopId)
	}
	if info.TableNum != "" {
		db = db.Where("table_num = ?", info.TableNum)
	}
	if info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+info.Remark+"%")
	}
	if info.Uid != nil {
		db = db.Where("uid = ?", info.Uid)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["date_add"] = true
	orderMap["date_update"] = true
	orderMap["shop_id"] = true
	orderMap["table_num"] = true
	orderMap["uid"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeCyTables).Error
	return beeCyTables, total, err
}

func (beeCyTableService *BeeCyTableService) GetBeeCyTableQrCode(id string, shopUserId int) ([]byte, error) {
	info, err := beeCyTableService.GetBeeCyTable(id, shopUserId)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	wxAppSrv, err := service.GetWxAppSrvByUserId(ctx, int64(shopUserId))
	if err != nil {
		return nil, err
	}
	return wxAppSrv.GetWxAppQrCode(ctx, &weapp.UnlimitedQRCode{
		Scene:      fmt.Sprintf("shopId=%d,id=%d,key=%s", cast.ToInt64(info.ShopId), cast.ToInt64(info.Id), info.Key),
		Page:       "pages/index/index",
		CheckPath:  true,
		EnvVersion: "",
		Width:      1280,
		AutoColor:  true,
		LineColor:  weapp.Color{},
		IsHyaline:  false,
	})
}
