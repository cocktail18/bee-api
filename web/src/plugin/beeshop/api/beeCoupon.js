import service from '@/utils/request'

// @Tags BeeCoupon
// @Summary 创建优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeCoupon true "创建优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeCoupon/createBeeCoupon [post]
export const createBeeCoupon = (data) => {
  return service({
    url: '/bee-shop/beeCoupon/createBeeCoupon',
    method: 'post',
    data
  })
}

// @Tags BeeCoupon
// @Summary 删除优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeCoupon true "删除优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeCoupon/deleteBeeCoupon [delete]
export const deleteBeeCoupon = (params) => {
  return service({
    url: '/bee-shop/beeCoupon/deleteBeeCoupon',
    method: 'delete',
    params
  })
}

// @Tags BeeCoupon
// @Summary 批量删除优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeCoupon/deleteBeeCoupon [delete]
export const deleteBeeCouponByIds = (params) => {
  return service({
    url: '/bee-shop/beeCoupon/deleteBeeCouponByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeCoupon
// @Summary 更新优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeCoupon true "更新优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeCoupon/updateBeeCoupon [put]
export const updateBeeCoupon = (data) => {
  return service({
    url: '/bee-shop/beeCoupon/updateBeeCoupon',
    method: 'put',
    data
  })
}

// @Tags BeeCoupon
// @Summary 用id查询优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeCoupon true "用id查询优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeCoupon/findBeeCoupon [get]
export const findBeeCoupon = (params) => {
  return service({
    url: '/bee-shop/beeCoupon/findBeeCoupon',
    method: 'get',
    params
  })
}

// @Tags BeeCoupon
// @Summary 分页获取优惠券列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取优惠券列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeCoupon/getBeeCouponList [get]
export const getBeeCouponList = (params) => {
  return service({
    url: '/bee-shop/beeCoupon/getBeeCouponList',
    method: 'get',
    params
  })
}
