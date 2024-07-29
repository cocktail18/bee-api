import service from '@/utils/request'

// @Tags BeeWxConfig
// @Summary 创建微信配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeWxConfig true "创建微信配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeWxConfig/createBeeWxConfig [post]
export const createBeeWxConfig = (data) => {
  return service({
    url: '/bee-shop/beeWxConfig/createBeeWxConfig',
    method: 'post',
    data
  })
}

// @Tags BeeWxConfig
// @Summary 删除微信配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeWxConfig true "删除微信配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeWxConfig/deleteBeeWxConfig [delete]
export const deleteBeeWxConfig = (params) => {
  return service({
    url: '/bee-shop/beeWxConfig/deleteBeeWxConfig',
    method: 'delete',
    params
  })
}

// @Tags BeeWxConfig
// @Summary 批量删除微信配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除微信配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeWxConfig/deleteBeeWxConfig [delete]
export const deleteBeeWxConfigByIds = (params) => {
  return service({
    url: '/bee-shop/beeWxConfig/deleteBeeWxConfigByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeWxConfig
// @Summary 更新微信配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeWxConfig true "更新微信配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeWxConfig/updateBeeWxConfig [put]
export const updateBeeWxConfig = (data) => {
  return service({
    url: '/bee-shop/beeWxConfig/updateBeeWxConfig',
    method: 'put',
    data
  })
}

// @Tags BeeWxConfig
// @Summary 用id查询微信配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeWxConfig true "用id查询微信配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeWxConfig/findBeeWxConfig [get]
export const findBeeWxConfig = (params) => {
  return service({
    url: '/bee-shop/beeWxConfig/findBeeWxConfig',
    method: 'get',
    params
  })
}

// @Tags BeeWxConfig
// @Summary 分页获取微信配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取微信配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeWxConfig/getBeeWxConfigList [get]
export const getBeeWxConfigList = (params) => {
  return service({
    url: '/bee-shop/beeWxConfig/getBeeWxConfigList',
    method: 'get',
    params
  })
}
