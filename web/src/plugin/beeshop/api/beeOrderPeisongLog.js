import service from '@/utils/request'

// @Tags BeeOrderPeisongLog
// @Summary 创建beeOrderPeisongLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderPeisongLog true "创建beeOrderPeisongLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeOrderPeisongLog/createBeeOrderPeisongLog [post]
export const createBeeOrderPeisongLog = (data) => {
    return service({
        url: '/bee-shop/beeOrderPeisongLog/createBeeOrderPeisongLog',
        method: 'post',
        data
    })
}

// @Tags BeeOrderPeisongLog
// @Summary 删除beeOrderPeisongLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderPeisongLog true "删除beeOrderPeisongLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrderPeisongLog/deleteBeeOrderPeisongLog [delete]
export const deleteBeeOrderPeisongLog = (params) => {
    return service({
        url: '/bee-shop/beeOrderPeisongLog/deleteBeeOrderPeisongLog',
        method: 'delete',
        params
    })
}

// @Tags BeeOrderPeisongLog
// @Summary 批量删除beeOrderPeisongLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除beeOrderPeisongLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrderPeisongLog/deleteBeeOrderPeisongLog [delete]
export const deleteBeeOrderPeisongLogByIds = (params) => {
    return service({
        url: '/bee-shop/beeOrderPeisongLog/deleteBeeOrderPeisongLogByIds',
        method: 'delete',
        params
    })
}

// @Tags BeeOrderPeisongLog
// @Summary 更新beeOrderPeisongLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderPeisongLog true "更新beeOrderPeisongLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeOrderPeisongLog/updateBeeOrderPeisongLog [put]
export const updateBeeOrderPeisongLog = (data) => {
    return service({
        url: '/bee-shop/beeOrderPeisongLog/updateBeeOrderPeisongLog',
        method: 'put',
        data
    })
}

// @Tags BeeOrderPeisongLog
// @Summary 用id查询beeOrderPeisongLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeOrderPeisongLog true "用id查询beeOrderPeisongLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeOrderPeisongLog/findBeeOrderPeisongLog [get]
export const findBeeOrderPeisongLog = (params) => {
    return service({
        url: '/bee-shop/beeOrderPeisongLog/findBeeOrderPeisongLog',
        method: 'get',
        params
    })
}

// @Tags BeeOrderPeisongLog
// @Summary 分页获取beeOrderPeisongLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取beeOrderPeisongLog表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeOrderPeisongLog/getBeeOrderPeisongLogList [get]
export const getBeeOrderPeisongLogList = (params) => {
    return service({
        url: '/bee-shop/beeOrderPeisongLog/getBeeOrderPeisongLogList',
        method: 'get',
        params
    })
}
