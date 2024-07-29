import service from '@/utils/request'

// @Tags BeeShopGoodsProp
// @Summary 创建商品属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsProp true "创建商品属性"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeShopGoodsProp/createBeeShopGoodsProp [post]
export const createBeeShopGoodsProp = (data) => {
  return service({
    url: '/bee-shop/beeShopGoodsProp/createBeeShopGoodsProp',
    method: 'post',
    data
  })
}

// @Tags BeeShopGoodsProp
// @Summary 删除商品属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsProp true "删除商品属性"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsProp/deleteBeeShopGoodsProp [delete]
export const deleteBeeShopGoodsProp = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsProp/deleteBeeShopGoodsProp',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsProp
// @Summary 批量删除商品属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品属性"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsProp/deleteBeeShopGoodsProp [delete]
export const deleteBeeShopGoodsPropByIds = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsProp/deleteBeeShopGoodsPropByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsProp
// @Summary 更新商品属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsProp true "更新商品属性"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeShopGoodsProp/updateBeeShopGoodsProp [put]
export const updateBeeShopGoodsProp = (data) => {
  return service({
    url: '/bee-shop/beeShopGoodsProp/updateBeeShopGoodsProp',
    method: 'put',
    data
  })
}

// @Tags BeeShopGoodsProp
// @Summary 用id查询商品属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeShopGoodsProp true "用id查询商品属性"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeShopGoodsProp/findBeeShopGoodsProp [get]
export const findBeeShopGoodsProp = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsProp/findBeeShopGoodsProp',
    method: 'get',
    params
  })
}

// @Tags BeeShopGoodsProp
// @Summary 分页获取商品属性列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品属性列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeShopGoodsProp/getBeeShopGoodsPropList [get]
export const getBeeShopGoodsPropList = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsProp/getBeeShopGoodsPropList',
    method: 'get',
    params
  })
}
