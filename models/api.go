/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/5
*/

package models

import "gorm.io/gorm"

// 路由

type APIPath struct {
	gorm.Model
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
	Desc   string `json:"desc"  binding:"required"`
	MenuId uint64 `gorm:"default:1;comment:'菜单外键'" json:"menu_id"`
	Menu   Menu   `gorm:"foreignkey:MenuId" json:"menu"`
}

func (*APIPath) TableName() string {
	return "api_path"
}
