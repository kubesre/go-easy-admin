/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package users

import "go-easy-admin/models"

type InterfaceUsers interface {
	Register(user *models.User) error
	UserInfo(id string) (*models.User, error)
	UserSearchList(username string, limit, page int) (*models.UserList, error)
	UserList(limit, page int) (*models.UserList, error)
}
type userInfo struct{}

func NewUserInfo() InterfaceUsers {
	return &userInfo{}
}
