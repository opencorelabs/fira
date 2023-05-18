package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/opencorelabs/fira/internal/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidToken       = errors.New("invalid token")
	StandardRejectionCode = status.Error(codes.Unauthenticated, "invalid authorization token")
	// PublicRoutes are available without any authentication
	PublicRoutes = map[string]struct{}{
		"/protos.fira.v1.FiraService/LoginAccount":          {},
		"/protos.fira.v1.FiraService/CreateAccount":         {},
		"/protos.fira.v1.FiraService/VerifyAccount":         {},
		"/protos.fira.v1.FiraService/BeginPasswordReset":    {},
		"/protos.fira.v1.FiraService/CompletePasswordReset": {},
		"/protos.fira.v1.FiraService/GetApiInfo":            {},
	}
	// AccountRoutes are available only with Account authentication
	AccountRoutes = map[string]struct{}{
		"/protos.fira.v1.FiraService/GetAccount": {},
	}
	// AppRoutes are available only with developer.App authentication
	AppRoutes = map[string]struct{}{
		"/protos.fira.v1.FiraService/CreateLinkSession": {},
		"/protos.fira.v1.FiraService/GetLinkSession":    {},
	}
)

func JWTInterceptor(log logging.Provider, accountJWTManager JWTManager, appJWTManager JWTManager) grpc.UnaryServerInterceptor {
	logger := log.Logger().Named("jwt-interceptor").Sugar()

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		_, ok := PublicRoutes[info.FullMethod]
		if ok {
			return handler(ctx, req)
		}

		md, mdOk := metadata.FromIncomingContext(ctx)
		if !mdOk {
			return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}

		values := md.Get("authorization")
		if len(values) == 0 {
			logger.Debug("authorization token is not provided")
			return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
		}

		token := values[0]
		if strings.HasPrefix(token, "Bearer ") {
			token = strings.TrimPrefix(token, "Bearer ")
		}

		var jwtManager JWTManager

		_, isAccountScope := AccountRoutes[info.FullMethod]
		_, isAppScope := AppRoutes[info.FullMethod]

		if isAccountScope {
			jwtManager = accountJWTManager
		} else if isAppScope {
			jwtManager = appJWTManager
		}

		logger.Debugw("jwt manager running", "manager", jwtManager, "isaccountscope", isAccountScope, "isappscope", isAppScope)

		if jwtManager == nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization scope")
		}

		var verifyErr error
		ctx, verifyErr = jwtManager.Verify(ctx, token)
		if verifyErr != nil {
			return nil, StandardRejectionCode
		}

		return handler(ctx, req)
	}
}
