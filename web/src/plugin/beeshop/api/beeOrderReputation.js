import service from '@/utils/request'

// @Tags BeeOrderReputation
// @Summary 创建商品评价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderReputation true "创建商品评价"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /beeOrderReputation/createBeeOrderReputation [post]
export const createBeeOrderReputation = (data) => {
  return service({
    url: '/bee-shop/beeOrderReputation/createBeeOrderReputation',
    method: 'post',
    data
  })
}

// @Tags BeeOrderReputation
// @Summary 删除商品评价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderReputation true "删除商品评价"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrderReputation/deleteBeeOrderReputation [delete]
export const deleteBeeOrderReputation = (params) => {
  return service({
    url: '/bee-shop/beeOrderReputation/deleteBeeOrderReputation',
    method: 'delete',
    params
  })
}

// @Tags BeeOrderReputation
// @Summary 批量删除商品评价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品评价"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /beeOrderReputation/deleteBeeOrderReputation [delete]
export const deleteBeeOrderReputationByIds = (params) => {
  return service({
    url: '/bee-shop/beeOrderReputation/deleteBeeOrderReputationByIds',
    method: 'delete',
    params
  })
}

// @Tags BeeOrderReputation
// @Summary 更新商品评价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BeeOrderReputation true "更新商品评价"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /beeOrderReputation/updateBeeOrderReputation [put]
export const updateBeeOrderReputation = (data) => {
  return service({
    url: '/bee-shop/beeOrderReputation/updateBeeOrderReputation',
    method: 'put',
    data
  })
}

// @Tags BeeOrderReputation
// @Summary 用id查询商品评价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BeeOrderReputation true "用id查询商品评价"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /beeOrderReputation/findBeeOrderReputation [get]
export const findBeeOrderReputation = (params) => {
  return service({
    url: '/bee-shop/beeOrderReputation/findBeeOrderReputation',
    method: 'get',
    params
  })
}

// @Tags BeeOrderReputation
// @Summary 分页获取商品评价列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品评价列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /beeOrderReputation/getBeeOrderReputationList [get]
export const getBeeOrderReputationList = (params) => {
  return service({
    url: '/bee-shop/beeOrderReputation/getBeeOrderReputationList',
    method: 'get',
    params
  })
}
