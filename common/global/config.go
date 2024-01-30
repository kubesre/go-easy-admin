/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package global

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
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
	// 监听配置变化,无需重启应用读取配置
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println(in.Name, in.Op)
		readConfig()
	})
	readConfig()
}

func readConfig() {
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
