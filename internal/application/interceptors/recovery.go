package interceptors

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryRecoveryInterceptor(l *zap.Logger) grpc.UnaryServerInterceptor {
	logger := l.Named("grpc-panic")
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p any) (err error) {
			logger.Error("panic recovered", zap.Any("panic", p), zap.Stack("stack"))
			return status.Errorf(codes.Internal, "internal service error")
		}),
	}
	return recovery.UnaryServerInterceptor(recoveryOpts...)
}
