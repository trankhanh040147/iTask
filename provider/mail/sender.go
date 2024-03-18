package mail

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

type EmailSender interface {
	SendEmail(
		subject string,
		body string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string) error
}

const (
	smptAuthAddress = "smtp.gmail.com"
	smtpServerAddr  = "smtp.gmail.com:587"
)

type gmailSender struct {
	name              string
	fromEmailAddr     string
	fromEmailPassword string
}

func NewGmailSender(name string, fromEmailAddr string, fromEmailPassword string) EmailSender {
	return &gmailSender{
		name:              name,
		fromEmailAddr:     fromEmailAddr,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *gmailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string) error {

	e := email.NewEmail()
	e.From = sender.name + "<" + sender.fromEmailAddr + ">"
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.Subject = subject
	e.HTML = []byte(content)

	for _, attachFile := range attachFiles {
		_, err := e.AttachFile(attachFile)
		if err != nil {
			return err
		}
	}
	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddr, sender.fromEmailPassword, smptAuthAddress)
	return e.Send(smtpServerAddr, smtpAuth)
}
