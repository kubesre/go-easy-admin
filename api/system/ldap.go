/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/6
*/

package system

import (
	"github.com/gin-gonic/gin"

	modSys "go-easy-admin/internal/model/system"
	"go-easy-admin/pkg/controller/system"
	"go-easy-admin/pkg/global"
)

type LdapInterface interface {
	Create(ctx *gin.Context)
	Info(ctx *gin.Context)
	Ping(ctx *gin.Context)
}

type sysLdap struct{}

func NewSysLdap() LdapInterface {
	return &sysLdap{}
}

func (s *sysLdap) Create(ctx *gin.Context) {
	body := new(modSys.Ldap)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err := system.NewSysLdap(ctx).Create(body); err != nil {
		global.ReturnContext(ctx).Failed("创建或更新失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("创建或更新成功", nil)
}

func (s *sysLdap) Info(ctx *gin.Context) {
	err, ldap := system.NewSysLdap(ctx).Info()
	if err != nil {
		global.ReturnContext(ctx).Failed("查询失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("查询成功", ldap)
}

func (s *sysLdap) Ping(ctx *gin.Context) {
	body := new(modSys.Ldap)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	err := system.NewSysLdap(ctx).Ping(body)
	if err != nil {
		global.ReturnContext(ctx).Failed("连接失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("连接成功", nil)
}
