package service

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/gin-gonic/gin"
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

func (s CyTableSrv) Token(c *gin.Context, tableId int64, key string) (*proto.CyTableTokenResp, error) {
	var info model.BeeCyTable
	if err := model.DecryptCyTableData(key, &info); err != nil {
		return nil, err
	}
	if info.Id != tableId {
		return nil, enum.ErrParamError
	}
	resp, err := GetUserSrv().CreateUserToken(c, info.UserId, info.Uid, "")
	if err != nil {
		return nil, err
	}
	return (*proto.CyTableTokenResp)(resp), nil
}

// CreateToken 扫码点餐需要一个虚拟用户token
func (s CyTableSrv) CreateToken(c *gin.Context, info *model.BeeCyTable) (string, error) {
	// 创建一个虚拟用户
	vUser := &model.BeeUser{
		BaseModel: common.BaseModel{
			UserId: info.UserId,
		},
		IsVirtual: true,
	}
	if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := GetUserSrv().CreateUser(tx, vUser); err != nil {
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

func (s CyTableSrv) AddOrder(c *gin.Context, goods []*proto.BeeOrderGoods) error {
	//@todo 参数校验，增加桌号信息
	var cyTable model.BeeCyTable
	if err := db.GetDB().Where("uid = ? and is_deleted=0", kit.GetUid(c)).Take(&cyTable).Error; err != nil {
		return err
	}
	_, err := GetOrderSrv().Create(c, goods, "", "", true, 0, 0, "", false)
	return err
}
