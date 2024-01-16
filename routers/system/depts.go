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

func InitDeptRouters(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/dept/addDept", system.AddDept)
		r.GET("/dept/getDeptList", system.ListDept)
		r.GET("/dept/getDeptInfo", system.InfoDept)
		r.POST("/dept/deleteDept", system.DelDept)
	}
	return r
}
