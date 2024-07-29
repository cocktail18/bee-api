import service from '@/utils/request'

// @Tags BeeUserBalanceLog
// @Summary 创建用户消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserBalanceLog true "创建用户消费记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeUserBalanceLog/createBeeUserBalanceLog [post]
export const createBeeUserBalanceLog = (data) => {
  return service({
    url: '/bee-shop/beeUserBalanceLog/createBeeUserBalanceLog',
    method: 'post',
    data
  })
}

// @Tags BeeUserBalanceLog
// @Summary 删除用户消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserBalanceLog true "删除用户消费记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUserBalanceLog/deleteBeeUserBalanceLog [delete]
export const deleteBeeUserBalanceLog = (params) => {
  return service({
    url: '/bee-shop/beeUserBalanceLog/deleteBeeUserBalanceLog',
    method: 'delete',
    params
  })
}

// @Tags BeeUserBalanceLog
// @Summary 批量删除用户消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户消费记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUserBalanceLog/deleteBeeUserBalanceLog [delete]
export const deleteBeeUserBalanceLogByIds = (params) => {
  return service({
    url: '/bee-shop/beeUserBalanceLog/deleteBeeUserBalanceLogByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeUserBalanceLog
// @Summary 更新用户消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserBalanceLog true "更新用户消费记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeUserBalanceLog/updateBeeUserBalanceLog [put]
export const updateBeeUserBalanceLog = (data) => {
  return service({
    url: '/bee-shop/beeUserBalanceLog/updateBeeUserBalanceLog',
    method: 'put',
    data
  })
}

// @Tags BeeUserBalanceLog
// @Summary 用id查询用户消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeUserBalanceLog true "用id查询用户消费记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeUserBalanceLog/findBeeUserBalanceLog [get]
export const findBeeUserBalanceLog = (params) => {
  return service({
    url: '/bee-shop/beeUserBalanceLog/findBeeUserBalanceLog',
    method: 'get',
    params
  })
}

// @Tags BeeUserBalanceLog
// @Summary 分页获取用户消费记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户消费记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeUserBalanceLog/getBeeUserBalanceLogList [get]
export const getBeeUserBalanceLogList = (params) => {
  return service({
    url: '/bee-shop/beeUserBalanceLog/getBeeUserBalanceLogList',
    method: 'get',
    params
  })
}
