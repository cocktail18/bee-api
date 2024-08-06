package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeCommentService struct{}

// CreateBeeComment 创建评论表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCommentService *BeeCommentService) CreateBeeComment(beeComment *bee.BeeComment) (err error) {
	beeComment.DateAdd = utils.NowPtr()
	beeComment.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeComment).Error
	return err
}

// DeleteBeeComment 删除评论表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCommentService *BeeCommentService) DeleteBeeComment(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeComment{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeCommentByIds 批量删除评论表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCommentService *BeeCommentService) DeleteBeeCommentByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeComment{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeComment 更新评论表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCommentService *BeeCommentService) UpdateBeeComment(beeComment bee.BeeComment, shopUserId int) (err error) {
	beeComment.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeComment{}).Where("id = ? and user_id = ?", beeComment.Id, shopUserId).Updates(&beeComment).Error
	return err
}

// GetBeeComment 根据id获取评论表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCommentService *BeeCommentService) GetBeeComment(id string, shopUserId int) (beeComment bee.BeeComment, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeComment).Error
	return
}

// GetBeeCommentInfoList 分页获取评论表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeCommentService *BeeCommentService) GetBeeCommentInfoList(info beeReq.BeeCommentSearch, shopUserId int) (list []bee.BeeComment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeComment{})
	db = db.Where("user_id = ?", shopUserId)
	var beeComments []bee.BeeComment
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartDateAdd != nil && info.EndDateAdd != nil {
		db = db.Where("date_add BETWEEN ? AND ? ", info.StartDateAdd, info.EndDateAdd)
	}
	if info.StartDateUpdate != nil && info.EndDateUpdate != nil {
		db = db.Where("date_update BETWEEN ? AND ? ", info.StartDateUpdate, info.EndDateUpdate)
	}
	if info.Uid != nil {
		db = db.Where("uid = ?", info.Uid)
	}
	if info.RefId != nil {
		db = db.Where("ref_id = ?", info.RefId)
	}
	if info.ShopId != nil {
		db = db.Where("shop_id = ?", info.ShopId)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.Mobile != "" {
		db = db.Where("mobile = ?", info.Mobile)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["uid"] = true
	orderMap["ref_id"] = true
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

	err = db.Find(&beeComments).Error
	return beeComments, total, err
}
