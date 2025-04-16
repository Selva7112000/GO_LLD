package single_responsibility

import (
	"fmt"
)

type BadUser struct {
	Name  string
	Email string
}

func (u *BadUser) Save() {
	fmt.Println("\nSaving the data to the database", u.Name)
}

func (u *BadUser) SendWelcomeEmail() {
	fmt.Printf("\nHi %v welcome to the team", u.Name)
}

func BadSingleResponsibility() {
	data := BadUser{Name: "Selva", Email: "selva7112000@gmail.com"}
	data.Save()
	data.SendWelcomeEmail()
}
