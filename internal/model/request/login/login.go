/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/7
*/

package login

type ReqLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"  aes:"true" binding:"required"`
}
