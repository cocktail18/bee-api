package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/util"
	"sync"
	"time"
)

type TokenSrv struct {
	BaseSrv
}

var tokenSrvOnce sync.Once
var tokenSrvInstance *TokenSrv

func GetTokenSrv() *TokenSrv {
	tokenSrvOnce.Do(func() {
		tokenSrvInstance = &TokenSrv{}
		go func() {
			ticket := time.NewTicker(time.Second * 10)
			for {
				<-ticket.C
				_ = tokenSrvInstance.CleanExpireToken()
			}
		}()
	})
	return tokenSrvInstance
}

func (srv *TokenSrv) RefreshToken(token string) error {
	return db.GetDB().Where("token = ?", token).Update("expire_at", time.Now().Add(time.Hour*24)).Error
}

func (srv *TokenSrv) CleanExpireToken() error {
	return db.GetDB().Where("expire_at < ?", time.Now().Unix()).Delete(&model.BeeToken{}).Error
}

func (srv *TokenSrv) CreateToken(c context.Context, uid int64, sessionKey string) (string, error) {
	info := model.BeeToken{}
	info.Token = util.GetUUIDStr()
	info.Uid = uid
	info.SessionKey = sessionKey
	info.BaseModel = *kit.GetInsertBaseModel(c)
	info.ExpireAt = time.Now().Add(time.Hour * 24 * 10).Unix()
	err := db.GetDB().Create(&info).Error
	return info.Token, err
}

func (srv *TokenSrv) GetUserIdFromToken(token string) (int64, error) {
	var tokenInfo model.BeeToken
	err := db.GetDB().Where("token = ?", token).Take(&tokenInfo).Error
	if err != nil {
		return 0, err
	}
	return tokenInfo.Uid, nil
}

func (srv *TokenSrv) GetTokenInfoFromToken(token string) (*model.BeeToken, error) {
	var tokenInfo = &model.BeeToken{}
	err := db.GetDB().Where("token = ?", token).Take(tokenInfo).Error
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
