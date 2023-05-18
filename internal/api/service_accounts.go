package api

import (
	"context"
	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/auth/verification"
	"github.com/opencorelabs/fira/internal/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountService struct {
	authRegistry         auth.Registry
	logger               *zap.SugaredLogger
	jwtManager           auth.JWTManager
	verificationProvider verification.Provider
}

func NewAccountService(
	log logging.Provider,
	authReg auth.Registry,
	manager auth.JWTManager,
	verificationProvider verification.Provider,
) *AccountService {
	return &AccountService{
		authRegistry:         authReg,
		jwtManager:           manager,
		verificationProvider: verificationProvider,
		logger:               log.Logger().Named("api-service").Sugar(),
	}
}

func (s *AccountService) decodeLoginCredential(c *v1.LoginCredential) (map[string]string, auth.Backend, error) {
	var credType auth.CredentialsType
	credential := make(map[string]string)
	switch c.CredentialType {
	case v1.AccountCredentialType_ACCOUNT_CREDENTIAL_TYPE_EMAIL:
		credType = auth.CredentialsTypeEmailPassword
		credential["email"] = c.EmailCredential.Email
		credential["password"] = c.EmailCredential.Password
	case v1.AccountCredentialType_ACCOUNT_CREDENTIAL_TYPE_OAUTH_GITHUB:
		credType = auth.CredentialsTypeOAuth
		credential["provider"] = "github"
		credential["token"] = c.GithubCredential.ClientId
		credential["code"] = c.GithubCredential.Code
		credential["redirect_uri"] = c.GithubCredential.RedirectUri
	default:
		return nil, nil, status.Errorf(codes.InvalidArgument, "invalid credential type")
	}

	backend, getErr := s.authRegistry.GetBackend(credType)
	if getErr != nil {
		s.logger.Errorw("failed to get auth backend",
			"error", getErr,
			"credential_type", credType,
		)
		return nil, nil, status.Errorf(codes.Internal, "internal error")
	}

	return credential, backend, nil
}

func getAccountStatus(acct *auth.Account) v1.AccountRegistrationStatus {
	if acct.Valid {
		return v1.AccountRegistrationStatus_ACCOUNT_REGISTRATION_STATUS_OK
	}
	switch acct.CredentialsType {
	case auth.CredentialsTypeEmailPassword:
		return v1.AccountRegistrationStatus_ACCOUNT_REGISTRATION_STATUS_VERIFY_EMAIL
	}
	return v1.AccountRegistrationStatus_ACCOUNT_REGISTRATION_STATUS_UNSPECIFIED
}

func getAccountNamespace(ns v1.AccountNamespace) auth.AccountNamespace {
	switch ns {
	case v1.AccountNamespace_ACCOUNT_NAMESPACE_DEVELOPER:
		return auth.AccountNamespaceDeveloper
	case v1.AccountNamespace_ACCOUNT_NAMESPACE_CONSUMER:
		return auth.AccountNamespaceConsumer
	default:
		return auth.AccountNamespaceNone
	}
}

func (s *AccountService) CreateAccount(ctx context.Context, request *v1.CreateAccountRequest) (*v1.CreateAccountResponse, error) {
	credential, backend, decodeErr := s.decodeLoginCredential(request.Credential)
	if decodeErr != nil {
		return nil, decodeErr
	}

	account, regErr := backend.Register(ctx, getAccountNamespace(request.Namespace), credential)

	if regErr != nil {
		s.logger.Errorw("failed to register account",
			"error", regErr,
			"credential_type", request.Credential.CredentialType,
		)
		return nil, status.Errorf(codes.Unauthenticated, "registration error")
	}

	jwt, jwtErr := s.jwtManager.Generate(ctx, account)
	if jwtErr != nil {
		s.logger.Errorw("failed to generate jwt", "error", jwtErr)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &v1.CreateAccountResponse{
		Status: getAccountStatus(account),
		Jwt:    jwt,
	}, nil
}

func (s *AccountService) VerifyAccount(ctx context.Context, request *v1.VerifyAccountRequest) (*v1.VerifyAccountResponse, error) {
	var verifier verification.Verifier
	switch request.Type {
	case v1.VerificationType_VERIFICATION_TYPE_EMAIL:
		verifier = s.verificationProvider.Email()
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid verification type")
	}

	acct, verifErr := verifier.VerifyWithToken(ctx, getAccountNamespace(request.Namespace), request.Token)
	if verifErr != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid verification token")
	}
	resp := &v1.VerifyAccountResponse{
		Status: getAccountStatus(acct),
	}
	if acct.Valid {
		jwt, jwtErr := s.jwtManager.Generate(ctx, acct)
		if jwtErr != nil {
			s.logger.Errorw("failed to generate jwt", "error", jwtErr)
			return nil, status.Errorf(codes.Internal, "internal error")
		}
		resp.Jwt = jwt
	}

	return resp, nil
}

func (s *AccountService) LoginAccount(ctx context.Context, request *v1.LoginAccountRequest) (*v1.LoginAccountResponse, error) {
	credential, backend, decodeErr := s.decodeLoginCredential(request.Credential)
	if decodeErr != nil {
		return nil, decodeErr
	}

	account, loginErr := backend.Authenticate(ctx, getAccountNamespace(request.Namespace), credential)
	if loginErr != nil {
		s.logger.Debugw("failed to authenticate account", "error", loginErr)
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	jwt, jwtErr := s.jwtManager.Generate(ctx, account)
	if jwtErr != nil {
		s.logger.Errorw("failed to generate jwt", "error", jwtErr)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &v1.LoginAccountResponse{
		Status: getAccountStatus(account),
		Jwt:    jwt,
	}, nil
}

func (s *AccountService) BeginPasswordReset(ctx context.Context, request *v1.BeginPasswordResetRequest) (*v1.BeginPasswordResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *AccountService) CompletePasswordReset(ctx context.Context, request *v1.CompletePasswordResetRequest) (*v1.CompletePasswordResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *AccountService) GetAccount(ctx context.Context, request *v1.GetAccountRequest) (*v1.GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func accountToApi(account *auth.Account) *v1.Account {
	var apiNamespace v1.AccountNamespace
	switch account.Namespace {
	case auth.AccountNamespaceDeveloper:
		apiNamespace = v1.AccountNamespace_ACCOUNT_NAMESPACE_DEVELOPER
	case auth.AccountNamespaceConsumer:
		apiNamespace = v1.AccountNamespace_ACCOUNT_NAMESPACE_CONSUMER
	default:
		apiNamespace = v1.AccountNamespace_ACCOUNT_NAMESPACE_UNSPECIFIED
	}

	var apiCredentialType v1.AccountCredentialType
	switch account.CredentialsType {
	case auth.CredentialsTypeEmailPassword:
		apiCredentialType = v1.AccountCredentialType_ACCOUNT_CREDENTIAL_TYPE_EMAIL
	default:
		apiCredentialType = v1.AccountCredentialType_ACCOUNT_CREDENTIAL_TYPE_UNSPECIFIED
	}

	return &v1.Account{
		Id:             account.ID,
		Namespace:      apiNamespace,
		CredentialType: apiCredentialType,
		Name:           account.Name,
		Email:          account.Email,
		AvatarUrl:      account.AvatarURL,
		CreatedAt:      timestamppb.New(account.CreatedAt),
		UpdatedAt:      timestamppb.New(account.UpdatedAt),
	}
}
