package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeePeiSongService struct{}

// CreateBeePeiSong 创建配送信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePeiSongService *BeePeiSongService) CreateBeePeiSong(beePeiSong *bee.BeePeiSong) (err error) {
	beePeiSong.DateAdd = utils.NowPtr()
	beePeiSong.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beePeiSong).Error
	return err
}

// DeleteBeePeiSong 删除配送信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePeiSongService *BeePeiSongService) DeleteBeePeiSong(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeePeiSong{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeePeiSongByIds 批量删除配送信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePeiSongService *BeePeiSongService) DeleteBeePeiSongByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeePeiSong{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeePeiSong 更新配送信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePeiSongService *BeePeiSongService) UpdateBeePeiSong(beePeiSong bee.BeePeiSong, shopUserId int) (err error) {
	beePeiSong.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeePeiSong{}).Where("id = ? and user_id = ?", beePeiSong.Id, shopUserId).Updates(&beePeiSong).Error
	return err
}

// GetBeePeiSong 根据id获取配送信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePeiSongService *BeePeiSongService) GetBeePeiSong(id string, shopUserId int) (beePeiSong bee.BeePeiSong, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beePeiSong).Error
	return
}

// GetBeePeiSongInfoList 分页获取配送信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (beePeiSongService *BeePeiSongService) GetBeePeiSongInfoList(info beeReq.BeePeiSongSearch, shopUserId int) (list []bee.BeePeiSong, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeePeiSong{})
	db = db.Where("user_id = ?", shopUserId)
	var beePeiSongs []bee.BeePeiSong
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Fwf1Type != nil {
		db = db.Where("fwf1_type = ?", info.Fwf1Type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
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

	err = db.Find(&beePeiSongs).Error
	return beePeiSongs, total, err
}
