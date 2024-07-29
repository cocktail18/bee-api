import service from '@/utils/request'

// @Tags BeeUser
// @Summary 创建beeUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUser true "创建beeUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeUser/createBeeUser [post]
export const createBeeUser = (data) => {
  return service({
    url: '/bee-shop/beeUser/createBeeUser',
    method: 'post',
    data
  })
}

// @Tags BeeUser
// @Summary 删除beeUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUser true "删除beeUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUser/deleteBeeUser [delete]
export const deleteBeeUser = (params) => {
  return service({
    url: '/bee-shop/beeUser/deleteBeeUser',
    method: 'delete',
    params
  })
}

// @Tags BeeUser
// @Summary 批量删除beeUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除beeUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUser/deleteBeeUser [delete]
export const deleteBeeUserByIds = (params) => {
  return service({
    url: '/bee-shop/beeUser/deleteBeeUserByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeUser
// @Summary 更新beeUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUser true "更新beeUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeUser/updateBeeUser [put]
export const updateBeeUser = (data) => {
  return service({
    url: '/bee-shop/beeUser/updateBeeUser',
    method: 'put',
    data
  })
}

// @Tags BeeUser
// @Summary 用id查询beeUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeUser true "用id查询beeUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeUser/findBeeUser [get]
export const findBeeUser = (params) => {
  return service({
    url: '/bee-shop/beeUser/findBeeUser',
    method: 'get',
    params
  })
}

// @Tags BeeUser
// @Summary 分页获取beeUser表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取beeUser表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeUser/getBeeUserList [get]
export const getBeeUserList = (params) => {
  return service({
    url: '/bee-shop/beeUser/getBeeUserList',
    method: 'get',
    params
  })
}
