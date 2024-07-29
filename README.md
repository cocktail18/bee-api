# bee-api

#### 介绍
微信小程序-餐饮点餐外卖商城的go后端实现，代码完全开源，商用请联系作者授权

针对餐饮行业推出的一套完整的餐饮解决方案，实现了用户在线点餐下单、外卖、叫号排队、支付、配送等功能，完美的使餐饮行业更高效便捷！

小程序端：https://github.com/woniudiancang/bee 


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
3.  mv server/config.yml.demo server/config.yml # 复制配置文件
4.  将 db-list 下面的数据库账号密码改成你的， 注意：alias-name 不要改
5.  cd server && go run main.go  # 启动后端
6.  cd web && npm run serve # 启动前端
7.  点击前往初始化，根据需要配置账号密码、数据库账号密码
8.  重启下后端服务
9.  将小程序的 config.js 的 subDomain 改成对应的subDomain, 默认是 cocktailBeeOrder
10.  将小程序的 miniprogram_npm/apifm-wxapi/index.js 的 API_BASE_URL 改为对应的域名， 默认开发环境：http://127.0.0.1:18083
11.  默认api端口：18083，管理后台：8888，可以使用nginx进行代理

#### 使用说明

1.  配置小程序APPID/SECRET （必须）
    管理后台-》bee 商城-》微信配置-》小程序配置 （没有则新增一条）
2.  配置微信支付
    管理后台-》bee 商城-》微信配置-》微信支付配置 （没有则新增一条）
3.  启动前端项目
4.  api项目已经继承在管理后台项目中，不需要单独启动

#### 后台截图demo
![](imgs/demo01.jpg)
![](imgs/demo02.jpg)
![](imgs/demo03.jpg)
![](imgs/demo04.jpg)
![](imgs/demo05.jpg)
![](imgs/demo06.jpg)
![](imgs/demo07.jpg)
![](imgs/demo08.jpg)


#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

