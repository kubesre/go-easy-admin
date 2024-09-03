/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/8
*/

package middles

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-easy-admin/pkg/controller/system"
	"go-easy-admin/pkg/global"
)

func RbacMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctxUser := ctx.GetString("username")
		if ctxUser == "admin" {
			ctx.Next()
		} else {
			// 获取用户ID
			userID := ctx.Keys["id"]
			if userID == nil {
				global.ReturnContext(ctx).Failed("failed", "用户未登录")
				ctx.Abort()
				return
			}
			id, _ := strconv.Atoi(fmt.Sprintf("%d", userID))
			// 获取用户信息
			// TODO 缓存角色信息
			err, userInfo := system.NewSysUser(ctx).Get(id)
			if err != nil {
				global.ReturnContext(ctx).Failed("failed", "用户不存在")
				ctx.Abort()
				return
			}
			var sub []uint
			for _, role := range userInfo.Roles {
				if role.Status != 1 {
					continue
				}
				sub = append(sub, role.ID)
			}
			// 获取请求路径
			obj := ctx.FullPath()
			// 获取请求方法
			act := ctx.Request.Method
			for _, s := range sub {
				success, _ := global.CasbinCacheEnforcer.Enforce(strconv.Itoa(int(s)), obj, act)
				if success {
					ctx.Next()
					return
				}
			}
			global.ReturnContext(ctx).Failed("failed", "权限不足")
			ctx.Abort()
			return

		}
	}
}
