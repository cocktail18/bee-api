import service from '@/utils/request'

// @Tags BeeShopGoodsImages
// @Summary 创建商品图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsImages true "创建商品图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeShopGoodsImages/createBeeShopGoodsImages [post]
export const createBeeShopGoodsImages = (data) => {
  return service({
    url: '/bee-shop/beeShopGoodsImages/createBeeShopGoodsImages',
    method: 'post',
    data
  })
}

// @Tags BeeShopGoodsImages
// @Summary 删除商品图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsImages true "删除商品图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsImages/deleteBeeShopGoodsImages [delete]
export const deleteBeeShopGoodsImages = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsImages/deleteBeeShopGoodsImages',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsImages
// @Summary 批量删除商品图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsImages/deleteBeeShopGoodsImages [delete]
export const deleteBeeShopGoodsImagesByIds = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsImages/deleteBeeShopGoodsImagesByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsImages
// @Summary 更新商品图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsImages true "更新商品图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeShopGoodsImages/updateBeeShopGoodsImages [put]
export const updateBeeShopGoodsImages = (data) => {
  return service({
    url: '/bee-shop/beeShopGoodsImages/updateBeeShopGoodsImages',
    method: 'put',
    data
  })
}

// @Tags BeeShopGoodsImages
// @Summary 用id查询商品图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeShopGoodsImages true "用id查询商品图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeShopGoodsImages/findBeeShopGoodsImages [get]
export const findBeeShopGoodsImages = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsImages/findBeeShopGoodsImages',
    method: 'get',
    params
  })
}

// @Tags BeeShopGoodsImages
// @Summary 分页获取商品图列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品图列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeShopGoodsImages/getBeeShopGoodsImagesList [get]
export const getBeeShopGoodsImagesList = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsImages/getBeeShopGoodsImagesList',
    method: 'get',
    params
  })
}
