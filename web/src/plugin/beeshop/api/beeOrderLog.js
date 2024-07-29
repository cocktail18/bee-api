import service from '@/utils/request'

// @Tags BeeOrderLog
// @Summary 创建订单日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderLog true "创建订单日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeOrderLog/createBeeOrderLog [post]
export const createBeeOrderLog = (data) => {
  return service({
    url: '/bee-shop/beeOrderLog/createBeeOrderLog',
    method: 'post',
    data
  })
}

// @Tags BeeOrderLog
// @Summary 删除订单日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderLog true "删除订单日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrderLog/deleteBeeOrderLog [delete]
export const deleteBeeOrderLog = (params) => {
  return service({
    url: '/bee-shop/beeOrderLog/deleteBeeOrderLog',
    method: 'delete',
    params
  })
}

// @Tags BeeOrderLog
// @Summary 批量删除订单日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrderLog/deleteBeeOrderLog [delete]
export const deleteBeeOrderLogByIds = (params) => {
  return service({
    url: '/bee-shop/beeOrderLog/deleteBeeOrderLogByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeOrderLog
// @Summary 更新订单日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderLog true "更新订单日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeOrderLog/updateBeeOrderLog [put]
export const updateBeeOrderLog = (data) => {
  return service({
    url: '/bee-shop/beeOrderLog/updateBeeOrderLog',
    method: 'put',
    data
  })
}

// @Tags BeeOrderLog
// @Summary 用id查询订单日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeOrderLog true "用id查询订单日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeOrderLog/findBeeOrderLog [get]
export const findBeeOrderLog = (params) => {
  return service({
    url: '/bee-shop/beeOrderLog/findBeeOrderLog',
    method: 'get',
    params
  })
}

// @Tags BeeOrderLog
// @Summary 分页获取订单日志列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订单日志列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeOrderLog/getBeeOrderLogList [get]
export const getBeeOrderLogList = (params) => {
  return service({
    url: '/bee-shop/beeOrderLog/getBeeOrderLogList',
    method: 'get',
    params
  })
}
