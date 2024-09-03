/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/6
*/

package system

import (
	"strconv"

	"github.com/gin-gonic/gin"

	reqSystem "go-easy-admin/internal/model/request/system"
	"go-easy-admin/pkg/controller/system"
	"go-easy-admin/pkg/global"
)

type RoleInterface interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type sysRole struct{}

func NewSysRole() RoleInterface {
	return &sysRole{}
}

func (sr *sysRole) Create(ctx *gin.Context) {
	body := new(reqSystem.CreateRoleReq)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err := system.NewSysRole(ctx).Create(body); err != nil {
		global.ReturnContext(ctx).Failed("创建失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("创建成功", nil)

}

func (sr *sysRole) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := system.NewSysRole(ctx).Delete(id); err != nil {
		global.ReturnContext(ctx).Failed("删除失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("删除成功", nil)
}

func (sr *sysRole) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	body := new(reqSystem.CreateRoleReq)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err := system.NewSysRole(ctx).Update(id, body); err != nil {
		global.ReturnContext(ctx).Failed("更新失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("更新成功", nil)
}

func (sr *sysRole) List(ctx *gin.Context) {
	if err, data := system.NewSysRole(ctx).List(); err != nil {
		global.ReturnContext(ctx).Failed("查询失败", err.Error())
		return
	} else {
		global.ReturnContext(ctx).Successful("查询成功", data)
	}
}

func (sr *sysRole) Get(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err, data := system.NewSysRole(ctx).Get(id); err != nil {
		global.ReturnContext(ctx).Failed("查询失败", err.Error())
		return
	} else {
		global.ReturnContext(ctx).Successful("查询成功", data)
	}
}
