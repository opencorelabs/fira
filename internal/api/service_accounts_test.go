package api

import (
	"context"
	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
	"github.com/opencorelabs/fira/internal/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FiraApiSuite) Test_CreateAccount_EmailPassword() {
	resp, err := s.createAccount()

	s.Require().NoError(err, "failed to create account")
	s.Require().NotNil(resp, "response is nil")

	s.Equal(v1.AccountRegistrationStatus_ACCOUNT_REGISTRATION_STATUS_VERIFY_EMAIL, resp.Status, "status is not verify email")
	s.NotEmpty(resp.Jwt, "jwt is empty")
}

func (s *FiraApiSuite) Test_VerifyAccount_EmailPassword() {
	_, createErr := s.createAccount()
	s.Require().NoErrorf(createErr, "failed to create account")

	resp, verifErr := s.verifyAccount("pnwx@opencorelabs.org")

	s.Require().NoErrorf(verifErr, "failed to verify account")
	s.Require().NotNil(resp, "response is nil")

	s.Equal(v1.AccountRegistrationStatus_ACCOUNT_REGISTRATION_STATUS_OK, resp.Status, "status is not ok")
}

func (s *FiraApiSuite) Test_LoginAccount_EmailPassword() {
	s.verifiedAccount()

	resp, loginErr := s.loginAccount()

	s.Require().NoErrorf(loginErr, "failed to login account")
	s.Require().NotNil(resp, "response is nil")
	s.Require().Equal(v1.AccountRegistrationStatus_ACCOUNT_REGISTRATION_STATUS_OK, resp.Status, "status is not ok")
	s.Require().NotEmpty(resp.Jwt, "jwt is empty")
}

func (s *FiraApiSuite) Test_LoginAccount_BeforeVerification_EmailPassword() {
	_, createErr := s.createAccount()
	s.Require().NoErrorf(createErr, "failed to create account")

	resp, loginErr := s.loginAccount()

	s.Require().NoErrorf(loginErr, "failed to login account")
	s.Require().NotNil(resp, "response is nil")
	// status should tell user to verify email
	s.Require().Equal(v1.AccountRegistrationStatus_ACCOUNT_REGISTRATION_STATUS_VERIFY_EMAIL, resp.Status, "status is not verify email")
	s.Require().NotEmpty(resp.Jwt, "jwt is not empty")
}

func (s *FiraApiSuite) Test_LoginAccount_NoAccountExists() {
	req := &v1.LoginAccountRequest{
		Credential: &v1.LoginCredential{
			CredentialType: v1.AccountCredentialType_ACCOUNT_CREDENTIAL_TYPE_EMAIL,
			EmailCredential: &v1.CredentialTypeEmail{
				Email:    "idunno@opencorelabs.org",
				Password: "simplepassword",
			},
		},
	}

	_, respErr := s.api.LoginAccount(context.Background(), req)
	stat := status.Convert(respErr)
	s.Equal(codes.Unauthenticated, stat.Code(), "status code is not unauthenticated")
}

func (s *FiraApiSuite) Test_LoginAccount_WrongPassword() {
	s.verifiedAccount()

	req := &v1.LoginAccountRequest{
		Credential: &v1.LoginCredential{
			CredentialType: v1.AccountCredentialType_ACCOUNT_CREDENTIAL_TYPE_EMAIL,
			EmailCredential: &v1.CredentialTypeEmail{
				Email:    "pnwx@opencorelabs.org",
				Password: "simplepassword-wrong",
			},
		},
	}

	_, respErr := s.api.LoginAccount(context.Background(), req)
	stat := status.Convert(respErr)
	s.Equal(codes.Unauthenticated, stat.Code(), "status code is not unauthenticated")
}

func (s *FiraApiSuite) createAccount() (*v1.CreateAccountResponse, error) {
	req := &v1.CreateAccountRequest{
		Namespace: v1.AccountNamespace_ACCOUNT_NAMESPACE_CONSUMER,
		Credential: &v1.LoginCredential{
			CredentialType: v1.AccountCredentialType_ACCOUNT_CREDENTIAL_TYPE_EMAIL,
			EmailCredential: &v1.CredentialTypeEmail{
				Email:    "pnwx@opencorelabs.org",
				Password: "simplepassword",
			},
		},
	}

	return s.api.CreateAccount(context.TODO(), req)
}

func (s *FiraApiSuite) verifyAccount(email string) (*v1.VerifyAccountResponse, error) {
	// read out the verification token
	usr, findErr := s.acctStore.FindByCredentials(context.TODO(), auth.AccountNamespaceConsumer, map[string]string{
		"email": email,
	})
	s.Require().NoErrorf(findErr, "failed to find account")
	tok, hasTok := usr.Credentials["logging_verification_token"]
	s.Require().Truef(hasTok, "account does not have verification token")

	// verify the account
	req := &v1.VerifyAccountRequest{
		Token:     tok,
		Namespace: v1.AccountNamespace_ACCOUNT_NAMESPACE_CONSUMER,
		Type:      v1.VerificationType_VERIFICATION_TYPE_EMAIL,
	}

	return s.api.VerifyAccount(context.Background(), req)
}

func (s *FiraApiSuite) loginAccount() (*v1.LoginAccountResponse, error) {
	loginReq := &v1.LoginAccountRequest{
		Namespace: v1.AccountNamespace_ACCOUNT_NAMESPACE_CONSUMER,
		Credential: &v1.LoginCredential{
			CredentialType: v1.AccountCredentialType_ACCOUNT_CREDENTIAL_TYPE_EMAIL,
			EmailCredential: &v1.CredentialTypeEmail{
				Email:    "pnwx@opencorelabs.org",
				Password: "simplepassword",
			},
		},
	}

	return s.api.LoginAccount(context.Background(), loginReq)
}

func (s *FiraApiSuite) verifiedAccount() {
	_, createErr := s.createAccount()
	s.Require().NoErrorf(createErr, "failed to create account")
	_, verifErr := s.verifyAccount("pnwx@opencorelabs.org")
	s.Require().NoErrorf(verifErr, "failed to verify account")
}

func (s *FiraApiSuite) acctCtx() (*auth.Account, context.Context) {
	s.verifiedAccount()
	acct, acctErr := s.acctStore.FindByCredentials(context.TODO(), auth.AccountNamespaceConsumer, map[string]string{
		"email": "pnwx@opencorelabs.org",
	})
	s.Require().NoErrorf(acctErr, "failed to find account")
	s.Require().NotNil(acct, "account is nil")
	return acct, auth.WithAccount(context.Background(), acct)
}
