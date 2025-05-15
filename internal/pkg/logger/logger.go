package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

type MetaData map[string]interface{}

type Logger interface {
	Info(ctx context.Context, msg string, meta MetaData)
	Error(ctx context.Context, msg string, err error, meta MetaData)
	Sync()
}

type logger struct {
	logger *zap.Logger
}

func NewLogger() Logger {
	zap, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Error init logger: ", err)
		panic("Logger initialization error")
	}

	return &logger{
		logger: zap,
	}
}

func (l *logger) Info(ctx context.Context, msg string, metaData MetaData) {
	zapFields := l.getFields(ctx, metaData)
	l.logger.Info(msg, zapFields...)
}

func (l *logger) Error(ctx context.Context, msg string, err error, metaData MetaData) {
	var errStr string
	zapFields := l.getFields(ctx, metaData)
	if err != nil {
		errStr = err.Error()
	}

	errorMsg := msg + " - " + errStr
	l.logger.Error(errorMsg, zapFields...)
}

func (l *logger) getFields(ctx context.Context, metaData MetaData) []zap.Field {
	zapFields := []zap.Field{}

	if metaData == nil {
        metaData = make(MetaData)
    }

	if traceID, ok := ctx.Value("trace_id").(string); ok {
		zapFields = append(zapFields, zap.String("trace_id", traceID))
	}
	if url, ok := ctx.Value("url").(string); ok {
		zapFields = append(zapFields, zap.String("url", url))
	}
	if method, ok := ctx.Value("method").(string); ok {
		zapFields = append(zapFields, zap.String("method", method))
	}
	if remoteIP, ok := ctx.Value("remote_ip").(string); ok {
		zapFields = append(zapFields, zap.String("remote_ip", remoteIP))
	}

	for key, value := range metaData {
		zapFields = append(zapFields, zap.Any(key, value))
	}

	return zapFields
}

func (l *logger) Sync() {
	l.logger.Sync()
}