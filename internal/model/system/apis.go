/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/2
*/

package system

import "go-easy-admin/internal/model"

type APIs struct {
	model.BaseModel
	Path     string `gorm:"not null;unique" json:"path" binding:"required"`
	Method   string `json:"method" binding:"required"`
	Desc     string `json:"desc"  binding:"required"`
	ApiGroup string `json:"apiGroup" binding:"required"`
	//Menus    []Menu `gorm:"many2many:sys_menu_api;" json:"menus"`
}

func (*APIs) TableName() string {
	return "system_apis"
}
