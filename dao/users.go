package dao

import (
	"errors"
	"go-easy-admin/common/global"
	"go-easy-admin/models"
	"gorm.io/gorm"
)

// 用户相关

type InterfaceUsers interface {
	ExitUser(userName, password string) (bool, uint)
	Register(user *models.User) error
	UserInfo(id uint) (*models.User, error)
	UserList(username string, limit, page int) (*models.UserList, error)
	GetUserFromUserName(userName string) (*models.User, error)
}

type userInfo struct{}

func NewUserInterface() InterfaceUsers {
	return &userInfo{}
}

// 判断用户是否存在，用户登录

func (u *userInfo) ExitUser(userName, password string) (bool, uint) {
	var user models.User
	err := global.GORM.Where("username = ? AND password = ? AND status = ?", userName, password, true).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, 0
	}
	return true, user.ID
}

// 用户注册

func (u *userInfo) Register(user *models.User) error {
	if err := global.GORM.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// 用户详情

func (u *userInfo) UserInfo(id uint) (*models.User, error) {
	var user models.User
	err := global.GORM.Where("id = ?", id).Preload("Role").Preload("Dept").First(&user).Error
	return &user, err
}

// 用户列表

func (u *userInfo) UserList(username string, limit, page int) (*models.UserList, error) {
	// 定义分页起始位置
	startSet := (page - 1) * limit
	var (
		userList []models.User
		total    int64
	)
	if err := global.GORM.Model(&models.User{}).Where("username LIKE ?", username).Count(&total).
		Limit(limit).Offset(startSet).Order("id desc").Find(&userList).Error; err != nil {
		return nil, err
	}
	return &models.UserList{
		Items: userList,
		Total: total,
	}, nil
}

// 用户查询

func (u *userInfo) GetUserFromUserName(userName string) (*models.User, error) {
	var user models.User
	err := global.GORM.Where("username = ?", userName).Preload("Role").Preload("Dept").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
