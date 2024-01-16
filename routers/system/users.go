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

func InitUserRouters(r *gin.RouterGroup) gin.IRoutes {
	{
		r.GET("/user/getUserInfo", system.GetUserInfo)
		r.GET("/user/getUserList", system.UserList)
		r.POST("/user/updateUser", system.UserUpdate)
		r.POST("/user/addUser", system.UserAdd)
	}
	return r
}
