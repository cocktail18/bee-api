import service from '@/utils/request'

// @Tags BeeShopGoodsSku
// @Summary 创建商品sku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsSku true "创建商品sku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeShopGoodsSku/createBeeShopGoodsSku [post]
export const createBeeShopGoodsSku = (data) => {
  return service({
    url: '/bee-shop/beeShopGoodsSku/createBeeShopGoodsSku',
    method: 'post',
    data
  })
}

// @Tags BeeShopGoodsSku
// @Summary 删除商品sku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsSku true "删除商品sku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsSku/deleteBeeShopGoodsSku [delete]
export const deleteBeeShopGoodsSku = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsSku/deleteBeeShopGoodsSku',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsSku
// @Summary 批量删除商品sku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品sku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsSku/deleteBeeShopGoodsSku [delete]
export const deleteBeeShopGoodsSkuByIds = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsSku/deleteBeeShopGoodsSkuByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsSku
// @Summary 更新商品sku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsSku true "更新商品sku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeShopGoodsSku/updateBeeShopGoodsSku [put]
export const updateBeeShopGoodsSku = (data) => {
  return service({
    url: '/bee-shop/beeShopGoodsSku/updateBeeShopGoodsSku',
    method: 'put',
    data
  })
}

// @Tags BeeShopGoodsSku
// @Summary 用id查询商品sku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeShopGoodsSku true "用id查询商品sku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeShopGoodsSku/findBeeShopGoodsSku [get]
export const findBeeShopGoodsSku = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsSku/findBeeShopGoodsSku',
    method: 'get',
    params
  })
}

// @Tags BeeShopGoodsSku
// @Summary 分页获取商品sku列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品sku列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeShopGoodsSku/getBeeShopGoodsSkuList [get]
export const getBeeShopGoodsSkuList = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsSku/getBeeShopGoodsSkuList',
    method: 'get',
    params
  })
}
