package main

import (
	CONFIG "refund/config"
	CRON "refund/cron"
	DB "refund/database"
)

// get all refund/authorised payments from razorpay and capture them, create order

func main() {
	CONFIG.LoadConfig()
	DB.ConnectDatabase()
	CRON.Start()
}
