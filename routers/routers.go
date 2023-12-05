/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-easy-admin/common/global"
	"go-easy-admin/controllers/casbin"
	"go-easy-admin/controllers/dept"
	"go-easy-admin/controllers/menu"
	"go-easy-admin/controllers/role"
	"go-easy-admin/controllers/users"
	"go-easy-admin/middles"
	"time"
)

func BaseRouters() *gin.Engine {
	r := gin.New()
	// 自定义日志格式
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("[%s | method: %s | path: %s | host: %s | proto: %s | code: %d | %s | %s ]\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.ClientIP,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
		)
	}))
	// 初始化JWT认证中间件
	authMiddleware, err := middles.InitAuth()
	if err != nil {
		global.TPLogger.Error("初始化JWT认证中间件失败：", err)
		panic(err)
	}
	// 健康检查
	r.GET("/health", func(ctx *gin.Context) {
		global.ReturnContext(ctx).Successful("success", "success")
		return
	})
	// 不需要做鉴权的接口 PublicGroup
	PublicGroup := r.Group("/api/v1")
	{
		PublicGroup.POST("/login", authMiddleware.LoginHandler)
		PublicGroup.POST("/register", users.Register)
	}
	// 需要做鉴权的接口
	PrivateGroup := r.Group("/api/v1")
	// 鉴权
	PrivateGroup.Use(gin.Recovery()).Use(middles.Cors()).Use(authMiddleware.MiddlewareFunc()).
		Use(middles.OperationLog()).Use(middles.CasbinMiddle())
	{
		PrivateGroup.GET("/userinfo/", users.GetUserInfo)
		PrivateGroup.GET("/user/list", users.ListUser)
		PrivateGroup.GET("/role/", role.RolesInfo)
		PrivateGroup.POST("/role/add", role.AddRole)
		PrivateGroup.POST("/role/update/", role.UpdateRole)
		PrivateGroup.POST("/role/bind_menu", role.AddRelationRoleAndMenu)
		PrivateGroup.POST("/role/del", role.DelRole)
		PrivateGroup.POST("/dept/add", dept.AddDept)
		PrivateGroup.GET("/dept/list", dept.ListDept)
		PrivateGroup.GET("/dept/info/", dept.InfoDept)
		PrivateGroup.POST("/dept/del", dept.DelDept)
		PrivateGroup.POST("/menu/add", menu.AddMenus)
		PrivateGroup.GET("/menu/list", menu.ListMenus)
		PrivateGroup.POST("/policy/add", casbin.AddCasbin)
		PrivateGroup.POST("/policy/del", casbin.DelPolicy)
		PrivateGroup.GET("/policy/list", casbin.ListPolicy)
	}
	r.NoRoute(func(ctx *gin.Context) {
		global.ReturnContext(ctx).Failed("fail", "该接口未开放")
		return
	})
	return r
}
