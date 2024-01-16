/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package system

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/controllers/system"
)

func InitMenusRouters(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/menu/createMenu", system.AddMenus)
		r.GET("/menu/getMenuList", system.ListMenus)
	}
	return r
}
