/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/2
*/

package system

// 权限相关

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
