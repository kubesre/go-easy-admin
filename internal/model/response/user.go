/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/5
*/

package response

import (
	"go-easy-admin/internal/model"
	"go-easy-admin/internal/model/system"
)

type User struct {
	model.BaseModel
	Username string `json:"userName" gorm:"index;comment:用户登录名"`       // 用户登录名
	Password string `json:"-"  gorm:"comment:用户登录密码"`                  // 用户登录密码
	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
	Avatar   string `gorm:"column:avatar;default:https://img1.baidu.com/it/u=2206814125,3628191178&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500;comment:'用户头像';size:128" json:"avatar"`
	Email    string `gorm:"column:email;comment:'邮箱';size:128" json:"email"`
	Phone    string `gorm:"column:phone;comment:'手机号码';size:11" json:"phone"`
	Status   uint   `gorm:"type:tinyint(1);default:1;comment:'用户状态(正常/禁用, 默认正常)'" json:"status"`
	Roles    []Role `gorm:"many2many:sys_user_role;" json:"roles" copier:"-"`
}

type Role struct {
	model.BaseModel
	Name   string `gorm:"column:name;comment:'角色名称';size:128" json:"name"`
	Desc   string `gorm:"column:desc;comment:'角色描述';size:128" json:"desc"`
	Status uint   `gorm:"type:tinyint(1);default:1;comment:'用户状态(正常/禁用, 默认正常)'" json:"status"`
	Users  []User `gorm:"many2many:sys_user_role;" json:"users"`
	Menus  []Menu `gorm:"many2many:sys_role_menu;" json:"menus"`
}
type Menu struct {
	model.BaseModel
	Name      string `gorm:"comment:'菜单名称';size:64" json:"name"`
	NameCode  string `gorm:"column:name_code;comment:'前端路径name';size:64" json:"name_code"`
	IsShow    uint   `gorm:"column:is_show;type:tinyint(1);default:2;comment:'状态(1隐藏/2显示, 默认正常)'" json:"is_show"`
	Icon      string `gorm:"comment:'菜单图标';size:64" json:"icon"`
	Path      string `gorm:"comment:'菜单访问路径';size:64" json:"path"`
	Sort      int    `gorm:"default:0;type:int(3);comment:'菜单顺序(同级菜单, 从0开始, 越小显示越靠前)'" json:"sort"`
	ParentId  uint   `gorm:"default:0;comment:'父菜单编号(编号为0时表示根菜单)'" json:"parent_id"`
	Component string `gorm:"comment:'前端路径';size:64" json:"component"`
	Children  []Menu `gorm:"-" json:"children"  copier:"-"`
	Roles     []Role `gorm:"many2many:sys_role_menu;" json:"roles"  copier:"-"`
	APIs      []APIs `gorm:"many2many:sys_menu_api;" json:"apis"  copier:"-"`
}

type APIs struct {
	model.BaseModel
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
	Desc   string `json:"desc"  binding:"required"`
	MenuId uint64 `gorm:"default:1;comment:'菜单外键'" json:"menu_id"`
	Menu   Menu   `gorm:"foreignkey:MenuId" json:"menu"`
	Menus  []Menu `gorm:"many2many:sys_menu_api;" json:"menus"`
}

type LoginUser struct {
	model.BaseModel
	Username string        `json:"userName" gorm:"index;comment:用户登录名"`       // 用户登录名
	Password string        `json:"-"  gorm:"comment:用户登录密码"`                  // 用户登录密码
	NickName string        `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
	Avatar   string        `gorm:"column:avatar;default:https://img1.baidu.com/it/u=2206814125,3628191178&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500;comment:'用户头像';size:128" json:"avatar"`
	Email    string        `gorm:"column:email;comment:'邮箱';size:128" json:"email"`
	Phone    string        `gorm:"column:phone;comment:'手机号码';size:11" json:"phone"`
	Status   uint          `gorm:"type:tinyint(1);default:1;comment:'用户状态(正常/禁用, 默认正常)'" json:"status"`
	Roles    []uint        `gorm:"many2many:sys_user_role;" json:"roles" copier:"-"`
	Menus    []system.Menu `gorm:"-" json:"menus" copier:"-"`
}
