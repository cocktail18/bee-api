package bee

import (
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
	"github.com/spf13/cast"
)

type BeeOrderPeisongService struct{}

// CreateBeeOrderPeisong 创建beeOrderPeisong表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongService *BeeOrderPeisongService) CreateBeeOrderPeisong(beeOrderPeisong *bee.BeeOrderPeisong) (err error) {
	beeOrderPeisong.DateAdd = utils.NowPtr()
	beeOrderPeisong.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Create(beeOrderPeisong).Error
	return err
}

// DeleteBeeOrderPeisong 删除beeOrderPeisong表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongService *BeeOrderPeisongService) DeleteBeeOrderPeisong(id string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeOrderPeisong{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeOrderPeisongByIds 批量删除beeOrderPeisong表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongService *BeeOrderPeisongService) DeleteBeeOrderPeisongByIds(ids []string, shopUserId int) (err error) {
	err = GetBeeDB().Model(&bee.BeeOrderPeisong{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeOrderPeisong 更新beeOrderPeisong表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongService *BeeOrderPeisongService) UpdateBeeOrderPeisong(beeOrderPeisong bee.BeeOrderPeisong, shopUserId int) (err error) {
	beeOrderPeisong.DateUpdate = utils.NowPtr()
	err = GetBeeDB().Model(&bee.BeeOrderPeisong{}).Where("id = ? and user_id = ?", beeOrderPeisong.Id, shopUserId).Updates(&beeOrderPeisong).Error
	return err
}

// Cancel 取消配送
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongService *BeeOrderPeisongService) Cancel(beeOrderPeisong beeReq.BeeOrderPeisongCancel, shopUserId int) (err error) {
	ctx, err := getContext(shopUserId)
	if err != nil {
		return err
	}
	return service.GetOrderSrv().CancelDelivery(ctx, cast.ToInt64(beeOrderPeisong.Id), beeOrderPeisong.ReasonId, beeOrderPeisong.Reason)
}

// Notify 通知供应商配送
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongService *BeeOrderPeisongService) Notify(beeOrderPeisong bee.BeeOrderPeisong, shopUserId int) (err error) {
	ctx, err := getContext(shopUserId)
	if err != nil {
		return err
	}
	return service.GetOrderSrv().NotifyDelivery(ctx, cast.ToInt64(beeOrderPeisong.OrderId))
}

// GetBeeOrderPeisong 根据id获取beeOrderPeisong表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongService *BeeOrderPeisongService) GetBeeOrderPeisong(id string, shopUserId int) (beeOrderPeisong bee.BeeOrderPeisong, err error) {
	err = GetBeeDB().Where("id = ? and user_id = ?", id, shopUserId).First(&beeOrderPeisong).Error
	return
}

// GetBeeOrderPeisongInfoList 分页获取beeOrderPeisong表记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeOrderPeisongService *BeeOrderPeisongService) GetBeeOrderPeisongInfoList(info beeReq.BeeOrderPeisongSearch, shopUserId int) (list []bee.BeeOrderPeisong, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := GetBeeDB().Model(&bee.BeeOrderPeisong{})
	db = db.Where("user_id = ?", shopUserId)
	var beeOrderPeisongs []bee.BeeOrderPeisong
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderId != nil {
		db = db.Where("order_id = ?", info.OrderId)
	}
	if info.PeisongOrderId != "" {
		db = db.Where("peisong_order_id = ?", info.PeisongOrderId)
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

	err = db.Find(&beeOrderPeisongs).Error
	return beeOrderPeisongs, total, err
}

func (beeOrderPeisongService *BeeOrderPeisongService) GetBeeOrderPeisongDetail(peisongId int, shopUserId int) (*proto.QueryDeliveryResult, error) {
	ctx, err := getContext(shopUserId)
	if err != nil {
		return nil, err
	}
	return service.GetOrderSrv().GetPeisongDetail(ctx, int64(peisongId))
}
