import service from '@/utils/request'

// @Tags BeeUserCoupon
// @Summary 创建用户优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserCoupon true "创建用户优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeUserCoupon/createBeeUserCoupon [post]
export const createBeeUserCoupon = (data) => {
  return service({
    url: '/bee-shop/beeUserCoupon/createBeeUserCoupon',
    method: 'post',
    data
  })
}

// @Tags BeeUserCoupon
// @Summary 删除用户优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserCoupon true "删除用户优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUserCoupon/deleteBeeUserCoupon [delete]
export const deleteBeeUserCoupon = (params) => {
  return service({
    url: '/bee-shop/beeUserCoupon/deleteBeeUserCoupon',
    method: 'delete',
    params
  })
}

// @Tags BeeUserCoupon
// @Summary 批量删除用户优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUserCoupon/deleteBeeUserCoupon [delete]
export const deleteBeeUserCouponByIds = (params) => {
  return service({
    url: '/bee-shop/beeUserCoupon/deleteBeeUserCouponByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeUserCoupon
// @Summary 更新用户优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserCoupon true "更新用户优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeUserCoupon/updateBeeUserCoupon [put]
export const updateBeeUserCoupon = (data) => {
  return service({
    url: '/bee-shop/beeUserCoupon/updateBeeUserCoupon',
    method: 'put',
    data
  })
}

// @Tags BeeUserCoupon
// @Summary 用id查询用户优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeUserCoupon true "用id查询用户优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeUserCoupon/findBeeUserCoupon [get]
export const findBeeUserCoupon = (params) => {
  return service({
    url: '/bee-shop/beeUserCoupon/findBeeUserCoupon',
    method: 'get',
    params
  })
}

// @Tags BeeUserCoupon
// @Summary 分页获取用户优惠券列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户优惠券列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeUserCoupon/getBeeUserCouponList [get]
export const getBeeUserCouponList = (params) => {
  return service({
    url: '/bee-shop/beeUserCoupon/getBeeUserCouponList',
    method: 'get',
    params
  })
}
