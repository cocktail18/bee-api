import service from '@/utils/request'

// @Tags BeeOrderGoods
// @Summary 创建订单商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderGoods true "创建订单商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeOrderGoods/createBeeOrderGoods [post]
export const createBeeOrderGoods = (data) => {
  return service({
    url: '/bee-shop/beeOrderGoods/createBeeOrderGoods',
    method: 'post',
    data
  })
}

// @Tags BeeOrderGoods
// @Summary 删除订单商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderGoods true "删除订单商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrderGoods/deleteBeeOrderGoods [delete]
export const deleteBeeOrderGoods = (params) => {
  return service({
    url: '/bee-shop/beeOrderGoods/deleteBeeOrderGoods',
    method: 'delete',
    params
  })
}

// @Tags BeeOrderGoods
// @Summary 批量删除订单商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrderGoods/deleteBeeOrderGoods [delete]
export const deleteBeeOrderGoodsByIds = (params) => {
  return service({
    url: '/bee-shop/beeOrderGoods/deleteBeeOrderGoodsByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeOrderGoods
// @Summary 更新订单商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderGoods true "更新订单商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeOrderGoods/updateBeeOrderGoods [put]
export const updateBeeOrderGoods = (data) => {
  return service({
    url: '/bee-shop/beeOrderGoods/updateBeeOrderGoods',
    method: 'put',
    data
  })
}

// @Tags BeeOrderGoods
// @Summary 用id查询订单商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeOrderGoods true "用id查询订单商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeOrderGoods/findBeeOrderGoods [get]
export const findBeeOrderGoods = (params) => {
  return service({
    url: '/bee-shop/beeOrderGoods/findBeeOrderGoods',
    method: 'get',
    params
  })
}

// @Tags BeeOrderGoods
// @Summary 分页获取订单商品信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订单商品信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeOrderGoods/getBeeOrderGoodsList [get]
export const getBeeOrderGoodsList = (params) => {
  return service({
    url: '/bee-shop/beeOrderGoods/getBeeOrderGoodsList',
    method: 'get',
    params
  })
}
