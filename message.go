package emailist

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func NewMessage(text, html string) *Message {
	return &Message{TextMessage: text, HtmlMessage: html}
}

// TODO send message here
func (c *Config) SendMail(from *Identity, to []string, m *Message) error {
	if m.TextMessage == "" && m.HtmlMessage == "" {
		return fmt.Errorf("no message to send")
	}

	if os.Getenv("DEBUG") == "" {
		logrus.Info("--TEST MESSAGE--")
		logrus.Info(m.TextMessage)

		logrus.Info("--------------------")

		logrus.Info("--HTML MESSAGE-")
		logrus.Info(m.TextMessage)
	}
	return nil
}
