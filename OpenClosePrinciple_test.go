package main

import (
	"fmt"
	"testing"

	ocp "go-solid/OCP"
)

func Test_OpenClosePrinciple(t *testing.T) {
	fmt.Println("OpenClosePrinciple()")

	emailNotification := &ocp.EmailNotification{}
	notificationSender := &ocp.NotificationSender{
		NotificationService: emailNotification,
	}
	notificationSender.SendNotification("Hello World.")

	smsNotification := &ocp.SmsNotification{}
	notificationSender.NotificationService = smsNotification
	notificationSender.SendNotification("Hello World.")
}
