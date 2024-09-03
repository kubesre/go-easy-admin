/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/8
*/

package system

import (
	"strconv"

	"github.com/gin-gonic/gin"

	reqSystem "go-easy-admin/internal/model/request/system"
	"go-easy-admin/pkg/controller/system"
	"go-easy-admin/pkg/global"
)

type ApisInterface interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetApiGroup(ctx *gin.Context)
}

type sysApis struct{}

func NewSysApis() ApisInterface {
	return &sysApis{}
}

func (sa *sysApis) Create(ctx *gin.Context) {
	body := new(reqSystem.CreateAPIsReq)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err := system.NewSysApis(ctx).Create(body); err != nil {
		global.ReturnContext(ctx).Failed("创建失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("创建成功", nil)
}

func (sa *sysApis) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := system.NewSysApis(ctx).Delete(id); err != nil {
		global.ReturnContext(ctx).Failed("删除失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("删除成功", nil)
}

func (sa *sysApis) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	body := new(reqSystem.UpdateAPIsReq)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err := system.NewSysApis(ctx).Update(id, body); err != nil {
		global.ReturnContext(ctx).Failed("更新失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("更新成功", nil)
}

func (sa *sysApis) List(ctx *gin.Context) {
	if err, data := system.NewSysApis(ctx).List(); err != nil {
		global.ReturnContext(ctx).Failed("查询失败", err.Error())
		return
	} else {
		global.ReturnContext(ctx).Successful("查询成功", data)
	}
}

func (sa *sysApis) Get(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err, data := system.NewSysApis(ctx).Get(id); err != nil {
		global.ReturnContext(ctx).Failed("查询失败", err.Error())
		return
	} else {
		global.ReturnContext(ctx).Successful("查询成功", data)
	}

}
func (sa *sysApis) GetApiGroup(ctx *gin.Context) {
	if err, data := system.NewSysApis(ctx).GetApiGroup(); err != nil {
		global.ReturnContext(ctx).Failed("删除失败", err.Error())
		return
	} else {
		global.ReturnContext(ctx).Successful("查询成功", data)
	}
}
