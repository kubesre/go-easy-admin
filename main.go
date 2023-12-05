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
