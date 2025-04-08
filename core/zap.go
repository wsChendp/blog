package core

import (
	"log"
	"server/global"

	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() *zap.Logger {  // 初始化日志记录器
	zapConfig := global.Config.Zap

	writeSyncer := getLogWriter(zapConfig.Filename, zapConfig.MaxSize, zapConfig.MaxBackups, zapConfig.MaxAge)
	
	if zapConfig.IsConsolePrint {
		writeSyncer = zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout))
	}

	encoder := getEncoder()

	var logLevel zapcore.Level
	err := logLevel.UnmarshalText([]byte(zapConfig.Level))
	if err != nil {
		log.Fatal("zap log level unmarshal failed", zap.Error(err))
	}

	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger
}

func getLogWriter(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	// 日志文件配置，处理日志滚动记录，切割归档
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.TimeKey = "time"
	return zapcore.NewJSONEncoder(encoderConfig)
}