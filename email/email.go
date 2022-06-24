package email

import (
	"github.com/thgamejam/pkg/conf"
	"gopkg.in/gomail.v2"
)

type emailService struct {
	email *gomail.Dialer
	user  string
}

func NewEmailService(c *conf.Email) (*emailService, error) {
	d := gomail.NewDialer(c.Host, int(c.Port), c.Username, c.Password)

	return &emailService{
		email: d,
		user:  c.Username,
	}, nil
}

// SendEmail 发送邮件from是发送人名称，subject是主题，mailto是收件人
func (e *emailService) SendEmail(from, subject, body string, mailTo ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(e.user, from))
	m.SetHeader("Subject", subject)
	m.SetHeader("To", mailTo...)
	m.SetBody("text/html", body)
	err := e.email.DialAndSend(m)
	return err
}
