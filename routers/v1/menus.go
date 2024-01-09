/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package v1

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/controllers/menu"
)

func InitMenusRouters(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/menu/add", menu.AddMenus)
		r.GET("/menu/list", menu.ListMenus)
	}
	return r
}
