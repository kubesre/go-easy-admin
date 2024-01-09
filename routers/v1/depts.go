/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package v1

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/controllers/dept"
)

func InitDeptRouters(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/dept/add", dept.AddDept)
		r.GET("/dept/list", dept.ListDept)
		r.GET("/dept/info", dept.InfoDept)
		r.POST("/dept/del", dept.DelDept)
	}
	return r
}
