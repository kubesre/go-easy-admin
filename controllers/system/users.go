/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package system

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/common/global"
	mod "go-easy-admin/models/system"
	service "go-easy-admin/service/system"
)

// 用户注册

func Register(ctx *gin.Context) {
	params := new(mod.User)
	if err := ctx.ShouldBind(&params); err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	err := service.NewUserInfo().Register(params)
	if err != nil {
		global.TPLogger.Error(err)
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", "用户注册成功！！！")
	return
}

// 用户详情

func GetUserInfo(ctx *gin.Context) {
	idStr, _ := ctx.GetQuery("id")
	data, err := service.NewUserInfo().UserInfo(idStr)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)
}

// 用户列表

func UserList(ctx *gin.Context) {
	params := new(struct {
		Name  string `form:"name"`
		Limit int    `form:"limit"`
		Page  int    `form:"page"`
	})
	if err := ctx.ShouldBindQuery(&params); err != nil {
		global.TPLogger.Error("用户查询数据绑定失败：", err)
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	data, err := service.NewUserInfo().UserList(params.Name, params.Limit, params.Page)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", data)

}

// 用户更新

func UserUpdate(ctx *gin.Context) {
	params := new(mod.User)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("用户更新数据绑定失败：", err)
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	err := service.NewUserInfo().UserUpdate(params)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", "用户更新成功")
}

// 用户添加

func UserAdd(ctx *gin.Context) {
	params := new(mod.User)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("用户添加绑定失败：", err)
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	err := service.NewUserInfo().UserAdd(params)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", "用户添加成功")
}
