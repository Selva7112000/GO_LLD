package decorator

import "fmt"

type BadCoffee struct {
	HasMilk  bool
	HasSugar bool
}

func (b *BadCoffee) Cost() float64 {
	cost := 5.0
	if b.HasMilk {
		cost += 1.5
	}
	if b.HasSugar {
		cost += 0.5
	}
	return cost
}

func (b *BadCoffee) Description() string {
	desc := "Simple Coffee"
	if b.HasMilk {
		desc += ", Milk"
	}
	if b.HasSugar {
		desc += ", Sugar"
	}
	return desc
}

func BadDecorator() {
	coffee := &BadCoffee{HasMilk: true, HasSugar: true}
	fmt.Println(coffee.Description(), "=", coffee.Cost())
}
