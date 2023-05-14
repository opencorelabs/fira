package api

import (
	"context"
	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
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

func (s *FiraApiSuite) createAccount() (*v1.CreateAccountResponse, error) {
	req := &v1.CreateAccountRequest{
		Namespace: v1.AccountNamespace_ACCOUNT_NAMESPACE_DEVELOPER,
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
	usr, findErr := s.acctStore.FindByCredentials(context.TODO(), map[string]string{
		"email": email,
	})
	s.Require().NoErrorf(findErr, "failed to find account")
	tok, hasTok := usr.Credentials["logging_verification_token"]
	s.Require().Truef(hasTok, "account does not have verification token")

	// verify the account
	req := &v1.VerifyAccountRequest{
		Token: tok,
		Type:  v1.VerificationType_VERIFICATION_TYPE_EMAIL,
	}

	return s.api.VerifyAccount(context.Background(), req)
}
