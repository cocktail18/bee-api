package api

import (
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/gin-gonic/gin"
)

type YuyueAPi struct {
	BaseApi
}

func (api YuyueAPi) Join(c *gin.Context) {
	//yuyueId: 377
	//extJsonStr: {"姓名":"asdfa","联系电话":"1231","到店时间":"2022/11/1上午9:27:55","用餐人数":"1-2人"}
	//	@todo
	api.Fail(c, enum.ResCodeEmpty, "todo")
}
