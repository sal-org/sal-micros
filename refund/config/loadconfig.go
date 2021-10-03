package config

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig - load .env file from given path for local, else will be getting from env var
func LoadConfig() {
	// load .env file from given path for local, else will be getting from env var
	if len(os.Getenv("lambda")) == 0 {
		err := godotenv.Load(".test-env")
		if err != nil {
			panic("Error loading .env file")
		}
	}

	DBConfig = os.Getenv("DB_CONFIG")
	RazorpayAuth = os.Getenv("RAZORPAY_AUTH")
	SESAccessKey = os.Getenv("SES_ACCESS_KEY")
	SESSecretKey = os.Getenv("SES_SECRET_KEY")
}
