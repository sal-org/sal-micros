package main

import (
	CONFIG "agora/config"
	CRON "agora/cron"
	DB "agora/database"
)

// generate agora meeting token ids (since expiry is 24 hours, doing here and not while creating appointments)

func main() {
	CONFIG.LoadConfig()
	DB.ConnectDatabase()
	CRON.Start()
}
