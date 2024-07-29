import service from '@/utils/request'

// @Tags BeeCashLog
// @Summary 创建用户现金消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeCashLog true "创建用户现金消费记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeCashLog/createBeeCashLog [post]
export const createBeeCashLog = (data) => {
  return service({
    url: '/bee-shop/beeCashLog/createBeeCashLog',
    method: 'post',
    data
  })
}

// @Tags BeeCashLog
// @Summary 删除用户现金消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeCashLog true "删除用户现金消费记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeCashLog/deleteBeeCashLog [delete]
export const deleteBeeCashLog = (params) => {
  return service({
    url: '/bee-shop/beeCashLog/deleteBeeCashLog',
    method: 'delete',
    params
  })
}

// @Tags BeeCashLog
// @Summary 批量删除用户现金消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户现金消费记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeCashLog/deleteBeeCashLog [delete]
export const deleteBeeCashLogByIds = (params) => {
  return service({
    url: '/bee-shop/beeCashLog/deleteBeeCashLogByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeCashLog
// @Summary 更新用户现金消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeCashLog true "更新用户现金消费记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeCashLog/updateBeeCashLog [put]
export const updateBeeCashLog = (data) => {
  return service({
    url: '/bee-shop/beeCashLog/updateBeeCashLog',
    method: 'put',
    data
  })
}

// @Tags BeeCashLog
// @Summary 用id查询用户现金消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeCashLog true "用id查询用户现金消费记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeCashLog/findBeeCashLog [get]
export const findBeeCashLog = (params) => {
  return service({
    url: '/bee-shop/beeCashLog/findBeeCashLog',
    method: 'get',
    params
  })
}

// @Tags BeeCashLog
// @Summary 分页获取用户现金消费记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户现金消费记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeCashLog/getBeeCashLogList [get]
export const getBeeCashLogList = (params) => {
  return service({
    url: '/bee-shop/beeCashLog/getBeeCashLogList',
    method: 'get',
    params
  })
}
