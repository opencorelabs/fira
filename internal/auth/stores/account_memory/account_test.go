package account_memory_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/auth/stores/account_memory"
	"github.com/stretchr/testify/suite"
	"testing"
)

type InMemorySuite struct {
	suite.Suite
	store auth.AccountStore
}

func TestInMemorySuite(t *testing.T) {
	suite.Run(t, new(InMemorySuite))
}

func (s *InMemorySuite) BeforeTest(_, _ string) {
	s.store = account_memory.New()
}

func (s *InMemorySuite) TestAccountStore_Create_SetsID() {
	acct := &auth.Account{
		Namespace:       auth.AccountNamespaceConsumer,
		CredentialsType: auth.CredentialsTypeEmailPassword,
	}
	cerr := s.store.Create(context.Background(), acct, map[string]string{"email": "test@test.net"})
	s.Require().NoError(cerr)
	s.Require().NotEmpty(acct.ID)
}

func (s *InMemorySuite) TestAccountStore_Create_DisallowsDuplicates() {
	acct := testAcct()
	distinctCreds := map[string]string{"email": "test@test.net"}
	cerr := s.store.Create(context.Background(), acct, distinctCreds)
	s.Require().NoError(cerr)

	acct2 := testAcct()
	cerr = s.store.Create(context.Background(), acct2, distinctCreds)
	s.Require().ErrorIs(cerr, auth.ErrAccountExists)
}

func (s *InMemorySuite) TestAccountStore_FindAccountByID() {
	acct := testAcct()
	cerr := s.store.Create(context.Background(), acct, map[string]string{})
	s.Require().NoError(cerr)

	acct2, err := s.store.FindAccountByID(context.Background(), auth.AccountNamespaceConsumer, acct.ID)
	s.Require().NoError(err)
	s.Equal(acct, acct2)
}

func (s *InMemorySuite) Test_AccountStore_Update() {
	acct := testAcct()
	cerr := s.store.Create(context.Background(), acct, map[string]string{})
	s.Require().NoError(cerr)

	acct.Credentials["email"] = "different@test.net"

	err := s.store.Update(context.Background(), acct)
	s.Require().NoError(err)

	acct2, err := s.store.FindAccountByID(context.Background(), auth.AccountNamespaceConsumer, acct.ID)
	s.Require().NoError(err)
	s.Equal(acct, acct2)
}

func (s *InMemorySuite) TestAccountStore_FindByCredentials() {
	acct := testAcct()
	distinctCreds := map[string]string{"email": "test@test.net"}
	cerr := s.store.Create(context.Background(), acct, distinctCreds)
	s.Require().NoError(cerr)

	acct2, err := s.store.FindByCredentials(context.Background(), auth.AccountNamespaceConsumer, distinctCreds)
	s.Require().NoError(err)
	s.Equal(acct, acct2)
}

func (s *InMemorySuite) TestAccountStore_FindByCredentials_Missing() {
	acct := testAcct()
	distinctCreds := map[string]string{"email": "test@test.net"}
	cerr := s.store.Create(context.Background(), acct, distinctCreds)
	s.Require().NoError(cerr)

	missingCreds := map[string]string{"email": "uhh@test.net"}
	acct2, err := s.store.FindByCredentials(context.Background(), auth.AccountNamespaceConsumer, missingCreds)
	s.Require().ErrorIs(err, auth.ErrNoAccount)
	s.Nil(acct2)
}

func (s *InMemorySuite) TestAccountStore_FindByID_Missing() {
	acct := testAcct()
	distinctCreds := map[string]string{"email": "test@test.net"}
	cerr := s.store.Create(context.Background(), acct, distinctCreds)
	s.Require().NoError(cerr)

	acct2, err := s.store.FindAccountByID(context.Background(), auth.AccountNamespaceConsumer, uuid.NewString())
	s.Require().ErrorIs(err, auth.ErrNoAccount)
	s.Nil(acct2)
}

func testAcct() *auth.Account {
	return &auth.Account{
		Namespace:       auth.AccountNamespaceConsumer,
		CredentialsType: auth.CredentialsTypeEmailPassword,
		Credentials: map[string]string{
			"email":    "test@test.net",
			"password": "oooh-you-shoulda-hashed-this",
		},
	}
}
