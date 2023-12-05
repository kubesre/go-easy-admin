package app

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go-easy-admin/common/global"
	"go-easy-admin/dao"
	"go-easy-admin/middles"
	"go-easy-admin/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", viper.GetString("server.address"),
			viper.GetInt("server.port")),
		Handler:        routers.BaseRouters(),
		MaxHeaderBytes: 1 << 20,
	}
	// 打印服务启动参数
	global.TPLogger.Info("服务启动配置：", srv.Addr)
	// 开启日志记录
	goroutineNum := viper.GetInt("operation.GoroutineNum")
	if goroutineNum == 0 {
		goroutineNum = 3
	}
	for i := 0; i < goroutineNum; i++ {
		go dao.OperationLogService.SaveOperationLogChannel(middles.OperationLogChan)
	}
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
	global.TPLogger.Info("shutdown server...")
	// 执行延迟停止
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	global.TPLogger.Info("server exiting...")
}

func init() {
	// 初始化配置文件
	global.InitConfig()
	// 初始化日志
	global.InitLog()
	// 初始化数据库
	global.InitMysql()
	// 初始化表
	global.InitMysqlTables()
	// 初始化casbin
	global.InitCasbinEnforcer()
}
