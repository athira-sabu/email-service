package service

import (
	"github.com/athira-sabu/email-service/models"
	"github.com/athira-sabu/email-service/providers"
)

type EmailClient struct {
	Provider EmailProvider
}

type EmailProvider interface {
	SendEmail(email *models.Email) error
}

func (client *EmailClient) SendEmail(email *models.Email) error {
	return client.Provider.SendEmail(email)
}

func NewEmailClient(config map[string]string) (*EmailClient, error) {
	mailgunProvider, err := providers.NewMailgun(config["MAILGUN_API_KEY"], config["MAILGUN_DOMAIN"])
	if err != nil {
		return nil, err
	}

	return &EmailClient{
		Provider: mailgunProvider,
	}, nil

}
