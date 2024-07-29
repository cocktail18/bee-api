import service from '@/utils/request'

// @Tags BeeShopGoodsContent
// @Summary 创建商品详情描述
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsContent true "创建商品详情描述"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeShopGoodsContent/createBeeShopGoodsContent [post]
export const createBeeShopGoodsContent = (data) => {
  return service({
    url: '/bee-shop/beeShopGoodsContent/createBeeShopGoodsContent',
    method: 'post',
    data
  })
}

// @Tags BeeShopGoodsContent
// @Summary 删除商品详情描述
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsContent true "删除商品详情描述"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsContent/deleteBeeShopGoodsContent [delete]
export const deleteBeeShopGoodsContent = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsContent/deleteBeeShopGoodsContent',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsContent
// @Summary 批量删除商品详情描述
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品详情描述"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopGoodsContent/deleteBeeShopGoodsContent [delete]
export const deleteBeeShopGoodsContentByIds = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsContent/deleteBeeShopGoodsContentByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeShopGoodsContent
// @Summary 更新商品详情描述
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopGoodsContent true "更新商品详情描述"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeShopGoodsContent/updateBeeShopGoodsContent [put]
export const updateBeeShopGoodsContent = (data) => {
  return service({
    url: '/bee-shop/beeShopGoodsContent/updateBeeShopGoodsContent',
    method: 'put',
    data
  })
}

// @Tags BeeShopGoodsContent
// @Summary 用id查询商品详情描述
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeShopGoodsContent true "用id查询商品详情描述"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeShopGoodsContent/findBeeShopGoodsContent [get]
export const findBeeShopGoodsContent = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsContent/findBeeShopGoodsContent',
    method: 'get',
    params
  })
}

// @Tags BeeShopGoodsContent
// @Summary 分页获取商品详情描述列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品详情描述列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeShopGoodsContent/getBeeShopGoodsContentList [get]
export const getBeeShopGoodsContentList = (params) => {
  return service({
    url: '/bee-shop/beeShopGoodsContent/getBeeShopGoodsContentList',
    method: 'get',
    params
  })
}
