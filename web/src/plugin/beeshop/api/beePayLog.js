import service from '@/utils/request'

// @Tags BeePayLog
// @Summary 创建支付流水
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeePayLog true "创建支付流水"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beePayLog/createBeePayLog [post]
export const createBeePayLog = (data) => {
  return service({
    url: '/bee-shop/beePayLog/createBeePayLog',
    method: 'post',
    data
  })
}

// @Tags BeePayLog
// @Summary 删除支付流水
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeePayLog true "删除支付流水"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beePayLog/deleteBeePayLog [delete]
export const deleteBeePayLog = (params) => {
  return service({
    url: '/bee-shop/beePayLog/deleteBeePayLog',
    method: 'delete',
    params
  })
}

// @Tags BeePayLog
// @Summary 批量删除支付流水
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除支付流水"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beePayLog/deleteBeePayLog [delete]
export const deleteBeePayLogByIds = (params) => {
  return service({
    url: '/bee-shop/beePayLog/deleteBeePayLogByIds',
    method: 'delete',
    params
  })
}

// @Tags BeePayLog
// @Summary 更新支付流水
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeePayLog true "更新支付流水"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beePayLog/updateBeePayLog [put]
export const updateBeePayLog = (data) => {
  return service({
    url: '/bee-shop/beePayLog/updateBeePayLog',
    method: 'put',
    data
  })
}

// @Tags BeePayLog
// @Summary 用id查询支付流水
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeePayLog true "用id查询支付流水"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beePayLog/findBeePayLog [get]
export const findBeePayLog = (params) => {
  return service({
    url: '/bee-shop/beePayLog/findBeePayLog',
    method: 'get',
    params
  })
}

// @Tags BeePayLog
// @Summary 分页获取支付流水列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取支付流水列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beePayLog/getBeePayLogList [get]
export const getBeePayLogList = (params) => {
  return service({
    url: '/bee-shop/beePayLog/getBeePayLogList',
    method: 'get',
    params
  })
}

// @Tags BeePayLog
// @Summary 分页获取支付流水列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取支付流水总额"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beePayLog/getBeePayTotal [get]
export const getBeePayTotal = (params) => {
  return service({
    url: '/bee-shop/beePayLog/getBeePayTotal',
    method: 'get',
    params
  })
}
