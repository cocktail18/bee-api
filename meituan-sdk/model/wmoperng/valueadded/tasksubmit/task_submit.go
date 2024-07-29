package tasksubmit

/**
* 任务提交
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const task_submit_url = "/wmoper/ng/valueadded/async/taskSubmit"

type TaskSubmitRequest struct {
    /**
    *  业务请求参数 json格式，具体参数说明如下 对于外卖商家端来说，此接口的封装逻辑是： 商家端通过通过用户中心提供的手机号换取用户id（ getUserByMobileWithMsg ）。 然后调用通过用户id发券接口发券（ thrift接口sendFullDiscountCoupon ）。 1 根据系统标签获取用户列表,json {"labelId":xx} 2 根据自定义标签获取用户列表，json 如： { "useType":1,// 用户类型 "sex":"1",//性别 "orderNum":[2,3],// 频次 "timeStart":"2020-10-31", "timeEnd":"2020-11-01", "commentContext":"1,2,3,4" } 3.建券并发券，json 具体属性如下： 参数名 类型 是否必须 示例值 描述 wmPoiId String 否 12233 APP方门店ID appId int 否 123 三方ID couponName String 是 优惠券 优惠券名称 price int 是 2 优惠券金额单位元 limitPrice int 是 6 优惠券门槛金额单位元 validTime long 是 3600 有效期，单位秒 comment String 是 备注 备注 tel String task_type=5时，必传。 13000000000 支持11位大陆手机号和港澳台手机号。 手机号格式为{地区码}_{手机号}；如果是大陆手机号（86），手机号格式为{手机号}。 备注：即将废弃的手机号发券接口已新增 注意： 1. 手机号需要校验大陆手机号和港澳台手机号。 2. 对于服务商侧接口来说， appId应该改成developerId，wmPoiId应该改成epoiid，但从外卖底层接口能力来说，即便输入的参数，也会按请求参数中的公共参数为准，所以，对于研发侧来说，透传字段；对于文档展示来说，将隐藏着两个字段的展示。 
    */
    TaskParam string `json:"taskParam"`
    /**
    *  任务名字 当task_type=5时，task_name 与 OpenFullDiscountCouponSendReq 的 couponName 字段值一致。 
    */
    TaskName string `json:"taskName"`
    /**
    *  任务备注 当task_type=5时，task_comment 与OpenFullDiscountCouponSendReq 的 comment 字段值一致。 
    */
    TaskComment string `json:"taskComment"`
    /**
    *  枚举: 2-根据系统标签获取用户列表 3-根据自定以标签获取用户列表 4-建券并发券 5-手机号发券 
    */
    TaskType int32 `json:"taskType"`
}
type TaskSubmitResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 任务id。此次提交任务的ID，开放平台可以回调用户url推送数据，用户也可以根据这个ID去调用具体的任务查询接口。
    */
    Data int64 `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *TaskSubmitRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TaskSubmitResponse, error) {
    resp, err := client.InvokeApi(task_submit_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TaskSubmitResponse
    err = json.Unmarshal(resp, &response)
    if err != nil {
        return nil, err
    }

    if response.IsSuccess() {
        return &response, nil
    } else {
        return &response, &api_error.ApiError{Code: response.Code, Msg: response.Msg, TraceId: response.TraceId}
    }
}

func (response *TaskSubmitResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

