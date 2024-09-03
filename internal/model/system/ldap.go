/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/6
*/

package system

import (
	"encoding/json"
	"errors"

	"database/sql/driver"

	"go-easy-admin/internal/model"
)

// Ldap 用户登录ldap配置

type JSON []byte

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return string(j), nil
}

func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

func (j *JSON) UnmarshalToJSON(i interface{}) error {
	err := json.Unmarshal(*j, i)
	return err
}

type Ldap struct {
	model.BaseModel
	Address   string `gorm:"column:address;type:varchar(256);not null" json:"address"`
	DN        string `gorm:"column:dn" json:"dn"`
	AdminUser string `gorm:"column:admin_user;not null" json:"admin_user"`
	Password  string `gorm:"column:password" json:"password"`
	OU        string `gorm:"column:ou" json:"ou"`
	Filter    string `gorm:"column:filter;not null" json:"filter"`
	Mapping   JSON   `gorm:"column:mapping;type:json;comment:'属性映射'" json:"mapping"`
	SSL       uint   `gorm:"type:tinyint(1);default:2;comment:'状态(正常/禁用, 默认禁用)'" json:"ssl"`
	Status    uint   `gorm:"type:tinyint(1);default:2;comment:'状态(正常/禁用, 默认禁用)'" json:"status"`
}

func (*Ldap) TableName() string {
	return "sys_ldap"
}
