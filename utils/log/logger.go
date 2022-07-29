package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	writer := zapcore.AddSync(os.Stdout)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)
	log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func Log(data ...interface{}) {
	var lvl = zapcore.ErrorLevel
	if len(data) > 0 {
		if fmt.Sprintf("%T", data[0]) == "string" {
			lvl, _ = zapcore.ParseLevel(strings.ToLower(data[0].(string)))
			data = data[1:]
		}
	}

	switch lvl {
	case zapcore.DebugLevel:
		log.Debug(fmt.Sprint(data...))
	case zapcore.InfoLevel:
		log.Info(fmt.Sprint(data...))
	case zapcore.WarnLevel:
		log.Warn(fmt.Sprint(data...))
	case zapcore.ErrorLevel:
		log.Error(fmt.Sprint(data...))
	case zapcore.FatalLevel:
		log.Fatal(fmt.Sprint(data...))
	}
	return
}

func Info(args ...interface{}) {
	log.Info(fmt.Sprint(args...))
}

func Debug(args ...interface{}) {
	log.Debug(fmt.Sprint(args...))
}

func Error(args ...interface{}) {
	log.Error(fmt.Sprint(args...))
}

func Fatal(args ...interface{}) {
	log.Fatal(fmt.Sprint(args...))
}
