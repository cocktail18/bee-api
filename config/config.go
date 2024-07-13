package config

import (
	"gitee.com/stuinfer/bee-api/util"
	"github.com/spf13/viper"
)

var AppConfigIns = &AppConfig{}

type AppConfig struct {
	App     *App
	DB      *AppDBConfig
	Default *DefaultConfig
}
type App struct {
	Listen       string
	DfsHost      string
	PayNotifyUrl string
}

type DefaultConfig struct {
	Wx      *WxConfig
	SysUser *SysUser `yaml:"sysUser"`
}

type SysUser struct {
	Domain   string `yaml:"domain"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type WxConfig struct {
	AppId  string `yaml:"appid"`
	Secret string `yaml:"secret"`
}

type AppDBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Drop     bool
	NeedInit bool `yaml:"needInit"`
}

func InitConfig() {
	// 加载配置文件
	configFilePath := "./config.yml"
	if ok, err := util.FileExists(configFilePath); err != nil {
		panic(err)
	} else if !ok {
		configFilePath = "./../config.yml"
	}
	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 将配置映射到结构体
	if err := viper.Unmarshal(AppConfigIns); err != nil {
		panic(err)
	}
}
