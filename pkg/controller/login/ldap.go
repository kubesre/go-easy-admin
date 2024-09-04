/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/7
*/

package login

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/jinzhu/copier"

	reqLogin "go-easy-admin/internal/model/request/login"
	reqSystem "go-easy-admin/internal/model/request/system"
	"go-easy-admin/internal/model/response"
	modeSystem "go-easy-admin/internal/model/system"
	"go-easy-admin/pkg/controller/system"
	"go-easy-admin/pkg/global"
)

func (sl *sysLogin) LdapLogin(request *reqLogin.ReqLogin) (error, interface{}) {
	var ld *ldap.Conn
	err, req := system.NewSysLdap(sl.ctx).Get()
	if err != nil {
		return err, nil
	}
	if req.SSL == 1 {
		ld, err = ldap.DialURL("ldaps://"+req.Address, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	} else {
		ld, err = ldap.DialURL("ldap://" + req.Address)
	}
	if err != nil {
		return err, nil
	}
	defer ld.Close()
	if ld != nil {
		if err = ld.Bind(req.AdminUser, req.Password); err != nil {
			return global.OtherErr(errors.New("连接失败" + err.Error())), nil
		}
	}
	searchRequest := ldap.NewSearchRequest(req.DN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(req.Filter, request.Username), []string{}, nil)
	sr, err := ld.Search(searchRequest)
	if err != nil {
		return err, nil
	}
	if len(sr.Entries) != 1 {
		return global.OtherErr(errors.New("user does not exist or too many entries returned")), nil
	}
	// ldap普通用户登录
	userDN := sr.Entries[0].DN
	if err = ld.Bind(userDN, request.Password); err != nil {
		return global.OtherErr(errors.New("登录失败"), err.Error()), nil
	}
	var userMap map[string]interface{}
	var createUser reqSystem.CreateUserReq
	if err = json.Unmarshal(req.Mapping, &userMap); err != nil {
		return global.OtherErr(errors.New("ldap映射关系错误"), err.Error()), nil
	}
	// 处理映射关系
	for k, v := range userMap {
		userMap[k] = sr.Entries[0].GetAttributeValue(v.(string))
	}
	jsonData, _ := json.Marshal(userMap)
	_ = json.Unmarshal(jsonData, &createUser)
	// 登录成功,将用户信息入库
	ok, userInfo := system.NewSysUser(sl.ctx).GetByUsername(request.Username)
	if ok {
		_ = system.NewSysUser(sl.ctx).Create(&createUser)
		ok, userInfo = system.NewSysUser(sl.ctx).GetByUsername(request.Username)
		if !ok && userInfo != nil {
			return nil, userInfo
		}
	}
	if userInfo != nil {
		return nil, userInfo
	}
	return errors.New("登录失败"), nil
}

func GetLoginUserResource(id int, ctx context.Context) (error, interface{}) {
	// 基础用户信息  LoginUser
	var user modeSystem.User
	if err := global.GORM.WithContext(ctx).Model(&modeSystem.User{}).
		Preload("Roles").
		Preload("Roles.Menus").
		Where("id = ?", id).
		First(&user).Error; err != nil {
		global.GeaLogger.Error("查询用户失败: ", err)
		return errors.New("查询用户失败"), nil
	}
	loginUser := new(response.LoginUser)

	if err := copier.Copy(&loginUser, user); err != nil {
		return global.OtherErr(errors.New("转换角色数据失败")), nil
	}
	for i := range user.Roles {
		loginUser.Roles = append(loginUser.Roles, user.Roles[i].ID)
	}
	// 菜单信息
	var menus []modeSystem.Menu
	for i := range user.Roles {
		menus = append(menus, user.Roles[i].Menus...)
	}
	afterRemoveMenu := removeMenu(menus)
	loginUser.Menus = buildMenuTree(afterRemoveMenu, 0)
	return nil, loginUser
}

func removeMenu(menus []modeSystem.Menu) (resMenus []modeSystem.Menu) {
	menuList := make(map[uint]modeSystem.Menu, 0)
	for _, item := range menus {
		if _, ok := menuList[item.ID]; !ok {
			menuList[item.ID] = item
		}
	}
	for _, item := range menuList {
		resMenus = append(resMenus, item)
	}
	return resMenus
}

func buildMenuTree(menus []modeSystem.Menu, parentID uint) []modeSystem.Menu {
	var tree []modeSystem.Menu
	for _, menu := range menus {
		if menu.ParentId == parentID {
			menu.Children = buildMenuTree(menus, menu.ID)
			tree = append(tree, menu)
		}
	}
	return tree
}
