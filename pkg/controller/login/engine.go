/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/7
*/

package login

import (
	"context"

	reqLogin "go-easy-admin/internal/model/request/login"
)

type SysLogin interface {
	LdapLogin(request *reqLogin.ReqLogin) (error, interface{})
	GeneralLogin(request *reqLogin.ReqLogin) (error, interface{})
}
type sysLogin struct {
	ctx context.Context
}

func NewSysLogin(ctx context.Context) SysLogin {
	return &sysLogin{ctx: ctx}
}
