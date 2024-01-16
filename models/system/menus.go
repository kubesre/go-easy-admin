/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package system

import (
	"gorm.io/gorm"
)

// 菜单

type Menu struct {
	gorm.Model
	Name      string `gorm:"comment:'菜单名称';size:64" json:"name"`
	NameCode  string `gorm:"column:name_code;comment:'前端路径name';size:64" json:"name_code"`
	IsShow    uint   `gorm:"column:is_show;type:tinyint(1);default:2;comment:'状态(1隐藏/2显示, 默认正常)'" json:"is_show"`
	Icon      string `gorm:"comment:'菜单图标';size:64" json:"icon"`
	Path      string `gorm:"comment:'菜单访问路径';size:64" json:"path"`
	Sort      int    `gorm:"default:0;type:int(3);comment:'菜单顺序(同级菜单, 从0开始, 越小显示越靠前)'" json:"sort"`
	ParentId  uint   `gorm:"default:0;comment:'父菜单编号(编号为0时表示根菜单)'" json:"parent_id"`
	Creator   string `gorm:"comment:'创建人';size:64" json:"creator"`
	Component string `gorm:"comment:'前端路径';size:64" json:"component"`
	Children  []Menu `gorm:"-" json:"children"`
	Roles     []Role `gorm:"many2many:relation_role_menu;" json:"roles"`
}

func (*Menu) TableName() string {
	return "menu"
}
