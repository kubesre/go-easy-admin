/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package v1

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/controllers/users"
)

func InitUserRouters(r *gin.RouterGroup) gin.IRoutes {
	{
		r.GET("/user/info", users.GetUserInfo)
		r.GET("/user/list", users.UserList)
		r.POST("/user/update", users.UserUpdate)
		r.POST("/user/add", users.UserAdd)
	}
	return r
}
