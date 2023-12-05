/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package middles

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/common/global"
	"go-easy-admin/dao"
	"strconv"
	"strings"
)

func CasbinMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取username
		ctxUser := c.GetString("username")
		// TODO 从缓存中获取用户相关的信息，例如：role、dept、menu
		// 从数据库中获取用户角色信息sub
		usersInfo, err := dao.NewUserInterface().GetUserFromUserName(ctxUser)
		// TODO 这里需要修改为rile_id
		sub := usersInfo.RoleId
		//获取请求路径 这里要注意一下 请求的api不能使用 xxx/:id 这种的了
		obj := strings.Split(c.Request.RequestURI, "?")[0]
		// 获取请求方法
		act := c.Request.Method
		success, err := global.CasbinEnforcer.Enforce(strconv.Itoa(int(sub)), obj, act)
		if err != nil || !success {
			global.TPLogger.Error("权限验证失败：", err, success)
			global.ReturnContext(c).Failed("failed", "权限验证失败")
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
