/*
auth: AnRuo
source: 云原生运维圈
time: 2023/12/1
*/

package menu

import "go-easy-admin/models"

// 菜单相关

type InterfaceMenus interface {
	AddMenus(menu *models.Menu) error
	MenusList() ([]models.Menu, error)
}

type menusInfo struct{}

func NewMenusInterface() InterfaceMenus {
	return &menusInfo{}
}
