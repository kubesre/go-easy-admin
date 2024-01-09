/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package dao

import (
	"errors"
	"go-easy-admin/common/global"
	"go-easy-admin/models"
	"go-easy-admin/utils"
	"gorm.io/gorm"
)

// 用户相关

type InterfaceUsers interface {
	ExitUser(userName, password string) (bool, uint, uint)
	Register(user *models.User) error
	UserInfo(id uint) (*models.User, error)
	UserList(username string, limit, page int) (*models.UserList, error)
	GetUserFromUserName(userName string) (*models.User, error)
	UserUpdate(userData *models.User) error
	UserAdd(user *models.User) error
}

type userInfo struct{}

func NewUserInterface() InterfaceUsers {
	return &userInfo{}
}

// 判断用户是否存在，用户登录

func (u *userInfo) ExitUser(userName, password string) (bool, uint, uint) {
	var user models.User
	encryptPassword, err := utils.EncryptAES(password)
	if err != nil {
		return false, 0, 0
	}
	err = global.GORM.Where("username = ? AND password = ? AND status = ?", userName, encryptPassword, 1).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, 0, 0
	}
	return true, user.ID, user.RoleId
}

// 用户注册

func (u *userInfo) Register(user *models.User) error {
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
	if err := global.GORM.Model(&models.User{}).Where("username LIKE ?", "%"+username+"%").Preload("Role").
		Preload("Dept").Count(&total).
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

// 用户更新

func (u *userInfo) UserUpdate(userData *models.User) error {
	if err := global.GORM.Model(&models.User{}).Where("id = ?", userData.ID).Updates(&userData).Error; err != nil {
		return err
	}
	return nil
}

// 用户添加

func (u *userInfo) UserAdd(user *models.User) error {
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
