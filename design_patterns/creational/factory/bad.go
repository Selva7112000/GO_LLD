package factory

import "fmt"

type BadEmailNotifier struct{}

func (e *BadEmailNotifier) Send(message string) {
	fmt.Println("Sending Email:", message)
}

type BadSMSNotifier struct{}

func (s *BadSMSNotifier) Send(message string) {
	fmt.Println("Sending SMS:", message)
}

func BadFactoryPattern() {
	notificationType := "sms"

	var notifier interface{}

	if notificationType == "email" {
		notifier = GEmailNotifier{}
	} else if notificationType == "sms" {
		notifier = BadSMSNotifier{}
	} else {
		fmt.Println("Unknown notifier type")
		return
	}

	switch n := notifier.(type) {
	case GEmailNotifier:
		n.Send("Email message!")
	case BadSMSNotifier:
		n.Send("SMS message!")
	}
}
