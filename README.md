# bee-api

#### 介绍
微信小程序-餐饮点餐外卖商城的go后端实现，代码完全开源，商用请联系作者授权

针对餐饮行业推出的一套完整的餐饮解决方案，实现了用户在线点餐下单、外卖、叫号排队、支付、配送等功能，完美的使餐饮行业更高效便捷！

- 小程序端（原版）：https://github.com/woniudiancang/bee 
- 修改后的小程序端（建议使用）：https://github.com/cocktail18/bee


##### api 在线体验地址：
[https://bee.fbsst.top](https://bee.fbsst.top)

subDomain: `cocktailBeeOrder`

##### 管理后台在线体验地址：
[http://bee-admin.fbsst.top](http://bee-admin.fbsst.top)
账号：bee-demo  密码：123456

#### 软件架构

- api 使用 gin + gorm + mysql + redis
- 后台使用 gva 框架 https://www.gin-vue-admin.com


#### 目录说明
- server：管理后台代码
- web: 管理后台前端代码
- bee-api: api服务端代码


#### 安装教程

1.  安装go 1.22
2.  安装mysql or mariadb
3.  cd server && mv config.yaml.demo config.yaml && go run main.go  # 启动后端，如果没有科学上网先执行： `go env -w GOPROXY=https://goproxy.cn,direct`
4.  cd web && npm run serve # 启动前端
5.  点击前往初始化，根据需要配置账号密码、数据库账号密码
6.  重启下后端服务，(使用修改后的[小程序端](https://github.com/cocktail18/bee)可以跳过步奏7、8)
7.  将小程序的 config.js 的 subDomain 改成对应的subDomain, 默认是 cocktailBeeOrder
8.  将小程序的 miniprogram_npm/apifm-wxapi/index.js 的 API_BASE_URL 改为对应的域名， 默认开发环境：http://127.0.0.1:18083
9.  默认api端口：18083，管理后台：8888，可以使用nginx进行代理（建议）

#### 使用说明

1.  配置小程序APPID/SECRET （必须）
    管理后台-》bee 商城-》微信配置-》小程序配置 （没有则新增一条）
2.  配置微信支付
    管理后台-》bee 商城-》微信配置-》微信支付配置 （没有则新增一条）
3.  启动前端项目
4.  api项目已经继承在管理后台项目中，不需要单独启动


#### FAQ
1. 启动没有看到菜单

   a. 检查用户授权，进入后台-》超级管理员-》设置权限，勾选需要的权限，注意角色api 的也要勾选，如果不确定需要什么权限，直接全选即可
2. 小程序文件上传报错

   a.  搜索小程序的 miniprogram_npm/apifm-wxapi/index.js，将 `url: 'https://oss.apifm.com/upload2',` 替换成 `url: http://127.0.0.1:18083/upload2,`
3. 怎样使用云存储

    a.  修改配置文件 config.yaml   
```yaml
system:
    oss-type: local # 将这个改成你想要的云存储类型，目前支持：tencent-cos、

# 修改对应的云存储配置，比如说qiniu 
qiniu:
  zone: ZoneHuaDong
  bucket: YOUR_BUCKET_NAME
  img-path: http://xxxx.com
  access-key: YOUR_ACCESS_KEY
  secret-key: YOUR_SECRET_KEY
  use-https: false
  use-cdn-domains: false
```
目前支持的类型：
qiniu
tencent-cos
aliyun-oss
huawei-obs
aws-s3

4. 获取配送费失败，请检查配送配置: 无效的门店编号

    商店信息里面找到 `达达门店编号` 配置，填上即可

5. 如何配置dada配送

    a. 到达达官网申请开通商户号以及开发者账号 [http://newopen.imdada.cn/#/](http://newopen.imdada.cn/#/) 

    b. 在 `bee商城/运费相关/配送供应商配置` 配置页面增加一条配置,内容为上一步奏获得的 app_key 、app_secret 、app_id，测试的时候可以打开debug开关，防止无效扣费

    c. 在 `bee商城/商城基本信息/商店信息` 配置页面找到对应的商铺，`生鲜配送` 配置填`dada`,`达达门店编号`配置填第一步获得的达达门店编号，默认是用户支付之后就通知配送商，如果想要手动通知，将`配送接单需商家确认`配置设置为true

   d. 在 `bee商城/运费相关/配送信息` 配置页新增一条记录，内容为你想要的配送费配置, 注意不是运费模板模块，目前运费模板暂时没效果

6. dada配送、微信支付没有回调
   
    a.  这是因为本地用的是内网地址，外网无法访问，需要使用内网穿透或者发布到有外网的机器，同时需要修改`config.yaml` 配置文件， 参考：
    ```
    bee-shop:
       host: "改成你的域名"
    ```

   

### 后台截图de
![](imgs/demo01.jpg)
![](imgs/demo02.jpg)
![](imgs/demo03.jpg)
![](imgs/demo04.jpg)
![](imgs/demo05.jpg)
![](imgs/demo06.jpg)
![](imgs/demo07.jpg)
![](imgs/demo08.jpg)


#### QQ交流群 963437155
<img src="imgs/qqgroup.jpg" alt="qq交流群：963437155" style="width: 200px; height: auto;">


#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

