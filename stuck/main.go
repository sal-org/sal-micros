package main

import (
	CONFIG "stuck/config"
	CRON "stuck/cron"
)

// get all stuck/authorised payments from razorpay and capture them, create order

func main() {
	CONFIG.LoadConfig()
	CRON.Start()
}
