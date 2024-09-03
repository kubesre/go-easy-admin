/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/2
*/

package model

import (
	"time"

	"gorm.io/gorm"
)

// gorm 基础字段

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"` // 删除时间
	CreateBy  string         `gorm:"column:create_by;comment:'创建来源'" json:"create_by"`
}
