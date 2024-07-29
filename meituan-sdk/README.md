## 一、环境说明 
美团开放平台Go语言版本SDK，支持go 1.13.7及以上版本。

## 二、引入方式
下载sdk包，添加到项目依赖

## 三、调用示例

以接口[门店本地验券历史](https://developer.meituan.com/docs/api/tuangou-coupon-queryLocalListByDate) 为例， SDK提供了***CouponQueryListByDateRequest***来封装请求，您调用此接口的代码可以参考：

```go
import (
  "meituan/gosdk/model/tuangoung/coupon/couponquerylistbydate"
  "meituan/gosdk/mtclient"
)
// 使用developerId和signKey初始化client，建议使用单例
client := mtclient.NewDefaultClient(100567, "***signKey***")
appAuthToken := "xxxxxxxxxx"

var request couponquerylistbydate.CouponQueryListByDateRequest
request.Date = "2022-01-01"
request.Offset = 0
request.Limit = 10

response, err := request.DoInvoke(client, appAuthToken)

if err != nil {
	fmt.Printf("接口调用失败:%v\n", err)
} else {
	if response.IsSuccess() {
		data := response.Data
		fmt.Printf("接口调用成功，得到响应:%v\n", data)
	} else {
		fmt.Printf("接口调用响应异常, code: %s, msg: %s\n", response.Code, response.Msg)
	}
}
```

## 四、第三方业务授权获取和刷新Token

```go
import (
"fmt"
"meituan/gosdk/mtclient"
"meituan/gosdk/oauth"
)

client := mtclient.NewDefaultClient(100567, "***signKey***")
getTokenRequest := oauth.NewGetTokenRequest(33, "***authCode***")
response, err := getTokenRequest.GetToken(client)
if err != nil {
	fmt.Printf("接口调用失败:%v\n", err)
} else {
	fmt.Printf("接口调用成功，得到响应:%v\n", response.Data)
}
```

```go
import (
"fmt"
"meituan/gosdk/mtclient"
"meituan/gosdk/oauth"
)

client := mtclient.NewDefaultClient(100567, "***signKey***")
refreshTokenRequest := oauth.NewRefreshTokenRequest(33, "***refreshToken***")
response, err := getTokenRequest.TokenRefresh(client)
if err != nil {
	fmt.Printf("接口调用失败:%v\n", err)
} else {
	fmt.Printf("接口调用成功，得到响应:%v\n", response.Data)
}
```