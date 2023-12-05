/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package menu

import (
	"errors"
	"go-easy-admin/common/global"
	"go-easy-admin/dao"
	"go-easy-admin/models"
)

// 添加菜单

func (m *menusInfo) AddMenus(menu *models.Menu) error {
	err := dao.NewMenusInterface().AddMens(menu)
	if err != nil {
		global.TPLogger.Error("创建菜单失败：", err)
		return errors.New("创建菜单失败")
	}
	return nil
}

// 菜单列表

func (m *menusInfo) MenusList() ([]models.Menu, error) {
	data, err := dao.NewMenusInterface().MenusList()
	if err != nil {
		return nil, errors.New("获取菜单列表失败")
	}
	return data, nil
}
