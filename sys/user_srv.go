package sys

import (
	"gitee.com/stuinfer/bee-api/common"
	config2 "gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/model/sys"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"time"
)

type UserSrv struct {
}

var userSrv = &UserSrv{}

func GetUserSrv() *UserSrv {
	return userSrv
}

func (srv *UserSrv) Create(username, pass string, domain *string) (*sys.SysUserModel, error) {
	if domain == nil {
		_domain := lo.RandomString(10, append(lo.LowerCaseLettersCharset, lo.NumbersCharset...))
		domain = &_domain
	}

	salt := lo.RandomString(10, append(lo.LowerCaseLettersCharset, lo.NumbersCharset...))

	user := &sys.SysUserModel{
		BaseModel: common.BaseModel{
			DateAdd:    common.JsonTime(time.Now()),
			DateUpdate: common.JsonTime(time.Now()),
		},
		Domain:   *domain,
		Salt:     salt,
		Username: username,
		Password: srv.getPass(pass, salt),
	}
	if config2.AppConfigIns.Default.SysUser.Username == username { //demo
		user.Id = 100
	}
	err := db.GetDB().Create(user).Error
	if err != nil {
		return nil, err
	}
	user.UserId = user.Id
	err = db.GetDB().Save(user).Error

	return user, errors.Wrap(srv.InitBeeData(user.Id), "初始化数据失败")
}

func (srv *UserSrv) getPass(pass, salt string) string {
	return util.Md5(salt + pass + salt)
}

func (srv *UserSrv) GetByPass(username, pass string) (*sys.SysUserModel, error) {
	var user sys.SysUserModel
	err := db.GetDB().Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	if user.Password != srv.getPass(pass, user.Salt) {
		return nil, errors.New("账号或者密码错误")
	}
	return &user, nil
}

func (srv *UserSrv) UpdatePass(user *sys.SysUserModel, pass string) error {
	user.Password = srv.getPass(pass, user.Salt)
	user.DateUpdate = common.JsonTime(time.Now())
	return db.GetDB().Save(user).Error
}

func (srv *UserSrv) GetByDomain(domain string) (*sys.SysUserModel, error) {
	var user sys.SysUserModel
	err := db.GetDB().Where("domain = ?", domain).First(&user).Error
	return &user, err
}
