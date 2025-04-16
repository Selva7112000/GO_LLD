package observer

import "fmt"

// Observer interface
type Observer interface {
	Update(message string)
}

// Subject (Publisher) interface
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyAll(message string)
}

// Concrete Subject
type NewsPublisher struct {
	subscribers []Observer
}

func (n *NewsPublisher) Register(o Observer) {
	n.subscribers = append(n.subscribers, o)
}

func (n *NewsPublisher) Unregister(o Observer) {
	for i, sub := range n.subscribers {
		if sub == o {
			n.subscribers = append(n.subscribers[:i], n.subscribers[i+1:]...)
			break
		}
	}
}

func (n *NewsPublisher) NotifyAll(message string) {
	for _, sub := range n.subscribers {
		sub.Update(message)
	}
}

// Concrete Observer 1
type EmailSubscriber struct {
	Email string
}

func (e *EmailSubscriber) Update(message string) {
	fmt.Println("Email to", e.Email, "received:", message)
}

// Concrete Observer 2
type SMSSubscriber struct {
	Phone string
}

func (s *SMSSubscriber) Update(message string) {
	fmt.Println("SMS to", s.Phone, "received:", message)
}

func GoodObserver() {
	publisher := &NewsPublisher{}

	emailSub := &EmailSubscriber{Email: "user@example.com"}
	smsSub := &SMSSubscriber{Phone: "+911234567890"}

	publisher.Register(emailSub)
	publisher.Register(smsSub)

	publisher.NotifyAll("🚀 New Article Published: Go Design Patterns")

	publisher.Unregister(emailSub)

	publisher.NotifyAll("📢 Update: Observer Pattern Explained in Depth!")
}
