
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
├─middles  --------------------中间件
├─models  ---------------------数据库表以及请求参数定义
├─routers  ---------------------路由
├─service  ---------------------业务逻辑
=======
## 中间件casbin
```shell
go get github.com/casbin/gorm-adapter/v3
go get github.com/casbin/casbin/v2
```

## 功能概述

![image-20231207164738020](https://gitee.com/root_007/md_file_image/raw/master/202312071647162.png)

> 其他功能前端还没有完成，暂时使用该图占位使用。

## 快速开始

### 拉取代码

```bash
git clone  https://github.com/kubesre/go-easy-admin.git`
``

### 修改配置文件

```bash
cd  go-easy-admin/config
mv  applicatino-example.yaml   application.yaml
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
```

### 启动服务

```bash
go run main.go
```
=======
注意： casbin在数据库中的数据，可以删除整条，但是不要将整条置为null
