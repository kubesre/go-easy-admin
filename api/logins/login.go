/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/7
*/

package logins

import (
	"github.com/gin-gonic/gin"

	"go-easy-admin/pkg/controller/login"
	"go-easy-admin/pkg/global"
)

type LoginInterface interface {
	//LdapLogin(ctx *gin.Context)
	//GeneralLogin(ctx *gin.Context)
	GetLoginUserResource(ctx *gin.Context)
}

type sysLogin struct{}

func NewSysLogin() LoginInterface {
	return &sysLogin{}
}

//func (sl *sysLogin) LdapLogin(ctx *gin.Context) {
//	body := new(reqLogin.ReqLogin)
//	if err := ctx.ShouldBindJSON(&body); err != nil {
//		global.ReturnContext(ctx).Failed("参数错误", nil)
//		return
//	}
//	if err, data := login.NewSysLogin(ctx).LdapLogin(body); err != nil {
//		global.ReturnContext(ctx).Failed("登录失败", nil)
//		return
//	} else {
//		global.ReturnContext(ctx).Successful("登录成功", data)
//	}
//}
//
//func (sl *sysLogin) GeneralLogin(ctx *gin.Context) {
//	body := new(reqLogin.ReqLogin)
//	if err := ctx.ShouldBindJSON(&body); err != nil {
//		global.ReturnContext(ctx).Failed("参数错误", nil)
//		return
//	}
//	utils.TagAes(body)
//	if err, data := login.NewSysLogin(ctx).GeneralLogin(body); err != nil {
//		global.ReturnContext(ctx).Failed("登录失败", nil)
//		return
//	} else {
//		global.ReturnContext(ctx).Successful("登录成功", data)
//	}
//}

func (sl *sysLogin) GetLoginUserResource(ctx *gin.Context) {
	var id uint
	if v, exists := ctx.Get("id"); exists {
		if uv, ok := v.(uint); ok {
			id = uv
		}
	} else {
		global.ReturnContext(ctx).Failed("用户未登录", nil)
		return
	}

	if err, data := login.GetLoginUserResource(int(id), ctx); err != nil {
		global.ReturnContext(ctx).Failed("获取失败", nil)
		return
	} else {
		global.ReturnContext(ctx).Successful("获取成功", data)
	}

}
