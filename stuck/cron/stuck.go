package cron

import (
	"strconv"
	"strings"
	UTIL "stuck/util"
)

func stuckPayments() {

	// get all payments for last 3 hours
	skip := 0
	for true {
		razorPayResponse := UTIL.GetOrdersFromRazorpay(strconv.Itoa(skip), 72) // check for previous 3 hour
		if razorPayResponse.Count > 0 {
			skip += razorPayResponse.Count
			for _, razorPayItem := range razorPayResponse.Items {
				if strings.EqualFold(razorPayItem.Status, "authorized") {
					// capture payment of authorised one
					UTIL.CapturePayment(razorPayItem.Description, razorPayItem.ID, strconv.Itoa(razorPayItem.Amount))
				}
			}
		} else {
			break
		}
	}

	// check if still any payments are in authorised state and send mail
	skip = 0
	authorizedPayments := false
	for true {
		razorPayResponse := UTIL.GetOrdersFromRazorpay(strconv.Itoa(skip), 72) // 3 days
		if razorPayResponse.Count > 0 {
			skip += razorPayResponse.Count
			for _, razorPayItem := range razorPayResponse.Items {
				if strings.EqualFold(razorPayItem.Status, "authorized") {
					authorizedPayments = true
					break
				}
			}
		} else {
			break
		}
		if authorizedPayments {
			break
		}
	}
	if authorizedPayments {
		UTIL.SendSESMail("Stuck Orders", "There are some authorised payments in razorpay", "dravid.rahul1526@gmail.com")
	}
}
