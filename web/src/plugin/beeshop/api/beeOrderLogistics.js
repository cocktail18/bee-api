import service from '@/utils/request'

// @Tags BeeOrderLogistics
// @Summary 创建用户订单地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderLogistics true "创建用户订单地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeOrderLogistics/createBeeOrderLogistics [post]
export const createBeeOrderLogistics = (data) => {
  return service({
    url: '/bee-shop/beeOrderLogistics/createBeeOrderLogistics',
    method: 'post',
    data
  })
}

// @Tags BeeOrderLogistics
// @Summary 删除用户订单地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderLogistics true "删除用户订单地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrderLogistics/deleteBeeOrderLogistics [delete]
export const deleteBeeOrderLogistics = (params) => {
  return service({
    url: '/bee-shop/beeOrderLogistics/deleteBeeOrderLogistics',
    method: 'delete',
    params
  })
}

// @Tags BeeOrderLogistics
// @Summary 批量删除用户订单地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户订单地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrderLogistics/deleteBeeOrderLogistics [delete]
export const deleteBeeOrderLogisticsByIds = (params) => {
  return service({
    url: '/bee-shop/beeOrderLogistics/deleteBeeOrderLogisticsByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeOrderLogistics
// @Summary 更新用户订单地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderLogistics true "更新用户订单地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeOrderLogistics/updateBeeOrderLogistics [put]
export const updateBeeOrderLogistics = (data) => {
  return service({
    url: '/bee-shop/beeOrderLogistics/updateBeeOrderLogistics',
    method: 'put',
    data
  })
}

// @Tags BeeOrderLogistics
// @Summary 用id查询用户订单地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeOrderLogistics true "用id查询用户订单地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeOrderLogistics/findBeeOrderLogistics [get]
export const findBeeOrderLogistics = (params) => {
  return service({
    url: '/bee-shop/beeOrderLogistics/findBeeOrderLogistics',
    method: 'get',
    params
  })
}

// @Tags BeeOrderLogistics
// @Summary 分页获取用户订单地址列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户订单地址列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeOrderLogistics/getBeeOrderLogisticsList [get]
export const getBeeOrderLogisticsList = (params) => {
  return service({
    url: '/bee-shop/beeOrderLogistics/getBeeOrderLogisticsList',
    method: 'get',
    params
  })
}
