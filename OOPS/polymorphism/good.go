package polymorphism

import "fmt"

type Vehicle interface {
	Start()
	Stop()
	GetBrand() string
	GetModel() string
}

type BaseVehicle struct {
	Brand string
	Model string
	Year  int
}

func (b *BaseVehicle) GetBrand() string {
	return b.Brand
}

func (b *BaseVehicle) GetModel() string {
	return b.Model
}

type Car struct {
	BaseVehicle
	NumberOfDoors  int
	NumberOfWheels int
}

func (c *Car) Start() {
	fmt.Println("car has been started")
}

func (c *Car) Stop() {
	fmt.Println("car has been stopped")
}

type Bike struct {
	BaseVehicle
	NumberOfWheels int
}

func (b *Bike) Start() {
	fmt.Println("bike has been started")
}

func (b *Bike) Stop() {
	fmt.Println("bike has been stopped")
}

type Plane struct {
	BaseVehicle
	NumberOfDoors int
}

func (p *Plane) Start() {
	fmt.Println("plane has been started")
}

func (p *Plane) Stop() {
	fmt.Println("plane has been stopped")
}

func GoodPolymorphism() {
	vehicles := []Vehicle{
		&Car{BaseVehicle: BaseVehicle{Brand: "Ford", Model: "Focus", Year: 2008}, NumberOfDoors: 5},
		&Bike{BaseVehicle: BaseVehicle{Brand: "Honda", Model: "Scoopy", Year: 2018}},
		&Plane{BaseVehicle: BaseVehicle{Brand: "Boeing", Model: "747", Year: 2015}, NumberOfDoors: 16},
	}

	for _, vehicle := range vehicles {
		fmt.Printf("Inspecting %s %s (%T)\n", vehicle.GetBrand(), vehicle.GetModel(), vehicle)
		vehicle.Start()
		vehicle.Stop()
	}
}
