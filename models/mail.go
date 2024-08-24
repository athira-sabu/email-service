package models

type Email struct {
	Sender    string
	Recipient []string
	Subject   string
	Body      string
}
