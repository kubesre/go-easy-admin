/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package dept

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/common/global"
	"go-easy-admin/models"
	"go-easy-admin/service/dept"
)

// 创建部门

func AddDept(ctx *gin.Context) {
	params := new(models.Dept)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("创建部门参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	if err := dept.NewDeptInterface().AddDept(params); err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "部门创建成功")
}

// 查看部门列表

func ListDept(ctx *gin.Context) {
	data, err := dept.NewDeptInterface().DeptList()
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)
}

// 部门详情

func InfoDept(ctx *gin.Context) {
	didStr, _ := ctx.GetQuery("id")
	data, err := dept.NewDeptInterface().DeptInfo(didStr)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)
}

// 删除部门

func DelDept(ctx *gin.Context) {
	params := new(struct {
		DeptID int `json:"dept_id"`
	})
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("删除部门参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err := dept.NewDeptInterface().DelDept(params.DeptID)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "删除部门成功！！！")

}
