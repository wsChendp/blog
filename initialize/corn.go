package initialize

import (
	"os"
	"server/global"
	"server/task"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

// 瀹炵幇cron.Logger鎺ュ彛
func (l *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg, zap.Any("keysAndValues", keysAndValues))
}

func (l *ZapLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logger.Error(msg, zap.Error(err), zap.Any("keysAndValues", keysAndValues))
}

func NewZapLogger() *ZapLogger {
	return &ZapLogger{logger: global.Log}
}

func InitCron() *cron.Cron {
	c := cron.New(cron.WithLogger(NewZapLogger()))

	err := task.RegisterScheduledTasks(c)
	if err != nil {
		global.Log.Error("瀹氭椂浠诲姟娉ㄥ唽澶辫触", zap.Error(err))
		os.Exit(1)
	}
	return c
}
