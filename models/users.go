/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primarykey" json:"id"`
	UID      string `gorm:"column:uid;comment:'用戶uid'" json:"uid"`
	UserName string `gorm:"column:username;comment:'用户名';size:128;uniqueIndex:uk_username;not null;unique" json:"username"  form:"username"`
	Password string `gorm:"column:password;comment:'用户密码';size:128" json:"password"  form:"password"`
	Phone    string `gorm:"column:phone;comment:'手机号码';size:11" json:"phone"`
	Email    string `gorm:"column:email;comment:'邮箱';size:128" json:"email"`
	NickName string `gorm:"column:nick_name;comment:'用户昵称';size:128" json:"nick_name"`
	Avatar   string `gorm:"column:avatar;default:https://img1.baidu.com/it/u=2206814125,3628191178&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500;comment:'用户头像';size:128" json:"avatar"`
	Status   uint   `gorm:"type:tinyint(1);default:1;comment:'用户状态(正常/禁用, 默认正常)'" json:"status"`
	RoleId   uint   `gorm:"column:role_id;default:1;comment:'角色id外键'" json:"role_id,omitempty"`
	Role     Role   `gorm:"foreignkey:RoleId" json:"role"`
	DeptId   uint64 `gorm:"default:1;comment:'部门id外键'" json:"dept_id"`
	Dept     Dept   `gorm:"foreignkey:DeptId" json:"dept"`
	CreateBy string `gorm:"column:create_by;comment:'创建来源'" json:"create_by"`
}

func (*User) TableName() string {
	return "user"
}

// 用户列表

type UserList struct {
	Items []User `json:"items"`
	Total int64  `json:"total"`
}
