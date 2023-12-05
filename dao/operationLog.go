package dao

import (
	"go-easy-admin/common/global"
	"go-easy-admin/models"
	"time"
)

type operationLogService struct{}

var OperationLogService operationLogService

//处理OperationLogChan将日志记录到数据库

func (s *operationLogService) SaveOperationLogChannel(olc <-chan *models.OperationLog) {
	Logs := make([]models.OperationLog, 0)
	//5s 自动同步一次
	duration := 5 * time.Second
	timer := time.NewTimer(duration)
	defer timer.Stop()
	for {
		select {
		case log := <-olc:
			Logs = append(Logs, *log)
			// 每5条记录到数据库
			if len(Logs) > 5 {
				global.GORM.Create(&Logs)
				Logs = make([]models.OperationLog, 0)
				timer.Reset(duration) // 入库重置定时器
			}
		case <-timer.C: //5s 自动同步一次
			if len(Logs) > 0 {
				global.GORM.Create(&Logs)
				Logs = make([]models.OperationLog, 0)
			}
			timer.Reset(duration) // 入库重置定时器
		}
	}
}
