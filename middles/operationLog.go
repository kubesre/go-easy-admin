package middles

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-easy-admin/models"
	"time"
)

// 操纵日志

var OperationLogChan = make(chan *models.OperationLog, 30)

func OperationLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		if viper.GetInt("operation.ActiveLog") == 1 {
			// 开始时间
			startTime := time.Now()
			// 处理请求
			c.Next()
			// 结束时间
			endTime := time.Now()
			// 执行耗时
			timeCost := endTime.Sub(startTime).Milliseconds()
			// 获取当前操作用户
			var userName string
			ctxUser, _ := c.Get("username")
			if ctxUser == "" {
				userName = "未登录"
				c.Abort()
			} else {
				userName, _ = ctxUser.(string)
			}
			// 获取访问路径
			path := c.FullPath()
			// 获取请求方式
			method := c.Request.Method
			operationLog := models.OperationLog{
				Username:   userName,
				Ip:         c.ClientIP(),
				IpLocation: "",
				Method:     method,
				Path:       path,
				Remark:     "",
				Status:     c.Writer.Status(),
				StartTime:  startTime,
				TimeCost:   timeCost,
			}
			OperationLogChan <- &operationLog
		} else {
			c.Next()
		}
	}
}
