/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package system

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/common/global"
	"go-easy-admin/controllers/system"
	"go-easy-admin/middles"
)

// 基础路由

func InitBaseRouters(r *gin.RouterGroup) gin.IRoutes {
	authMiddleware, err := middles.InitAuth()
	if err != nil {
		global.TPLogger.Error("初始化JWT认证中间件失败：", err)
		panic(err)
	}
	{
		r.POST("/login", authMiddleware.LoginHandler)     // 登录
		r.POST("/logout", authMiddleware.LogoutHandler)   // 退出
		r.POST("/refresh", authMiddleware.RefreshHandler) // 刷新令牌
		r.POST("/register", system.Register)              // 注册
	}
	return r
}
