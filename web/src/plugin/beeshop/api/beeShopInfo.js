import service from '@/utils/request'

// @Tags BeeShopInfo
// @Summary 创建商店信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopInfo true "创建商店信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeShopInfo/createBeeShopInfo [post]
export const createBeeShopInfo = (data) => {
  return service({
    url: '/bee-shop/beeShopInfo/createBeeShopInfo',
    method: 'post',
    data
  })
}

// @Tags BeeShopInfo
// @Summary 删除商店信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopInfo true "删除商店信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopInfo/deleteBeeShopInfo [delete]
export const deleteBeeShopInfo = (params) => {
  return service({
    url: '/bee-shop/beeShopInfo/deleteBeeShopInfo',
    method: 'delete',
    params
  })
}

// @Tags BeeShopInfo
// @Summary 批量删除商店信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商店信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeShopInfo/deleteBeeShopInfo [delete]
export const deleteBeeShopInfoByIds = (params) => {
  return service({
    url: '/bee-shop/beeShopInfo/deleteBeeShopInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeShopInfo
// @Summary 更新商店信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeShopInfo true "更新商店信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeShopInfo/updateBeeShopInfo [put]
export const updateBeeShopInfo = (data) => {
  return service({
    url: '/bee-shop/beeShopInfo/updateBeeShopInfo',
    method: 'put',
    data
  })
}

// @Tags BeeShopInfo
// @Summary 用id查询商店信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeShopInfo true "用id查询商店信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeShopInfo/findBeeShopInfo [get]
export const findBeeShopInfo = (params) => {
  return service({
    url: '/bee-shop/beeShopInfo/findBeeShopInfo',
    method: 'get',
    params
  })
}

// @Tags BeeShopInfo
// @Summary 分页获取商店信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商店信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeShopInfo/getBeeShopInfoList [get]
export const getBeeShopInfoList = (params) => {
  return service({
    url: '/bee-shop/beeShopInfo/getBeeShopInfoList',
    method: 'get',
    params
  })
}
