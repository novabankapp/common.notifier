package email

import (
	"fmt"
	"net/smtp"
)

type MailService interface {
	SendEmail(to []string, message string) (bool, error)
}

type SMTPConfig struct {
	FromEmail string
	Host      string
	Username  string
	Password  string
	Port      int16
}

type smtpMailService struct {
	config SMTPConfig
}

func NewSmtpService(config SMTPConfig) MailService {
	return &smtpMailService{
		config: config,
	}
}
func (s smtpMailService) SendEmail(to []string, message string) (bool, error) {
	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)
	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	err := smtp.SendMail(addr, auth, s.config.FromEmail, to, []byte(message))

	if err != nil {
		return false, err
	}
	return true, nil
}
