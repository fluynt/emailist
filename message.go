package emailist

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wneessen/go-mail"
)

func NewMessage(subject, text, html string) *Message {
	return &Message{Subject: subject, TextMessage: text, HtmlMessage: html}
}

// TODO send message here
func (c *Config) SendMail(i *Identity, to []string, m *Message) error {
	if m.TextMessage == "" && m.HtmlMessage == "" {
		return fmt.Errorf("no message to send")
	}

	if os.Getenv("DEBUG") == "" {
		logrus.Info("--TEST MESSAGE--")
		logrus.Info(m.TextMessage)
		logrus.Info("--------------------")
		logrus.Info("--HTML MESSAGE-")
		logrus.Info(m.TextMessage)
		return nil
	}

	mm := mail.NewMsg()
	mm.From(i.From())
	mm.SetDate()
	mm.Subject(m.Subject)
	mm.To(to...)

	if m.HtmlMessage != "" {
		mm.SetBodyString(mail.TypeTextHTML, m.HtmlMessage)
		if m.TextMessage != "" {
			mm.AddAlternativeString(mail.TypeTextPlain, m.TextMessage)
		}
	} else {
		mm.SetBodyString(mail.TypeTextPlain, m.TextMessage)
	}

	host := c.Host
	if host == "" {
		host = "locallhost"
	}

	port := 25
	if c.Port != 0 {
		port = c.Port
	}
	cli, err := mail.NewClient(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}

	if c.User != "" {
		cli.SetSMTPAuth(mail.SMTPAuthPlain)
		cli.SetUsername(c.User)
		cli.SetPassword(c.Password)
	}
	if c.TLS {
		cli.SetTLSPolicy(mail.TLSMandatory)
	}
	if c.SSL {
		cli.SetSSL(true)
	}

	err = cli.DialAndSend(mm)
	if err != nil {
		return err
	}

	cli.Close()

	return nil
}
