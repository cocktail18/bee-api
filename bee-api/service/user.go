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
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
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

func (srv *UserSrv) GetUserInfo(c context.Context, token string) (*model.BeeUser, error) {
	tokenInfo, err := GetTokenSrv().GetTokenInfoFromToken(token)
	if err != nil {
		return nil, err
	}
	beeUserInfo, err := srv.GetUserInfoByUid(c, tokenInfo.Uid)
	if err != nil {
		return nil, err
	}
	beeUserInfo.SessionKey = tokenInfo.SessionKey
	return beeUserInfo, nil
}

func (srv *UserSrv) GetUserInfoByUid(c context.Context, uid int64) (*model.BeeUser, error) {
	var userInfo model.BeeUser
	err := db.GetDB().Where("id = ?", uid).Take(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (srv *UserSrv) Modify(user *model.BeeUser) error {
	user.DateUpdate = common.JsonTime(time.Now())
	return db.GetDB().Where("id = ?", user.Id).Save(user).Error
}

func (srv *UserSrv) Login(c context.Context, code string) (*proto.AuthorizeResp, error) {
	wxSrv, err := GetWxAppSrv(c)
	if err != nil {
		return nil, err
	}
	authResp, err := wxSrv.Authorize(c, code)
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
			if err := srv.CreateUser(c, tx, &user); err != nil {
				return err
			}
			uid = user.Id
			mapper.Uid = uid
			mapper.UserId = kit.GetUserId(c)
			mapper.Source = enum.BeeUserSourceWx
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

	return srv.CreateUserToken(c, kit.GetUserId(c), uid, authResp.OpenID, authResp.SessionKey)
}

func (srv *UserSrv) CreateUser(c context.Context, tx *gorm.DB, user *model.BeeUser) error {
	var cur = struct {
		MaxShowUid int64 `json:"max_show_uid"`
	}{}
	minLevel, err := srv.GetMinLevel(c)
	if err != nil {
		return err
	}
	if err := tx.Set("gorm:query_option", "FOR UPDATE").
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
	if err := tx.Create(&user).Error; err != nil {
		return err
	}
	if err := tx.Create(&model.BeeUserAmount{
		BaseModel:         *kit.GetInsertBaseModel(c),
		Uid:               user.Id,
		Balance:           decimal.Zero,
		Freeze:            decimal.Zero,
		FxCommisionPaying: decimal.Zero,
		Growth:            decimal.Zero,
		Score:             decimal.Zero,
		ScoreLastRound:    decimal.Zero,
		TotalPayAmount:    decimal.Zero,
		TotalPayNumber:    decimal.Zero,
		TotalScore:        decimal.Zero,
		TotalWithdraw:     decimal.Zero,
		TotalConsumed:     decimal.Zero,
		Pwd:               "",
		Salt:              "",
	}).Error; err != nil {
		return err
	}
	userLevel := &model.BeeUserLevel{
		BaseModel: common.BaseModel{},
		Uid:       user.Id,
		Level:     util.IF(minLevel.UpgradeAmount.LessThanOrEqual(decimal.NewFromInt(0)), minLevel.Id, 0),
		PayAmount: decimal.Decimal{},
	}
	return tx.Create(userLevel).Error
}

func (srv *UserSrv) CreateUserToken(c context.Context, userId int64, uid int64, openId string, sessionKey string) (*proto.AuthorizeResp, error) {
	token, err := GetTokenSrv().CreateToken(c, uid, sessionKey)
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

func (srv *UserSrv) Amount(c context.Context, userId int64) (*proto.GetUserAmountResp, error) {
	var userAmount model.BeeUserAmount
	if err := db.GetDB().Where("user_id = ?", userId).Take(&userAmount).Error; err != nil {
		return nil, err
	}
	return &proto.GetUserAmountResp{
		BeeUserAmount: userAmount,
	}, nil
}

func (srv *UserSrv) CashLog(c context.Context, userId int64, req *proto.CashLogReq) (*proto.CashLogResp, error) {
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

func (srv *UserSrv) PayLogs(c context.Context, userId int64, req *proto.PayLogsReq) (proto.PayLogsResp, error) {
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

func (srv *UserSrv) GetDynamicUserCode(c context.Context) (string, error) {
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

func (srv *UserSrv) LevelList(c context.Context) (*proto.GetLevelListResp, error) {
	var list []*model.BeeLevel
	var cnt int64
	dbIns := db.GetDB().Model(&model.BeeLevel{}).Where("user_id = ?", kit.GetUserId(c))
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

func (srv *UserSrv) GetMinLevel(c context.Context) (*model.BeeLevel, error) {
	var item = &model.BeeLevel{}
	if err := db.GetDB().Where("user_id = ?", kit.GetUserId(c)).Order("upgrade_amount asc").First(item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.BeeLevel{}, nil
		}
		return nil, err
	}
	return item, nil
}

func (srv *UserSrv) GetLevelByAmount(c context.Context, amount decimal.Decimal) (*model.BeeLevel, error) {
	var item = &model.BeeLevel{}
	if err := db.GetDB().Where("user_id = ? and upgrade_amount < ?", kit.GetUserId(c), amount).Order("upgrade_amount desc").First(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (srv *UserSrv) RechargeSendRule(c context.Context) ([]*model.RechargeSendRule, error) {
	var list []*model.RechargeSendRule
	if err := db.GetDB().Where("user_id = ?", kit.GetUserId(c)).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (srv *UserSrv) GetUserLevel(c context.Context, uid int64) (*model.BeeUserLevel, error) {
	item := &model.BeeUserLevel{}
	if err := db.GetDB().Where("uid = ?", uid).Take(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (srv *UserSrv) IncrUserLevelAmount(c context.Context, tx *gorm.DB, uid int64, payAmount decimal.Decimal) error {
	item := &model.BeeUserLevel{}
	if err := tx.Where("uid = ?", uid).Take(item).Error; err != nil {
		return err
	}
	item.PayAmount = item.PayAmount.Add(payAmount)
	newLevelInfo, err := srv.GetLevelByAmount(c, item.PayAmount)
	if err == nil {
		item.Level = newLevelInfo.Level
	}
	return tx.Save(item).Error
}

func (srv *UserSrv) BindWxMobile(c context.Context, req *proto.BindWxMobileReq) (string, error) {
	wxSrv, err := GetWxAppSrv(c)
	if err != nil {
		return "", err
	}
	var phoneNum string
	//phoneRes, err := wxSrv.WeAppClient.NewPhonenumber().GetPhoneNumber(&phonenumber.GetPhoneNumberRequest{Code: req.Code})
	//if err == nil && phoneRes.GetResponseError() == nil {
	//	phoneNum = phoneRes.Data.PhoneNumber
	//} else {
	mobileRes, err := wxSrv.WeAppClient.DecryptMobile(kit.GetSessionKey(c), req.EncryptedData, req.Iv)
	if err != nil {
		return "", err
	}
	phoneNum = mobileRes.PhoneNumber
	//}

	userMobile := &model.BeeUserMobile{
		BaseModel: *kit.GetInsertBaseModel(c),
		Uid:       kit.GetUid(c),
		Mobile:    phoneNum,
	}
	oldItem := &model.BeeUserMobile{}
	err = db.GetDB().Where("uid = ?", kit.GetUserId(c)).Take(oldItem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.GetDB().Create(userMobile).Error
	} else if err == nil {
		err = db.GetDB().Save(userMobile).Error
	}
	if err != nil {
		return "", err
	}

	if err := db.GetDB().Model(&model.BeeUser{}).Where("id = ?", kit.GetUid(c)).Updates(map[string]interface{}{
		"mobile":      phoneNum,
		"date_update": time.Now(),
	}).Error; err != nil {
		return "", err
	}
	return phoneNum, nil
}

func (srv *UserSrv) GetUserWxOpenId(c context.Context) (string, error) {
	item := &model.BeeUserMapper{}
	if err := db.GetDB().Where("uid = ? and source = ?", kit.GetUserId(c), enum.BeeUserSourceWx).Take(item).Error; err != nil {
		return "", err
	}
	return item.OpenId, nil
}
