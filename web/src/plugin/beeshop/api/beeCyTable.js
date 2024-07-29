import service from '@/utils/request'

// @Tags BeeCyTable
// @Summary 创建桌号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeCyTable true "创建桌号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeCyTable/createBeeCyTable [post]
export const createBeeCyTable = (data) => {
  return service({
    url: '/bee-shop/beeCyTable/createBeeCyTable',
    method: 'post',
    data
  })
}

// @Tags BeeCyTable
// @Summary 删除桌号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeCyTable true "删除桌号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeCyTable/deleteBeeCyTable [delete]
export const deleteBeeCyTable = (params) => {
  return service({
    url: '/bee-shop/beeCyTable/deleteBeeCyTable',
    method: 'delete',
    params
  })
}

// @Tags BeeCyTable
// @Summary 批量删除桌号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除桌号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeCyTable/deleteBeeCyTable [delete]
export const deleteBeeCyTableByIds = (params) => {
  return service({
    url: '/bee-shop/beeCyTable/deleteBeeCyTableByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeCyTable
// @Summary 更新桌号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeCyTable true "更新桌号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeCyTable/updateBeeCyTable [put]
export const updateBeeCyTable = (data) => {
  return service({
    url: '/bee-shop/beeCyTable/updateBeeCyTable',
    method: 'put',
    data
  })
}

// @Tags BeeCyTable
// @Summary 用id查询桌号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeCyTable true "用id查询桌号信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeCyTable/findBeeCyTable [get]
export const findBeeCyTable = (params) => {
  return service({
    url: '/bee-shop/beeCyTable/findBeeCyTable',
    method: 'get',
    params
  })
}

// @Tags BeeCyTable
// @Summary 分页获取桌号信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取桌号信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeCyTable/getBeeCyTableList [get]
export const getBeeCyTableList = (params) => {
  return service({
    url: '/bee-shop/beeCyTable/getBeeCyTableList',
    method: 'get',
    params
  })
}
