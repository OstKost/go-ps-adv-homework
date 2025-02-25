package mail

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"go-ps-adv-homework/configs"
	"net/smtp"
)

func SendMail(config configs.EmailConfig, from, to, subject, text string) error {
	e := email.NewEmail()
	e.From = from       // Отправитель
	e.To = []string{to} // Получатель
	e.Subject = subject
	e.Text = []byte(text)
	smtpStr := config.SMTPHost + ":" + config.SMTPPort
	auth := smtp.PlainAuth("", config.Address, config.Password, config.SMTPHost)
	// Настройки TLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         config.SMTPHost,
	}
	err := e.SendWithTLS(smtpStr, auth, tlsConfig)
	return err
}
