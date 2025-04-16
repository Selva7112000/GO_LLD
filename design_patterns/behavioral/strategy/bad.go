package strategy

import "fmt"

// No strategy interface; hardcoded types
type BadCart struct {
	UserType string
}

func (c *BadCart) Checkout(amount float64) float64 {
	if c.UserType == "regular" {
		return amount * 0.9
	} else if c.UserType == "premium" {
		return amount * 0.8
	} else if c.UserType == "vip" {
		return amount * 0.7
	}
	return amount
}

func BadStrategy() {
	cart := &BadCart{UserType: "regular"}
	fmt.Println("Regular:", cart.Checkout(1000))

	cart.UserType = "premium"
	fmt.Println("Premium:", cart.Checkout(1000))

	cart.UserType = "vip"
	fmt.Println("VIP:", cart.Checkout(1000))
}
