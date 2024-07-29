import service from '@/utils/request'

// @Tags BeeUploadFile
// @Summary 创建用户上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUploadFile true "创建用户上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeUploadFile/createBeeUploadFile [post]
export const createBeeUploadFile = (data) => {
  return service({
    url: '/bee-shop/beeUploadFile/createBeeUploadFile',
    method: 'post',
    data
  })
}

// @Tags BeeUploadFile
// @Summary 删除用户上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUploadFile true "删除用户上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUploadFile/deleteBeeUploadFile [delete]
export const deleteBeeUploadFile = (params) => {
  return service({
    url: '/bee-shop/beeUploadFile/deleteBeeUploadFile',
    method: 'delete',
    params
  })
}

// @Tags BeeUploadFile
// @Summary 批量删除用户上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUploadFile/deleteBeeUploadFile [delete]
export const deleteBeeUploadFileByIds = (params) => {
  return service({
    url: '/bee-shop/beeUploadFile/deleteBeeUploadFileByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeUploadFile
// @Summary 更新用户上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUploadFile true "更新用户上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeUploadFile/updateBeeUploadFile [put]
export const updateBeeUploadFile = (data) => {
  return service({
    url: '/bee-shop/beeUploadFile/updateBeeUploadFile',
    method: 'put',
    data
  })
}

// @Tags BeeUploadFile
// @Summary 用id查询用户上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeUploadFile true "用id查询用户上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeUploadFile/findBeeUploadFile [get]
export const findBeeUploadFile = (params) => {
  return service({
    url: '/bee-shop/beeUploadFile/findBeeUploadFile',
    method: 'get',
    params
  })
}

// @Tags BeeUploadFile
// @Summary 分页获取用户上传文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户上传文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeUploadFile/getBeeUploadFileList [get]
export const getBeeUploadFileList = (params) => {
  return service({
    url: '/bee-shop/beeUploadFile/getBeeUploadFileList',
    method: 'get',
    params
  })
}
