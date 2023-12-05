/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package global

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// 全局CasbinEnforcer

var CasbinEnforcer *casbin.Enforcer

// 初始化casbin策略管理器

func InitCasbinEnforcer() {
	e, err := mysqlCasbin()
	if err != nil {
		TPLogger.Error("初始化casbin策略管理器失败：", err)
		panic(err)
	}
	e.EnableAutoSave(true)
	CasbinEnforcer = e
}

// 定义casbin

func mysqlCasbin() (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapterByDB(GORM)
	if err != nil {
		TPLogger.Error("casbin adapter gorm failed: ", err)
		return nil, err
	}
	e, err := casbin.NewEnforcer("config/rbac_model.conf", a)
	if err != nil {
		return nil, err
	}
	if err = e.LoadPolicy(); err != nil {
		return nil, err
	}
	return e, nil
}
