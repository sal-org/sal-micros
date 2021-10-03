package main

import (
	CONFIG "newslotday/config"
	CRON "newslotday/cron"
	DB "newslotday/database"
)

// add 1 day slots to counsellors/listeners/therapist based on their weekly schedule

func main() {
	CONFIG.LoadConfig()
	DB.ConnectDatabase()
	CRON.Start()
}
