/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/6
*/

package system

type CreateRoleReq struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Status uint   `json:"status"`
	Users  []int  `json:"users"`
	Menus  []int  `json:"menus"`
}
