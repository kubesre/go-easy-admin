/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/9
*/

package system

import (
	"context"
	"fmt"
	"strconv"

	gormadapter "github.com/casbin/gorm-adapter/v3"

	"go-easy-admin/internal/model/system"
	"go-easy-admin/pkg/global"
)

// casbin 权限

type SysRBAC interface {
	Create(apiIDs []int, roleID int) error
	GetRbacByRoleID(roleID int) []int
	DeleteByAPIsID(apiID int) error
	UpdateByAPI(api *system.APIs) error
}
type sysRbac struct {
	tips string
	ctx  context.Context
}

func NewSysRBAC(ctx context.Context) SysRBAC {
	return &sysRbac{ctx: ctx, tips: "权限"}
}

// 添加权限

func (sr *sysRbac) Create(apiIDs []int, roleID int) error {
	var casbinRules []gormadapter.CasbinRule
	for _, id := range apiIDs {
		err, api := NewSysApis(sr.ctx).Get(id)
		if err != nil {
			continue
		}
		if api.Desc == "" {
			api.Desc = "系统添加描述信息"
		}
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    fmt.Sprintf("%d", roleID),
			V1:    api.Path,
			V2:    api.Method,
			V3:    api.Desc,
			V4:    sr.ctx.Value("username").(string),
			V5:    fmt.Sprintf("%d", api.ID),
		})
	}
	// 清空角色对应权限
	sr.DeleteByRoleID(roleID)
	if err := global.GORM.WithContext(sr.ctx).Create(&casbinRules).Error; err != nil {
		return err
	}
	freshRBAC()
	return nil

}

func (sr *sysRbac) UpdateByAPI(api *system.APIs) error {
	var updateCasbinRule = gormadapter.CasbinRule{
		V1: api.Path,
		V2: api.Method,
		V3: api.Desc,
	}
	if err := global.GORM.Model(&gormadapter.CasbinRule{}).WithContext(sr.ctx).
		Where("v5 = ?", api.ID).
		Updates(&updateCasbinRule).Error; err != nil {
		return err
	}
	freshRBAC()
	return nil
}

func (sr *sysRbac) DeleteByRoleID(roleID int) {
	// 根据ID查找权限
	var casbinRules []gormadapter.CasbinRule
	global.GORM.Model(&casbinRules).WithContext(sr.ctx).Where("v0 = ?", roleID).Find(&casbinRules)
	if len(casbinRules) < 1 {
		return
	}
	for _, rule := range casbinRules {
		_, _ = global.CasbinCacheEnforcer.RemovePolicy(rule.V0, rule.V1, rule.V2, rule.V3, rule.V4)
	}
	return
}

func (sr *sysRbac) GetRbacByRoleID(roleID int) []int {
	var (
		casbinRule []gormadapter.CasbinRule
		apiIDs     []int
	)
	global.GORM.Model(&casbinRule).WithContext(sr.ctx).Where("v0 = ?", roleID).Find(&casbinRule)
	if len(casbinRule) < 1 {
		return nil
	}
	for _, rbac := range casbinRule {
		id, _ := strconv.Atoi(rbac.V5)
		apiIDs = append(apiIDs, id)
	}
	return apiIDs
}

func (sr *sysRbac) DeleteByAPIsID(apiID int) error {
	var casbinRule gormadapter.CasbinRule
	if err := global.GORM.Model(&casbinRule).WithContext(sr.ctx).Where("v5 = ?", apiID).First(&casbinRule).Error; err != nil {
		return err
	}
	_, err := global.CasbinCacheEnforcer.RemovePolicy(casbinRule.V0, casbinRule.V1, casbinRule.V2, casbinRule.V3, casbinRule.V4, casbinRule.V5)
	freshRBAC()
	return err
}

func freshRBAC() {
	_ = global.CasbinCacheEnforcer.LoadPolicy()
}
