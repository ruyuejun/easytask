# 项目概述

本项目是使用gin，xorm等技术开发的RESTAPI服务器。    

修改配置（因为不能上传真实配置文件）：
- src/config文件夹下的`conf.go`中`ENV`设置为0
- 注释掉`conf.go`中`init`函数的`case1`部分

项目启动：
```
cd src
go run main.go
```
访问：localhost:3000/demo  

## 项目框架

`go version`:1.12，使用`go mod`包管理方式

- gin：核心框架
- xorm：orm框架
- log4go：日志框架
- go-redis：redis框架

## 项目结构

- bin   二进制文件目录
- pkg   安装文件目录
- src   源码文件目录
    - main.go       入口文件
    - go.mod        包管理文件
    - go.sum        包版本管理文件
    - route         路由目录
    - controller    控制层
    - service       服务层
    - model         模型层
    - middleware    中间件
    - utils         工具包
    
