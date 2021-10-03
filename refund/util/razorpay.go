package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	CONFIG "refund/config"
	CONSTANT "refund/constant"
	MODEL "refund/model"
	"strings"
)

// RefundRazorpayPayment - refund amount from razorpay transaction
func RefundRazorpayPayment(transactionID string, amount float64) (string, bool) {
	refundBodyBytes, _ := json.Marshal(map[string]interface{}{
		"amount": amount,
	})

	req, err := http.NewRequest("POST", CONSTANT.RazorPayURL+"/payments/"+transactionID+"/refund", bytes.NewBuffer(refundBodyBytes))
	if err != nil {
		fmt.Println("RefundRazorpayPayment", err)
		return err.Error(), false
	}
	req.Header.Add("Authorization", CONFIG.RazorpayAuth)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("RefundRazorpayPayment", err)
		return err.Error(), false
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("RefundRazorpayPayment", err)
		return err.Error(), false
	}

	if strings.Contains(string(body), "error") {
		razorpayRefundErrorResponse := MODEL.RazorpayRefundErrorResponse{}
		err = json.Unmarshal(body, &razorpayRefundErrorResponse)
		if err != nil {
			fmt.Println("RefundRazorpayPayment", err)
			return razorpayRefundErrorResponse.Error.Description, false
		}
	}

	return "", true
}
