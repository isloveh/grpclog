package grpcLogs

import (
	"github.com/sirupsen/logrus"
	"os"
)

func (l *Logs) hookInit() {
	// 设置日志为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	logrus.SetOutput(os.Stdout)
	// 设置日志级别
	logrus.SetLevel(l.Level)
}
