/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-easy-admin/common/global"
	"go-easy-admin/middles"
	"go-easy-admin/routers/system"
	"time"
)

func BaseRouters() *gin.Engine {
	r := gin.New()
	// 自定义日志格式
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("[%s | method: %s | path: %s | host: %s | proto: %s | code: %d | %s | %s ]\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.ClientIP,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
		)
	}))
	// 开启全部跨域允许
	r.Use(middles.Cors())
	// 健康检查
	r.GET("/health", func(ctx *gin.Context) {
		global.ReturnContext(ctx).Successful("success", "success")
		return
	})
	// 不需要做鉴权的接口 PublicGroup
	PublicGroup := r.Group("/api/base")
	{
		system.InitBaseRouters(PublicGroup)
	}
	// 需要做鉴权的接口
	PrivateGroup := r.Group("/api/system")
	// 鉴权
	//PrivateGroup.Use(gin.Recovery()).
	//	Use(middles.OperationLog()).Use(middles.CasbinMiddle())
	{
		system.InitUserRouters(PrivateGroup)
		system.InitRolesRouters(PrivateGroup)
		system.InitDeptRouters(PrivateGroup)
		system.InitMenusRouters(PrivateGroup)
		system.InitPolicyRouters(PrivateGroup)
		system.InitLogRouters(PrivateGroup)
	}
	r.NoRoute(func(ctx *gin.Context) {
		global.ReturnContext(ctx).Failed("fail", "该接口未开放")
		return
	})
	return r
}
