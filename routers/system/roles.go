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

func InitRolesRouters(r *gin.RouterGroup) gin.IRoutes {
	{
		r.GET("/role/getRoleInfo", system.RolesInfo)
		r.POST("/role/addRole", system.AddRole)
		r.POST("/role/updateRole", system.UpdateRole)
		r.POST("/role/addBindMenu", system.AddRelationRoleAndMenu)
		r.POST("/role/deleteRole", system.DelRole)
		r.GET("/role/getRoleList", system.ListRole)
	}
	return r
}
