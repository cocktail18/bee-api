import service from '@/utils/request'

// @Tags BeeShopGoodsAddition
// @Summary 创建商品额外信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsAddition true "创建商品额外信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeShopGoodsAddition/createBeeShopGoodsAddition [post]
export const createBeeShopGoodsAddition = (data) => {
  return service({
    url: '/bee-shop/beeShopGoodsAddition/createBeeShopGoodsAddition',
    method: 'post',
    data
  })
}

// @Tags BeeShopGoodsAddition
// @Summary 删除商品额外信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsAddition true "删除商品额外信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsAddition/deleteBeeShopGoodsAddition [delete]
export const deleteBeeShopGoodsAddition = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsAddition/deleteBeeShopGoodsAddition',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsAddition
// @Summary 批量删除商品额外信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品额外信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsAddition/deleteBeeShopGoodsAddition [delete]
export const deleteBeeShopGoodsAdditionByIds = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsAddition/deleteBeeShopGoodsAdditionByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsAddition
// @Summary 更新商品额外信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsAddition true "更新商品额外信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeShopGoodsAddition/updateBeeShopGoodsAddition [put]
export const updateBeeShopGoodsAddition = (data) => {
  return service({
    url: '/bee-shop/beeShopGoodsAddition/updateBeeShopGoodsAddition',
    method: 'put',
    data
  })
}

// @Tags BeeShopGoodsAddition
// @Summary 用id查询商品额外信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeShopGoodsAddition true "用id查询商品额外信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeShopGoodsAddition/findBeeShopGoodsAddition [get]
export const findBeeShopGoodsAddition = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsAddition/findBeeShopGoodsAddition',
    method: 'get',
    params
  })
}

// @Tags BeeShopGoodsAddition
// @Summary 分页获取商品额外信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品额外信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeShopGoodsAddition/getBeeShopGoodsAdditionList [get]
export const getBeeShopGoodsAdditionList = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsAddition/getBeeShopGoodsAdditionList',
    method: 'get',
    params
  })
}
