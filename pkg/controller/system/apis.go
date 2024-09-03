/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/8
*/

package system

import (
	"context"

	"github.com/jinzhu/copier"

	reqSystem "go-easy-admin/internal/model/request/system"
	"go-easy-admin/internal/model/system"
	"go-easy-admin/pkg/global"
)

type SysApis interface {
	Create(req *reqSystem.CreateAPIsReq) error
	Delete(id int) error
	Update(id int, req *reqSystem.UpdateAPIsReq) error
	List() (error, interface{})
	Get(id int) (error, *system.APIs)
	GetApiGroup() (error, []string)
}
type sysApis struct {
	tips string
	ctx  context.Context
}

func NewSysApis(ctx context.Context) SysApis {
	return &sysApis{ctx: ctx, tips: "路由"}
}

func (sa *sysApis) Create(req *reqSystem.CreateAPIsReq) error {
	api := new(system.APIs)
	if err := copier.Copy(&api, req); err != nil {
		return global.OtherErr(err, "转换路由数据失败")
	}
	api.CreateBy = sa.ctx.Value("username").(string)
	if err := global.GORM.WithContext(sa.ctx).Create(api).Error; err != nil {
		return global.CreateErr(sa.tips, err)
	}
	return nil
}

func (sa *sysApis) Delete(id int) error {
	err, api := sa.Get(id)
	if err != nil {
		return err
	}
	if err = global.GORM.WithContext(sa.ctx).Delete(&api).Error; err != nil {
		return global.DeleteErr(sa.tips, err)
	}
	return NewSysRBAC(sa.ctx).DeleteByAPIsID(id)
}

func (sa *sysApis) Update(id int, req *reqSystem.UpdateAPIsReq) error {
	api := new(system.APIs)
	if err := copier.Copy(&api, req); err != nil {
		return global.OtherErr(err, "转换路由数据失败")
	}
	api.CreateBy = sa.ctx.Value("username").(string)
	if err := global.GORM.WithContext(sa.ctx).Model(&system.APIs{}).Where("id = ?", id).Updates(api).Error; err != nil {
		return global.UpdateErr(sa.tips, err)
	}
	_, resApi := sa.Get(id)
	return NewSysRBAC(sa.ctx).UpdateByAPI(resApi)

}

func (sa *sysApis) List() (error, interface{}) {
	var resApis []system.APIs
	if err := global.GORM.WithContext(sa.ctx).Model(&system.APIs{}).
		Find(&resApis).Error; err != nil {
		return global.GetErr(sa.tips, err), nil
	}
	return nil, &resApis
}

func (sa *sysApis) Get(id int) (error, *system.APIs) {
	api := new(system.APIs)
	if err := global.GORM.WithContext(sa.ctx).Model(&system.APIs{}).Where("id = ?", id).First(&api).Error; err != nil {
		return global.GetErr(sa.tips, err), nil
	}
	return nil, api
}

func (sa *sysApis) GetApiGroup() (error, []string) {
	var apiGroups []string
	if err := global.GORM.WithContext(sa.ctx).Model(&system.APIs{}).Distinct().Pluck("api_group", &apiGroups).Error; err != nil {
		return global.GetErr(sa.tips, err), nil
	}
	return nil, apiGroups
}
