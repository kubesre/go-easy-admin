/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/2
*/

// 路由

package router

import (
	"github.com/gin-gonic/gin"

	routerLogin "go-easy-admin/internal/router/v1/logins"
	routerSys "go-easy-admin/internal/router/v1/system"
	"go-easy-admin/pkg/middles"
)

func RegisterRouters() *gin.Engine {
	r := gin.New()
	r.Use(middles.LogHandlerFunc(), gin.Recovery(), middles.Cors())
	authMiddleware, err := middles.InitAuth()
	if err != nil {
		panic(err)
	}
	// 健康检查
	r.GET("/health", func(ctx *gin.Context) {
		return
	})
	PrivateGroup := r.Group("")
	PrivateGroup.Use(authMiddleware.MiddlewareFunc(), middles.RbacMiddle())
	//PrivateGroup.Use(authMiddleware.MiddlewareFunc())
	{
		UserGroup := PrivateGroup.Group("/sys/user")
		{
			routerSys.User(UserGroup, authMiddleware)
		}
		MenuGroup := PrivateGroup.Group("/sys/menu")
		{
			routerSys.Menu(MenuGroup)
		}
		RoleGroup := PrivateGroup.Group("/sys/role")
		{
			routerSys.Role(RoleGroup)
		}
		ApisGroup := PrivateGroup.Group("/sys/apis")
		{
			routerSys.Apis(ApisGroup)
		}
		LdapGroup := PrivateGroup.Group("/sys/ldap")
		{
			routerSys.Ldap(LdapGroup)
		}
		RbacGroup := PrivateGroup.Group("/sys/rbac")
		{
			routerSys.RBAC(RbacGroup)
		}
		LoginUserResource := PrivateGroup.Group("/sys/login")
		{
			routerLogin.Resource(LoginUserResource)
		}
	}

	PublicGroup := r.Group("")
	{
		LoginGroup := PublicGroup.Group("/sys/login")
		{
			routerLogin.Login(LoginGroup, authMiddleware)
		}
	}

	r.NoRoute(func(ctx *gin.Context) {
		return
	})
	return r
}
