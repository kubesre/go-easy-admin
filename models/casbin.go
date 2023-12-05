/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/2
*/

package models

// 权限相关

//type CasbinModel struct {
//	PType    string `json:"p_type" gorm:"column:p_type" description:"策略类型"`
//	RoleId string `json:"role_id" gorm:"column:v0" description:"角色名称"`
//	Path     string `json:"path" gorm:"column:v1" description:"api路径"`
//	Method   string `json:"method" gorm:"column:v2" description:"访问方法"`
//	Desc     string `json:"desc" gorm:"column:v3" description:"说明"`
//}
//
//// 不要手动进行casbin初始化
//
//func (c *CasbinModel) TableName() string {
//	return "casbin_rule"
//}

type CasbinPolicy struct {
	PType  string `json:"p_type" binding:"required"`
	RoleID string `json:"role_id" binding:"required"`
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
	Desc   string `json:"desc"  binding:"required"`
}

type CasbinPolicyList struct {
	Items []CasbinPolicy `json:"items"`
	Total int            `json:"total"`
}
