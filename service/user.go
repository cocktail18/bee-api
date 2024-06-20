package service

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"sync"
	"time"
)

type UserSrv struct {
	BaseSrv
}

var userSrvOnce sync.Once
var userSrvInstance *UserSrv

func GetUserSrv() *UserSrv {
	userSrvOnce.Do(func() {
		userSrvInstance = &UserSrv{}

	})
	return userSrvInstance
}

func (srv *UserSrv) GetUserInfo(token string) (*model.BeeUser, error) {
	uid, err := GetTokenSrv().GetUserIdFromToken(token)
	if err != nil {
		return nil, err
	}
	var userInfo model.BeeUser
	err = db.GetDB().Where("id = ?", uid).Take(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (srv *UserSrv) Modify(user *model.BeeUser) error {
	user.DateUpdate = common.JsonTime(time.Now())
	return db.GetDB().Where("id = ?", user.Id).Save(user).Error
}

func (srv *UserSrv) Login(c *gin.Context, code string) (*proto.AuthorizeResp, error) {
	wxSrv, err := GetWxAppSrv(c)
	if err != nil {
		return nil, err
	}
	authResp, err := wxSrv.Authorize(code)
	if err != nil {
		return nil, err
	}
	if err = authResp.GetResponseError(); err != nil {
		return nil, err
	}
	var mapper model.BeeUserMapper
	err = db.GetDB().Where("open_id", authResp.OpenID).Take(&mapper).Error
	var uid int64
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 用户不存在，新建
		err = db.GetDB().Transaction(func(tx *gorm.DB) error {
			var user model.BeeUser
			now := time.Now()
			user.DateAdd = common.JsonTime(now)
			user.DateLogin = common.JsonTime(now)
			user.UserId = kit.GetUserId(c)
			if err := srv.CreateUser(tx, &user); err != nil {
				return err
			}
			uid = user.Id
			mapper.Uid = uid
			mapper.UserId = kit.GetUserId(c)
			mapper.Source = 1
			mapper.OpenId = authResp.OpenID
			mapper.UnionId = authResp.UnionID
			mapper.DateAdd = common.JsonTime(now)
			mapper.DateUpdate = common.JsonTime(now)
			err = db.GetDB().Create(&mapper).Error
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		uid = mapper.Uid
	}

	return srv.CreateUserToken(c, kit.GetUserId(c), uid, authResp.OpenID)
}

func (srv *UserSrv) CreateUser(tx *gorm.DB, user *model.BeeUser) error {
	var cur = struct {
		MaxShowUid int64 `json:"max_show_uid"`
	}{}
	if err := db.GetDB().Set("gorm:query_option", "FOR UPDATE").
		Select("max(`show_uid`) as `max_show_uid`").Model(&model.BeeUser{}).Take(&cur).Error; err != nil {
		return err
	}
	if cur.MaxShowUid == 0 {
		cur.MaxShowUid = 1000000
	}
	now := time.Now()
	user.ShowUid = cur.MaxShowUid + int64(util.RandInt(1, 10))
	user.DateAdd = common.JsonTime(now)
	user.DateUpdate = common.JsonTime(now)
	user.DateLogin = common.JsonTime(now)
	return tx.Create(&user).Error
}

func (srv *UserSrv) CreateUserToken(c *gin.Context, userId int64, uid int64, openId string) (*proto.AuthorizeResp, error) {
	token, err := GetTokenSrv().CreateToken(c, uid)
	if err != nil {
		return nil, err
	}
	resp := &proto.AuthorizeResp{}
	resp.Openid = openId
	resp.Token = token
	resp.Uid = uid
	resp.UserId = userId
	return resp, nil
}

func (srv *UserSrv) RecordIsSendRegisterCoupons(userId int64) error {
	return db.GetDB().Where("id = ?", userId).Update("is_send_register_coupons", 1).Error
}

func (srv *UserSrv) RecordBirthdayProcessSuccessYear(userId int64) error {
	return db.GetDB().Where("id = ?", userId).Update("birthday_process_success_year", carbon.Now().Year()).Error
}

func (srv *UserSrv) Amount(c *gin.Context, userId int64) (*proto.GetUserAmountResp, error) {
	var userAmount model.BeeUserAmount
	if err := db.GetDB().Where("user_id = ?", userId).Take(&userAmount).Error; err != nil {
		return nil, err
	}
	return &proto.GetUserAmountResp{
		BeeUserAmount: userAmount,
	}, nil
}

func (srv *UserSrv) CashLog(c *gin.Context, userId int64, req *proto.CashLogReq) (*proto.CashLogResp, error) {
	//@todo 类型筛选
	var (
		cnt      int64
		pageSize int64 = util.IF(req.PageSize == 0, 20, req.PageSize)
		page     int64 = util.IF(req.Page == 0, 1, req.Page)
		list     []*model.BeeCashLog
	)
	dbIns := db.GetDB().Where("uid = ?", userId)
	if err := dbIns.Count(&cnt).Error; err != nil {
		return nil, err
	}
	if err := dbIns.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&list).Error; err != nil {
		return nil, err
	}
	lo.ForEach(list, func(item *model.BeeCashLog, _ int) {
		item.FillData()
	})
	return &proto.CashLogResp{
		PageRespCommon: proto.GetPageCommon(cnt, pageSize),
		Result:         list,
	}, nil
}

func (srv *UserSrv) PayLogs(c *gin.Context, userId int64, req *proto.PayLogsReq) (proto.PayLogsResp, error) {
	//@todo 类型筛选
	var (
		pageSize int64 = util.IF(req.PageSize == 0, 20, req.PageSize)
		page     int64 = util.IF(req.Page == 0, 1, req.Page)
		list     []*model.BeePayLog
	)
	dbIns := db.GetDB().Where("uid = ?", userId)
	if err := dbIns.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&list).Error; err != nil {
		return nil, err
	}
	lo.ForEach(list, func(item *model.BeePayLog, _ int) {
		item.FillData()
	})
	return proto.PayLogsResp(list), nil
}

func (srv *UserSrv) GetDynamicUserCode(c *gin.Context) (string, error) {
	item := &model.BeeUserDynamicCode{}
	now := time.Now()
	err := db.GetDB().Where("uid = ?", kit.GetUserId(c)).Take(item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//暂无
		item = &model.BeeUserDynamicCode{
			BaseModel: *kit.GetInsertBaseModel(c),
			Uid:       kit.GetUserId(c),
			Code:      lo.RandomString(1, []rune("123456789")) + lo.RandomString(20, lo.NumbersCharset),
			ExpireAt:  now.Add(time.Minute * 10).Unix(),
		}
	} else if err != nil {
		return "", err
	} else {
		if item.ExpireAt > now.Unix() { // 还没过期
			return item.Code, nil
		}
		item.ExpireAt = now.Add(time.Minute * 10).Unix()
		item.Code = lo.RandomString(1, []rune("123456789")) + lo.RandomString(20, lo.NumbersCharset)
		item.DateUpdate = common.JsonTime(now)
	}
	//生成新的
	return item.Code, db.GetDB().Save(item).Error
}

func (srv *UserSrv) LevelList(c *gin.Context) (*proto.GetLevelListResp, error) {
	var list []*model.BeeUserLevel
	var cnt int64
	dbIns := db.GetDB().Where("user_id = ?", kit.GetUserId(c))
	if err := dbIns.Count(&cnt).Error; err != nil {
		return nil, err
	}
	if err := dbIns.Find(&list).Error; err != nil {
		return nil, err
	}
	return &proto.GetLevelListResp{
		PageRespCommon: proto.GetPageCommon(cnt, 1000),
		Result:         list,
	}, nil
}

func (srv *UserSrv) RechargeSendRule(c *gin.Context) ([]*model.RechargeSendRule, error) {
	var list []*model.RechargeSendRule
	if err := db.GetDB().Where("user_id = ?", kit.GetUserId(c)).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
