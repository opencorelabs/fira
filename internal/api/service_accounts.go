package api

import (
	"context"
	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/auth/verification"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FiraApiService) decodeLoginCredential(c *v1.LoginCredential) (map[string]string, auth.Backend, error) {
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

func (s *FiraApiService) getAccountStatus(acct *auth.Account) v1.AccountRegistrationStatus {
	if acct.Valid {
		return v1.AccountRegistrationStatus_ACCOUNT_REGISTRATION_STATUS_OK
	}
	switch acct.CredentialsType {
	case auth.CredentialsTypeEmailPassword:
		return v1.AccountRegistrationStatus_ACCOUNT_REGISTRATION_STATUS_VERIFY_EMAIL
	}
	return v1.AccountRegistrationStatus_ACCOUNT_REGISTRATION_STATUS_UNSPECIFIED
}

func (s *FiraApiService) CreateAccount(ctx context.Context, request *v1.CreateAccountRequest) (*v1.CreateAccountResponse, error) {
	credential, backend, decodeErr := s.decodeLoginCredential(request.Credential)
	if decodeErr != nil {
		return nil, decodeErr
	}

	account, regErr := backend.Register(ctx, credential)

	if regErr != nil {
		s.logger.Errorw("failed to register account",
			"error", regErr,
			"credential_type", request.Credential.CredentialType,
		)
		return nil, status.Errorf(codes.Unauthenticated, "registration error")
	}

	jwt, jwtErr := s.jwtManager.Generate(account.ID)
	if jwtErr != nil {
		s.logger.Errorw("failed to generate jwt", "error", jwtErr)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &v1.CreateAccountResponse{
		Status: s.getAccountStatus(account),
		Jwt:    jwt,
	}, nil
}

func (s *FiraApiService) VerifyAccount(ctx context.Context, request *v1.VerifyAccountRequest) (*v1.VerifyAccountResponse, error) {
	var verifier verification.Verifier
	switch request.Type {
	case v1.VerificationType_VERIFICATION_TYPE_EMAIL:
		verifier = s.verificationProvider.Email()
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid verification type")
	}

	acct, verifErr := verifier.VerifyWithToken(ctx, request.Token)
	if verifErr != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid verification token")
	}
	resp := &v1.VerifyAccountResponse{
		Status: s.getAccountStatus(acct),
	}
	if acct.Valid {
		jwt, jwtErr := s.jwtManager.Generate(acct.ID)
		if jwtErr != nil {
			s.logger.Errorw("failed to generate jwt", "error", jwtErr)
			return nil, status.Errorf(codes.Internal, "internal error")
		}
		resp.Jwt = jwt
	}

	return resp, nil
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
