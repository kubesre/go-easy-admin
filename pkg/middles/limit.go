/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/27
*/

package middles

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"

	"go-easy-admin/pkg/global"
)

func RateLimitMiddle(fillInterval time.Duration, capacity int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucket(fillInterval, capacity)
	return func(ctx *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			global.ReturnContext(ctx).Failed("failed", "访问限流")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
