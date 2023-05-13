package verification

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/logging"
	"go.uber.org/zap"
	"time"
)

// LoggingVerifier is a verifier that logs the token instead of sending it, useful for debug
type LoggingVerifier struct {
	logger   *zap.Logger
	accounts auth.AccountStoreProvider
}

func NewLoggingVerifier(provider logging.Provider, storeProvider auth.AccountStoreProvider) Verifier {
	return &LoggingVerifier{
		logger:   provider.Logger().Named("logging-verifier"),
		accounts: storeProvider,
	}
}

func (l *LoggingVerifier) SendVerificationToken(ctx context.Context, account *auth.Account) (map[string]string, error) {
	tok := uuid.NewString()
	l.logger.Info("verification token for email",
		zap.String("email", account.Credentials["email"]),
		zap.String("token", tok),
	)
	return map[string]string{
		"logging_verification_token":           tok,
		"logging_verification_token_timestamp": time.Now().Format(time.RFC3339Nano),
	}, nil
}

func (l *LoggingVerifier) VerifyWithToken(ctx context.Context, token string) (*auth.Account, error) {
	cred := map[string]string{
		"logging_verification_token": token,
	}
	user, userErr := l.accounts.AccountStore().FindByCredentials(ctx, cred)
	if userErr != nil {
		return nil, auth.ErrNoAccount
	}

	user.Valid = true
	user.Credentials["logging_verification_token"] = ""
	user.Credentials["logging_verification_token_timestamp"] = ""

	err := l.accounts.AccountStore().Update(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error updating account: %w", err)
	}
	return user, nil
}
