package email

import "context"

type Sender interface {
	SendOne(ctx context.Context, from, to, subject, templateName string, templateVariables map[string]string) error
}

type Provider interface {
	Sender() Sender
}
