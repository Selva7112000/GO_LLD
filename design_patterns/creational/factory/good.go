package factory

import "fmt"

type GoodNotifer interface {
	Send(message string)
}

type GEmailNotifier struct{}

func (e *GEmailNotifier) Send(message string) {
	fmt.Println("Sending Email ", message)
}

type GSMSNotifier struct{}

func (s *GSMSNotifier) Send(message string) {
	fmt.Println("Sending SMS ", message)
}

func GetNotifier(notifierType string) GoodNotifer {
	if notifierType == "Email" {
		return &GEmailNotifier{}
	} else if notifierType == "SMS" {
		return &GSMSNotifier{}
	}
	return nil
}

func GoodFactoryPattern() {
	notifier := GetNotifier("Email")
	if notifier != nil {
		notifier.Send("Hello via Email")
	}

	notifier = GetNotifier("SMS")
	if notifier != nil {
		notifier.Send("Hello via SMS")
	}
}
