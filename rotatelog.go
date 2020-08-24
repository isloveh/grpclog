package grpcLogs

import (
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func (l *Logs) rotateInit() {
	path := l.Path
	/*
		`WithLinkName` 为最新的日志建立软连接
		`WithRotationTime` 设置日志分割的时间，隔多久分割一次
		 *** WithMaxAge 和 WithRotationCount二者只能设置一个
		`WithMaxAge` 设置文件清理前的最长保存时间
		`WithRotationCount` 设置文件清理前最多保存的个数
	*/
	// 下面配置日志每隔一天轮转一个新文件，保留最近7天的日志文件，多余的自动清理掉。
	writer, _ := rotate.New(
		path+".%Y%m%d%H%M",                          // 文件名格式
		rotate.WithLinkName(path),                   // 文件存储路径
		rotate.WithMaxAge(l.WithMaxAge),             // time.Duration(7)*time.Hour*24 日志文件七天内有效
		rotate.WithRotationTime(l.WithRotationTime), // time.Duration(1)*time.Hour*24 日志文件按照一天一分割
	)
	logrus.SetOutput(writer)
}
