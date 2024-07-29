import service from '@/utils/request'

// @Tags BeeLogistics
// @Summary 创建运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeLogistics true "创建运费模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeLogistics/createBeeLogistics [post]
export const createBeeLogistics = (data) => {
  return service({
    url: '/bee-shop/beeLogistics/createBeeLogistics',
    method: 'post',
    data
  })
}

// @Tags BeeLogistics
// @Summary 删除运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeLogistics true "删除运费模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeLogistics/deleteBeeLogistics [delete]
export const deleteBeeLogistics = (params) => {
  return service({
    url: '/bee-shop/beeLogistics/deleteBeeLogistics',
    method: 'delete',
    params
  })
}

// @Tags BeeLogistics
// @Summary 批量删除运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除运费模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeLogistics/deleteBeeLogistics [delete]
export const deleteBeeLogisticsByIds = (params) => {
  return service({
    url: '/bee-shop/beeLogistics/deleteBeeLogisticsByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeLogistics
// @Summary 更新运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeLogistics true "更新运费模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeLogistics/updateBeeLogistics [put]
export const updateBeeLogistics = (data) => {
  return service({
    url: '/bee-shop/beeLogistics/updateBeeLogistics',
    method: 'put',
    data
  })
}

// @Tags BeeLogistics
// @Summary 用id查询运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeLogistics true "用id查询运费模版"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeLogistics/findBeeLogistics [get]
export const findBeeLogistics = (params) => {
  return service({
    url: '/bee-shop/beeLogistics/findBeeLogistics',
    method: 'get',
    params
  })
}

// @Tags BeeLogistics
// @Summary 分页获取运费模版列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取运费模版列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeLogistics/getBeeLogisticsList [get]
export const getBeeLogisticsList = (params) => {
  return service({
    url: '/bee-shop/beeLogistics/getBeeLogisticsList',
    method: 'get',
    params
  })
}
