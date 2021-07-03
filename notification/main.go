package main

import (
	CONFIG "notification/config"
	CRON "notification/cron"
	DB "notification/database"
)

// send notification to devices through onesignal

func main() {
	CONFIG.LoadConfig()
	DB.ConnectDatabase()
	CRON.Start()
}
