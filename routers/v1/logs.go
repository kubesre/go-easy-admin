/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package v1

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/controllers/operationLogs"
)

func InitLogRouters(r *gin.RouterGroup) gin.IRoutes {
	{
		r.GET("/log/list", operationLogs.GetOperationLogList)
	}
	return r
}
