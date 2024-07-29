import service from '@/utils/request'

// @Tags BeeUserMapper
// @Summary 创建beeUserMapper表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserMapper true "创建beeUserMapper表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeUserMapper/createBeeUserMapper [post]
export const createBeeUserMapper = (data) => {
  return service({
    url: '/bee-shop/beeUserMapper/createBeeUserMapper',
    method: 'post',
    data
  })
}

// @Tags BeeUserMapper
// @Summary 删除beeUserMapper表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserMapper true "删除beeUserMapper表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUserMapper/deleteBeeUserMapper [delete]
export const deleteBeeUserMapper = (params) => {
  return service({
    url: '/bee-shop/beeUserMapper/deleteBeeUserMapper',
    method: 'delete',
    params
  })
}

// @Tags BeeUserMapper
// @Summary 批量删除beeUserMapper表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除beeUserMapper表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUserMapper/deleteBeeUserMapper [delete]
export const deleteBeeUserMapperByIds = (params) => {
  return service({
    url: '/bee-shop/beeUserMapper/deleteBeeUserMapperByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeUserMapper
// @Summary 更新beeUserMapper表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserMapper true "更新beeUserMapper表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeUserMapper/updateBeeUserMapper [put]
export const updateBeeUserMapper = (data) => {
  return service({
    url: '/bee-shop/beeUserMapper/updateBeeUserMapper',
    method: 'put',
    data
  })
}

// @Tags BeeUserMapper
// @Summary 用id查询beeUserMapper表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeUserMapper true "用id查询beeUserMapper表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeUserMapper/findBeeUserMapper [get]
export const findBeeUserMapper = (params) => {
  return service({
    url: '/bee-shop/beeUserMapper/findBeeUserMapper',
    method: 'get',
    params
  })
}

// @Tags BeeUserMapper
// @Summary 分页获取beeUserMapper表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取beeUserMapper表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeUserMapper/getBeeUserMapperList [get]
export const getBeeUserMapperList = (params) => {
  return service({
    url: '/bee-shop/beeUserMapper/getBeeUserMapperList',
    method: 'get',
    params
  })
}
