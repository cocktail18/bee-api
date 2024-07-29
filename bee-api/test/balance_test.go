package test

import (
	db2 "gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/service"
	"gitee.com/stuinfer/bee-api/sys"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestBalanceSrv_GetAmount(t *testing.T) {
	InitTestConfig()
	containerName, dsn, err := util.InitDockerMysqlForTest()
	assert.Nil(t, err)
	defer func() {
		util.StopDockerMysqlForTest(containerName)
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	dbIns, err := db2.NewDB(dsn)
	assert.Nil(t, err)

	//mock.Patch(db2.GetDB, func() *gorm.DB {
	//	return dbIns.GetDB()
	//})
	patches := gomonkey.ApplyFuncReturn(db2.GetDB, dbIns.GetDB())
	defer patches.Reset()
	model.InitDB()
	sys.InitDB()

	testAmount(t)
}

func testAmount(t *testing.T) {
	Convey("余额测试", t, func() {
		defaultUser := &model.BeeUser{}
		err := db2.GetDB().Transaction(func(tx *gorm.DB) error {
			return service.GetUserSrv().CreateUser(GetTestContext(), tx, defaultUser)
		})
		So(err, ShouldBeNil)
		amount, err := service.GetBalanceSrv().GetAmount(defaultUser.Id)
		So(err, ShouldBeNil)
		So(amount.Balance.String(), ShouldEqual, decimal.Zero.String())
		ctx := GetTestContextByUserInfo(defaultUser)
		addBalance := decimal.NewFromInt(100)
		decrBalance := decimal.NewFromInt(-10)
		newBalance, err := service.GetBalanceSrv().OperAmount(ctx, defaultUser.Id, enum.BalanceTypeBalance, addBalance, util.GetUUIDStr(), "test")
		So(err, ShouldBeNil)
		So(newBalance.Balance.String(), ShouldEqual, addBalance.String())
		amount, err = service.GetBalanceSrv().GetAmount(defaultUser.Id)
		So(err, ShouldBeNil)
		So(amount.Balance.String(), ShouldEqual, addBalance.String())

		newBalance, err = service.GetBalanceSrv().OperAmount(ctx, defaultUser.Id, enum.BalanceTypeBalance, decrBalance, util.GetUUIDStr(), "test")
		So(err, ShouldBeNil)
		So(newBalance.Balance.String(), ShouldEqual, addBalance.Add(decrBalance).String())
		amount, err = service.GetBalanceSrv().GetAmount(defaultUser.Id)
		So(err, ShouldBeNil)
		So(amount.Balance.String(), ShouldEqual, addBalance.Add(decrBalance).String())
	})
}
