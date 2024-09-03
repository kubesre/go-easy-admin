/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/8
*/

package global

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var CasbinCacheEnforcer *casbin.SyncedCachedEnforcer

func InitCasbinEnforcer() {
	e, err := MysqlCasbin()
	if err != nil {
		GeaLogger.Error("初始化casbin策略管理器失败：", err)
		panic(err)
	}
	e.EnableAutoSave(true)
	CasbinCacheEnforcer = e
}

// 定义casbin

func MysqlCasbin() (*casbin.SyncedCachedEnforcer, error) {
	a, err := gormadapter.NewAdapterByDB(GORM)
	if err != nil {
		GeaLogger.Error("casbin adapter gorm failed: ", err)
		return nil, err
	}
	//e, err := casbin.NewEnforcer("rbac_model.conf", a)
	e, err := casbin.NewSyncedCachedEnforcer("rbac_model.conf", a)
	if err != nil {
		return nil, err
	}
	e.SetExpireTime(60 * 60)
	if err = e.LoadPolicy(); err != nil {
		return nil, err
	}
	return e, nil
}
