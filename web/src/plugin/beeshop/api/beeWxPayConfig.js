import service from '@/utils/request'

// @Tags BeeWxPayConfig
// @Summary 创建微信支付配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeWxPayConfig true "创建微信支付配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeWxPayConfig/createBeeWxPayConfig [post]
export const createBeeWxPayConfig = (data) => {
  return service({
    url: '/bee-shop/beeWxPayConfig/createBeeWxPayConfig',
    method: 'post',
    data
  })
}

// @Tags BeeWxPayConfig
// @Summary 删除微信支付配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeWxPayConfig true "删除微信支付配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeWxPayConfig/deleteBeeWxPayConfig [delete]
export const deleteBeeWxPayConfig = (params) => {
  return service({
    url: '/bee-shop/beeWxPayConfig/deleteBeeWxPayConfig',
    method: 'delete',
    params
  })
}

// @Tags BeeWxPayConfig
// @Summary 批量删除微信支付配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除微信支付配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeWxPayConfig/deleteBeeWxPayConfig [delete]
export const deleteBeeWxPayConfigByIds = (params) => {
  return service({
    url: '/bee-shop/beeWxPayConfig/deleteBeeWxPayConfigByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeWxPayConfig
// @Summary 更新微信支付配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeWxPayConfig true "更新微信支付配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeWxPayConfig/updateBeeWxPayConfig [put]
export const updateBeeWxPayConfig = (data) => {
  return service({
    url: '/bee-shop/beeWxPayConfig/updateBeeWxPayConfig',
    method: 'put',
    data
  })
}

// @Tags BeeWxPayConfig
// @Summary 用id查询微信支付配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeWxPayConfig true "用id查询微信支付配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeWxPayConfig/findBeeWxPayConfig [get]
export const findBeeWxPayConfig = (params) => {
  return service({
    url: '/bee-shop/beeWxPayConfig/findBeeWxPayConfig',
    method: 'get',
    params
  })
}

// @Tags BeeWxPayConfig
// @Summary 分页获取微信支付配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取微信支付配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeWxPayConfig/getBeeWxPayConfigList [get]
export const getBeeWxPayConfigList = (params) => {
  return service({
    url: '/bee-shop/beeWxPayConfig/getBeeWxPayConfigList',
    method: 'get',
    params
  })
}
