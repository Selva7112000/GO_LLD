package observer

import "fmt"

// BadPublisher directly knows about specific subscribers and calls them explicitly
type BadPublisher struct {
	emailSubscriber *BadEmailSubscriber
	smsSubscriber   *BadSMSSubscriber
}

func (p *BadPublisher) NotifyAll(message string) {
	if p.emailSubscriber != nil {
		p.emailSubscriber.ReceiveEmail(message)
	}
	if p.smsSubscriber != nil {
		p.smsSubscriber.ReceiveSMS(message)
	}
}

// Tightly coupled concrete subscriber
type BadEmailSubscriber struct {
	Email string
}

func (e *BadEmailSubscriber) ReceiveEmail(message string) {
	fmt.Println("Email sent to", e.Email, "with message:", message)
}

// Another tightly coupled subscriber
type BadSMSSubscriber struct {
	Phone string
}

func (s *BadSMSSubscriber) ReceiveSMS(message string) {
	fmt.Println("SMS sent to", s.Phone, "with message:", message)
}

func BadObserver() {
	emailSub := &BadEmailSubscriber{Email: "baduser@example.com"}
	smsSub := &BadSMSSubscriber{Phone: "+911234567890"}

	publisher := &BadPublisher{
		emailSubscriber: emailSub,
		smsSubscriber:   smsSub,
	}

	publisher.NotifyAll("🚫 Tightly Coupled Notification Sent")
}
