/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package system

import (
	"errors"
	"go-easy-admin/common/global"
	dao "go-easy-admin/dao/system"
	mod "go-easy-admin/models/system"
	"strconv"
)

type InterfaceDept interface {
	AddDept(dept *mod.Dept) error
	DeptList() ([]mod.Dept, error)
	DeptInfo(did string) ([]mod.Dept, error)
	DelDept(did int) error
}

type deptInfo struct{}

func NewDeptInterface() InterfaceDept {
	return &deptInfo{}
}

// 创建部门

func (d *deptInfo) AddDept(dept *mod.Dept) error {
	err := dao.NewDeptInterface().AddDept(dept)
	if err != nil {
		global.TPLogger.Error("创建部门失败： ", err)
		return errors.New("创建部门失败")
	}
	global.TPLogger.Info("创建部门成功！！！")
	return nil
}

// 部门列表

func (d *deptInfo) DeptList() ([]mod.Dept, error) {
	data, err := dao.NewDeptInterface().DeptList()
	if err != nil {
		global.TPLogger.Error("查看部门列表失败: ", err)
		return nil, errors.New("查看部门列表失败")
	}
	return data, nil
}

// 部门详情

func (d *deptInfo) DeptInfo(did string) ([]mod.Dept, error) {
	didInt, err := strconv.Atoi(did)
	if err != nil {
		global.TPLogger.Error("查看部门详情失败：", err)
		return nil, errors.New("查看部门详情失败")
	}
	data, err := dao.NewDeptInterface().DeptInfo(didInt)
	if err != nil {
		global.TPLogger.Error("查看部门详情失败: ", err)
		return nil, errors.New("查看部门详情失败")
	}
	return data, nil
}

// 删除部门

func (d *deptInfo) DelDept(did int) error {
	err := dao.NewDeptInterface().DelDept(did)
	if err != nil {
		global.TPLogger.Error("删除部门失败：", err)
		return errors.New("删除部门失败")
	}
	return nil
}
