import service from '@/utils/request'

// @Tags BeePeiSong
// @Summary 创建配送信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeePeiSong true "创建配送信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beePeiSong/createBeePeiSong [post]
export const createBeePeiSong = (data) => {
  return service({
    url: '/bee-shop/beePeiSong/createBeePeiSong',
    method: 'post',
    data
  })
}

// @Tags BeePeiSong
// @Summary 删除配送信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeePeiSong true "删除配送信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beePeiSong/deleteBeePeiSong [delete]
export const deleteBeePeiSong = (params) => {
  return service({
    url: '/bee-shop/beePeiSong/deleteBeePeiSong',
    method: 'delete',
    params
  })
}

// @Tags BeePeiSong
// @Summary 批量删除配送信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除配送信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beePeiSong/deleteBeePeiSong [delete]
export const deleteBeePeiSongByIds = (params) => {
  return service({
    url: '/bee-shop/beePeiSong/deleteBeePeiSongByIds',
    method: 'delete',
    params
  })
}

// @Tags BeePeiSong
// @Summary 更新配送信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeePeiSong true "更新配送信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beePeiSong/updateBeePeiSong [put]
export const updateBeePeiSong = (data) => {
  return service({
    url: '/bee-shop/beePeiSong/updateBeePeiSong',
    method: 'put',
    data
  })
}

// @Tags BeePeiSong
// @Summary 用id查询配送信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeePeiSong true "用id查询配送信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beePeiSong/findBeePeiSong [get]
export const findBeePeiSong = (params) => {
  return service({
    url: '/bee-shop/beePeiSong/findBeePeiSong',
    method: 'get',
    params
  })
}

// @Tags BeePeiSong
// @Summary 分页获取配送信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取配送信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beePeiSong/getBeePeiSongList [get]
export const getBeePeiSongList = (params) => {
  return service({
    url: '/bee-shop/beePeiSong/getBeePeiSongList',
    method: 'get',
    params
  })
}
