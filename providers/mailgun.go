package providers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/athira-sabu/email-service/models"
	"github.com/mailgun/mailgun-go/v4"
)

type MailgunService struct {
	APIKey string
	Domain string
}

func (mg *MailgunService) SendEmail(mail *models.Email) error {
	mailGun := mailgun.NewMailgun(mg.Domain, mg.APIKey)

	message := mailGun.NewMessage(
		mail.Sender,
		mail.Subject,
		mail.Body,
		mail.Recipient...,
	)

	message.SetHtml(mail.Body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := mailGun.Send(ctx, message)

	if err != nil {
		log.Printf("Failed to send email: %v\n", err)
		return err
	}

	log.Printf("Email sent successfully! ID: %s Resp: %s\n", id, resp)
	return nil
}

func NewMailgun(apiKey string, domain string) (*MailgunService, error) {
	if apiKey == "" || domain == "" {
		return nil, fmt.Errorf("apiKey and domain cannot be empty")
	}
	return &MailgunService{
		APIKey: apiKey,
		Domain: domain,
	}, nil
}
