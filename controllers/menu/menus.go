/*
auth: AnRuo
source: 云原生运维圈
time: 2023/12/1
*/

package menu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-easy-admin/common/global"
	"go-easy-admin/models"
	"go-easy-admin/service/menu"
)

// 添加菜单

func AddMenus(ctx *gin.Context) {
	params := new(models.Menu)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("添加菜单参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err := menu.NewMenusInterface().AddMenus(params)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "添加菜单成功！！！")
}

// 获取菜单列表

func ListMenus(ctx *gin.Context) {
	fmt.Println("获取菜单列表", ctx.GetString("username"))
	data, err := menu.NewMenusInterface().MenusList()
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)
}
