import service from '@/utils/request'

// @Tags BeeConfig
// @Summary 创建公用配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeConfig true "创建公用配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeConfig/createBeeConfig [post]
export const createBeeConfig = (data) => {
  return service({
    url: '/bee-shop/beeConfig/createBeeConfig',
    method: 'post',
    data
  })
}

// @Tags BeeConfig
// @Summary 删除公用配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeConfig true "删除公用配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeConfig/deleteBeeConfig [delete]
export const deleteBeeConfig = (params) => {
  return service({
    url: '/bee-shop/beeConfig/deleteBeeConfig',
    method: 'delete',
    params
  })
}

// @Tags BeeConfig
// @Summary 批量删除公用配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除公用配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeConfig/deleteBeeConfig [delete]
export const deleteBeeConfigByIds = (params) => {
  return service({
    url: '/bee-shop/beeConfig/deleteBeeConfigByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeConfig
// @Summary 更新公用配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeConfig true "更新公用配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeConfig/updateBeeConfig [put]
export const updateBeeConfig = (data) => {
  return service({
    url: '/bee-shop/beeConfig/updateBeeConfig',
    method: 'put',
    data
  })
}

// @Tags BeeConfig
// @Summary 用id查询公用配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeConfig true "用id查询公用配置表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeConfig/findBeeConfig [get]
export const findBeeConfig = (params) => {
  return service({
    url: '/bee-shop/beeConfig/findBeeConfig',
    method: 'get',
    params
  })
}

// @Tags BeeConfig
// @Summary 分页获取公用配置表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取公用配置表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeConfig/getBeeConfigList [get]
export const getBeeConfigList = (params) => {
  return service({
    url: '/bee-shop/beeConfig/getBeeConfigList',
    method: 'get',
    params
  })
}
