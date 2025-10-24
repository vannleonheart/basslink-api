package basslink

import (
	"context"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type MailgunClient struct {
	config *MailgunConfig
	mg     *mailgun.MailgunImpl
}

type MailgunConfig struct {
	Domain string `json:"domain"`
	ApiKey string `json:"api_key"`
	From   string `json:"from"`
}

func NewMailgunClient(config *MailgunConfig) *MailgunClient {
	mg := mailgun.NewMailgun(config.Domain, config.ApiKey)
	return &MailgunClient{
		config: config,
		mg:     mg,
	}
}

func (c *MailgunClient) SendEmail(to string, subject string, body *string) (string, error) {
	message := mailgun.NewMessage(c.config.From, subject, "", to)
	message.SetHTML(*body)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	_, id, err := c.mg.Send(ctx, message)

	return id, err
}
