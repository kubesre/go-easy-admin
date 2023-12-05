/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package models

import "gorm.io/gorm"

// 部门

type Dept struct {
	gorm.Model
	Name     string  `gorm:"comment:'部门名称';size:64" json:"name"`
	Sort     int     `gorm:"default:0;type:int(3);comment:'排序'" json:"sort"`
	ParentId uint    `gorm:"default:0;comment:'父级部门(编号为0时表示根)'" json:"parent_id"`
	Children []*Dept `gorm:"-" json:"children"` // 下属部门集合
	Users    []User  `gorm:"foreignkey:DeptId"`
}

func (*Dept) TableName() string {
	return "dept"
}
