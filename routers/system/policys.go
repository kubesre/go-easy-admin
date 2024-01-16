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

func InitPolicyRouters(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/policy/createPolicy", system.AddCasbin)
		r.POST("/policy/deletePolicy", system.DelPolicy)
		r.GET("/policy/getPolicyList", system.ListPolicy)
	}
	return r
}
