/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/2
*/

package system

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	reqSystem "go-easy-admin/internal/model/request/system"
	"go-easy-admin/internal/model/system"
	"go-easy-admin/pkg/global"
)

type SysUser interface {
	Create(req *reqSystem.CreateUserReq) error
	Delete(id int) error
	Update(id int, req *reqSystem.UpdateUserReq) error
	List(userName string, limit, page int) (error, interface{})
	Get(id int) (error, *system.User)
	GetByUsername(username string) (bool, *system.User)
	GetByUsernameAndPwd(username, password string) (bool, *system.User)
}
type sysUser struct {
	ctx context.Context
}

func NewSysUser(ctx context.Context) SysUser {
	return &sysUser{ctx: ctx}
}

func (su *sysUser) Create(req *reqSystem.CreateUserReq) error {
	user := new(system.User)
	if err := copier.Copy(&user, req); err != nil {
		global.GeaLogger.Error("转换用户数据失败: ", err)
		return errors.New("转换用户数据失败")
	}
	if v, ok := su.ctx.Value("username").(string); ok {
		user.CreateBy = v
	} else {
		user.CreateBy = "LDAP"
	}
	if err := global.GORM.WithContext(su.ctx).Create(user).Error; err != nil {
		global.GeaLogger.Error("创建用户失败: ", err)
		return errors.New("创建用户失败")
	}
	// 处理多对多关系
	if len(req.Roles) > 0 {
		if err := su.association(user, req.Roles); err != nil {
			return err
		}
	}
	return nil
}

func (su *sysUser) Delete(id int) error {
	var users []*system.User
	if err := global.GORM.WithContext(su.ctx).Where("id = ?", id).First(&users).Error; err != nil {
		global.GeaLogger.Error("删除用户失败: ", err)
		return errors.New("删除用户失败")
	}
	su.clear(users...)
	if err := global.GORM.WithContext(su.ctx).Delete(&users, id).Error; err != nil {
		global.GeaLogger.Error("删除用户失败: ", err)
		return errors.New("删除用户失败")
	}
	return nil
}

func (su *sysUser) Update(id int, req *reqSystem.UpdateUserReq) error {
	user := new(system.User)
	err, u := su.Get(id)
	if err != nil {
		return err
	}
	if err = copier.Copy(&user, req); err != nil {
		global.GeaLogger.Error("转换用户数据失败: ", err)
		return errors.New("转换用户数据失败")
	}
	if err = global.GORM.WithContext(su.ctx).Model(&system.User{}).Where("id = ?", id).Updates(&user).Error; err != nil {
		global.GeaLogger.Error("更新用户失败: ", err.Error)
		return errors.New("更新用户失败")
	}
	if len(req.Roles) > 0 {
		if err = su.association(u, req.Roles); err != nil {
			return err
		}
	}
	return nil
}

func (su *sysUser) List(userName string, limit, page int) (error, interface{}) {
	startSet := (page - 1) * limit
	resUser := new(struct {
		Items []system.User
		Total int64
	})
	if err := global.GORM.WithContext(su.ctx).Model(&system.User{}).Where("username LIKE ?", "%"+userName+"%").
		Count(&resUser.Total).Preload("Roles").
		Limit(limit).Offset(startSet).Find(&resUser.Items).Error; err != nil {
		global.GeaLogger.Error("获取用户失败: ", err)
		return errors.New("获取用户失败"), nil
	}
	return nil, &resUser
}

func (su *sysUser) Get(id int) (error, *system.User) {
	var user system.User
	if err := global.GORM.WithContext(su.ctx).Model(system.User{}).
		Preload("Roles").
		Preload("Roles.Menus").
		Where("id = ?", id).
		First(&user).Error; err != nil {
		global.GeaLogger.Error("查询用户失败: ", err)
		return errors.New("查询用户失败"), nil
	}
	// 获取子菜单
	menuList := make([]system.Menu, 0)
	for i := range user.Roles {
		for j := range user.Roles[i].Menus {
			menuList = append(menuList, user.Roles[i].Menus[j])
		}
	}
	return nil, &user
}

func (su *sysUser) GetByUsername(username string) (bool, *system.User) {
	var user *system.User
	if err := global.GORM.WithContext(su.ctx).Model(&system.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, nil
	}
	return false, user
}

func (su *sysUser) GetByUsernameAndPwd(username, password string) (bool, *system.User) {
	var user *system.User
	if err := global.GORM.WithContext(su.ctx).Model(&system.User{}).Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, user
}

func (su *sysUser) association(user *system.User, roleIDs []int) error {
	var roles []system.Role
	if err := global.GORM.WithContext(su.ctx).Where("id IN ?", roleIDs).Find(&roles).Error; err != nil {
		global.GeaLogger.Error("查询角色失败: ", err)
		return errors.New("查询角色失败")
	}
	// 检查是否所有传入的角色 ID 都存在
	if len(roles) != len(roleIDs) {
		global.GeaLogger.Error("部分角色不存在")
		return errors.New("部分角色不存在")
	}

	if err := global.GORM.WithContext(su.ctx).Model(user).Association("Roles").Replace(roles); err != nil {
		global.GeaLogger.Error("关联角色失败: ", err)
		return errors.New("关联角色失败")
	}
	return nil
}

func (su *sysUser) clear(users ...*system.User) {
	for _, user := range users {
		err := global.GORM.WithContext(su.ctx).Model(&user).Association("Roles").Clear()
		if err != nil {
			continue
		}
	}
}
