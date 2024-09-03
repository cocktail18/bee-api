package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
)

type BeeUserQueueService struct{}

// CreateBeeUserQueue 创建beeUserQueue表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserQueueService *BeeUserQueueService) CreateBeeUserQueue(beeUserQueue *bee.BeeUserQueue) (err error) {
	beeUserQueue.DateAdd = utils.NowPtr()
	beeUserQueue.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeUserQueue).Error
	return err
}

// DeleteBeeUserQueue 删除beeUserQueue表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserQueueService *BeeUserQueueService) DeleteBeeUserQueue(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeUserQueue{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeUserQueueByIds 批量删除beeUserQueue表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserQueueService *BeeUserQueueService) DeleteBeeUserQueueByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeUserQueue{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeUserQueue 更新beeUserQueue表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserQueueService *BeeUserQueueService) UpdateBeeUserQueue(beeUserQueue bee.BeeUserQueue, shopUserId int) (err error) {
	beeUserQueue.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeUserQueue{}).Where("id = ? and user_id = ?", beeUserQueue.Id, shopUserId).Updates(&beeUserQueue).Error
	return err
}

// GetBeeUserQueue 根据id获取beeUserQueue表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserQueueService *BeeUserQueueService) GetBeeUserQueue(id string, shopUserId int) (beeUserQueue bee.BeeUserQueue, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeUserQueue).Error
	return
}

// GetBeeUserQueueInfoList 分页获取beeUserQueue表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeUserQueueService *BeeUserQueueService) GetBeeUserQueueInfoList(info beeReq.BeeUserQueueSearch, shopUserId int) (list []bee.BeeUserQueue, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeUserQueue{})
	db = db.Where("user_id = ?", shopUserId)
	var beeUserQueues []bee.BeeUserQueue
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Uid != nil {
		db = db.Where("uid = ?", info.Uid)
	}
	if info.QueueId != nil {
		db = db.Where("queue_id = ?", info.QueueId)
	}
	if info.Number != nil {
		db = db.Where("number = ?", info.Number)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["number"] = true
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

	err = db.Find(&beeUserQueues).Error
	return beeUserQueues, total, err
}
