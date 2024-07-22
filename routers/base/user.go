package base

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 基础路由

func InitBaseRouters(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	{
		r.POST("/base/user/login", authMiddleware.LoginHandler) // 登录
	}
	return r
}
