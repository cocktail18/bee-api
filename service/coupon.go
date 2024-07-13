package service

import (
	"context"
	"fmt"
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
	"time"
)

type CouponSrv struct {
	BaseSrv
}

var couponSrvOnce sync.Once
var couponSrvInstance *CouponSrv

func GetCouponSrv() *CouponSrv {
	couponSrvOnce.Do(func() {
		couponSrvInstance = &CouponSrv{}
	})
	return couponSrvInstance
}

func (srv *CouponSrv) GetMyCouponListByStatus(userId int64, status enum.CouponStatus) ([]*model.BeeUserCoupon, error) {
	//@todo 分页
	var list []*model.BeeUserCoupon
	err := db.GetDB().Where("uid =? and status=?", userId, status).Find(&list).Error
	return list, err
}

func (srv *CouponSrv) GetMyCouponStatistics(userId int64) (*proto.MyCouponStatisticsResp, error) {
	//@todo 待领取
	var list []*model.BeeUserCoupon
	err := db.GetDB().Where("uid =?", userId).Find(&list).Error
	resp := &proto.MyCouponStatisticsResp{}
	expireCouponIds := make([]int64, 0, 10)
	for _, coupon := range list {
		switch coupon.Status {
		case 0:
			if coupon.IsExpire() {
				expireCouponIds = append(expireCouponIds, coupon.Id)
				resp.Invalid++
			} else {
				resp.CanUse++
			}
		case 1:
			fallthrough
		case 2:
			resp.Used++
		case 3:
			resp.Invalid++
		}
	}
	db.GetDB().Where("id in (?)", expireCouponIds).Update("status", 3)
	return resp, err
}

func (srv *CouponSrv) GetCoupons(showInFront int) ([]*model.BeeCoupon, error) {
	var couponList []*model.BeeCoupon
	err := db.GetDB().Where(map[string]interface{}{
		"show_in_front": showInFront,
		"status":        0,
	}).Find(&couponList).Error
	return couponList, err
}

func (srv *CouponSrv) FetchCoupon(c context.Context, userInfo *model.BeeUser, id int64) error {
	coupon, err := srv.getCoupon(id)
	if err != nil {
		return errors.Wrap(err, "获取优惠券信息失败")
	}
	if coupon.Status != 0 {
		return enum.NewBussErr(nil, 30001, "优惠券已经下架")
	}
	if coupon.NumberLeft <= 0 {
		return enum.NewBussErr(nil, 30005, "优惠券已经被领完了")
	}
	//SendBirthday         bool     `gorm:"column:send_birthday;type:tinyint(1);comment:生日赠送" json:"sendBirthday"`
	//SendInviteM          bool     `gorm:"column:send_invite_m;type:tinyint(1);comment:邀请新人注册赠送" json:"sendInviteM"`
	//SendInviteS          bool     `gorm:"column:send_invite_s;type:tinyint(1);comment:被邀请赠送" json:"sendInviteS"`
	//SendRegister         bool     `gorm:"column:send_register;type:tinyint(1);comment:注册赠送" json:"sendRegister"`
	var uniqStr = ""
	var couponLog = &model.BeeUserCouponLog{}

	if coupon.NeedSignedContinuous > 0 {
		//连续签到多少天可得
		lastSignLog, err := GetScoreSrv().GetLastSignLog(c)
		if err != nil {
			return err
		}
		if lastSignLog == nil {
			return enum.ErrIneligible
		}
		if int64(lastSignLog.Continues) < coupon.NeedSignedContinuous {
			return enum.ErrIneligible
		}
		signBeginDate := time.Now().AddDate(0, 0, -1*int(coupon.NeedSignedContinuous))
		uniqStr = fmt.Sprintf("%s-%d", signBeginDate.Format("20060102"), lastSignLog.Continues/int(coupon.NeedSignedContinuous))
	}
	if coupon.NumberPersonMax > 0 {
		var cnt int64
		if err := db.GetDB().Model(&model.BeeUserCouponLog{}).Where("type = ?", enum.CouponLogTypeReceive).Count(&cnt).Error; err != nil {
			return err
		}
		if cnt >= coupon.NumberPersonMax {
			return enum.NewBussErr(nil, 30006, "超过最大领取次数了")
		}
	}

	if coupon.SendBirthday {
		if userInfo.Birthday != time.Now().Format("2006-01-02") {
			return enum.NewBussErr(nil, 30002, "仅生日当天可领取")
		}
		uniqStr = time.Now().Format("2006")
	}

	if db.GetDB().Where("uniq = ?", uniqStr).Take(couponLog); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		couponLog.CouponId = coupon.Id
		couponLog.Uniq = uniqStr
		couponLog.BaseModel = *kit.GetInsertBaseModel(c)
		couponLog.Remark = ""
		couponLog.Uid = userInfo.Id
		couponLog.Type = enum.CouponLogTypeReceive
	} else {
		return enum.ErrReceived
	}

	dateStart := common.JsonTime(time.Now())
	expireMills := int64(0)
	switch coupon.DateStartType { //，1立即，2次日，0指定时间
	case 1:
	case 2:
		dateStart = common.JsonTime(time.Now().AddDate(0, 0, 1))
	case 0:
		dateStart = coupon.DateStart
	}

	switch coupon.DateEndType {
	case enum.CouponDateEndTypeDelay:
		expireMills = time.Time(coupon.DateEnd).AddDate(0, 0, coupon.DateEndDays).UnixMilli()
	case enum.CouponDateEndTypeFixed:
		expireMills = time.Time(coupon.DateEnd).UnixMilli() - time.Now().UnixMilli()

	}

	// 检查完毕，开始领取
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(couponLog).Error; err != nil {
			return err
		}
		rs := tx.Model(&model.BeeCoupon{}).Where("id = ? and number_left > 0", coupon.Id).Updates(map[string]interface{}{
			"number_left": gorm.Expr("number_left - 1"),
			"date_update": time.Now(),
		})
		if rs.Error != nil {
			return err
		}
		if rs.RowsAffected != 1 {
			return enum.NewBussErr(nil, 30005, "优惠券已经被领完了")
		}
		if err := tx.Create(&model.BeeUserCoupon{
			BaseModel:    *kit.GetInsertBaseModel(c),
			Uid:          userInfo.Id,
			DateStart:    dateStart,
			ExpiryMillis: expireMills,
			//Money:         coupon.NeedAmount,
			MoneyHreshold: coupon.MoneyHreshold,
			MoneyMin:      coupon.MoneyMin,
			MoneyMax:      coupon.MoneyMax,
			MoneyType:     coupon.MoneyType,
			Name:          coupon.Name,
			OnlyFreight:   coupon.OnlyFreight,
			Pid:           0,
			ShopId:        coupon.ShopId,
			Source:        "用户领取",
			Status:        enum.CouponStatusNormal,
		}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (srv *CouponSrv) getCoupon(id int64) (*model.BeeCoupon, error) {
	var data model.BeeCoupon
	err := db.GetDB().Where("id=?", id).Find(&data).Error
	return &data, err
}

func (srv *CouponSrv) GetUserCoupon(userId, id int64) (*model.BeeUserCoupon, error) {
	var data model.BeeUserCoupon
	err := db.GetDB().Where("id=? and uid  =?", id, userId).Take(&data).Error
	return &data, err
}

func (srv *CouponSrv) GetUserCouponByIds(c context.Context, userId int64, ids []int64) ([]*model.BeeUserCoupon, error) {
	var data []*model.BeeUserCoupon
	err := db.GetDB().Where("id in ? and uid  =?", ids, userId).Find(&data).Error
	return data, err
}

func (srv *CouponSrv) UseCoupon(c context.Context, tx *gorm.DB, userId int64, id int64) error {
	err := tx.Model(&model.BeeUserCoupon{}).Where("id = ? and uid  =?", id, userId).Updates(map[string]interface{}{
		"status":      enum.CouponStatusUsed,
		"date_update": time.Now(),
	}).Error
	return err
}

func (srv *CouponSrv) ReturnCoupon(c context.Context, tx *gorm.DB, userId int64, id int64) error {
	err := tx.Model(&model.BeeUserCoupon{}).Where("id = ? and uid  =?", id, userId).Updates(map[string]interface{}{
		"status":      enum.CouponStatusUsing,
		"date_update": time.Now(),
	}).Error
	return err
}

func (srv *CouponSrv) CouponDetail(c context.Context, id int64) (*proto.CouponDetailResp, error) {
	var couponDetail model.BeeCoupon

	err := db.GetDB().Where("id=?", id).Find(&couponDetail).Error
	if err != nil {
		return nil, err
	}
	return &proto.CouponDetailResp{Info: &couponDetail}, nil
}
