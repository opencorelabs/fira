package auth_test

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/auth/stores/account_memory"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"testing"
	"time"
)

type AccountJWTManagerSuite struct {
	suite.Suite
	acctStore auth.AccountStore
	mgr       auth.JWTManager
	secrets   [][]byte
	secretsFn auth.SecretsFn
}

func TestAccountJWTManagerSuite(t *testing.T) {
	suite.Run(t, new(AccountJWTManagerSuite))
}

func (s *AccountJWTManagerSuite) BeforeTest(suiteName, testName string) {
	s.acctStore = account_memory.New()
	s.secrets = [][]byte{[]byte("secret1")}
	s.secretsFn = func(context.Context) [][]byte {
		return s.secrets
	}
	s.mgr = auth.NewAccountJWTManager(s.secretsFn, time.Minute, s, s)
}

func (s *AccountJWTManagerSuite) AccountStore() auth.AccountStore {
	return s.acctStore
}

func (s *AccountJWTManagerSuite) Logger() *zap.Logger {
	l, _ := zap.NewDevelopment()
	return l
}

func (s *AccountJWTManagerSuite) TestGenerate() {
	tok, tokErr := s.mgr.Generate(context.Background(), &auth.Account{
		ID:        "test",
		Namespace: auth.AccountNamespaceConsumer,
	})
	s.Require().NoError(tokErr)
	s.Require().NotEmpty(tok)

	p := jwt.NewParser()
	claims := &auth.FiraClaims{}
	pTok, _, err := p.ParseUnverified(tok, claims)
	s.Require().NoError(err)
	s.Require().NotNil(pTok)
	s.Equal(claims.AccountID, "test")
	s.Equal(claims.AccountNamespace, string(auth.AccountNamespaceConsumer))
}

func (s *AccountJWTManagerSuite) TestVerify() {
	acct := &auth.Account{Namespace: auth.AccountNamespaceConsumer, Valid: true}
	createErr := s.acctStore.Create(context.Background(), acct, map[string]string{})
	s.NoErrorf(createErr, "failed to create account")

	tok, tokErr := s.mgr.Generate(context.Background(), acct)
	s.Require().NoError(tokErr)
	s.Require().NotEmpty(tok)

	ctx, ctxErr := s.mgr.Verify(context.Background(), tok)
	s.Require().NoError(ctxErr)
	s.Require().NotNil(ctx)

	acctFromCtx, ok := auth.AccountFromContext(ctx)
	s.Require().True(ok)
	s.Require().NotNil(acctFromCtx)
	s.Require().Equal(acct.ID, acctFromCtx.ID)
}

func (s *AccountJWTManagerSuite) TestVerify_WithRotatedSecrets() {
	acct := &auth.Account{Namespace: auth.AccountNamespaceConsumer, Valid: true}
	createErr := s.acctStore.Create(context.Background(), acct, map[string]string{})
	s.NoErrorf(createErr, "failed to create account")

	tok, tokErr := s.mgr.Generate(context.Background(), acct)
	s.Require().NoError(tokErr)
	s.Require().NotEmpty(tok)

	// "rotate" secrets by adding a new one
	s.secrets = [][]byte{[]byte("secret1"), []byte("secret2")}

	_, ctxErr := s.mgr.Verify(context.Background(), tok)
	s.Require().NoError(ctxErr)

	// "rotate" secrets by removing the old one
	s.secrets = [][]byte{[]byte("secret2")}

	acctCtx, secondCtxErr := s.mgr.Verify(context.Background(), tok)
	s.Nil(acctCtx, "account should not be in context")
	s.Require().ErrorContains(secondCtxErr, "invalid token")
}
