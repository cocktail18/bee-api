package service

import (
	"context"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/proto"
	"github.com/pkg/errors"
	"sync"
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

func (srv *CouponSrv) FetchCoupon(userInfo *model.BeeUser, id int64) error {
	coupon, err := srv.getCoupon(id)
	if err != nil {
		return errors.Wrap(err, "获取优惠券信息失败")
	}
	if coupon.Status != 0 {
		return enum.NewBussErr(nil, 30001, "优惠券已经下架")
	}
	//SendBirthday         bool     `gorm:"column:send_birthday;type:tinyint(1);comment:生日赠送" json:"sendBirthday"`
	//SendInviteM          bool     `gorm:"column:send_invite_m;type:tinyint(1);comment:邀请新人注册赠送" json:"sendInviteM"`
	//SendInviteS          bool     `gorm:"column:send_invite_s;type:tinyint(1);comment:被邀请赠送" json:"sendInviteS"`
	//SendRegister         bool     `gorm:"column:send_register;type:tinyint(1);comment:注册赠送" json:"sendRegister"`
	if coupon.SendBirthday { //@todo 生日校验
	}
	if coupon.SendRegister {
		if userInfo.IsSendRegisterCoupons {
			return enum.NewBussErr(nil, 30002, "该优惠券已领取过了")
		}
	}

	// 检查完毕，开始领取
	if coupon.SendRegister {
		err = GetUserSrv().RecordIsSendRegisterCoupons(userInfo.UserId)
		if err != nil {
			return errors.Wrap(err, "记录获取注册优惠券失败")
		}
	}
	if coupon.SendBirthday {
		err = GetUserSrv().RecordBirthdayProcessSuccessYear(userInfo.UserId)
		if err != nil {
			return errors.Wrap(err, "记录获取生日优惠券失败")
		}
	}
	return nil
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

func (srv *CouponSrv) CouponDetail(c context.Context, id int64) (*proto.CouponDetailResp, error) {
	var couponDetail model.BeeCoupon

	err := db.GetDB().Where("id=?", id).Find(&couponDetail).Error
	if err != nil {
		return nil, err
	}
	return &proto.CouponDetailResp{Info: &couponDetail}, nil
}
