/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package v1

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/controllers/casbin"
)

func InitPolicyRouters(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/policy/add", casbin.AddCasbin)
		r.POST("/policy/del", casbin.DelPolicy)
		r.GET("/policy/list", casbin.ListPolicy)
	}
	return r
}
