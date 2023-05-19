package account_psql_test

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/auth/stores/account_psql"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type PostgresSuite struct {
	suite.Suite
	store auth.AccountStore
}

func TestPostgresSuite(t *testing.T) {
	suite.Run(t, new(PostgresSuite))
}

func (s *PostgresSuite) BeforeTest(_, _ string) {
	pg, err := pgx.Connect(context.Background(), "postgres://postgres:docker@localhost:5432/fira?sslmode=disable")
	s.Require().NoError(err)
	_, err = pg.Exec(context.Background(), "TRUNCATE accounts")
	s.Require().NoError(err)

	s.store = account_psql.New()
}

func (s *PostgresSuite) TestAccountStore_Create_SetsID() {
	acct := &auth.Account{
		Namespace:       auth.AccountNamespaceConsumer,
		CredentialsType: auth.CredentialsTypeEmailPassword,
	}
	cerr := s.store.Create(context.Background(), acct, map[string]string{"email": "test@test.net"})
	s.Require().NoError(cerr)
	s.Require().NotEmpty(acct.ID)
}

func (s *PostgresSuite) TestAccountStore_Create_DisallowsDuplicates() {
	acct := testAcct()
	distinctCreds := map[string]string{"email": "test@test.net"}
	cerr := s.store.Create(context.Background(), acct, distinctCreds)
	s.Require().NoError(cerr)

	acct2 := testAcct()
	cerr = s.store.Create(context.Background(), acct2, distinctCreds)
	s.Require().ErrorIs(cerr, auth.ErrAccountExists)
}

func (s *PostgresSuite) TestAccountStore_FindAccountByID() {
	acct := testAcct()
	cerr := s.store.Create(context.Background(), acct, map[string]string{})
	s.Require().NoError(cerr)

	acct2, err := s.store.FindAccountByID(context.Background(), auth.AccountNamespaceConsumer, acct.ID)
	s.Require().NoError(err)
	s.Equal(acct.ID, acct2.ID)
}

func (s *PostgresSuite) Test_AccountStore_Update() {
	acct := testAcct()
	cerr := s.store.Create(context.Background(), acct, map[string]string{})
	s.Require().NoError(cerr)

	acct.Credentials["email"] = "different@test.net"

	err := s.store.Update(context.Background(), acct)
	s.Require().NoError(err)

	acct2, err := s.store.FindAccountByID(context.Background(), auth.AccountNamespaceConsumer, acct.ID)
	s.Require().NoError(err)
	s.Equal(acct.ID, acct2.ID)
}

func (s *PostgresSuite) TestAccountStore_FindByCredentials() {
	acct := testAcct()
	distinctCreds := map[string]string{"email": "test@test.net"}
	cerr := s.store.Create(context.Background(), acct, distinctCreds)
	s.Require().NoError(cerr)

	acct2, err := s.store.FindByCredentials(context.Background(), auth.AccountNamespaceConsumer, distinctCreds)
	s.Require().NoError(err)
	s.Equal(acct.ID, acct2.ID)
}

func (s *PostgresSuite) TestAccountStore_FindByCredentials_Missing() {
	acct := testAcct()
	distinctCreds := map[string]string{"email": "test@test.net"}
	cerr := s.store.Create(context.Background(), acct, distinctCreds)
	s.Require().NoError(cerr)

	missingCreds := map[string]string{"email": "uhh@test.net"}
	acct2, err := s.store.FindByCredentials(context.Background(), auth.AccountNamespaceConsumer, missingCreds)
	s.Require().ErrorIs(err, auth.ErrNoAccount)
	s.Nil(acct2)
}

func (s *PostgresSuite) TestAccountStore_FindByID_Missing() {
	acct := testAcct()
	distinctCreds := map[string]string{"email": "test@test.net"}
	cerr := s.store.Create(context.Background(), acct, distinctCreds)
	s.Require().NoError(cerr)

	acct2, err := s.store.FindAccountByID(context.Background(), auth.AccountNamespaceConsumer, "1234")
	s.Require().ErrorIs(err, auth.ErrNoAccount)
	s.Nil(acct2)
}

// TODO: assert update column values
// TODO: assert read column values

func testAcct() *auth.Account {
	return &auth.Account{
		Namespace:       auth.AccountNamespaceConsumer,
		CredentialsType: auth.CredentialsTypeEmailPassword,
		Credentials: map[string]string{
			"email":    "test@test.net",
			"password": "oooh-you-shoulda-hashed-this",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
