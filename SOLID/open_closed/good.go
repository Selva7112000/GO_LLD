package open_closed

import "fmt"

type GoodInvoice interface {
	Calculate(float64) float64
}

type RegularDiscount struct{}

func (r *RegularDiscount) Calculate(amount float64) float64 {
	return amount * 0.1
}

type PremiumDiscount struct{}

func (r *PremiumDiscount) Calculate(amount float64) float64 {
	return amount * 0.2
}

type VipDiscount struct{}

func (v *VipDiscount) Calculate(amount float64) float64 {
	return amount * 0.3
}

func GoodOpenClosed() {
	amount := 1000.0

	var strategy GoodInvoice
	strategy = &RegularDiscount{}
	fmt.Println("Regular Discount", strategy.Calculate(amount))

	strategy = &PremiumDiscount{}
	fmt.Println("Premium Discount", strategy.Calculate(amount))

	strategy = &VipDiscount{}
	fmt.Println("Premium Discount", strategy.Calculate(amount))
}
