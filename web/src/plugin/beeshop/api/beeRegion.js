import service from '@/utils/request'

// @Tags BeeRegion
// @Summary 创建地址库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeRegion true "创建地址库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeRegion/createBeeRegion [post]
export const createBeeRegion = (data) => {
  return service({
    url: '/bee-shop/beeRegion/createBeeRegion',
    method: 'post',
    data
  })
}

// @Tags BeeRegion
// @Summary 删除地址库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeRegion true "删除地址库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeRegion/deleteBeeRegion [delete]
export const deleteBeeRegion = (params) => {
  return service({
    url: '/bee-shop/beeRegion/deleteBeeRegion',
    method: 'delete',
    params
  })
}

// @Tags BeeRegion
// @Summary 批量删除地址库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除地址库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeRegion/deleteBeeRegion [delete]
export const deleteBeeRegionByIds = (params) => {
  return service({
    url: '/bee-shop/beeRegion/deleteBeeRegionByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeRegion
// @Summary 更新地址库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeRegion true "更新地址库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeRegion/updateBeeRegion [put]
export const updateBeeRegion = (data) => {
  return service({
    url: '/bee-shop/beeRegion/updateBeeRegion',
    method: 'put',
    data
  })
}

// @Tags BeeRegion
// @Summary 用id查询地址库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeRegion true "用id查询地址库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeRegion/findBeeRegion [get]
export const findBeeRegion = (params) => {
  return service({
    url: '/bee-shop/beeRegion/findBeeRegion',
    method: 'get',
    params
  })
}

// @Tags BeeRegion
// @Summary 分页获取地址库列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取地址库列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeRegion/getBeeRegionList [get]
export const getBeeRegionList = (params) => {
  return service({
    url: '/bee-shop/beeRegion/getBeeRegionList',
    method: 'get',
    params
  })
}

// GetBeeRegionCityTree 获取省市区树
// @Tags BeeRegion
// @Summary 获取省市区树
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeRegion/getBeeRegionCityTree [get]
export const getBeeRegionCityTree = (params) => {
  return service({
    url: '/bee-shop/beeRegion/getBeeRegionCityTree',
    method: 'get',
    params
  })
}
