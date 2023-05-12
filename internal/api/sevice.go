package api

import (
	"context"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
)

type FiraApiService struct {
	v1.UnimplementedFiraServiceServer
	authRegistry auth.Registry
	logger       *zap.SugaredLogger
	jwtManager   *auth.JWTManager
}

func New(log logging.Provider, authReg auth.Registry, manager *auth.JWTManager) v1.FiraServiceServer {
	return &FiraApiService{
		authRegistry: authReg,
		jwtManager:   manager,
		logger:       log.Logger().Named("api-service").Sugar(),
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

func (s *FiraApiService) CreateAccount(ctx context.Context, request *v1.CreateAccountRequest) (*v1.CreateAccountResponse, error) {
	var credType auth.CredentialsType
	credential := make(map[string]string)
	switch request.Credential.CredentialType {
	case v1.AccountCredentialType_ACCOUNT_CREDENTIAL_TYPE_EMAIL:
		credType = auth.CredentialsTypeEmailPassword
		credential["email"] = request.Credential.EmailCredential.Email
		credential["password"] = request.Credential.EmailCredential.Password
	case v1.AccountCredentialType_ACCOUNT_CREDENTIAL_TYPE_OAUTH_GITHUB:
		credType = auth.CredentialsTypeOAuth
		credential["provider"] = "github"
		credential["token"] = request.Credential.GithubCredential.ClientId
		credential["code"] = request.Credential.GithubCredential.Code
		credential["redirect_uri"] = request.Credential.GithubCredential.RedirectUri
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid credential type")
	}

	backend, getErr := s.authRegistry.GetBackend(credType)
	if getErr != nil {
		s.logger.Errorw("failed to get auth backend",
			"error", getErr,
			"credential_type", credType,
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	account, regErr := backend.Register(ctx, credential)

	if regErr != nil {
		s.logger.Errorw("failed to register account",
			"error", regErr,
			"credential_type", credType,
		)
		return nil, status.Errorf(codes.Unauthenticated, "registration error")
	}

	jwt, jwtErr := s.jwtManager.Generate(account.ID)
	if jwtErr != nil {
		s.logger.Errorw("failed to generate jwt",
			"error", jwtErr,
			"credential_type", credType,
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &v1.CreateAccountResponse{
		Status: v1.AccountRegistrationStatus_ACCOUNT_REGISTRATION_STATUS_OK,
		Jwt:    jwt,
	}, nil
}

func (s *FiraApiService) VerifyAccount(ctx context.Context, request *v1.VerifyAccountRequest) (*v1.VerifyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *FiraApiService) LoginAccount(ctx context.Context, request *v1.LoginAccountRequest) (*v1.LoginAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *FiraApiService) BeginPasswordReset(ctx context.Context, request *v1.BeginPasswordResetRequest) (*v1.BeginPasswordResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *FiraApiService) CompletePasswordReset(ctx context.Context, request *v1.CompletePasswordResetRequest) (*v1.CompletePasswordResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *FiraApiService) GetAccount(ctx context.Context, request *v1.GetAccountRequest) (*v1.GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *FiraApiService) CreateLinkSession(ctx context.Context, request *v1.CreateLinkSessionRequest) (*v1.CreateLinkSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *FiraApiService) GetLinkSession(ctx context.Context, request *v1.GetLinkSessionRequest) (*v1.GetLinkSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}
