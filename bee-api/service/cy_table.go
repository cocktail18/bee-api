package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/util"
	"gorm.io/gorm"
	"sync"
	"time"
)

type CyTableSrv struct {
	BaseSrv
}

var cyTableSvcOnce sync.Once
var cyTableSvcInstance *CyTableSrv

func GetCyTableSrv() *CyTableSrv {
	cyTableSvcOnce.Do(func() {
		cyTableSvcInstance = &CyTableSrv{}
	})
	return cyTableSvcInstance
}

func (s CyTableSrv) Token(c context.Context, tableId int64, key string) (*proto.CyTableTokenResp, error) {
	var info model.BeeCyTable
	if err := db.GetDB().Where("id = ? and is_deleted = 0", tableId).Where("`key` = ?", key).Take(&info).Error; err != nil {
		return nil, enum.ErrParamError
	}
	var resp *proto.AuthorizeResp
	if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		if info.Uid == 0 {
			// 创建一个虚拟用户
			vUser := &model.BeeUser{
				BaseModel: common.BaseModel{
					UserId: info.UserId,
				},
				IsVirtual: true,
			}
			if err := GetUserSrv().CreateUser(c, tx, vUser); err != nil {
				return err
			}
			info.Uid = vUser.Id
			if err := db.GetDB().Model(&model.BeeCyTable{}).Where("id = ?", tableId).Updates(map[string]interface{}{
				"uid": info.Uid,
			}).Error; err != nil {
				return err
			}
		}
		var err error
		resp, err = GetUserSrv().CreateUserToken(c, info.UserId, info.Uid, "", "", 0)
		return err
	}); err != nil {
		return nil, err
	}
	return (*proto.CyTableTokenResp)(resp), nil
}

// CreateToken 扫码点餐需要一个虚拟用户token
func (s CyTableSrv) CreateToken(c context.Context, info *model.BeeCyTable) (string, error) {
	// 创建一个虚拟用户
	vUser := &model.BeeUser{
		BaseModel: common.BaseModel{
			UserId: info.UserId,
		},
		IsVirtual: true,
	}
	if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := GetUserSrv().CreateUser(c, tx, vUser); err != nil {
			return err
		}
		info.Uid = vUser.Id
		info.DateAdd = common.JsonTime(time.Now())
		info.DateUpdate = common.JsonTime(time.Now())
		if err := db.GetDB().Create(info).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", err
	}

	key := model.EncryptCyTableData(info)
	return key, nil
}

func (s CyTableSrv) AddOrder(c context.Context, ip string, goods []*proto.BeeOrderGoods) error {
	//@todo 参数校验，增加桌号信息
	var cyTable model.BeeCyTable
	if err := db.GetDB().Where("uid = ? and is_deleted=0", kit.GetUid(c)).Take(&cyTable).Error; err != nil {
		return err
	}
	req := &proto.CreateOrderReq{
		GoodsJsonStr: util.ToJsonWithoutErr(goods, "[]"),
		PeisongType:  "zq",
		ShopIdZt:     cyTable.ShopId,
		TableNum:     cyTable.TableNum,
		GoodsType:    "0",
	}
	_, err := GetOrderSrv().CreateOrder(c, ip, req)
	return err
}
