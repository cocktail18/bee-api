import service from '@/utils/request'

// @Tags BeeUserAddress
// @Summary 创建用户地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserAddress true "创建用户地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeUserAddress/createBeeUserAddress [post]
export const createBeeUserAddress = (data) => {
  return service({
    url: '/bee-shop/beeUserAddress/createBeeUserAddress',
    method: 'post',
    data
  })
}

// @Tags BeeUserAddress
// @Summary 删除用户地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserAddress true "删除用户地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUserAddress/deleteBeeUserAddress [delete]
export const deleteBeeUserAddress = (params) => {
  return service({
    url: '/bee-shop/beeUserAddress/deleteBeeUserAddress',
    method: 'delete',
    params
  })
}

// @Tags BeeUserAddress
// @Summary 批量删除用户地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUserAddress/deleteBeeUserAddress [delete]
export const deleteBeeUserAddressByIds = (params) => {
  return service({
    url: '/bee-shop/beeUserAddress/deleteBeeUserAddressByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeUserAddress
// @Summary 更新用户地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserAddress true "更新用户地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeUserAddress/updateBeeUserAddress [put]
export const updateBeeUserAddress = (data) => {
  return service({
    url: '/bee-shop/beeUserAddress/updateBeeUserAddress',
    method: 'put',
    data
  })
}

// @Tags BeeUserAddress
// @Summary 用id查询用户地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeUserAddress true "用id查询用户地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeUserAddress/findBeeUserAddress [get]
export const findBeeUserAddress = (params) => {
  return service({
    url: '/bee-shop/beeUserAddress/findBeeUserAddress',
    method: 'get',
    params
  })
}

// @Tags BeeUserAddress
// @Summary 分页获取用户地址列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户地址列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeUserAddress/getBeeUserAddressList [get]
export const getBeeUserAddressList = (params) => {
  return service({
    url: '/bee-shop/beeUserAddress/getBeeUserAddressList',
    method: 'get',
    params
  })
}
