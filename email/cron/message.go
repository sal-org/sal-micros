package cron

import (
	CONSTANT "email/constant"
	DB "email/database"
	UTIL "email/util"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var wg sync.WaitGroup

func sendEmails() {
	if !strings.EqualFold(DB.QueryRowSQL("select email_cron_status from "+CONSTANT.CronStatusTable+" limit 1"), "0") { // run only if no other email cron is active
		return
	}
	defer DB.ExecuteSQL("update " + CONSTANT.CronStatusTable + " set email_cron_status = 0")

	DB.ExecuteSQL("update " + CONSTANT.CronStatusTable + " set email_cron_status = 1")
	startTime := time.Now()
	for {
		if time.Now().Sub(startTime).Minutes() < 10 { // run cron only for 10 min
			// get all emails which are not sent
			emails, ok := DB.SelectProcess("select * from " + CONSTANT.EmailsTable + " where status = " + CONSTANT.EmailInProgress + " limit 100")
			if !ok || len(emails) == 0 { // stop if no emails found
				break
			}

			// send emails
			for _, email := range emails {
				wg.Add(1)
				go sendSESMail(email["title"], email["body"], email["email"], email["type"])
			}

			emailIDs := UTIL.ExtractValuesFromArrayMap(emails, "email_id")

			// update messages to sent status
			DB.ExecuteSQL("update " + CONSTANT.EmailsTable + " set status = " + CONSTANT.EmailSent + " where email_id in ('" + strings.Join(emailIDs, "','") + "')")

			wg.Wait()
		} else {
			break
		}
	}
}

func sendSESMail(title, body, email, emailType string) {
	defer wg.Done()
	fromEmailID := CONSTANT.FromEmailIDTransactional
	if strings.EqualFold(emailType, "2") {
		fromEmailID = CONSTANT.FromEmailIDPromotional
	}
	// start a new aws session
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	// start a new ses session
	svc := ses.New(sess, &aws.Config{
		Credentials: credentials.NewStaticCredentials(CONSTANT.S3AccessKey, CONSTANT.S3SecretKey, ""),
		Region:      aws.String("ap-south-1"),
	})

	params := &ses.SendEmailInput{
		Destination: &ses.Destination{ // Required
			ToAddresses: []*string{
				aws.String(email), // Required
			},
		},
		Message: &ses.Message{ // Required
			Body: &ses.Body{ // Required
				Html: &ses.Content{
					Data:    aws.String(body), // Required
					Charset: aws.String("UTF-8"),
				},
			},
			Subject: &ses.Content{ // Required
				Data:    aws.String(title), // Required
				Charset: aws.String("UTF-8"),
			},
		},
		Source: aws.String(fromEmailID),
	}

	//end email
	output, err := svc.SendEmail(params)
	fmt.Println(err, output.String())
}
