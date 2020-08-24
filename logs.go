package grpcLogs

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
)

type Logs struct {
	Path             string        `json:"path"`             // 日志存储路径
	WithMaxAge       time.Duration `json:"withMaxAge"`       // 日志存储最长时间
	WithRotationTime time.Duration `json:"withRotationTime"` // 日志分割时间
	IsWrite          bool          `json:"isWrite"`          // 是否写文件 false 不写入　true 写入
	Level            logrus.Level  `json:"level"`            // 日志写入等级
}

func (l *Logs) init() {
	l.hookInit()
	if l.IsWrite {
		l.rotateInit()
	}
}

func (l *Logs) Info(message string) {
	l.init()
	logrus.Info(message)
}

func (l *Logs) Warn(message string) {
	l.init()
	logrus.Warn(message)
}

func (l *Logs) Debug(message string) {
	l.init()
	logrus.Debug(message)
}

func (l *Logs) Error(code int, err error, errorType string) {
	l.init()
	funcName, path, line := l.runtimeTrace(err)
	logrus.WithFields(logrus.Fields{
		"line":      line,
		"func":      funcName,
		"path":      path,
		"code":      code,
		"errorType": errorType,
	}).Error(err)
}

func (l *Logs) Fatal(err error) {
	l.init()
	funcName, path, line := l.runtimeTrace(err)
	logrus.WithFields(logrus.Fields{
		"line": line,
		"func": funcName,
		"path": path,
	}).Fatal(err)
}

func (l *Logs) Panic(err error) {
	l.init()
	funcName, path, line := l.runtimeTrace(err)
	logrus.WithFields(logrus.Fields{
		"line": line,
		"func": funcName,
		"path": path,
	}).Panic(err)
}

func (l *Logs) runtimeTrace(err error) (string, string, int) {
	var line int
	var funcName string
	var path string
	frame := tracerr.Wrap(err).StackTrace()
	if frame != nil {
		if len(frame) > 0 {
			line = frame[0].Line
			path = frame[0].Path
			funcName = frame[0].Func
		}
	}
	return funcName, path, line
}
