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
	MediaURL = os.Getenv("MediaURL")
	OneSignalAppIDForClient = os.Getenv("ONESIGNAL_APP_ID_FOR_CLIENT")
	OneSignalApiKeyForClient = os.Getenv("ONESIGNAL_API_KEY_FOR_CLIENT")
	OneSignalAppIDForTherapist = os.Getenv("ONESIGNAL_APP_ID_FOR_THERAPIST")
	OneSignalApiKeyForTherapist = os.Getenv("ONESIGNAL_API_KEY_FOR_THERAPIST")

}
