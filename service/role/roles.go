/*
auth: AnRuo
source: 云原生运维圈
time: 2023/12/1
*/

package role

import (
	"errors"
	"go-easy-admin/common/global"
	"go-easy-admin/dao"
	"go-easy-admin/models"
	"strconv"
)

// 创建角色

func (r *roleInfo) AddRole(role *models.Role) error {
	err := dao.NewRolesInterface().AddRole(role)
	if err != nil {
		global.TPLogger.Error("创建角色失败：", err)
		return errors.New("创建角色失败")
	}
	return nil
}

// 获取角色详情

func (r *roleInfo) RoleInfo(rid string) (*models.Role, error) {
	ridInt, _ := strconv.Atoi(rid)
	data, err := dao.NewRolesInterface().RoleInfo(uint(ridInt))
	if err != nil {
		global.TPLogger.Error("获取角色详情失败：", err)
		return nil, errors.New("获取角色详情失败")
	}
	return data, nil
}

// 更新角色信息

func (r *roleInfo) UpdateRole(rid string, roleData *models.Role) error {
	ridInt, _ := strconv.Atoi(rid)
	err := dao.NewRolesInterface().UpdateRole(uint(ridInt), roleData)
	if err != nil {
		global.TPLogger.Error("更新角色信息失败：", err)
		return errors.New("更新角色信息失败")
	}
	return nil
}

// 创建角色对应的菜单

func (r *roleInfo) AddRelationRoleAndMenu(menuID, roleID []int) error {
	err := dao.NewRolesInterface().AddRelationRoleAndMenu(menuID, roleID)
	if err != nil {
		global.TPLogger.Error("创建角色对应的菜单失败：", err)
		return err
	}
	return nil
}

// 删除角色

func (r *roleInfo) DelRole(rid []int) error {
	err := dao.NewRolesInterface().DelRole(rid)
	if err != nil {
		global.TPLogger.Error("删除角色失败：", err)
		return errors.New("删除角色失败")
	}
	return nil
}
