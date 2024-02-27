package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLoggerConf struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	LogLevel   string
	AppMode    string
}

func NewZapLoggerConf() ZapLoggerConf {
	return ZapLoggerConf{
		Filename:   "sra.log",
		MaxSize:    200,
		MaxAge:     30,
		MaxBackups: 7,
		LogLevel:   "debug",
	}
}

func InitLogger(zcf ZapLoggerConf) error {
	writerSyncer := getLogWriter(zcf)
	encoder := getEncode()
	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(zcf.LogLevel)); err != nil {
		return err
	}
	var core zapcore.Core
	if zcf.AppMode == "debug" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writerSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writerSyncer, l)
	}
	lg := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
	return nil
}

func getEncode() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

func getLogWriter(zcf ZapLoggerConf) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   zcf.Filename,
		MaxSize:    zcf.MaxSize,
		MaxBackups: zcf.MaxBackups,
		MaxAge:     zcf.MaxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
