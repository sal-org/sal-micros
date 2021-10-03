package cron

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

// Start - start based on local or lambda
func Start() {
	if len(os.Getenv("lambda")) > 0 {
		// if lambda, set a cloudwatch cron
		lambda.Start(stuckPayments)
	} else {
		stuckPayments()
	}
}
