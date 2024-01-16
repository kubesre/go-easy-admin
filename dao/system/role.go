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
	"go-easy-admin/models/system"
)

// 角色相关

type InterfaceRoles interface {
	AddRole(role *system.Role) error
	RoleInfo(rid uint) (*system.Role, error)
	UpdateRole(rid uint, roleData *system.Role) error
	AddRelationRoleAndMenu(menuID, roleID []int) error
	DelRole(rid []int) error
	RoleList() (roleData []*system.Role, err error)
}

type rolesInfo struct{}

func NewRolesInterface() InterfaceRoles {
	return &rolesInfo{}
}

// 创建角色

func (r *rolesInfo) AddRole(role *system.Role) error {
	err := global.GORM.Create(&role).Error
	return err
}

// 获取角色详情

func (r *rolesInfo) RoleInfo(rid uint) (*system.Role, error) {
	var roleData *system.Role
	err := global.GORM.Model(&system.Role{}).Where("id = ?", rid).Preload("Menus").First(&roleData).Error
	return roleData, err
}

// 更新角色信息

func (r *rolesInfo) UpdateRole(rid uint, roleData *system.Role) error {
	err := global.GORM.Model(&system.Role{}).Where("id = ?", rid).Updates(&roleData).Error
	return err
}

// 创建角色对应的菜单

func (r *rolesInfo) AddRelationRoleAndMenu(menuID, roleID []int) error {
	var (
		menuList  []system.Menu
		roleList  []system.Role
		menuTotal int64
		roleTotal int64
	)
	// 先查询是否存在 角色和菜单
	global.GORM.Model(&system.Menu{}).Where("id IN (?)", menuID).Count(&menuTotal)
	if int(menuTotal) < len(menuID) {
		return errors.New("菜单不存在")
	}
	global.GORM.Model(&system.Role{}).Where("id IN (?)", roleID).Count(&roleTotal)
	if int(roleTotal) < len(roleID) {
		return errors.New("角色不存在")
	}
	err := global.GORM.Model(&roleList).Association("Menus").Append(menuList)
	if err != nil {
		return err
	}
	return nil
}

// 删除角色

func (r *rolesInfo) DelRole(rid []int) error {
	var (
		roleData  []system.Role
		roleTotal int64
	)
	global.GORM.Find(&roleData, rid).Count(&roleTotal)
	if len(roleData) < len(rid) {
		return errors.New("角色列表中有不存在的ID")
	}
	// 清空角色与菜单的关系
	if err := global.GORM.Model(&roleData).Association("Menus").Clear(); err != nil {
		return errors.New("清空角色与菜单的关系失败:" + err.Error())
	}
	// 删除角色
	if err := global.GORM.Delete(&roleData, rid).Error; err != nil {
		return err
	}
	return nil
}

// 角色列表

func (r *rolesInfo) RoleList() (roleData []*system.Role, err error) {
	if err = global.GORM.Model(&system.Role{}).Find(&roleData).Error; err != nil {
		return nil, err
	}
	return roleData, nil
}
