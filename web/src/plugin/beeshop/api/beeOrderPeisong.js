import service from '@/utils/request'

// @Tags BeeOrderPeisong
// @Summary 创建beeOrderPeisong表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderPeisong true "创建beeOrderPeisong表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeOrderPeisong/createBeeOrderPeisong [post]
export const createBeeOrderPeisong = (data) => {
    return service({
        url: '/bee-shop/beeOrderPeisong/createBeeOrderPeisong',
        method: 'post',
        data
    })
}

// @Tags BeeOrder
// @Summary 取消配送
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrder true "取消配送"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"取消成功"}"
// @Router /beeOrderPeisong/cancelBeeOrderPeisong [post]
export const cancelBeeOrderPeisong = (data) => {
    return service({
        url: '/bee-shop/beeOrderPeisong/cancelBeeOrderPeisong',
        method: 'post',
        data
    })
}

// @Tags BeeOrder
// @Summary 通知配送
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrder true "通知配送"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"通知成功"}"
// @Router /beeOrderPeisong/notifyBeeOrderPeisong [post]
export const notifyBeeOrderPeisong = (data) => {
    return service({
        url: '/bee-shop/beeOrderPeisong/notifyBeeOrderPeisong',
        method: 'post',
        data
    })
}

// @Tags BeeOrderPeisong
// @Summary 删除beeOrderPeisong表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderPeisong true "删除beeOrderPeisong表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrderPeisong/deleteBeeOrderPeisong [delete]
export const deleteBeeOrderPeisong = (params) => {
    return service({
        url: '/bee-shop/beeOrderPeisong/deleteBeeOrderPeisong',
        method: 'delete',
        params
    })
}

// @Tags BeeOrderPeisong
// @Summary 批量删除beeOrderPeisong表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除beeOrderPeisong表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrderPeisong/deleteBeeOrderPeisong [delete]
export const deleteBeeOrderPeisongByIds = (params) => {
    return service({
        url: '/bee-shop/beeOrderPeisong/deleteBeeOrderPeisongByIds',
        method: 'delete',
        params
    })
}

// @Tags BeeOrderPeisong
// @Summary 更新beeOrderPeisong表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderPeisong true "更新beeOrderPeisong表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeOrderPeisong/updateBeeOrderPeisong [put]
export const updateBeeOrderPeisong = (data) => {
    return service({
        url: '/bee-shop/beeOrderPeisong/updateBeeOrderPeisong',
        method: 'put',
        data
    })
}

// @Tags BeeOrderPeisong
// @Summary 用id查询beeOrderPeisong表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeOrderPeisong true "用id查询beeOrderPeisong表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeOrderPeisong/findBeeOrderPeisong [get]
export const findBeeOrderPeisong = (params) => {
    return service({
        url: '/bee-shop/beeOrderPeisong/findBeeOrderPeisong',
        method: 'get',
        params
    })
}

// @Tags BeeOrderPeisong
// @Summary 用id查询配送状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeOrderPeisong true "用id查询beeOrderPeisong表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeOrderPeisong/getBeeOrderPeisongDetail [get]
export const getBeeOrderPeisongDetail = (params) => {
    return service({
        url: '/bee-shop/beeOrderPeisong/getBeeOrderPeisongDetail',
        method: 'get',
        params
    })
}

// @Tags BeeOrderPeisong
// @Summary 分页获取beeOrderPeisong表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取beeOrderPeisong表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeOrderPeisong/getBeeOrderPeisongList [get]
export const getBeeOrderPeisongList = (params) => {
    return service({
        url: '/bee-shop/beeOrderPeisong/getBeeOrderPeisongList',
        method: 'get',
        params
    })
}
