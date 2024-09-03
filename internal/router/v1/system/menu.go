/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/5
*/

package system

import (
	"github.com/gin-gonic/gin"

	apiSystem "go-easy-admin/api/system"
)

func Menu(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/create", apiSystem.NewSysMenu().Create)
		r.POST("/delete/:id", apiSystem.NewSysMenu().Delete)
		r.POST("/update/:id", apiSystem.NewSysMenu().Update)
		r.GET("/list", apiSystem.NewSysMenu().List)
		r.GET("/get/:id", apiSystem.NewSysMenu().Get)
	}
	return r
}
