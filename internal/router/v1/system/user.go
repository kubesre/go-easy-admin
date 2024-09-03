/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/2
*/

package system

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	apiSystem "go-easy-admin/api/system"
)

func User(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	{
		r.POST("/create", apiSystem.NewSysUser().Create)
		r.POST("/delete/:id", apiSystem.NewSysUser().Delete)
		r.POST("/update/:id", apiSystem.NewSysUser().Update)
		r.GET("/list", apiSystem.NewSysUser().List)
		r.GET("/get/:id", apiSystem.NewSysUser().Get)
		r.POST("/logout", authMiddleware.LogoutHandler)
		r.POST("/refresh", authMiddleware.RefreshHandler)
	}
	return r
}
