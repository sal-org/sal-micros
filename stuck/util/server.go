package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	CONFIG "stuck/config"
	CONSTANT "stuck/constant"
	MODEL "stuck/model"
)

// CapturePayment - capture payment by hitting API endppoints - if success then payment captured and order created
func CapturePayment(ordUID string, transactionID string, amount string) {
	order := map[string]string{
		"order_id":       ordUID,
		"payment_method": "RAZORPAY",
		"payment_id":     transactionID,
	}

	orderBytes, _ := json.Marshal(order)

	for _, paymentURL := range CONSTANT.PaymentURLs {
		req, err := http.NewRequest("POST", CONFIG.APIURL+paymentURL, bytes.NewBuffer(orderBytes))
		if err != nil {
			fmt.Println("CapturePayment", err)
			return
		}

		req.Header.Add("Content-Type", "application/json")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("CapturePayment", err)
			return
		}

		serverResponse := MODEL.ServerResponse{}
		err = json.Unmarshal(body, &serverResponse)
		if err != nil {
			fmt.Println("CapturePayment", err)
			return
		}

		if strings.EqualFold(serverResponse.Meta.Status, CONSTANT.StatusCodeOk) || strings.EqualFold(serverResponse.Meta.Status, CONSTANT.StatusCodeCreated) {
			break
		}
	}
}
