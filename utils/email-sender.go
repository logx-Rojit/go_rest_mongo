package utils

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	HOST := os.Getenv("SMTP_HOST")
	PORT := os.Getenv("SMTP_PORT")
	EMAIL := os.Getenv("FROM")
	PASSWORD := os.Getenv("PASSWORD")
	SENDER_EMAIL := os.Getenv("SENDER_EMAIL")

	m := gomail.NewMessage()

	m.SetHeader("From", SENDER_EMAIL)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	port, err := strconv.Atoi(PORT)
	if err != nil {
		return nil
	}
	d := gomail.NewDialer(HOST, port, EMAIL, PASSWORD)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error is ", err)
		return err
	}
	return nil
}
