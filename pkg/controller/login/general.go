/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/7
*/

/* UserGeneralLogin*/

package login

import (
	"errors"

	reqLogin "go-easy-admin/internal/model/request/login"
	"go-easy-admin/pkg/controller/system"
	"go-easy-admin/pkg/global"
)

func (sl *sysLogin) GeneralLogin(request *reqLogin.ReqLogin) (error, interface{}) {
	ok, userInfo := system.NewSysUser(sl.ctx).GetByUsernameAndPwd(request.Username, request.Password)
	if ok && userInfo != nil {
		return nil, userInfo
	}
	return global.OtherErr(errors.New("登录失败")), nil
}
