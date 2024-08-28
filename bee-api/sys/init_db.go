package sys

import (
	"gitee.com/stuinfer/bee-api/common"
	config2 "gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/model/sys"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

func InitDB() {
	allModel := []interface{}{
		&sys.SysUserModel{},
	}
	if err := db.GetDB().AutoMigrate(allModel...); err != nil {
		panic(err)
	}
	if config2.AppConfigIns.DB != nil && config2.AppConfigIns.DB.Drop {
		logger.GetLogger().Info("清空数据库")
		for _, mod := range allModel {
			if err := db.GetDB().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(mod).Error; err != nil {
				panic(err)
			}
		}
		logger.GetLogger().Info("清空数据库成功")
	}
	if err := InitDemoData(); err != nil {
		panic(err)
	}
}

func InitDemoData() error {
	if config2.AppConfigIns.Default.SysUser != nil {
		_domain := config2.AppConfigIns.Default.SysUser.Domain
		userDefault, err := userSrv.GetByDomain(_domain)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userDefault, err = userSrv.Create(config2.AppConfigIns.Default.SysUser.Username, config2.AppConfigIns.Default.SysUser.Password, &_domain)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		} else {
			err = userSrv.UpdatePass(userDefault, config2.AppConfigIns.Default.SysUser.Password)
			if err != nil {
				return err
			}
		}
		if err := initWxAppConfig(userDefault.Id, config2.AppConfigIns.Default.Wx); err != nil {
			return err
		}
	}
	return nil
}

func initWxAppConfig(userId int64, wxConfig *config2.WxConfig) error {
	var curAppConfig = &model.BeeWxConfig{}
	err := db.GetDB().Where("user_id = ?", userId).First(curAppConfig).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		curAppConfig = &model.BeeWxConfig{
			BaseModel: common.BaseModel{
				UserId:     userId,
				DateAdd:    common.JsonTime(time.Now()),
				DateUpdate: common.JsonTime(time.Now()),
			},
			AppId:     wxConfig.AppId,
			AppSecret: wxConfig.Secret,
		}
		return db.GetDB().Create(curAppConfig).Error
	} else if err != nil {
		return err
	} else {
		return nil
	}
}
