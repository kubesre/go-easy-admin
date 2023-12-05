package global

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// 初始化配置文件

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("加载配置文件错误")
	}
	switch viper.GetString("server.model") {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	}
}
