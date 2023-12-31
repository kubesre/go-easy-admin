/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package role

import "go-easy-admin/models"

type InterfaceRole interface {
	AddRole(role *models.Role) error
	RoleInfo(rid string) (*models.Role, error)
	UpdateRole(rid string, roleData *models.Role) error
	AddRelationRoleAndMenu(menuID, roleID []int) error
	DelRole(rid []int) error
	RoleList() (roleData []*models.Role, err error)
}

type roleInfo struct{}

func NewRoleInterface() InterfaceRole {
	return &roleInfo{}
}
