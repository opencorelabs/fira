package app_psql_test

import (
	"context"
	"fmt"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/auth/stores/account_psql"
	"github.com/opencorelabs/fira/internal/developer"
	"github.com/opencorelabs/fira/internal/developer/stores/app_psql"
	"github.com/opencorelabs/fira/internal/persistence/psql"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PostgresSuite struct {
	suite.Suite
	psqlHelper   *psql.TestHelper
	accountStore auth.AccountStore
	store        developer.AppStore
	account1     *auth.Account
}

func TestPostgresSuite(t *testing.T) {
	suite.Run(t, new(PostgresSuite))
}

func (s *PostgresSuite) SetupSuite() {
	s.psqlHelper = psql.NewTestHelper(&s.Suite)
	s.psqlHelper.Migrate()
}

func (s *PostgresSuite) TearDownSuite() {
	s.psqlHelper.Close()
}

func (s *PostgresSuite) BeforeTest(_, _ string) {
	s.psqlHelper.Reset()

	s.accountStore = account_psql.New(s.psqlHelper, s.psqlHelper)
	s.store = app_psql.New(s.psqlHelper, s.psqlHelper)

	acct := &auth.Account{
		Namespace:       auth.AccountNamespaceConsumer,
		CredentialsType: auth.CredentialsTypeEmailPassword,
		Credentials:     map[string]string{"email": "test@test.net"},
	}
	acctErr := s.accountStore.Create(context.Background(), acct, acct.Credentials)
	s.Require().NoError(acctErr)

	s.account1 = acct
}

func (s *PostgresSuite) TestAppStore_CreateApp_SetsIDAndDates() {
	app := &developer.App{AccountID: s.account1.ID}
	cerr := s.store.CreateApp(context.Background(), app)
	s.Require().NoError(cerr)
	s.Require().NotEmpty(app.ID)
	s.Require().NotZero(app.CreatedAt)
	s.Require().NotZero(app.UpdatedAt)
}

func (s *PostgresSuite) TestAppStore_UpdateApp_UpdatesUpdatedAt() {
	app := &developer.App{AccountID: s.account1.ID}
	cerr := s.store.CreateApp(context.Background(), app)
	s.Require().NoError(cerr)
	updatedAt := app.UpdatedAt

	app.Name = "New name"
	cerr = s.store.UpdateApp(context.Background(), app)
	s.Require().NoError(cerr)
	s.Require().NotEqual(updatedAt, app.UpdatedAt)
}

func (s *PostgresSuite) TestAppStore_CreateThenGetByID() {
	app := &developer.App{AccountID: s.account1.ID, Name: "Test App"}
	s.Require().NoError(app.Rotate(developer.EnvironmentSandbox))
	cerr := s.store.CreateApp(context.Background(), app)
	s.Require().NoError(cerr)

	got, gerr := s.store.GetAppByID(context.Background(), app.ID)
	s.Require().NoError(gerr)
	s.Require().Equal(app.ID, got.ID)
	s.Require().Equal(app.Name, got.Name)
	s.Require().Equal(app.Tokens[developer.EnvironmentSandbox][0].Key, got.Tokens[developer.EnvironmentSandbox][0].Key)
	s.Require().Equal(app.AccountID, got.AccountID)
}

func (s *PostgresSuite) TestAppStore_GetByAccountID() {
	for i := 0; i < 10; i++ {
		app := &developer.App{AccountID: s.account1.ID, Name: fmt.Sprintf("Test App %d", i)}
		s.Require().NoError(app.Rotate(developer.EnvironmentSandbox))
		cerr := s.store.CreateApp(context.Background(), app)
		s.Require().NoError(cerr)
	}

	got, gerr := s.store.GetAppsByAccountID(context.Background(), s.account1.ID)
	s.Require().NoError(gerr)

	s.Require().Len(got, 10)

	for _, app := range got {
		s.Require().Equal(s.account1.ID, app.AccountID)
	}
}
