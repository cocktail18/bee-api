package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/common"
	config2 "gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/medivhzhan/weapp/v3"
	"github.com/medivhzhan/weapp/v3/auth"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"io"
	"sync"
	"time"
)

type WxAppSrv struct {
	BaseSrv
	sync.RWMutex
	Appid       string
	Secret      string
	WeAppClient *weapp.Client
	WeAppAuth   *auth.Auth
}

func GetWxAppSrv(c context.Context) (*WxAppSrv, error) {
	//@todo cache
	userId := kit.GetUserId(c)
	return GetWxAppSrvByUserId(c, userId)
}

func GetWxAppSrvByUserId(c context.Context, userId int64) (*WxAppSrv, error) {
	var wxConfig model.BeeWxConfig
	if err := db.GetDB().Where("user_id = ? and is_deleted = 0", userId).Order("id desc").First(&wxConfig).Error; err != nil {
		return nil, errors.New("未配置微信小程序")
	}
	appSrv := &WxAppSrv{}
	if err := appSrv.init(wxConfig.AppId, wxConfig.AppSecret); err != nil {
		return nil, errors.New("微信小程序配置错误")
	}
	return appSrv, nil
}

func CreateOrSaveWxAppConfig(userId int64, wxConfig *config2.WxConfig) error {
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
	} else if err != nil {
		return err
	} else {
		if wxConfig.AppId == "" || wxConfig.Secret == "" {
			return nil
		}
		curAppConfig.AppId = wxConfig.AppId
		curAppConfig.AppSecret = wxConfig.Secret
		curAppConfig.DateUpdate = common.JsonTime(time.Now())
	}
	return db.GetDB().Save(curAppConfig).Error
}

func (srv *WxAppSrv) init(appId, secret string) error {
	srv.Appid = appId
	srv.Secret = secret
	srv.WeAppClient = weapp.NewClient(srv.Appid, srv.Secret)
	srv.WeAppAuth = srv.WeAppClient.NewAuth()
	return nil
}

func (srv *WxAppSrv) Authorize(c context.Context, code string) (*weapp.LoginResponse, error) {
	resp, err := srv.WeAppClient.Login(code)
	return resp, err
}

// GetWxAppQrCode 获取小程序码
func (srv *WxAppSrv) GetWxAppQrCode(c context.Context, req *weapp.UnlimitedQRCode) ([]byte, error) {
	res, commonErr, err := srv.WeAppClient.GetUnlimitedQRCode(req)
	if err != nil {
		return nil, err
	}
	if commonErr != nil && commonErr.ErrCode != 0 {
		return nil, commonErr.GetResponseError()
	}
	bs, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return bs, nil
}
