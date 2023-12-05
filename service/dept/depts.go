package dept

import (
	"errors"
	"go-easy-admin/common/global"
	"go-easy-admin/dao"
	"go-easy-admin/models"
	"strconv"
)

// 创建部门

func (d *deptInfo) AddDept(dept *models.Dept) error {
	err := dao.NewDeptInterface().AddDept(dept)
	if err != nil {
		global.TPLogger.Error("创建部门失败： ", err)
		return errors.New("创建部门失败")
	}
	global.TPLogger.Info("创建部门成功！！！")
	return nil
}

// 部门列表

func (d *deptInfo) DeptList() ([]models.Dept, error) {
	data, err := dao.NewDeptInterface().DeptList()
	if err != nil {
		global.TPLogger.Error("查看部门列表失败: ", err)
		return nil, errors.New("查看部门列表失败")
	}
	return data, nil
}

// 部门详情

func (d *deptInfo) DeptInfo(did string) ([]models.Dept, error) {
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
