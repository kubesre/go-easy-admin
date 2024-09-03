/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/5
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

type SysMenu interface {
	Create(req *reqSystem.CreateMenuReq) error
	Delete(id int) error
	Update(id int, req *reqSystem.CreateMenuReq) error
	List() (error, interface{})
	Get(id int) (error, *system.Menu)
}
type sysMenu struct {
	tips string
	ctx  context.Context
}

func NewSysMenu(ctx context.Context) SysMenu {
	return &sysMenu{ctx: ctx, tips: "菜单"}
}

func (sm *sysMenu) Create(req *reqSystem.CreateMenuReq) error {
	menu := new(system.Menu)
	if err := copier.Copy(&menu, req); err != nil {
		return global.OtherErr(errors.New("转换菜单数据失败"), err.Error())
	}
	menu.CreateBy = sm.ctx.Value("username").(string)
	if err := global.GORM.WithContext(sm.ctx).Create(&menu).Error; err != nil {
		return global.CreateErr(sm.tips, err)
	}
	//if len(req.APIs) > 0 {
	//	return sm.association(menu, req.APIs)
	//}
	return nil
}

// 级联删除菜单

//var MenuSlice []*system.Menu

func (sm *sysMenu) Delete(id int) error {
	menu := new(system.Menu)
	if err := global.GORM.WithContext(sm.ctx).Where("id = ?", id).First(&menu).Error; err != nil {
		// 菜单不存在
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return global.NotFoundErr(sm.tips, err)
		}
		return err
	}
	// 判断是否存在子级菜单
	if hasChildren(menu) {
		return global.OtherErr(errors.New("存在子级菜单"), "请先删除子级菜单")
	}
	if err := global.GORM.WithContext(sm.ctx).Delete(menu).Error; err != nil {
		return global.DeleteErr(sm.tips, err)
	}
	// 清空关联
	sm.clear(menu)
	return nil
}

func (sm *sysMenu) Update(id int, req *reqSystem.CreateMenuReq) error {
	menu := new(system.Menu)
	if err := copier.Copy(&menu, req); err != nil {
		return global.OtherErr(errors.New("转换菜单数据失败"), err.Error())
	}
	if err := global.GORM.WithContext(sm.ctx).Where("id = ?", id).Updates(&menu).Error; err != nil {
		return global.UpdateErr(sm.tips, err)
	}
	//err, m := sm.Get(id)
	//if err != nil {
	//	return err
	//}
	//if len(req.APIs) > 0 {
	//	return sm.association(m, req.APIs)
	//}
	return nil
}

func (sm *sysMenu) List() (error, interface{}) {
	var menus []system.Menu
	if err := global.GORM.WithContext(sm.ctx).Where("parent_id = ?", 0).Find(&menus).Error; err != nil {
		return global.GetErr(sm.tips, err), nil
	}
	for i := range menus {
		err := GetChildren(&menus[i])
		if err != nil {
			return err, nil
		}
	}
	return nil, menus
}
func (sm *sysMenu) Get(id int) (error, *system.Menu) {
	var menu system.Menu
	if err := global.GORM.WithContext(sm.ctx).Where("id = ?", id).First(&menu).Error; err != nil {
		return global.GetErr(sm.tips, err), nil
	}
	if err := GetChildren(&menu); err != nil {
		return err, nil
	}
	return nil, &menu
}

// 获取子菜单

func GetChildren(menu *system.Menu) error {
	if err := global.GORM.Where("parent_id = ?", menu.ID).
		Find(&menu.Children).Error; err != nil {
		return global.OtherErr(errors.New("获取子菜单失败"), err.Error())
	}
	//MenuSlice = append(MenuSlice, menu)
	for i := range menu.Children {
		if err := GetChildren(&menu.Children[i]); err != nil {
			return err
		}
	}
	return nil
}

func hasChildren(menu *system.Menu) bool {
	var count int64
	if err := global.GORM.Where("parent_id = ?", menu.ID).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true
		}
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

func (sm *sysMenu) association(menu *system.Menu, apisIDs []int) error {
	var apis []system.APIs
	if err := global.GORM.WithContext(sm.ctx).Where("id IN (?)", apisIDs).Find(&apis).Error; err != nil {
		return global.OtherErr(errors.New("获取APIs失败"), err.Error())
	}
	// 检查是否所有APIs都存在
	if len(apisIDs) != len(apis) {
		return global.OtherErr(errors.New("部分APIs不存在"))
	}
	if err := global.GORM.WithContext(sm.ctx).Model(menu).Association("APIs").Replace(apis); err != nil {
		return global.OtherErr(errors.New("关联APIs失败: " + err.Error()))
	}
	return nil
}

func (sm *sysMenu) clear(menus ...*system.Menu) {
	for _, menu := range menus {
		//_ = global.GORM.WithContext(sm.ctx).Model(&menu).Association("APIs").Clear()
		_ = global.GORM.WithContext(sm.ctx).Model(&menu).Association("Roles").Clear()
	}
	//MenuSlice = nil
}
