package service

import (
	"bytes"
	"context"
	"encoding/json"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/printer"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"sync"
	"text/template"
	"time"
)

type PrinterSrv struct {
	BaseSrv
}

var printerSrvOnce sync.Once
var printerSrvInstance *PrinterSrv

func GetPrinterSrv() *PrinterSrv {
	printerSrvOnce.Do(func() {
		printerSrvInstance = &PrinterSrv{}
	})
	return printerSrvInstance
}

func (srv *PrinterSrv) printOrder(ctx context.Context, item *model.BeeOrderPrintLog) error {
	//获取打印机信息
	var printers []*model.BeePrinter
	if err := db.GetDB().Where("user_id = ? and condition = ? and is_deleted=0 and `condition`=?", item.UserId, item.Condition).Find(&printers).Error; err != nil {
		return err
	}
	if len(printers) == 0 {
		logger.GetLogger().Debug("未配置打印机")
		return nil
	}
	orderDto, err := GetOrderSrv().GetOrderDetailByOrderId(ctx, item.OrderId)
	if err != nil {
		return err
	}
	userBalance, err := GetBalanceSrv().GetAmount(item.UserId)
	if err != nil {
		return err
	}

	userInfo, err := GetUserSrv().GetUserInfoByUid(ctx, item.UserId)
	if err != nil {
		return err
	}
	var extJson map[string]interface{}
	if orderDto.ExtJsonStr != "" {
		_ = json.Unmarshal([]byte(orderDto.ExtJsonStr), &extJson)
	}
	var nowStr = time.Now().Format("2006-01-02 15:04:05")
	for _, _printer := range printers {
		if _printer.ShopId > 0 && orderDto.ShopId != _printer.ShopId {
			logger.GetLogger().Debug("订单不属于该打印机", zap.Any("printer", _printer), zap.Any("order", orderDto))
			continue
		}

		content, err := srv.buildContent(ctx, _printer.Template, map[string]interface{}{
			"order":     orderDto,
			"userCash":  userBalance,
			"user":      userInfo,
			"logistics": orderDto.OrderLogistics,
			"goods":     orderDto.OrderGoods,
			"extJson":   extJson,
			"nowStr":    nowStr,
		})
		if err != nil {
			return err
		}
		if err := printer.GetPrinter(_printer).Print(_printer, content); err != nil {
			return err
		}
	}
	return nil
}

func (srv *PrinterSrv) buildContent(ctx context.Context, tpl string, data any) (string, error) {
	t := template.New("tmp")
	tplInstance, err := t.Parse(tpl)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)
	if err := tplInstance.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (srv *PrinterSrv) StartDaemon() {
	go func() {
		ticker := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-ticker.C:
				var items []*model.BeeOrderPrintLog
				if err := db.GetDB().Where("status = ? and is_deleted=0 and last_try_unix < ?",
					enum.OrderPrinterStatusWaiting, time.Now().Add(time.Second*-10).Unix()).Find(&items).Error; err != nil {
					logger.GetLogger().Error("查询待打印订单失败", zap.Error(err))
					continue
				}
				for _, item := range items {
					ctx := context.Background()
					if err := srv.printOrder(ctx, item); err != nil {
						logger.GetLogger().Error("打印订单失败", zap.Error(err))
						db.GetDB().Model(&model.BeeOrderPrintLog{}).Where("id = ?", item.Id).Updates(map[string]interface{}{
							"last_try_unix": time.Now().Unix(),
							"try_times":     item.TryTimes + 1,
							"err_msg":       lo.Substring(err.Error(), 0, 100),
							"status":        util.IF(item.TryTimes+1 >= 3, enum.OrderPrinterStatusErr, enum.OrderPrinterStatusWaiting),
						})
					} else {
						logger.GetLogger().Info("打印订单成功", zap.Any("item", item))
						db.GetDB().Model(&model.BeeOrderPrintLog{}).Where("id = ?", item.Id).Updates(map[string]interface{}{
							"last_try_unix": time.Now().Unix(),
							"try_times":     item.TryTimes + 1,
							"status":        enum.OrderPrinterStatusDone,
						})
					}
				}
			}
		}
	}()
}
