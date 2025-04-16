package decorator

import "fmt"

// Component interface
type Coffee interface {
	Cost() float64
	Description() string
}

// Concrete component
type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() float64 {
	return 5.0
}

func (s *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// Base decorator
type CoffeeDecorator struct {
	Coffee Coffee
}

func (d *CoffeeDecorator) Cost() float64 {
	return d.Coffee.Cost()
}

func (d *CoffeeDecorator) Description() string {
	return d.Coffee.Description()
}

// Concrete decorator: Milk
type MilkDecorator struct {
	Coffee
}

func (m *MilkDecorator) Cost() float64 {
	return m.Coffee.Cost() + 1.5
}

func (m *MilkDecorator) Description() string {
	return m.Coffee.Description() + ", Milk"
}

// Concrete decorator: Sugar
type SugarDecorator struct {
	Coffee
}

func (s *SugarDecorator) Cost() float64 {
	return s.Coffee.Cost() + 0.5
}

func (s *SugarDecorator) Description() string {
	return s.Coffee.Description() + ", Sugar"
}

// Test function
func GoodDecorator() {
	var coffee Coffee = &SimpleCoffee{}
	fmt.Println(coffee.Description(), "=", coffee.Cost())

	coffee = &MilkDecorator{Coffee: coffee}
	coffee = &SugarDecorator{Coffee: coffee}

	fmt.Println(coffee.Description(), "=", coffee.Cost())
}
