/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/9
*/

package system

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"go-easy-admin/pkg/controller/system"
	"go-easy-admin/pkg/global"
)

type SysRBAC interface {
	Create(ctx *gin.Context)
	GetRbacByRoleID(ctx *gin.Context)
}
type sysRbac struct {
}

func NewSysRBAC() SysRBAC {
	return &sysRbac{}
}

func (sr *sysRbac) Create(ctx *gin.Context) {
	body := new(struct {
		ApisID []int `json:"apis_id"`
	})
	roleID, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err := system.NewSysRBAC(ctx).Create(body.ApisID, roleID); err != nil {
		global.ReturnContext(ctx).Failed("创建失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("创建成功", nil)
}

func (sr *sysRbac) GetRbacByRoleID(ctx *gin.Context) {
	roleID, _ := strconv.Atoi(ctx.Param("id"))
	roleIDs := system.NewSysRBAC(ctx).GetRbacByRoleID(roleID)
	global.ReturnContext(ctx).Successful("获取已经授权角色成功", roleIDs)
}
