import service from '@/utils/request'

// @Tags BeeComment
// @Summary 创建评论表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeComment true "创建评论表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeComment/createBeeComment [post]
export const createBeeComment = (data) => {
  return service({
    url: '/bee-shop/beeComment/createBeeComment',
    method: 'post',
    data
  })
}

// @Tags BeeComment
// @Summary 删除评论表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeComment true "删除评论表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeComment/deleteBeeComment [delete]
export const deleteBeeComment = (params) => {
  return service({
    url: '/bee-shop/beeComment/deleteBeeComment',
    method: 'delete',
    params
  })
}

// @Tags BeeComment
// @Summary 批量删除评论表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评论表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeComment/deleteBeeComment [delete]
export const deleteBeeCommentByIds = (params) => {
  return service({
    url: '/bee-shop/beeComment/deleteBeeCommentByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeComment
// @Summary 更新评论表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeComment true "更新评论表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeComment/updateBeeComment [put]
export const updateBeeComment = (data) => {
  return service({
    url: '/bee-shop/beeComment/updateBeeComment',
    method: 'put',
    data
  })
}

// @Tags BeeComment
// @Summary 用id查询评论表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeComment true "用id查询评论表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeComment/findBeeComment [get]
export const findBeeComment = (params) => {
  return service({
    url: '/bee-shop/beeComment/findBeeComment',
    method: 'get',
    params
  })
}

// @Tags BeeComment
// @Summary 分页获取评论表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取评论表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeComment/getBeeCommentList [get]
export const getBeeCommentList = (params) => {
  return service({
    url: '/bee-shop/beeComment/getBeeCommentList',
    method: 'get',
    params
  })
}
