import service from '@/utils/request'

// @Tags BeeNotice
// @Summary 创建系统公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeNotice true "创建系统公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeNotice/createBeeNotice [post]
export const createBeeNotice = (data) => {
  return service({
    url: '/bee-shop/beeNotice/createBeeNotice',
    method: 'post',
    data
  })
}

// @Tags BeeNotice
// @Summary 删除系统公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeNotice true "删除系统公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeNotice/deleteBeeNotice [delete]
export const deleteBeeNotice = (params) => {
  return service({
    url: '/bee-shop/beeNotice/deleteBeeNotice',
    method: 'delete',
    params
  })
}

// @Tags BeeNotice
// @Summary 批量删除系统公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除系统公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeNotice/deleteBeeNotice [delete]
export const deleteBeeNoticeByIds = (params) => {
  return service({
    url: '/bee-shop/beeNotice/deleteBeeNoticeByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeNotice
// @Summary 更新系统公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeNotice true "更新系统公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeNotice/updateBeeNotice [put]
export const updateBeeNotice = (data) => {
  return service({
    url: '/bee-shop/beeNotice/updateBeeNotice',
    method: 'put',
    data
  })
}

// @Tags BeeNotice
// @Summary 用id查询系统公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeNotice true "用id查询系统公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeNotice/findBeeNotice [get]
export const findBeeNotice = (params) => {
  return service({
    url: '/bee-shop/beeNotice/findBeeNotice',
    method: 'get',
    params
  })
}

// @Tags BeeNotice
// @Summary 分页获取系统公告列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取系统公告列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeNotice/getBeeNoticeList [get]
export const getBeeNoticeList = (params) => {
  return service({
    url: '/bee-shop/beeNotice/getBeeNoticeList',
    method: 'get',
    params
  })
}
