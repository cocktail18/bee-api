package beeshop

import (
	"errors"
	"fmt"
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	system_model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func initBeeDict() {

	initByMap("交易类型", "BeeCashLogType", enum.BeeCashLogTypeMap)
	initByMap("收/支", "BeeCashLogBehavior", enum.BeeCashLogBehaviorMap)
	initByMap("优惠券状态", "CouponStatus", enum.CouponStatusMap)
	initByMap("货币类型", "BalanceType", enum.BalanceTypeMap)
	initByMap("用户来源", "BeeUserSource", enum.BeeUserSourceMap)
	initByMap("订单日志类型", "BeeOrderLogType", enum.OrderLogTypeMap)
	initByMap("订单状态", "OrderStatus", enum.OrderStatusMap)
	initByMap("订单类型", "OrderType", enum.OrderTypeMap)
	initByMap("订单评价", "OrderReputation", enum.OrderReputationMap)
	initByMap("运费类型", "BeeGoodsFeeType", enum.BeeGoodsFeeTypeMap)
	initByMap("商品状态", "BeeGoodsStatus", enum.GoodsStatusStrMap)
	initByMap("商品售后类型", "BeeGoodsAfterSale", enum.GoodsAfterSaleStrMap)
	initByMap("点餐订单状态", "BeeCyTableStatus", enum.CyTableStatusStrMap)
	initByMap("订单日志类型", "BeeOrderLogType", enum.OrderLogTypeMap)
	initByMap("运费模版类型", "BeeLogisticsType", enum.BeeLogisticsTypeMap)
	initByMap("运费模版单位", "BeeLogisticsUnit", enum.BeeLogisticsUnitMap)
	initByMap("支付流水状态", "BeePayLogStatus", enum.PayLogStatusMap)
	initByMap("打印机品牌", "BeePrinterBrand", enum.PrinterBrandMap)
	initByMap("打印条件", "BeePrinterCondition", enum.PrinterConditionMap)
	initByMap("配送类型", "BeeDelivery", enum.DeliveryTypeMap)
	initByMap("配送状态", "BeeOrderPeisongStatus", enum.OrderPeisongStatusMap)
	initByMap("配送取消订单理由", "BeeDeliveryCancelReason", enum.DeliveryCancelReasonMap)
	initByMap("配送收费类型", "BeePeiSongFeeTypeMap", enum.BeePeiSongFeeTypeMap)
}

func initByMap[T1 comparable, T2 any](name string, t string, m map[T1]T2) {
	svr := &system.DictionaryService{}
	svrDetail := &system.DictionaryDetailService{}
	status := true
	dict, err := svr.GetSysDictionary(t, 0, nil)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		dict = system_model.SysDictionary{
			Name:   name,
			Type:   t,
			Status: &status,
			Desc:   name + "字典",
		}
		err = svr.CreateSysDictionary(dict)
	}
	if err != nil {
		panic(err)
	}
	dict, err = svr.GetSysDictionary(t, 0, nil)
	if err != nil {
		panic(err)
	}

	_ = svr.UpdateSysDictionary(&system_model.SysDictionary{
		GVA_MODEL: global.GVA_MODEL{ID: dict.ID},
		Name:      name,
		Type:      t,
		Status:    &status,
		Desc:      name + "字典",
	})
	for key, value := range m {
		item, _ := svrDetail.GetDictionaryInfoByTypeValue(t, fmt.Sprintf("%d", key))
		if item.ID == 0 {
			_ = svrDetail.CreateSysDictionaryDetail(system_model.SysDictionaryDetail{
				Label:           cast.ToString(value),
				Value:           fmt.Sprintf("%d", key),
				Extend:          "",
				Status:          &status,
				Sort:            0,
				SysDictionaryID: int(dict.ID),
			})
		} else {
			item.Label = cast.ToString(value)
			item.Value = fmt.Sprintf("%d", key)
			_ = svrDetail.UpdateSysDictionaryDetail(&item)
		}
	}
}
