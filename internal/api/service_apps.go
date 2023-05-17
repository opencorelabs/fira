package api

import (
	"context"
	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
	"github.com/opencorelabs/fira/internal/developer"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppService struct {
	logger           *zap.Logger
	appJWTManager    developer.AppJWTManager
	appStoreProvider developer.AppStoreProvider
}

func (s *AppService) CreateApp(ctx context.Context, request *v1.CreateAppRequest) (*v1.CreateAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApp not implemented")
}

func (s *AppService) ListApps(ctx context.Context, request *v1.ListAppsRequest) (*v1.ListAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApps not implemented")
}

func (s *AppService) GetApp(ctx context.Context, request *v1.GetAppRequest) (*v1.GetAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApp not implemented")
}

func (s *AppService) RotateAppToken(ctx context.Context, request *v1.RotateAppTokenRequest) (*v1.RotateAppTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateAppToken not implemented")
}

func (s *AppService) InvalidateAppToken(ctx context.Context, request *v1.InvalidateAppTokenRequest) (*v1.InvalidateAppTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateAppToken not implemented")
}
