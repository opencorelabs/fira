package auth

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/opencorelabs/fira/internal/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

var (
	ErrInvalidToken       = errors.New("invalid token")
	StandardRejectionCode = status.Error(codes.Unauthenticated, "invalid authorization token")
	TodoJWTManager        = &JWTManager{
		secret:   []byte("dev_secret"),
		duration: time.Hour * 24,
	}
	PublicRoutes = map[string]struct{}{
		"/protos.fira.v1.FiraService/LoginAccount":  {},
		"/protos.fira.v1.FiraService/CreateAccount": {},
		"/protos.fira.v1.FiraService/VerifyAccount": {},
	}
)

type AccountClaims struct {
	jwt.RegisteredClaims
	AccountID string `json:"account_id"`
}

type JWTManager struct {
	secret   []byte
	duration time.Duration
}

func (m *JWTManager) Generate(accountID string) (string, error) {
	claims := AccountClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "fira",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		AccountID: accountID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secret)
}

func (m *JWTManager) Verify(tokenStr string) (*AccountClaims, error) {
	claims := &AccountClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return m.secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func JWTInterceptor(log logging.Provider, accounts AccountStoreProvider, manager *JWTManager) grpc.UnaryServerInterceptor {
	logger := log.Logger().Named("jwt-interceptor").Sugar()

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		logger.Debug("checking authorization for method", info.FullMethod)
		_, ok := PublicRoutes[info.FullMethod]
		if ok {
			logger.Debugf("public route %s allowed", info.FullMethod)
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
		claims, verifyErr := manager.Verify(token)
		if verifyErr != nil {
			logger.Debugw("token verification failed", "error", verifyErr)
			return nil, StandardRejectionCode
		}

		account, accountErr := accounts.AccountStore().FindAccountByID(ctx, claims.AccountID)
		if accountErr != nil {
			logger.Debugw("account not found", "account_id", claims.AccountID)
			return nil, StandardRejectionCode
		}

		if !account.Valid {
			logger.Debugw("account is not valid", "account_id", claims.AccountID)
			return nil, StandardRejectionCode
		}

		return handler(ctx, req)
	}
}
