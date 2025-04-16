package interface_segregation

import "fmt"

type BadWorker interface {
	Work()
	Eat()
}

type Human struct{}

func (h *Human) Work() {
	fmt.Println("Human is working")
}

func (h *Human) Eat() {
	fmt.Println("Human is eating")
}

type Robot struct{}

func (r *Robot) Work() {
	fmt.Println("Robot is working")
}

func (r *Robot) Eat() {
	panic("Robots don't eat")
}

func BadInterfaceSegregation() {
	var data BadWorker

	data = &Human{}
	data.Eat()
	data.Work()

	data = &Robot{}
	data.Eat()
	data.Work()
}
