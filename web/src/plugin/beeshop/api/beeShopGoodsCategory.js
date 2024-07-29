import service from '@/utils/request'

// @Tags BeeShopGoodsCategory
// @Summary 创建商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsCategory true "创建商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeShopGoodsCategory/createBeeShopGoodsCategory [post]
export const createBeeShopGoodsCategory = (data) => {
  return service({
    url: '/bee-shop/beeShopGoodsCategory/createBeeShopGoodsCategory',
    method: 'post',
    data
  })
}

// @Tags BeeShopGoodsCategory
// @Summary 删除商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsCategory true "删除商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsCategory/deleteBeeShopGoodsCategory [delete]
export const deleteBeeShopGoodsCategory = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsCategory/deleteBeeShopGoodsCategory',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsCategory
// @Summary 批量删除商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsCategory/deleteBeeShopGoodsCategory [delete]
export const deleteBeeShopGoodsCategoryByIds = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsCategory/deleteBeeShopGoodsCategoryByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsCategory
// @Summary 更新商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsCategory true "更新商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeShopGoodsCategory/updateBeeShopGoodsCategory [put]
export const updateBeeShopGoodsCategory = (data) => {
  return service({
    url: '/bee-shop/beeShopGoodsCategory/updateBeeShopGoodsCategory',
    method: 'put',
    data
  })
}

// @Tags BeeShopGoodsCategory
// @Summary 用id查询商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeShopGoodsCategory true "用id查询商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeShopGoodsCategory/findBeeShopGoodsCategory [get]
export const findBeeShopGoodsCategory = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsCategory/findBeeShopGoodsCategory',
    method: 'get',
    params
  })
}

// @Tags BeeShopGoodsCategory
// @Summary 分页获取商品分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeShopGoodsCategory/getBeeShopGoodsCategoryList [get]
export const getBeeShopGoodsCategoryList = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsCategory/getBeeShopGoodsCategoryList',
    method: 'get',
    params
  })
}
