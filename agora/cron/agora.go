package cron

import (
	CONSTANT "agora/constant"
	DB "agora/database"
	UTIL "agora/util"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
)

func generateAgoraMeetings() {
	if !strings.EqualFold(DB.QueryRowSQL("select agora_cron_status from "+CONSTANT.CronStatusTable+" limit 1"), "0") { // run only if no other email cron is active
		return
	}
	defer DB.ExecuteSQL("update " + CONSTANT.CronStatusTable + " set agora_cron_status = 0")

	DB.ExecuteSQL("update " + CONSTANT.CronStatusTable + " set agora_cron_status = 1")
	startTime := time.Now()
	for {
		if time.Now().Sub(startTime).Minutes() < 10 { // run cron only for 10 min
			// get all appointments which are yet to start within 24 hours and agora meeting link not generated
			appointments, ok := DB.SelectProcess("select * from " + CONSTANT.AppointmentsTable + " where status = " + CONSTANT.AppointmentToBeStarted + " and meeting_link is null")
			if !ok || len(appointments) == 0 { // stop if no appointments found
				break
			}

			// generate agaora meeting links
			for _, appointment := range appointments {
				if UTIL.BuildDateTime(appointment["date"], appointment["time"]).Sub(time.Now()).Hours() < 24 { // since meeting token will expire in 24 hours
					meetingToken, err := generateRtcToken("SAL", appointment["counsellor_id"], "userAccount", rtctokenbuilder.RolePublisher, uint32(UTIL.GetCurrentTime().Unix()))
					if err != nil {
						fmt.Println("generateAgoraMeetings", err)
						continue
					}
					DB.ExecuteSQL("update " + CONSTANT.AppointmentsTable + " set meeting_link = '" + meetingToken + "' where appointment_id = '" + appointment["appointment_id"] + "'")
				}
			}
		} else {
			break
		}
	}
}

func generateRtcToken(channelName, uidStr, tokentype string, role rtctokenbuilder.Role, expireTimestamp uint32) (rtcToken string, err error) {

	if tokentype == "userAccount" {
		rtcToken, err = rtctokenbuilder.BuildTokenWithUserAccount(CONSTANT.AgoraAppID, CONSTANT.AgoraAppCertificate, channelName, uidStr, role, expireTimestamp)
		return rtcToken, err

	} else if tokentype == "uid" {
		uid64, parseErr := strconv.ParseUint(uidStr, 10, 64)
		// check if conversion fails
		if parseErr != nil {
			err = fmt.Errorf("failed to parse uidStr: %s, to uint causing error: %s", uidStr, parseErr)
			return "", err
		}

		uid := uint32(uid64) // convert uid from uint64 to uint 32
		rtcToken, err = rtctokenbuilder.BuildTokenWithUID(CONSTANT.AgoraAppID, CONSTANT.AgoraAppCertificate, channelName, uid, role, expireTimestamp)
		return rtcToken, err

	} else {
		err = fmt.Errorf("failed to generate RTC token for Unknown Tokentype: %s", tokentype)
		return "", err
	}
}
