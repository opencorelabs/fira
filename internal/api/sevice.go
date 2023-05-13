package api

import (
	"context"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/auth/verification"
	"github.com/opencorelabs/fira/internal/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
)

type FiraApiService struct {
	v1.UnimplementedFiraServiceServer
	authRegistry         auth.Registry
	logger               *zap.SugaredLogger
	jwtManager           *auth.JWTManager
	verificationProvider verification.Provider
}

func New(
	log logging.Provider,
	authReg auth.Registry,
	manager *auth.JWTManager,
	verificationProvider verification.Provider,
) v1.FiraServiceServer {
	return &FiraApiService{
		authRegistry:         authReg,
		jwtManager:           manager,
		verificationProvider: verificationProvider,
		logger:               log.Logger().Named("api-service").Sugar(),
	}
}

func (*FiraApiService) GetApiInfo(context.Context, *v1.GetApiInfoRequest) (*v1.GetApiInfoResponse, error) {
	return &v1.GetApiInfoResponse{
		Version: &v1.GetApiInfoResponse_Version{
			Major: 1,
			Minor: 0,
			Patch: 0,
		},
	}, nil
}

func (s *FiraApiService) CreateLinkSession(ctx context.Context, request *v1.CreateLinkSessionRequest) (*v1.CreateLinkSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *FiraApiService) GetLinkSession(ctx context.Context, request *v1.GetLinkSessionRequest) (*v1.GetLinkSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}
