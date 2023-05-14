package api

import (
	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/auth/backends/email_password"
	"github.com/opencorelabs/fira/internal/auth/stores/in_memory"
	"github.com/opencorelabs/fira/internal/auth/verification"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"testing"
)

type FiraApiSuite struct {
	suite.Suite
	logger    *zap.Logger
	api       v1.FiraServiceServer
	acctStore auth.AccountStore
}

func TestFiraApiSuite(t *testing.T) {
	suite.Run(t, new(FiraApiSuite))
}

func (s *FiraApiSuite) SetupSuite() {
	s.logger, _ = zap.NewDevelopment()
}

func (s *FiraApiSuite) BeforeTest(_, _ string) {
	s.acctStore = in_memory.New()

	authReg := auth.NewDefaultRegistry()
	authReg.RegisterBackend(auth.CredentialsTypeEmailPassword, email_password.New(s, s))

	s.api = New(s, authReg, auth.TodoJWTManager, s)
}

func (s *FiraApiSuite) Logger() *zap.Logger {
	return s.logger
}

func (s *FiraApiSuite) AccountStore() auth.AccountStore {
	return s.acctStore
}

func (s *FiraApiSuite) Email() verification.Verifier {
	return verification.NewLoggingVerifier(s, s)
}