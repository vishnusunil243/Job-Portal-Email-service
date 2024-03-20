package email

import (
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(To string, message string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", To)
	m.SetHeader("Subject", "Subject")
	m.SetBody("text/plain", message)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))

	// Send the email
	if err := dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
