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
	Upload  *UploadConfig
}

type UploadConfig struct {
	OssType string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // Oss类型
	// oss
	Local      Local      `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu      Qiniu      `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS  AliyunOSS  `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`
	HuaWeiObs  HuaWeiObs  `mapstructure:"hua-wei-obs" json:"hua-wei-obs" yaml:"hua-wei-obs"`
	TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencent-cos" yaml:"tencent-cos"`
	AwsS3      AwsS3      `mapstructure:"aws-s3" json:"aws-s3" yaml:"aws-s3"`
}

type App struct {
	Listen       string
	DfsHost      string
	PayNotifyUrl string
	StorePath    string `yaml:"storePath"` //文件上传路径
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

func GetStorePath() string {
	return util.IF(AppConfigIns.App.StorePath == "", "uploads/file", AppConfigIns.App.StorePath)
}

func GetDfsHost(host string) string {
	return util.IF(AppConfigIns.App.DfsHost == "", "https://"+host, AppConfigIns.App.DfsHost)
}

func GetOssType() string {
	return util.IF(AppConfigIns.Upload.OssType == "", "local", AppConfigIns.Upload.OssType)
}
