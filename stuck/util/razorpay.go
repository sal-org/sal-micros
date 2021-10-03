package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	CONFIG "stuck/config"
	CONSTANT "stuck/constant"
	MODEL "stuck/model"
	"time"
)

// GetOrdersFromRazorpay - get payments from razorpay
func GetOrdersFromRazorpay(skip string, hours int) MODEL.RazorPayResponse {
	req, err := http.NewRequest("GET", CONSTANT.RazorPayURL, nil)
	if err != nil {
		fmt.Println("GetOrdersFromRazorpay", err)
		return MODEL.RazorPayResponse{}
	}

	req.Header.Add("Authorization", CONFIG.RazorpayAuth)
	q := req.URL.Query()
	q.Add("count", "100")
	q.Add("skip", skip)
	q.Add("from", strconv.FormatInt(time.Now().Add(-time.Duration(hours)*time.Hour).Unix(), 10))
	q.Add("to", strconv.FormatInt(time.Now().Unix(), 10))
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("GetOrdersFromRazorpay", err)
		return MODEL.RazorPayResponse{}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("GetOrdersFromRazorpay", err)
		return MODEL.RazorPayResponse{}
	}

	razorPayResponse := MODEL.RazorPayResponse{}
	err = json.Unmarshal(body, &razorPayResponse)
	if err != nil {
		fmt.Println("GetOrdersFromRazorpay", err)
		return MODEL.RazorPayResponse{}
	}

	return razorPayResponse
}
