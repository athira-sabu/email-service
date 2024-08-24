// main.go

package main

import (
	"log"

	"github.com/athira-sabu/email-service/config"
	"github.com/athira-sabu/email-service/models"
	"github.com/athira-sabu/email-service/service"
)

func main() {
	config, err := config.LoadEnv()
	if err != nil {
		log.Fatal("Unable to load the config/env file")
	}

	emailClient, _ := service.NewEmailClient(config)

	mail := models.Email{
		Sender:    "",
		Recipient: []string{""},
		Subject:   "A Great Email",
		Body:      "With great things to say.",
	}

	emailClient.SendEmail(&mail)
}
