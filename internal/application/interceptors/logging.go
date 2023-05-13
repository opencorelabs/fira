package interceptors

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"time"
)

func UnaryLoggingInterceptor(baseLogger *zap.Logger) grpc.UnaryServerInterceptor {
	loggerOpts := []zap.Option{
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.FatalLevel),
	}
	logger := baseLogger.Named("grpc").WithOptions(loggerOpts...)

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log := logger.With(
			zap.String("method", info.FullMethod),
		)
		level := zap.InfoLevel
		start := time.Now()
		resp, err = handler(ctx, req)
		end := time.Now()
		if err != nil {
			level = zap.ErrorLevel
			log = log.With(zap.Error(err))
		}
		respStatus := status.Convert(err)
		log = log.With(
			zap.String("status", respStatus.Code().String()),
			zap.Duration("duration", end.Sub(start)),
		)
		log.Log(level, "grpc.request")
		return
	}
}
