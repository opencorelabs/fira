package auth

import (
	"context"
	"errors"
	"strings"

	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
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
		v1.FiraService_LoginAccount_FullMethodName:          {},
		v1.FiraService_CreateAccount_FullMethodName:         {},
		v1.FiraService_VerifyAccount_FullMethodName:         {},
		v1.FiraService_BeginPasswordReset_FullMethodName:    {},
		v1.FiraService_CompletePasswordReset_FullMethodName: {},
		v1.FiraService_GetApiInfo_FullMethodName:            {},
	}
	// AccountRoutes are available only with Account authentication
	AccountRoutes = map[string]struct{}{
		v1.FiraService_GetAccount_FullMethodName:         {},
		v1.FiraService_CreateApp_FullMethodName:          {},
		v1.FiraService_GetApp_FullMethodName:             {},
		v1.FiraService_ListApps_FullMethodName:           {},
		v1.FiraService_RotateAppToken_FullMethodName:     {},
		v1.FiraService_InvalidateAppToken_FullMethodName: {},
	}
	// AppRoutes are available only with developer.App authentication
	AppRoutes = map[string]struct{}{
		v1.FiraService_CreateLinkSession_FullMethodName: {},
		v1.FiraService_GetLinkSession_FullMethodName:    {},
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
