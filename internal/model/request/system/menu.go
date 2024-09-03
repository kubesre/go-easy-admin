/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/5
*/

package system

type CreateMenuReq struct {
	Name      string `json:"name" binding:"required"`
	NameCode  string `json:"name_code"`
	IsShow    uint   `json:"is_show"`
	Icon      string `json:"icon"`
	Path      string `json:"path" binding:"required"`
	Sort      int    `json:"sort"`
	ParentId  uint   `json:"parent_id"`
	Component string `json:"component"`
	//APIs      []int  `json:"apis"`
}
