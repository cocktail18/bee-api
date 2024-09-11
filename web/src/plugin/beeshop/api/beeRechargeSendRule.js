
import service from '@/utils/request'

// @Tags BeeRechargeSendRule
// @Summary 创建beeRechargeSendRule表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeRechargeSendRule true "创建beeRechargeSendRule表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeRechargeSendRule/createBeeRechargeSendRule [post]
export const createBeeRechargeSendRule = (data) => {
    return service({
        url: '/bee-shop/beeRechargeSendRule/createBeeRechargeSendRule',
        method: 'post',
        data
    })
}

// @Tags BeeRechargeSendRule
// @Summary 删除beeRechargeSendRule表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeRechargeSendRule true "删除beeRechargeSendRule表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeRechargeSendRule/deleteBeeRechargeSendRule [delete]
export const deleteBeeRechargeSendRule = (params) => {
    return service({
        url: '/bee-shop/beeRechargeSendRule/deleteBeeRechargeSendRule',
        method: 'delete',
        params
    })
}

// @Tags BeeRechargeSendRule
// @Summary 批量删除beeRechargeSendRule表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除beeRechargeSendRule表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeRechargeSendRule/deleteBeeRechargeSendRule [delete]
export const deleteBeeRechargeSendRuleByIds = (params) => {
    return service({
        url: '/bee-shop/beeRechargeSendRule/deleteBeeRechargeSendRuleByIds',
        method: 'delete',
        params
    })
}

// @Tags BeeRechargeSendRule
// @Summary 更新beeRechargeSendRule表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeRechargeSendRule true "更新beeRechargeSendRule表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeRechargeSendRule/updateBeeRechargeSendRule [put]
export const updateBeeRechargeSendRule = (data) => {
    return service({
        url: '/bee-shop/beeRechargeSendRule/updateBeeRechargeSendRule',
        method: 'put',
        data
    })
}

// @Tags BeeRechargeSendRule
// @Summary 用id查询beeRechargeSendRule表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeRechargeSendRule true "用id查询beeRechargeSendRule表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeRechargeSendRule/findBeeRechargeSendRule [get]
export const findBeeRechargeSendRule = (params) => {
    return service({
        url: '/bee-shop/beeRechargeSendRule/findBeeRechargeSendRule',
        method: 'get',
        params
    })
}

// @Tags BeeRechargeSendRule
// @Summary 分页获取beeRechargeSendRule表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取beeRechargeSendRule表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeRechargeSendRule/getBeeRechargeSendRuleList [get]
export const getBeeRechargeSendRuleList = (params) => {
    return service({
        url: '/bee-shop/beeRechargeSendRule/getBeeRechargeSendRuleList',
        method: 'get',
        params
    })
}
