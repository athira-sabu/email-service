# email-service
This project is an email service that uses Mailgun to send emails. The service is configured using environment variables for the Mailgun API key and domain.

## Features
- Send emails using Mailgun
- Configure the service using environment variables

## Configuration
Create a .env file in the root of the project and add your Mailgun API key and domain:
````
MAILGUN_API_KEY=your-mailgun-api-key
MAILGUN_DOMAIN=your-mailgun-domain