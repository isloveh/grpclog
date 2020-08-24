package grpcLogs

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"testing"
	"time"
)

func TestLogs_Info(t *testing.T) {
	l := Logs{
		Path:             "./logs",
		WithMaxAge:       time.Duration(7) * time.Hour * 24,
		WithRotationTime: time.Duration(1) * time.Hour * 24,
		IsWrite:          true,
		Level:            logrus.DebugLevel,
	}
	l.Info("122334")
}

func TestLogs_Debug(t *testing.T) {
	l := Logs{
		Path:             "./logs",
		WithMaxAge:       time.Duration(7) * time.Hour * 24,
		WithRotationTime: time.Duration(1) * time.Minute * 1,
		IsWrite:          true,
		Level:            logrus.DebugLevel,
	}
	l.Debug("122334")
}

func TestLogs_Error(t *testing.T) {
	l := Logs{
		Path:             "./logs",
		WithMaxAge:       time.Duration(7) * time.Hour * 24,
		WithRotationTime: time.Duration(1) * time.Hour * 24,
		IsWrite:          true,
		Level:            logrus.DebugLevel,
	}
	l.Error(codes.Internal, xerrors.New("测试测试"), codes.Internal.String())
}

func TestLogs_Warn(t *testing.T) {
	l := Logs{
		Path:             "./logs",
		WithMaxAge:       time.Duration(7) * time.Hour * 24,
		WithRotationTime: time.Duration(1) * time.Hour * 24,
		IsWrite:          true,
		Level:            logrus.DebugLevel,
	}
	l.Warn("警告日志")
}

func TestLogs_Fatal(t *testing.T) {
	l := Logs{
		Path:             "./logs",
		WithMaxAge:       time.Duration(7) * time.Hour * 24,
		WithRotationTime: time.Duration(1) * time.Hour * 24,
		IsWrite:          true,
		Level:            logrus.DebugLevel,
	}
	l.Fatal(xerrors.New("测试测试"))
}

func TestLogs_Panic(t *testing.T) {
	l := Logs{
		Path:             "./logs",
		WithMaxAge:       time.Duration(7) * time.Hour * 24,
		WithRotationTime: time.Duration(1) * time.Hour * 24,
		IsWrite:          true,
		Level:            logrus.DebugLevel,
	}
	l.Panic(xerrors.New("测试"))
}
