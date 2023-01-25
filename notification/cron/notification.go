package cron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	CONFIG "notification/config"
	CONSTANT "notification/constant"
	DB "notification/database"
	MODEL "notification/model"
	UTIL "notification/util"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sendNotifications() {
	// if !strings.EqualFold(DB.QueryRowSQL("select notification_cron_status from "+CONSTANT.CronStatusTable+" limit 1"), "0") { // run only if no other notification cron is active
	// 	return
	// }
	// defer DB.ExecuteSQL("update " + CONSTANT.CronStatusTable + " set notification_cron_status = 0")

	// DB.ExecuteSQL("update " + CONSTANT.CronStatusTable + " set notification_cron_status = 1")
	startTime := time.Now()
	for {
		if time.Now().Sub(startTime).Minutes() < 10 { // run cron only for 10 min
			// get all notifications which are not sent
			notifications, ok := DB.SelectProcess("select * from " + CONSTANT.NotificationsTable + " where send_at <= now()+interval 330 minute and notification_status = " + CONSTANT.NotificationInProgress + " and onesignal_id != '' limit 100")
			if !ok || len(notifications) == 0 { // stop if no notifications found
				break
			}

			// send notifications
			for _, notification := range notifications {
				wg.Add(1)
				go sendNotification(notification["title"], notification["body"], notification["onesignal_id"], notification["type"])
			}

			notificationIDs := UTIL.ExtractValuesFromArrayMap(notifications, "notification_id")

			// update notifications to sent status
			DB.ExecuteSQL("update " + CONSTANT.NotificationsTable + " set notification_status = " + CONSTANT.NotificationSent + " where notification_id in ('" + strings.Join(notificationIDs, "','") + "')")

			wg.Wait()
		} else {
			break
		}
	}
}

func sendNotification(heading, content, notificationID, personType string) {
	defer wg.Done()

	var app_id string

	if personType == "3" {
		app_id = CONFIG.OneSignalAppIDForClient
	} else {
		app_id = CONFIG.OneSignalAppIDForTherapist
	}
	// sent to onesignal
	data := MODEL.OneSignalNotificationData{
		AppID:            app_id,
		Headings:         map[string]string{"en": heading},
		Contents:         map[string]string{"en": content},
		IncludePlayerIDs: []string{notificationID},
		Data:             map[string]string{},
	}
	byteData, _ := json.Marshal(data)
	resp, err := http.Post("https://onesignal.com/api/v1/notifications", "application/json", bytes.NewBuffer(byteData))
	if err != nil {
		fmt.Println("sendNotification", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("sendNotification", err)
		return
	}

	fmt.Println(data, string(body))
}
