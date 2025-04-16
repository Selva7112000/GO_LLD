package dependency_inversion

import "fmt"

type EmailNotifier struct{}

func (e EmailNotifier) SendEmail(message string) {
	fmt.Println("Email Sent", message)
}

type UserService struct {
	EmailNotifier
}

func (u *UserService) RegisterUser(userName string) {
	fmt.Println("User Registered", userName)
	u.SendEmail("Hi you're welcome " + userName)
}

func BadDependencyInversion() {
	email := EmailNotifier{}
	userService := UserService{EmailNotifier: email}
	userService.RegisterUser("selva")

}
