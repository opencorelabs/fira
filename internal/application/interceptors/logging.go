package interceptors

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// TODO: this logger kind of sucks, there is probably a better one out there...

func InterceptorLogger(l *zap.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, len(fields)/2)
		for i := 0; i < len(fields); i += 2 {
			i := logging.Fields(fields).Iterator()
			if i.Next() {
				k, v := i.At()
				f = append(f, zap.Any(k, v))
			}
		}
		logger := l.WithOptions(zap.AddCallerSkip(1)).With(f...)

		switch lvl {
		case logging.LevelDebug:
			logger.Debug(msg)
		case logging.LevelInfo:
			logger.Info(msg)
		case logging.LevelWarn:
			logger.Warn(msg)
		case logging.LevelError:
			logger.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func UnaryLoggingInterceptor(l *zap.Logger) grpc.UnaryServerInterceptor {
	logger := l.Named("grpc")
	return logging.UnaryServerInterceptor(InterceptorLogger(logger))
}
