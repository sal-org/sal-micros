package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

// send notification to devices through onesignal

type NotificationData struct {
	AppID            string            `json:"app_id"`
	Headings         map[string]string `json:"headings"`
	Contents         map[string]string `json:"contents"`
	IncludePlayerIDs []string          `json:"include_player_ids"`
	Data             map[string]string `json:"data"`
}

func main() {
	lambda.Start(Handler)
}

// Handler is started by lambda
func Handler(data NotificationData) {
	data.AppID = os.Getenv("app_id")
	sendNotification(data)
}

func sendNotification(data NotificationData) {
	byteData, _ := json.Marshal(data)
	resp, err := http.Post("https://onesignal.com/api/v1/notifications", "application/json", bytes.NewBuffer(byteData))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data, string(body))
}
