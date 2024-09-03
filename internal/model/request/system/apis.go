/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/8
*/

package system

type CreateAPIsReq struct {
	Path     string `json:"path" binding:"required"`
	Method   string `json:"method" binding:"required"`
	Desc     string `json:"desc"  binding:"required"`
	ApiGroup string `json:"apiGroup" binding:"required"`
	//Menus    []int  `json:"menus"`
}

type UpdateAPIsReq struct {
	Path     string `json:"path"`
	Method   string `json:"method"`
	Desc     string `json:"desc"`
	ApiGroup string `json:"apiGroup"`
	//Menus    []int  `json:"menus"`
}
