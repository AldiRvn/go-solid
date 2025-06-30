package ocp

import (
	"fmt"
	"log"
)

type NotificationService interface {
	SendNotification(message string) error
}

type NotificationSender struct {
	NotificationService NotificationService
}

func (ns *NotificationSender) SendNotification(message string) (err error) {
	log.Println("Sending Notification...")
	return ns.NotificationService.SendNotification(message)
}

// * -------------------- Email Notification Implementation ------------------- */
type EmailNotification struct{}

func (en *EmailNotification) SendNotification(message string) (err error) {
	fmt.Println("Email Notification:", message)
	return
}

// * --------------------- SMS Notification Implementation -------------------- */
type SmsNotification struct{}

func (sm *SmsNotification) SendNotification(message string) (err error) {
	fmt.Println("SMS Notification:", message)
	return
}
