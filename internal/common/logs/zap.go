package logs

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dcalsky/kong_service_demo/internal/common/logid"
)

var zapLogger *zap.Logger

func MustInit(level zapcore.LevelEnabler) {
	zapLogger = newZapLogger(level)
}

func newZapLogger(level zapcore.LevelEnabler) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.ConsoleSeparator = " "
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level),
	)
	options := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	}
	return zap.New(core, options...)
}

func Fatalf(ctx context.Context, template string, args ...any) {
	zapLogger.Fatal(logid.LogId(ctx) + " " + fmt.Sprintf(template, args...))
}

func Errorf(ctx context.Context, template string, args ...any) {
	zapLogger.Error(logid.LogId(ctx) + " " + fmt.Sprintf(template, args...))
}

func Warnf(ctx context.Context, template string, args ...any) {
	zapLogger.Warn(logid.LogId(ctx) + " " + fmt.Sprintf(template, args...))
}

func Infof(ctx context.Context, template string, args ...any) {
	zapLogger.Info(logid.LogId(ctx) + " " + fmt.Sprintf(template, args...))
}

func Debugf(ctx context.Context, template string, args ...any) {
	zapLogger.Debug(logid.LogId(ctx) + " " + fmt.Sprintf(template, args...))
}
