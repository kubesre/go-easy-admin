/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package global

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

type MyFormatter struct{}

var TPLogger *logrus.Logger

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	level := entry.Level
	var levelColor int
	switch level {
	case logrus.TraceLevel, logrus.DebugLevel:
		levelColor = gray // Cyan
	case logrus.InfoLevel:
		levelColor = blue // Green
	case logrus.WarnLevel:
		levelColor = yellow // Yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red // Red
	default:
		levelColor = blue // Reset color
	}

	funcVal := entry.Caller.Function
	fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
	msg := entry.Message
	time := entry.Time.Format("2006-01-02 15:04:05")
	return []byte(fmt.Sprintf("%s [ \033[%dm%s\033[0m ] [ %s ] [%s] %s\n", time, levelColor, level.String(), funcVal, fileVal, msg)), nil
}

func InitLog() {
	// 创建 Logrus 日志实例
	TPLogger = logrus.New()
	TPLogger.SetReportCaller(true)
	// 输出到标准输出
	TPLogger.SetOutput(os.Stdout)
	// 使用自定义日志格式
	TPLogger.SetFormatter(&MyFormatter{})
}
