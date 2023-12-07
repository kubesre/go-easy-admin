/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package role

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/common/global"
	"go-easy-admin/models"
	"go-easy-admin/service/role"
)

// 创建角色

func AddRole(ctx *gin.Context) {
	params := new(models.Role)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("创建角色参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	if err := role.NewRoleInterface().AddRole(params); err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "创建角色成功！！！")

}

// 获取角色详情

func RolesInfo(ctx *gin.Context) {
	idStr, _ := ctx.GetQuery("id")
	data, err := role.NewRoleInterface().RoleInfo(idStr)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)
}

// 更新角色

func UpdateRole(ctx *gin.Context) {
	idStr, _ := ctx.GetQuery("id")
	params := new(models.Role)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("更新角色参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err := role.NewRoleInterface().UpdateRole(idStr, params)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "更新角色成功")

}

// 创建角色对应的菜单

func AddRelationRoleAndMenu(ctx *gin.Context) {
	params := new(struct {
		MenuID []int `json:"menu_id" form:"menu_id"`
		RoleID []int `json:"role_id" form:"role_id"`
	})
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("创建角色对应的菜单参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err := role.NewRoleInterface().AddRelationRoleAndMenu(params.MenuID, params.RoleID)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "绑定成功")
}

// 删除角色

func DelRole(ctx *gin.Context) {
	params := new(struct {
		RoleID []int `json:"role_id" form:"role_id"`
	})
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("创建角色对应的菜单参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err := role.NewRoleInterface().DelRole(params.RoleID)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "删除角色成功")
}

// 角色列表

func ListRole(ctx *gin.Context) {
	data, err := role.NewRoleInterface().RoleList()
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)
}
