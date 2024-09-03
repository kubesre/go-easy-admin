/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/7
*/

package logins

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	apiLogin "go-easy-admin/api/logins"
)

func Login(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	{
		r.POST("/ldap", authMiddleware.LoginHandler)
		r.POST("/general", authMiddleware.LoginHandler)
	}
	return r
}

func Resource(r *gin.RouterGroup) gin.IRoutes {
	{
		r.GET("/info", apiLogin.NewSysLogin().GetLoginUserResource)
	}
	return r
}
