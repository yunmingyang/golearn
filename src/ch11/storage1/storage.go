package storage

import (
	"fmt"
	"log"
	"net/smtp"
)



func bytesInUse(username string) int64 { return 0 }

const (
	sender = "notifications@example.com"
	password = "correcthorsebatterystaple"
	hostname = "smtp.example.com"

	template = "Warning: you are using %d bytes of storage, %d%% of your quota."
)

func CheckQuota(username string) {
	const quota = 1000000000
	used := bytesInUse(username)
	percent := 100 * used / quota

	if percent < 90 {
		return
	}

	msg := fmt.Sprintf(template, used, percent)
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname + ":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %v", username, err)
	}
}