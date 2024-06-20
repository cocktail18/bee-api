package sys

import (
	config2 "gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/model/sys"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func InitDB() {
	allModel := []interface{}{
		&sys.SysUserModel{},
	}
	if err := db.GetDB().AutoMigrate(allModel...); err != nil {
		panic(err)
	}
	if config2.AppConfigIns.DB.Drop {
		logger.GetLogger().Infof("清空数据库")
		for _, model := range allModel {
			if err := db.GetDB().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(model).Error; err != nil {
				panic(err)
			}
		}
		logger.GetLogger().Infof("清空数据库成功")
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
		if err := service.CreateOrSaveWxAppConfig(userDefault.Id, config2.AppConfigIns.Default.Wx); err != nil {
			return err
		}
	}
	return nil
}
