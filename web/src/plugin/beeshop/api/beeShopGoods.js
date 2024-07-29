import service from '@/utils/request'

// @Tags BeeShopGoods
// @Summary 创建商品表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoods true "创建商品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeShopGoods/createBeeShopGoods [post]
export const createBeeShopGoods = (data) => {
  return service({
    url: '/bee-shop/beeShopGoods/createBeeShopGoods',
    method: 'post',
    data
  })
}

// @Tags BeeShopGoods
// @Summary 删除商品表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoods true "删除商品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoods/deleteBeeShopGoods [delete]
export const deleteBeeShopGoods = (params) => {
  return service({
    url: '/bee-shop/beeShopGoods/deleteBeeShopGoods',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoods
// @Summary 批量删除商品表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoods/deleteBeeShopGoods [delete]
export const deleteBeeShopGoodsByIds = (params) => {
  return service({
    url: '/bee-shop/beeShopGoods/deleteBeeShopGoodsByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoods
// @Summary 更新商品表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoods true "更新商品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeShopGoods/updateBeeShopGoods [put]
export const updateBeeShopGoods = (data) => {
  return service({
    url: '/bee-shop/beeShopGoods/updateBeeShopGoods',
    method: 'put',
    data
  })
}

// @Tags BeeShopGoods
// @Summary 用id查询商品表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeShopGoods true "用id查询商品表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeShopGoods/findBeeShopGoods [get]
export const findBeeShopGoods = (params) => {
  return service({
    url: '/bee-shop/beeShopGoods/findBeeShopGoods',
    method: 'get',
    params
  })
}

// @Tags BeeShopGoods
// @Summary 分页获取商品表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeShopGoods/getBeeShopGoodsList [get]
export const getBeeShopGoodsList = (params) => {
  return service({
    url: '/bee-shop/beeShopGoods/getBeeShopGoodsList',
    method: 'get',
    params
  })
}
