package application

import "github.com/opencorelabs/fira/internal/email"

func (a *App) Sender() email.Sender {
	a.initMtx.Lock()
	defer a.initMtx.Unlock()

	if a.emailSender == nil {
		a.emailSender = email.NewMailgunSender(a, a.cfg.MailgunDomain, a.cfg.MailgunApiKey)
	}

	return a.emailSender
}
