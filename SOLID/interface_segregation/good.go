package interface_segregation

import "fmt"

type Workable interface {
	Work()
}

type Eatable interface {
	Eat()
}

type HumanBeing struct{}

func (h *HumanBeing) Work() {
	fmt.Println("Human being is working")
}

func (h *HumanBeing) Eat() {
	fmt.Println("human being is eating")
}

type Robo struct{}

func (r *Robo) Work() {
	fmt.Println("Robo is working")
}

func GoodInterfaceSegregation() {
	var w Workable
	w = &HumanBeing{}
	w.Work()

	w = &Robo{}
	w.Work()

	var e Eatable

	e = &HumanBeing{}
	e.Eat()

	// e = Robot{}
}
