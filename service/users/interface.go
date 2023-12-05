package users

import "go-easy-admin/models"

type InterfaceUsers interface {
	Register(user *models.User) error
	UserInfo(id string) (*models.User, error)
	UserList(username string, limit, page int) (*models.UserList, error)
}
type userInfo struct{}

func NewUserInfo() InterfaceUsers {
	return &userInfo{}
}
