package main

import (
	CONFIG "email/config"
	CRON "email/cron"
	DB "email/database"
)

// send text emails to phones

func main() {
	CONFIG.LoadConfig()
	DB.ConnectDatabase()
	CRON.Start()
}
