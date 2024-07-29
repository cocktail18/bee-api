package baseinfoquery

/**
* 门店信息查询（选接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const base_info_query_url = "/foodmop/shop/baseInfo/query"

type BaseInfoQueryRequest struct {
    /**
    *  商家ERP门店编码，最长30个字符 
    */
    ErpShopId string `json:"erpShopId"`
}
type PicUrlListMap struct {
}
type ShopInfoDTO struct {
    /**
    * 商家ERP上门店编码，最长30个字符
    */
    ErpShopId string `json:"erpShopId"`
    /**
    * 商家POS机上门店编码，用于做团购核销门店映射使用，可以与ERP编码一致
    */
    PosShopId string `json:"posShopId"`
    /**
    * 门店品类，传门店品类查询接口返回的品类id
    */
    CatId int32 `json:"catId"`
    /**
    * 门店名称 1 字符≤50个 2 不包含别名/外文名/分店名/电话/营业时间等其他信息 3 尽量与门店招牌保持高度一致，包括中英文、大小写 4 不出现门店招牌上没有的省份名/城市名/道路名/商圈名等 5 不包含门店招牌无关的优惠促销、宣传语等广告信息
    */
    ShopName string `json:"shopName"`
    /**
    * 分店名称 1 字符≤50个 2 不可含有非法关键词 3 不可与门店名称完全一致 4 分店名称中不能包含网址、脚本 5 不可为英文字母和数字的组合 6 可以为地理词+**店的形式：如徐汇店、长宁龙之梦店、徐汇旗舰店、长宁龙之梦二店 7 非该种形式的分店名请咨询美团对接负责人
    */
    BranchShopName string `json:"branchShopName"`
    /**
    * 城市，传行政区划查询到的地区代码，只需传区级代码
    */
    AreaCode int32 `json:"areaCode"`
    /**
    * 门店地址，加密传输，加解密方式参见下面示例代码 测试联调时返回固定值:点评实验室XXX号 1 3个字符≤长度≤100个字符 2 无特殊符号 3 不可含有非法关键词 4 地址不可包含网址、脚本 5 不可仅包含商户名和分店名 6 地址不包含所在的国家/省份/城市信息 7 当商户为沿街单独门店，则地址写法如下： 1）道路名+门牌号。例如：定西路1112号 2）道路交叉口+方位米数 。例如：中十路与萨尔图大街交叉口西50米 3）道路参照地标+方位米数。例如：和平西街审美造型南行10米 8 当商户为楼中店，则地址在以上写法后加上 所在建筑物+楼层或编号。例如：长宁路1018号龙之梦购物中心8楼 9 “靠近（交叉路）”字段以括号形式展示在地址字段后。例如：定西路1112号（民生银行斜对面）
    */
    Address string `json:"address"`
    /**
    * 经度(以高德地图为标准) 测试联调时返回固定值: 47.60355 1 以一组数字的形式存储，数值小数点至少保留5位 2 纬度范围（-90，90），经度范围（-180，180）例如：121.423684, 31.21221 3 经纬度偏离<50米
    */
    Longitude float64 `json:"longitude"`
    /**
    * 纬度(以高德地图为标准) 测试联调时传固定值: -78.44295 1 以一组数字的形式存储，数值小数点至少保留5位 2 纬度范围（-90，90），经度范围（-180，180）例如：121.423684, 31.21221 3 经纬度偏离<50米
    */
    Latitude float64 `json:"latitude"`
    /**
    * 门店电话，加密传输，加解密方式参见示例代码 1 支持两个号码，以英文逗号隔开 2 不可出现除“-”之外的特殊字符 3 每个号码≤30个字符 4 *固话/固话转机类型的电话号码，境内区号由系统自动添加，无需人工填写 5 不得为我司商服电话 到餐：10105557 6 不可电话号码递增，递减或连续重复大于9个数字
    */
    Phone string `json:"phone"`
    /**
    * 营业时间，按如下格式填写 全天营业 周一至周日 全天 单段时间 周一至周日 09:30-21:30 单段时间但跨凌晨 ① 周一至周日 16:00 -0:00 0:00-9:00 ② 周一至周日 16:00 - 9:00 营业时间分段 周一至周五 09:00-12:00 13:00-17:30 营业日期分段 周一至周五 09:00-12:00 13:00-17:30 周六,周日 09:00-18:00 营业季节分段 春季、冬季 周五至周日 09:30-21:30 周一至周四 09:30-21:00 夏季、秋季 周五
    */
    BusinessHour string `json:"businessHour"`
    /**
    * 营业状态，正常-10、尚未营业(筹建中)-20、暂停营业-30、永久关闭-40
    */
    OpenStatus int32 `json:"openStatus"`
    /**
    * 图片信息，返回门店头图信息，key为头图的默认值1，value为图片地址，可以通过图片上传接口获取 1 清晰，招牌字体可简单辨认 2 尽量正对门头拍照，不要上传侧边照或者角度比较歪的照片
    */
    PicUrlListMap PicUrlListMap `json:"picUrlListMap"`
}
type BaseInfoQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ShopInfoDTO `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *BaseInfoQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*BaseInfoQueryResponse, error) {
    resp, err := client.InvokeApi(base_info_query_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response BaseInfoQueryResponse
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

func (response *BaseInfoQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

