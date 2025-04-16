package dependency_inversion

import "fmt"

type Notifier interface {
	Send(message string)
}

type GmailNotifier struct{}

func (g *GmailNotifier) Send(message string) {
	fmt.Println("Email sent", message)
}

type SMSNotifier struct{}

func (p *SMSNotifier) Send(message string) {
	fmt.Println("SMS Sent", message)
}

type UserServices struct {
	Notifier
}

func NewUserServices(n Notifier) *UserServices {
	return &UserServices{Notifier: n}
}

func (r *UserServices) RegisterUser(userName string) {
	fmt.Println("User registered", userName)
	r.Send("hi MR " + userName)
}

func GoodDependencyInversion() {
	emailNotifier := &GmailNotifier{}
	sMSNotifier := &SMSNotifier{}

	userService := NewUserServices(emailNotifier)
	userService.RegisterUser("selva")

	userService1 := NewUserServices(sMSNotifier)
	userService1.RegisterUser("kumar")

}
