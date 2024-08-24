package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (map[string]string, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	envMap := make(map[string]string)
	envMap["MAILGUN_API_KEY"] = os.Getenv("MAILGUN_API_KEY")
	envMap["MAILGUN_DOMAIN"] = os.Getenv("MAILGUN_DOMAIN")
	return envMap, err
}
