
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
├─app  ------------------------项目初始化操作
├─common  ------------------全局公用
├─config  ---------------------配置文件
├─controllers  ----------------控制层
├─dao  ------------------------数据库的CRUD
├─deployment  ---------------部署相关文件
├─doc  ----------------------项目文件相关说明
├─middles  --------------------中间件
├─models  ---------------------数据库表以及请求参数定义
├─routers  ---------------------路由
├─service  ---------------------业务逻辑
=======
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
cd  go-easy-admin/config
mv  application.yaml   application.yaml
cat  application.yaml
server:
  port: 8899
  address: 0.0.0.0
  name: go-easy-admin
  # # 生产环境建议使用release，debug：可以使用debug模式
  model: debug

mysql:
  DbHost: 127.0.0.1
  DbPort: 3306
  # 数据库名称 需要提前创建好
  DbName: go-easy-admin
  DbUser: root
  DbPwd: 123456
  MaxIdleConns: 10
  MaxOpenConns: 100
  # 是否开启debug，1 开启 0 关闭
  ActiveDebug: 1

jwt:
  realm: go-easy-admin
  # jwt加密因子
  key: anruo
  #  jwt token过期时间 单位为小时
  timeout: 1
  # jwt token刷新时间 单位为小时
  maxRefresh: 1
ldap:
  # ldap用户登录
  address: 127.0.0.1:389
  adminUser: cn=admin,dc=kubesre,dc=com
  baseDN: dc=kubesre,dc=com
  password: 123456
```

### 执行MySQL初始化脚本
`deployment/init.sql`文件为初始化`MySQL`数据库脚本，执行之后有对应的超级用户、角色、菜单等数据，方便项目启动之后进行测试

### 启动服务

```bash
go run main.go
```
