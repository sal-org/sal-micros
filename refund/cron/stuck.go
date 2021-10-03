package cron

import (
	CONSTANT "refund/constant"
	DB "refund/database"
	UTIL "refund/util"
	"strconv"
	"strings"
)

func refundPayments() {

	// get all refunds
	refunds, ok := DB.SelectProcess("select * from " + CONSTANT.RefundsTable + " where status = " + CONSTANT.RefundInProgress)
	if !ok {
		return
	}

	errors := []string{}

	for _, refund := range refunds {
		refundedAmount, _ := strconv.ParseFloat(refund["refunded_amount"], 64)
		err, ok := UTIL.RefundRazorpayPayment(refund["payment_id"], refundedAmount)
		if !ok {
			errors = append(errors, err)
		} else {
			DB.UpdateSQL(CONSTANT.RefundsTable, map[string]string{"refund_id": refund["refund_id"]}, map[string]string{"status": CONSTANT.RefundCompleted, "modified_at": UTIL.GetCurrentTime().String()})
		}
	}

	if len(errors) > 0 {
		UTIL.SendSESMail("Refund Errors", "These are some errors while refunding\n\n"+strings.Join(errors, "\n"), "dravid.rahul1526@gmail.com")
	}
}
