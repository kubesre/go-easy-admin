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
)

type InterfaceMenus interface {
	AddMenus(menu *mod.Menu) error
	MenusList() ([]mod.Menu, error)
}

type menusInfo struct{}

func NewMenusInterface() InterfaceMenus {
	return &menusInfo{}
}

// 添加菜单

func (m *menusInfo) AddMenus(menu *mod.Menu) error {
	err := dao.NewMenusInterface().AddMens(menu)
	if err != nil {
		global.TPLogger.Error("创建菜单失败：", err)
		return errors.New("创建菜单失败")
	}
	return nil
}

// 菜单列表

func (m *menusInfo) MenusList() ([]mod.Menu, error) {
	data, err := dao.NewMenusInterface().MenusList()
	if err != nil {
		return nil, errors.New("获取菜单列表失败")
	}
	return data, nil
}
