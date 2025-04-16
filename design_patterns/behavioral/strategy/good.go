package strategy

import "fmt"

// Strategy Interface
type DiscountStrategy interface {
	Calculate(amount float64) float64
}

// Concrete Strategies
type RegularDiscount struct{}

func (r *RegularDiscount) Calculate(amount float64) float64 {
	return amount * 0.9 // 10% off
}

type PremiumDiscount struct{}

func (p *PremiumDiscount) Calculate(amount float64) float64 {
	return amount * 0.8 // 20% off
}

type VIPDiscount struct{}

func (v *VIPDiscount) Calculate(amount float64) float64 {
	return amount * 0.7 // 30% off
}

// Context
type Cart struct {
	Strategy DiscountStrategy
}

func (c *Cart) Checkout(amount float64) float64 {
	return c.Strategy.Calculate(amount)
}

func GoodStrategy() {
	cart := &Cart{Strategy: &RegularDiscount{}}
	fmt.Println("Regular:", cart.Checkout(1000))

	cart.Strategy = &PremiumDiscount{}
	fmt.Println("Premium:", cart.Checkout(1000))

	cart.Strategy = &VIPDiscount{}
	fmt.Println("VIP:", cart.Checkout(1000))
}
