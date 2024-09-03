
import service from '@/utils/request'

// @Tags BeeQueue
// @Summary 创建beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeQueue true "创建beeQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeQueue/createBeeQueue [post]
export const createBeeQueue = (data) => {
    return service({
        url: '/bee-shop/beeQueue/createBeeQueue',
        method: 'post',
        data
    })
}

// @Tags BeeQueue
// @Summary 删除beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeQueue true "删除beeQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeQueue/deleteBeeQueue [delete]
export const deleteBeeQueue = (params) => {
    return service({
        url: '/bee-shop/beeQueue/deleteBeeQueue',
        method: 'delete',
        params
    })
}

// @Tags BeeQueue
// @Summary 批量删除beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除beeQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeQueue/deleteBeeQueue [delete]
export const deleteBeeQueueByIds = (params) => {
    return service({
        url: '/bee-shop/beeQueue/deleteBeeQueueByIds',
        method: 'delete',
        params
    })
}

// @Tags BeeQueue
// @Summary 更新beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeQueue true "更新beeQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeQueue/updateBeeQueue [put]
export const updateBeeQueue = (data) => {
    return service({
        url: '/bee-shop/beeQueue/updateBeeQueue',
        method: 'put',
        data
    })
}

// @Tags BeeQueue
// @Summary 用id查询beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeQueue true "用id查询beeQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeQueue/findBeeQueue [get]
export const findBeeQueue = (params) => {
    return service({
        url: '/bee-shop/beeQueue/findBeeQueue',
        method: 'get',
        params
    })
}

// @Tags BeeQueue
// @Summary 分页获取beeQueue表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取beeQueue表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeQueue/getBeeQueueList [get]
export const getBeeQueueList = (params) => {
    return service({
        url: '/bee-shop/beeQueue/getBeeQueueList',
        method: 'get',
        params
    })
}

// @Tags BeeQueue
// @Summary 叫号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeQueue true "更新beeQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeQueue/callBeeQueue [put]
export const callBeeQueue = (data) => {
    return service({
        url: '/bee-shop/beeQueue/callBeeQueue',
        method: 'put',
        data
    })
}


// @Tags BeeQueue
// @Summary 重新叫号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeQueue true "更新beeQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeQueue/reCallBeeQueue [put]
export const reCallBeeQueue = (data) => {
    return service({
        url: '/bee-shop/beeQueue/reCallBeeQueue',
        method: 'put',
        data
    })
}

// @Tags BeeQueue
// @Summary 重新叫号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeQueue true "更新beeQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeQueue/passBeeQueue [put]
export const passBeeQueue = (data) => {
    return service({
        url: '/bee-shop/beeQueue/passBeeQueue',
        method: 'put',
        data
    })
}

// @Tags BeeQueue
// @Summary 叫下一个号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeQueue true "更新beeQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeQueue/nextBeeQueue [put]
export const nextBeeQueue = (data) => {
    return service({
        url: '/bee-shop/beeQueue/nextBeeQueue',
        method: 'put',
        data
    })
}
