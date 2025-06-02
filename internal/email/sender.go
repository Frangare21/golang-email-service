package email

import (
	"fmt"
	gomail "gopkg.in/mail.v2"
	"os"
)

func SendMail(to, subject, body string) error {
	from := os.Getenv("EMAIL_FROM")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		getEnvAsInt("SMTP_PORT", 587),
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	return d.DialAndSend(m)
}

// Simple helper for int envs
func getEnvAsInt(name string, fallback int) int {
	val := os.Getenv(name)
	if val == "" {
		return fallback
	}
	var i int
	_, err := fmt.Sscanf(val, "%d", &i)
	if err != nil {
		return fallback
	}
	return i
}
