package cron

import (
	"fmt"
	"io/ioutil"
	CONFIG "message/config"
	CONSTANT "message/constant"
	DB "message/database"
	UTIL "message/util"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sendMessages() {
	// if !strings.EqualFold(DB.QueryRowSQL("select message_cron_status from "+CONSTANT.CronStatusTable+" limit 1"), "0") { // run only if no other message cron is active
	// 	return
	// }
	// defer DB.ExecuteSQL("update " + CONSTANT.CronStatusTable + " set message_cron_status = 0")

	// DB.ExecuteSQL("update " + CONSTANT.CronStatusTable + " set message_cron_status = 1")
	startTime := time.Now()
	for {
		if time.Now().Sub(startTime).Minutes() < 10 { // run cron only for 10 min
			// get all messages which are not sent
			messages, ok := DB.SelectProcess("select * from " + CONSTANT.MessagesTable + " where send_at <= now()+interval 330 minute and status = " + CONSTANT.MessageInProgress + " and  message_status = " + CONSTANT.MessageInProgress + " limit 100")
			if !ok || len(messages) == 0 { // stop if no messages found
				break
			}

			// send messages
			for _, message := range messages {
				wg.Add(1)
				go sendMessage(message["text"], message["route"], message["phone"])
			}

			messageIDs := UTIL.ExtractValuesFromArrayMap(messages, "message_id")

			// update messages to sent status
			DB.ExecuteSQL("update " + CONSTANT.MessagesTable + " set status = " + CONSTANT.MessageSent + ", message_status = " + CONSTANT.MessageSent + " where message_id in ('" + strings.Join(messageIDs, "','") + "')")

			wg.Wait()
		} else {
			break
		}
	}
}

func sendMessage(text, route, phone string) {
	defer wg.Done()
	resp, err := http.Get(buildMessageURL(text, route, phone))
	if err != nil {
		fmt.Println("sendMessage", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("sendMessage", err)
		return
	}

	fmt.Println("sendMessage", string(body))
}

func buildMessageURL(text, route, phone string) string {
	u, _ := url.Parse(CONSTANT.CorefactorsSendSMSEndpoint)

	v := url.Values{}
	v.Add("text", text)
	v.Add("key", CONFIG.CorefactorsAPIKey)
	v.Add("to", phone)
	v.Add("route", route)
	v.Add("from", CONSTANT.TextMessageFrom)

	u.RawQuery = v.Encode()

	return u.String()
}
