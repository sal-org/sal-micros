package main

import (
	CONFIG "message/config"
	CRON "message/cron"
	DB "message/database"
)

// send text messages to phones

func main() {
	CONFIG.LoadConfig()
	DB.ConnectDatabase()
	CRON.Start()
}
