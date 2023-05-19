package email

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/opencorelabs/fira/internal/logging"
	"go.uber.org/zap"
	"time"
)

type MailgunSender struct {
	logger *zap.SugaredLogger
	mg     *mailgun.MailgunImpl
}

func NewMailgunSender(loggingProvider logging.Provider, domain, apiKey string) Sender {
	mg := mailgun.NewMailgun(domain, apiKey)
	return &MailgunSender{
		logger: loggingProvider.Logger().Named("mailgun_sender").Sugar(),
		mg:     mg,
	}
}

func (m *MailgunSender) SendOne(ctx context.Context, from, to, subject, templateName string, templateVariables map[string]string) error {
	message := m.mg.NewMessage(from, subject, "", to)
	message.SetTemplate(templateName)
	for k, v := range templateVariables {
		err := message.AddTemplateVariable(k, v)
		if err != nil {
			return fmt.Errorf("failed to add template variable: %w", err)
		}
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, id, err := m.mg.Send(ctx, message)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	m.logger.Infow("sent email", "id", id, "resp", resp)

	return nil
}
