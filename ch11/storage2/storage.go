package storage

import (
	"fmt"
	"log"
	"net/smtp"
)



const (
	sender = "notifications@example.com"
	password = "correcthorsebatterystaple"
	hostname = "smtp.example.com"

	template = "Warning: you are using %d bytes of storage, %d%% of your quota."
)

var bytesInUseMock int64 = 0

var notifyUser = func (username, msg string)  {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname + ":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}

func bytesInUse(username string) int64 { return bytesInUseMock }

func CheckQuota(username string) {
	const quota = 1000000000
	used := bytesInUse(username)
	percent := 100 * used / quota

	if percent < 90 {
		return
	}

	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}