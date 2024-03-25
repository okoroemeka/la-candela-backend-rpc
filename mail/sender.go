package mail

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

const (
	smtpAuthServer = "smtp.gmail.com"
	smtpServerAdr  = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(subject, content string, to, cc, bcc, attachFiles []string) error
}

type GmailSender struct {
	name, FromEmailAddress, FromEmailPassword string
}

func NewGmailSender(name, fromEmailAddress, fromEmailPassword string) EmailSender {
	return &GmailSender{name: name, FromEmailAddress: fromEmailAddress, FromEmailPassword: fromEmailPassword}

}

func (sender *GmailSender) SendEmail(subject, content string, to, cc, bcc, attachFiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.FromEmailAddress)
	e.To = to
	e.Subject = subject
	e.Cc = cc
	e.Bcc = bcc
	e.HTML = []byte(content)

	for _, file := range attachFiles {
		_, err := e.AttachFile(file)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", file, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.FromEmailAddress, sender.FromEmailPassword, smtpAuthServer)

	return e.Send(smtpAuthServer, smtpAuth)
}
