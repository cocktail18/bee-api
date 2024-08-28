
import service from '@/utils/request'

// @Tags BeeDelivery
// @Summary 创建beeDelivery表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeDelivery true "创建beeDelivery表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeDelivery/createBeeDelivery [post]
export const createBeeDelivery = (data) => {
    return service({
        url: '/bee-shop/beeDelivery/createBeeDelivery',
        method: 'post',
        data
    })
}

// @Tags BeeDelivery
// @Summary 删除beeDelivery表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeDelivery true "删除beeDelivery表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeDelivery/deleteBeeDelivery [delete]
export const deleteBeeDelivery = (params) => {
    return service({
        url: '/bee-shop/beeDelivery/deleteBeeDelivery',
        method: 'delete',
        params
    })
}

// @Tags BeeDelivery
// @Summary 批量删除beeDelivery表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除beeDelivery表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeDelivery/deleteBeeDelivery [delete]
export const deleteBeeDeliveryByIds = (params) => {
    return service({
        url: '/bee-shop/beeDelivery/deleteBeeDeliveryByIds',
        method: 'delete',
        params
    })
}

// @Tags BeeDelivery
// @Summary 更新beeDelivery表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeDelivery true "更新beeDelivery表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeDelivery/updateBeeDelivery [put]
export const updateBeeDelivery = (data) => {
    return service({
        url: '/bee-shop/beeDelivery/updateBeeDelivery',
        method: 'put',
        data
    })
}

// @Tags BeeDelivery
// @Summary 用id查询beeDelivery表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeDelivery true "用id查询beeDelivery表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeDelivery/findBeeDelivery [get]
export const findBeeDelivery = (params) => {
    return service({
        url: '/bee-shop/beeDelivery/findBeeDelivery',
        method: 'get',
        params
    })
}

// @Tags BeeDelivery
// @Summary 分页获取beeDelivery表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取beeDelivery表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeDelivery/getBeeDeliveryList [get]
export const getBeeDeliveryList = (params) => {
    return service({
        url: '/bee-shop/beeDelivery/getBeeDeliveryList',
        method: 'get',
        params
    })
}
