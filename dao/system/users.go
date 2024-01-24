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
	"go-easy-admin/utils"
	"gorm.io/gorm"
)

// 用户相关

type InterfaceUsers interface {
	ExitUser(userName, password string) (*system.User, error)
	Register(user *system.User) error
	UserInfo(id uint) (*system.User, error)
	UserList(username string, limit, page int) (*system.UserList, error)
	GetUserFromUserName(userName string) (*system.User, error)
	UserUpdate(userData *system.User) error
	UserAdd(user *system.User) error
}

type userInfo struct{}

func NewUserInterface() InterfaceUsers {
	return &userInfo{}
}

// 判断用户是否存在，用户登录

func (u *userInfo) ExitUser(userName, password string) (*system.User, error) {
	var user *system.User
	encryptPassword, err := utils.EncryptAES(password)
	if err != nil {
		return nil, err
	}
	err = global.GORM.Where("username = ? AND password = ? AND status = ?", userName, encryptPassword, 1).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return user, err
}

// 用户注册

func (u *userInfo) Register(user *system.User) error {
	originPassword := user.Password
	encryptPassword, err := utils.EncryptAES(originPassword)
	if err != nil {
		return err
	}
	user.Password = encryptPassword
	if err := global.GORM.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// 用户详情

func (u *userInfo) UserInfo(id uint) (*system.User, error) {
	var user system.User
	err := global.GORM.Where("id = ?", id).Preload("Role").Preload("Dept").First(&user).Error
	return &user, err
}

// 用户列表

func (u *userInfo) UserList(username string, limit, page int) (*system.UserList, error) {
	// 定义分页起始位置
	startSet := (page - 1) * limit
	var (
		userList []system.User
		total    int64
	)
	if err := global.GORM.Model(&system.User{}).Where("username LIKE ?", "%"+username+"%").Preload("Role").
		Preload("Dept").Count(&total).
		Limit(limit).Offset(startSet).Order("id desc").Find(&userList).Error; err != nil {
		return nil, err
	}
	return &system.UserList{
		Items: userList,
		Total: total,
	}, nil
}

// 用户查询

func (u *userInfo) GetUserFromUserName(userName string) (*system.User, error) {
	var user system.User
	err := global.GORM.Where("username = ?", userName).Preload("Role").Preload("Dept").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 用户更新

func (u *userInfo) UserUpdate(userData *system.User) error {
	if err := global.GORM.Model(&system.User{}).Where("id = ?", userData.ID).Updates(&userData).Error; err != nil {
		return err
	}
	return nil
}

// 用户添加

func (u *userInfo) UserAdd(user *system.User) error {
	if user.Password != "" {
		originPassword := user.Password
		encryptPassword, err := utils.EncryptAES(originPassword)
		if err != nil {
			return err
		}
		user.Password = encryptPassword
	}
	if err := global.GORM.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
