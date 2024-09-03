
import service from '@/utils/request'

// @Tags BeeUserQueue
// @Summary 创建beeUserQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserQueue true "创建beeUserQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeUserQueue/createBeeUserQueue [post]
export const createBeeUserQueue = (data) => {
    return service({
        url: '/bee-shop/beeUserQueue/createBeeUserQueue',
        method: 'post',
        data
    })
}

// @Tags BeeUserQueue
// @Summary 删除beeUserQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserQueue true "删除beeUserQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUserQueue/deleteBeeUserQueue [delete]
export const deleteBeeUserQueue = (params) => {
    return service({
        url: '/bee-shop/beeUserQueue/deleteBeeUserQueue',
        method: 'delete',
        params
    })
}

// @Tags BeeUserQueue
// @Summary 批量删除beeUserQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除beeUserQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeUserQueue/deleteBeeUserQueue [delete]
export const deleteBeeUserQueueByIds = (params) => {
    return service({
        url: '/bee-shop/beeUserQueue/deleteBeeUserQueueByIds',
        method: 'delete',
        params
    })
}

// @Tags BeeUserQueue
// @Summary 更新beeUserQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeUserQueue true "更新beeUserQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeUserQueue/updateBeeUserQueue [put]
export const updateBeeUserQueue = (data) => {
    return service({
        url: '/bee-shop/beeUserQueue/updateBeeUserQueue',
        method: 'put',
        data
    })
}

// @Tags BeeUserQueue
// @Summary 用id查询beeUserQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeUserQueue true "用id查询beeUserQueue表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeUserQueue/findBeeUserQueue [get]
export const findBeeUserQueue = (params) => {
    return service({
        url: '/bee-shop/beeUserQueue/findBeeUserQueue',
        method: 'get',
        params
    })
}

// @Tags BeeUserQueue
// @Summary 分页获取beeUserQueue表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取beeUserQueue表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeUserQueue/getBeeUserQueueList [get]
export const getBeeUserQueueList = (params) => {
    return service({
        url: '/bee-shop/beeUserQueue/getBeeUserQueueList',
        method: 'get',
        params
    })
}
