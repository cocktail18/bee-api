import service from '@/utils/request'

// @Tags BeeUserAmount
// @Summary 创建用户资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserAmount true "创建用户资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeUserAmount/createBeeUserAmount [post]
export const createBeeUserAmount = (data) => {
  return service({
    url: '/bee-shop/beeUserAmount/createBeeUserAmount',
    method: 'post',
    data
  })
}

// @Tags BeeUserAmount
// @Summary 删除用户资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserAmount true "删除用户资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUserAmount/deleteBeeUserAmount [delete]
export const deleteBeeUserAmount = (params) => {
  return service({
    url: '/bee-shop/beeUserAmount/deleteBeeUserAmount',
    method: 'delete',
    params
  })
}

// @Tags BeeUserAmount
// @Summary 批量删除用户资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUserAmount/deleteBeeUserAmount [delete]
export const deleteBeeUserAmountByIds = (params) => {
  return service({
    url: '/bee-shop/beeUserAmount/deleteBeeUserAmountByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeUserAmount
// @Summary 更新用户资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserAmount true "更新用户资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeUserAmount/updateBeeUserAmount [put]
export const updateBeeUserAmount = (data) => {
  return service({
    url: '/bee-shop/beeUserAmount/updateBeeUserAmount',
    method: 'put',
    data
  })
}

// @Tags BeeUserAmount
// @Summary 用id查询用户资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeUserAmount true "用id查询用户资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeUserAmount/findBeeUserAmount [get]
export const findBeeUserAmount = (params) => {
  return service({
    url: '/bee-shop/beeUserAmount/findBeeUserAmount',
    method: 'get',
    params
  })
}

// @Tags BeeUserAmount
// @Summary 分页获取用户资产列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户资产列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeUserAmount/getBeeUserAmountList [get]
export const getBeeUserAmountList = (params) => {
  return service({
    url: '/bee-shop/beeUserAmount/getBeeUserAmountList',
    method: 'get',
    params
  })
}
