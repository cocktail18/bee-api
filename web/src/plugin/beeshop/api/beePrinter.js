import service from '@/utils/request'

// @Tags BeePrinter
// @Summary 创建beePrinter表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeePrinter true "创建beePrinter表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bee-shop/beePrinter/createBeePrinter [post]
export const createBeePrinter = (data) => {
    return service({
        url: '/bee-shop/beePrinter/createBeePrinter',
        method: 'post',
        data
    })
}

// @Tags BeePrinter
// @Summary 删除beePrinter表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeePrinter true "删除beePrinter表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bee-shop/beePrinter/deleteBeePrinter [delete]
export const deleteBeePrinter = (params) => {
    return service({
        url: '/bee-shop/beePrinter/deleteBeePrinter',
        method: 'delete',
        params
    })
}

// @Tags BeePrinter
// @Summary 批量删除beePrinter表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除beePrinter表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bee-shop/beePrinter/deleteBeePrinter [delete]
export const deleteBeePrinterByIds = (params) => {
    return service({
        url: '/bee-shop/beePrinter/deleteBeePrinterByIds',
        method: 'delete',
        params
    })
}

// @Tags BeePrinter
// @Summary 更新beePrinter表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeePrinter true "更新beePrinter表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bee-shop/beePrinter/updateBeePrinter [put]
export const updateBeePrinter = (data) => {
    return service({
        url: '/bee-shop/beePrinter/updateBeePrinter',
        method: 'put',
        data
    })
}

// @Tags BeePrinter
// @Summary 测试配置是否正确
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeePrinter true "测试配置是否正确"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /bee-shop/beePrinter/testBeePrinter [post]
export const testBeePrinter = (data) => {
    return service({
        url: '/bee-shop/beePrinter/testBeePrinter',
        method: 'post',
        data
    })
}

// @Tags BeePrinter
// @Summary 用id查询beePrinter表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeePrinter true "用id查询beePrinter表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bee-shop/beePrinter/findBeePrinter [get]
export const findBeePrinter = (params) => {
    return service({
        url: '/bee-shop/beePrinter/findBeePrinter',
        method: 'get',
        params
    })
}

// @Tags BeePrinter
// @Summary 分页获取beePrinter表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取beePrinter表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bee-shop/beePrinter/getBeePrinterList [get]
export const getBeePrinterList = (params) => {
    return service({
        url: '/bee-shop/beePrinter/getBeePrinterList',
        method: 'get',
        params
    })
}
