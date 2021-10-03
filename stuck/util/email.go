package util

import (
	"fmt"

	CONFIG "stuck/config"
	CONSTANT "stuck/constant"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

// SendSESMail - send mail
func SendSESMail(title, body, email string) {
	// start a new aws session
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	// start a new ses session
	svc := ses.New(sess, &aws.Config{
		Credentials: credentials.NewStaticCredentials(CONFIG.SESAccessKey, CONFIG.SESSecretKey, ""),
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
		Source: aws.String(CONSTANT.FromEmailIDPromotional),
	}

	//end email
	output, err := svc.SendEmail(params)
	fmt.Println(err, output.String())
}
