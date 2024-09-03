
<div align="center">
<h1>Go Easy Admin</h1>
</div><br>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.17-blue"/>
<img src="https://img.shields.io/badge/gin-1.9.0-lightBlue"/>
<img src="https://img.shields.io/badge/casbin-2.37.4-brightgreen"/>
<img src="https://img.shields.io/badge/viper-1.16.0-green"/>
<img src="https://img.shields.io/badge/gorm-1.25.2-red"/>

<p>基于Gin+Gorm实现非常简单的脚手架</p>
</div>

## 项目介绍

`go-easy-admin`是一个非常简单的`gin+gorm`脚手架，非常适合学习完`golang`基础的同学来进行练习使用。其中角色、权限都已经设计好，我们只需要关注业务接口即可。

## 目录结构

```bash
go-easy-admin
待补充
```

## 中间件casbin
```shell
go get github.com/casbin/gorm-adapter/v3
go get github.com/casbin/casbin/v2
```

## 快速开始

### 拉取代码

```bash
git clone  https://github.com/kubesre/go-easy-admin.git`
``

### 修改配置文件

```bash

cat  config.yaml
server:
  port: 8899
  address: 127.0.0.1
  name: go-easy-admin
  # # 生产环境建议使用 release，debug：可以使用debug模式
  model: release
  adminUser: admin
  adminPwd: 25285442ebc7d3a0c20047e01d341c31   # 密码为 123456

# 数据库配置
mysql:
  DbHost: 127.0.0.1
  DbPort: 3306
  # 数据库名称 需要提前创建好
  DbName: go-easy-admin
  DbUser: root
  DbPwd: pwd@123456
  MaxIdleConns: 10
  MaxOpenConns: 100
  # 是否开启debug，1 开启 0 关闭
  ActiveDebug: 0

# 密码加密
aes:
  key: go-easy-admin

jwt:
  realm: go-easy-admin
  # jwt加密因子
  key: anruo
  #  jwt token过期时间 单位为小时
  timeout: 100
  # jwt token刷新时间 单位为小时
  maxRefresh: 1
```

### 执行MySQL初始化脚本
script/*.sql

暂未补充完整
### 启动服务

```bash
go run main.go
```
