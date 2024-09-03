/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/5
*/

package system

import (
	"strconv"

	"github.com/gin-gonic/gin"

	reqSystem "go-easy-admin/internal/model/request/system"
	"go-easy-admin/pkg/controller/system"
	"go-easy-admin/pkg/global"
)

type MenuInterface interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type sysMenu struct{}

func NewSysMenu() MenuInterface {
	return &sysMenu{}
}

func (sm *sysMenu) Create(ctx *gin.Context) {
	body := new(reqSystem.CreateMenuReq)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err := system.NewSysMenu(ctx).Create(body); err != nil {
		global.ReturnContext(ctx).Failed("创建失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("创建成功", nil)
}

func (sm *sysMenu) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := system.NewSysMenu(ctx).Delete(id); err != nil {
		global.ReturnContext(ctx).Failed("删除失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("删除成功", nil)
}

func (sm *sysMenu) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	body := new(reqSystem.CreateMenuReq)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err := system.NewSysMenu(ctx).Update(id, body); err != nil {
		global.ReturnContext(ctx).Failed("更新失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("更新成功", nil)

}

func (sm *sysMenu) List(ctx *gin.Context) {
	err, menus := system.NewSysMenu(ctx).List()
	if err != nil {
		global.ReturnContext(ctx).Failed("查询失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("查询成功", menus)
}

func (sm *sysMenu) Get(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err, menu := system.NewSysMenu(ctx).Get(id)
	if err != nil {
		global.ReturnContext(ctx).Failed("查询失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("查询成功", menu)

}
