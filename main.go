/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-easy-admin/app"
)

func main() {
	gin.SetMode(viper.GetString("server.model"))
	// 启动服务
	app.Run()
}
