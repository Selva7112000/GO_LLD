package single_responsibility

import (
	"fmt"
)

type GoodUser struct {
	Name  string
	Email string
}

type UserRepository struct{}

func (user *UserRepository) Save(data GoodUser) {
	fmt.Printf("\nUser has been saved with user name %v", data.Name)
}

type EmailService struct{}

func (user *EmailService) SendWelcomeEmail(data GoodUser) {
	fmt.Printf("\nWelcome to the team buddy %v", data.Name)
}

func GoodSingleResponsibility() {
	data := GoodUser{Name: "Selva", Email: "selva7112000@gmail.com"}

	userRepo := UserRepository{}
	userRepo.Save(data)

	emailService := EmailService{}
	emailService.SendWelcomeEmail(data)
}
