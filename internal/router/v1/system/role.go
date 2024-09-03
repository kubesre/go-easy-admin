/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/6
*/

package system

import (
	"github.com/gin-gonic/gin"

	apiSystem "go-easy-admin/api/system"
)

func Role(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/create", apiSystem.NewSysRole().Create)
		r.POST("/delete/:id", apiSystem.NewSysRole().Delete)
		r.POST("/update/:id", apiSystem.NewSysRole().Update)
		r.GET("/list", apiSystem.NewSysRole().List)
		r.GET("/get/:id", apiSystem.NewSysRole().Get)
	}
	return r
}
