/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/2
*/

package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"go-easy-admin/internal/router"
	"go-easy-admin/pkg/global"
)

func Run() {
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", viper.GetString("server.address"),
			viper.GetInt("server.port")),
		Handler:        router.RegisterRouters(),
		MaxHeaderBytes: 1 << 20,
	}
	// 打印服务启动参数
	// 关闭服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	// 获取停止服务信号，kill  -9 获取不到
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 执行延迟停止
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}

}

func init() {
	global.InitConfig()
	global.InitSysTips()
	global.InitLog()
	global.InitMysql()
	global.InitCasbinEnforcer()
}
