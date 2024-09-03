/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/6
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

type SysRole interface {
	Create(req *reqSystem.CreateRoleReq) error
	Delete(id int) error
	Update(id int, req *reqSystem.CreateRoleReq) error
	List() (error, interface{})
	Get(id int) (error, *system.Role)
}
type sysRole struct {
	tips string
	ctx  context.Context
}

func NewSysRole(ctx context.Context) SysRole {
	return &sysRole{ctx: ctx, tips: "角色"}
}

func (sr *sysRole) Create(req *reqSystem.CreateRoleReq) error {
	role := new(system.Role)
	if err := copier.Copy(&role, req); err != nil {
		return global.OtherErr(errors.New("转换角色数据失败"))
	}
	role.CreateBy = sr.ctx.Value("username").(string)
	if err := global.GORM.WithContext(sr.ctx).Create(&role).Error; err != nil {
		return global.CreateErr(sr.tips, err)
	}
	return sr.association(role, req.Users, req.Menus)
}

func (sr *sysRole) Delete(id int) error {
	role := new(system.Role)
	if err := global.GORM.WithContext(sr.ctx).First(&role, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return global.NotFoundErr(sr.tips, err)
		}
		return err
	}
	if err := global.GORM.WithContext(sr.ctx).Delete(&role, id).Error; err != nil {
		return global.DeleteErr(sr.tips, err)
	}
	sr.clear(role)
	return nil
}

func (sr *sysRole) Update(id int, req *reqSystem.CreateRoleReq) error {
	role := new(system.Role)
	if err := copier.Copy(&role, req); err != nil {
		return global.OtherErr(errors.New("转换角色数据失败"))
	}
	if err := global.GORM.WithContext(sr.ctx).Where("id = ?", id).Updates(&role).Error; err != nil {
		return global.UpdateErr(sr.tips, err)
	}
	err, r := sr.Get(id)
	if err != nil {
		return err
	}
	return sr.association(r, req.Users, req.Menus)
}

func (sr *sysRole) List() (error, interface{}) {
	var resRole []system.Role
	if err := global.GORM.WithContext(sr.ctx).Model(&system.Role{}).
		Preload("Users").Find(&resRole).Error; err != nil {
		return global.GetErr(sr.tips, err), nil
	}
	return nil, &resRole
}
func (sr *sysRole) Get(id int) (error, *system.Role) {
	var roles system.Role
	if err := global.GORM.WithContext(sr.ctx).Model(&system.Role{}).Where("id = ?", id).Preload("Users").
		Preload("Menus").First(&roles).Error; err != nil {
		return global.GetErr(sr.tips, err), nil
	}
	return nil, &roles
}

func (sr *sysRole) association(role *system.Role, userIDs, menuIDs []int) error {
	var (
		users []system.User
		menus []system.Menu
	)
	if len(userIDs) > 0 {
		global.GORM.WithContext(sr.ctx).Where("id IN ?", userIDs).Find(&users)
		if len(userIDs) != len(users) {
			return global.OtherErr(errors.New("部分users不存在"))
		}
		if err := global.GORM.WithContext(sr.ctx).Model(role).Association("Users").Replace(users); err != nil {
			return global.OtherErr(err)
		}
	}
	if len(menuIDs) > 0 {
		global.GORM.WithContext(sr.ctx).Where("id IN ?", menuIDs).Find(&menus)
		if len(menuIDs) != len(menus) {
			return global.OtherErr(errors.New("部分menus不存在"))
		}
		if err := global.GORM.WithContext(sr.ctx).Model(role).Association("Menus").Replace(menus); err != nil {
			return global.OtherErr(err)
		}
	}
	return nil
}

func (sr *sysRole) clear(roles ...*system.Role) {
	for _, role := range roles {
		_ = global.GORM.WithContext(sr.ctx).Model(role).Association("Users").Clear()
		_ = global.GORM.WithContext(sr.ctx).Model(role).Association("Menus").Clear()
	}
}
