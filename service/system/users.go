/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package system

import (
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/spf13/viper"
	"go-easy-admin/common/global"
	dao "go-easy-admin/dao/system"
	mod "go-easy-admin/models/system"
	"strconv"
)

type InterfaceUsers interface {
	ExitUser(userName, password string) (*mod.User, error)
	LdapLogin(userName, password string) (*mod.User, error)
	UserInfo(id interface{}) (*mod.User, error)
	UserList(username string, limit, page int) (*mod.UserList, error)
	UserUpdate(userData *mod.User) error
	UserAdd(user *mod.User) error
}
type userInfo struct{}

func NewUserInfo() InterfaceUsers {
	return &userInfo{}
}

func (u *userInfo) ExitUser(userName, password string) (*mod.User, error) {
	data, err := dao.NewUserInterface().ExitUser(userName, password)
	if err != nil {
		global.TPLogger.Error("用户认证失败：", err)
		return nil, errors.New("用户认证失败")
	}
	return data, nil
}

// 用户详情

func (u *userInfo) UserInfo(id interface{}) (*mod.User, error) {
	idUint := fmt.Sprintf("%d", id)
	idInt, err := strconv.Atoi(idUint)
	if err != nil {
		global.TPLogger.Error("用户详情查看失败：", err)
		return nil, errors.New("用户详情查看失败")
	}
	data, err := dao.NewUserInterface().UserInfo(uint(idInt))
	if err != nil {
		global.TPLogger.Error("用户详情查看失败：", err)
		return nil, errors.New("用户详情查看失败")
	}
	return data, nil
}

// 用户列表

func (u *userInfo) UserList(username string, limit, page int) (*mod.UserList, error) {
	data, err := dao.NewUserInterface().UserList(username, limit, page)
	if err != nil {
		global.TPLogger.Error("UserList失败：", err)
		return nil, errors.New("UserList失败")
	}
	return data, nil
}

// 用户更新

func (u *userInfo) UserUpdate(userData *mod.User) error {
	err := dao.NewUserInterface().UserUpdate(userData)
	if err != nil {
		global.TPLogger.Error("用户更新失败：", err)
		return errors.New("用户更新失败")
	}
	return nil
}

// 用户添加

func (u *userInfo) UserAdd(user *mod.User) error {
	err := dao.NewUserInterface().UserAdd(user)
	if err != nil {
		global.TPLogger.Error("用户添加失败：", err)
		return errors.New("用户添加失败")
	}
	return nil
}

// ldap用户登录

func (u *userInfo) LdapLogin(userName, password string) (*mod.User, error) {
	// 连接ldap
	l, err := ldap.Dial("tcp", viper.GetString("ldap.address"))
	if err != nil {
		global.TPLogger.Error("连接ldap失败：", err)
		return nil, errors.New("连接ldap失败")
	}
	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: fmt.Sprintf("cn=%s,ou=user,%s", userName, viper.GetString("ldap.baseDN")),
		Password: password,
	})
	defer l.Close()
	if err != nil {
		global.TPLogger.Error("登录失败：", err)
		return nil, errors.New("登录失败")
	}
	var ldapUser *mod.User
	ldapUser, err = dao.NewUserInterface().ExitUser(userName, password)
	if err == nil && ldapUser != nil {
		global.TPLogger.Info("ldap用户已经存在数据库中，无需同步")
		return ldapUser, nil
	}
	_ = l.Bind(viper.GetString("ldap.adminUser"), viper.GetString("ldap.password"))
	searchRequest := ldap.NewSearchRequest(
		viper.GetString("ldap.baseDN"),
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(cn=%s)", userName),
		[]string{"uid", "cn", "mail", "sn", "telephoneNumber"},
		nil,
	)
	sr, err := l.Search(searchRequest)
	if err != nil || len(sr.Entries) != 1 {
		global.TPLogger.Error("用户不存在/用户名称重复：", err)
		return nil, errors.New("用户不存在/用户名称重复")
	}
	userEntry := sr.Entries[0]
	data := mod.User{
		UID:      userEntry.GetAttributeValue("uid"),
		UserName: userEntry.GetAttributeValue("cn"),
		Password: password,
		Phone:    userEntry.GetAttributeValue("telephoneNumber"),
		Email:    userEntry.GetAttributeValue("mail"),
		NickName: userEntry.GetAttributeValue("sn"),
		CreateBy: "ldap",
	}
	if err = dao.NewUserInterface().UserAdd(&data); err != nil {
		global.TPLogger.Error("用户写入数据库失败：", err)
		return nil, errors.New("登录失败")
	}
	// 执行这一步主要是为了获取用户ID 仅此而已
	ldapUser, err = dao.NewUserInterface().GetUserFromUserName(userName)
	if err != nil {
		global.TPLogger.Error("校验用户失败：", err)
		return nil, errors.New("登录失败")
	}
	return ldapUser, nil
}
