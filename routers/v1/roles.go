/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package v1

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/controllers/role"
)

func InitRolesRouters(r *gin.RouterGroup) gin.IRoutes {
	{
		r.GET("/role/info", role.RolesInfo)
		r.POST("/role/add", role.AddRole)
		r.POST("/role/update", role.UpdateRole)
		r.POST("/role/bind_menu", role.AddRelationRoleAndMenu)
		r.POST("/role/del", role.DelRole)
		r.GET("/role/list", role.ListRole)
	}
	return r
}
