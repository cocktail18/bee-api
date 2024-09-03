package bee

import (
	"gitee.com/stuinfer/bee-api/service"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
	"github.com/spf13/cast"
)

type BeeQueueService struct{}

// CreateBeeQueue 创建beeQueue表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeQueueService *BeeQueueService) CreateBeeQueue(beeQueue *bee.BeeQueue) (err error) {
	beeQueue.DateAdd = utils.NowPtr()
	beeQueue.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeQueue).Error
	return err
}

// DeleteBeeQueue 删除beeQueue表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeQueueService *BeeQueueService) DeleteBeeQueue(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeQueue{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeQueueByIds 批量删除beeQueue表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeQueueService *BeeQueueService) DeleteBeeQueueByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeQueue{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeQueue 更新beeQueue表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeQueueService *BeeQueueService) UpdateBeeQueue(beeQueue bee.BeeQueue, shopUserId int) (err error) {
	beeQueue.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeQueue{}).Where("id = ? and user_id = ?", beeQueue.Id, shopUserId).Updates(&beeQueue).Error
	return err
}

// GetBeeQueue 根据id获取beeQueue表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeQueueService *BeeQueueService) GetBeeQueue(id string, shopUserId int) (beeQueue bee.BeeQueue, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeQueue).Error
	return
}

// GetBeeQueueInfoList 分页获取beeQueue表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeQueueService *BeeQueueService) GetBeeQueueInfoList(info beeReq.BeeQueueSearch, shopUserId int) (list []bee.BeeQueue, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeQueue{})
	db = db.Where("user_id = ?", shopUserId)
	var beeQueues []bee.BeeQueue
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Id != nil {
		db = db.Where("id = ?", info.Id)
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

	err = db.Find(&beeQueues).Error
	return beeQueues, total, err
}

func (beeQueueService *BeeQueueService) CallBeeQueue(queue bee.BeeQueue, shopUserId int) error {
	ctx, err := getContext(shopUserId)
	if err != nil {
		return err
	}
	return service.GetQueueSrv().CallQueue(ctx, cast.ToInt64(queue.Id), cast.ToInt64(queue.CurNumber))
}

func (beeQueueService *BeeQueueService) PassBeeQueue(queue bee.BeeQueue, shopUserId int) error {
	ctx, err := getContext(shopUserId)
	if err != nil {
		return err
	}
	return service.GetQueueSrv().PassQueue(ctx, cast.ToInt64(queue.Id))
}

func (beeQueueService *BeeQueueService) NextBeeQueue(queue bee.BeeQueue, shopUserId int) error {
	ctx, err := getContext(shopUserId)
	if err != nil {
		return err
	}
	return service.GetQueueSrv().NextBeeQueue(ctx, cast.ToInt64(queue.Id))
}

func (beeQueueService *BeeQueueService) ReCallBeeQueue(queue bee.BeeQueue, shopUserId int) error {
	ctx, err := getContext(shopUserId)
	if err != nil {
		return err
	}
	return service.GetQueueSrv().ReCallBeeQueue(ctx, cast.ToInt64(queue.Id))
}
