# beego框架：mvc+路由配置

### 学习beego框架+rom映射

```
## 1.执行beego
go get github.com/astaxie/beego

## 2.安装beego开发工具 bee.exe
go get github.com/beego/bee

## 安装报错请参考
http://www.678fly.cn/d/3-go-mod-bee

## 3.在项目目录下创建项目 beego_project
bee new beego_project

## 4.执行项目
cd beego_project
bee run

## 5.执行成功，通过http访问网站
http://localhost:8080/

## 6.数据库:下载beego的orm包
go get github.com/astaxie/beego/orm

## 7.打印日志：下载日志包
go get github.com/astaxie/beego/logs


## 8.前台页面展示使用 jsonform (static/jsonform)
https://github.com/jsonform/jsonform
核心库：static/jsonform/lib/jsonform.js

## 9.go语言的json解析库
https://github.com/bitly/go-simplejson

## 10.beego 打包部署
1.命令进入项目文件夹目录下
2.set GOOS=  //设置环境                 （darwin/linux/windows）
  set GOARCH= //设置cpu                 (一般 amd64)
3.go build  会生成 项目名 的一个文件      (go项目一般执行)
3.bee pack  会生成一个 项目名.tar.gz 文件 (beego项目执行),
  打包完成，该文件独立于gopath环境变量，可以放到任意服务器启动项目
4.启动打包文件
  nohup ./项目名 &
5.查看是否开启：
  tail nohup.out

```

### 文件分层
```
MVC架构--Beego框架
MySql封装--ORM
登录&权限--Session应用
反向代理&负载均衡--Nginx高可用


beego_project

├── conf   配置文件夹
│   └── app.conf  配置│      
├── controller 控制器
│   └── default.go 控制器（业务逻辑层）
├── models 数据访问层
│   └── page.go beego框架的orm使用
├── router 路由器
│   └── router.go 路由配置
├── static
    └── css  
    └── img 
    └── js
├── tests 测试文件夹
    └── default_test.go  default功能测试
├── views
    └── index.tpl  index界面模板
├── main.go 主进程 在此运行 bee run ，开启项目

```


```
开发中遇到的问题：
1.orm:
    orm跟mysql的配置从conf中取；init()只能调用一次；否则会被覆盖
    type DataModel struct {
    	Did        int    `orm:"pk;auto;column(Did)"`
    }
    column必须写，否则会出现返回的数据字段不对应的问题
2.目前存在view传页面json未处理，所以页面没办法正常显示数据。

```
