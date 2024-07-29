import service from '@/utils/request'

// @Tags BeeBanner
// @Summary 创建商店banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeBanner true "创建商店banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeBanner/createBeeBanner [post]
export const createBeeBanner = (data) => {
  return service({
    url: '/bee-shop/beeBanner/createBeeBanner',
    method: 'post',
    data
  })
}

// @Tags BeeBanner
// @Summary 删除商店banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeBanner true "删除商店banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeBanner/deleteBeeBanner [delete]
export const deleteBeeBanner = (params) => {
  return service({
    url: '/bee-shop/beeBanner/deleteBeeBanner',
    method: 'delete',
    params
  })
}

// @Tags BeeBanner
// @Summary 批量删除商店banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商店banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeBanner/deleteBeeBanner [delete]
export const deleteBeeBannerByIds = (params) => {
  return service({
    url: '/bee-shop/beeBanner/deleteBeeBannerByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeBanner
// @Summary 更新商店banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeBanner true "更新商店banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeBanner/updateBeeBanner [put]
export const updateBeeBanner = (data) => {
  return service({
    url: '/bee-shop/beeBanner/updateBeeBanner',
    method: 'put',
    data
  })
}

// @Tags BeeBanner
// @Summary 用id查询商店banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeBanner true "用id查询商店banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeBanner/findBeeBanner [get]
export const findBeeBanner = (params) => {
  return service({
    url: '/bee-shop/beeBanner/findBeeBanner',
    method: 'get',
    params
  })
}

// @Tags BeeBanner
// @Summary 分页获取商店banner列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商店banner列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeBanner/getBeeBannerList [get]
export const getBeeBannerList = (params) => {
  return service({
    url: '/bee-shop/beeBanner/getBeeBannerList',
    method: 'get',
    params
  })
}
