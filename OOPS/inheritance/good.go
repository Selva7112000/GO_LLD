package inheritance

import (
	"fmt"
)

type GoodVehicle struct {
	Brand string
	Model string
	Year  string
}

func (v *GoodVehicle) Start() {
	fmt.Println("vehicle is starting")
}

func (v *GoodVehicle) Stop() {
	fmt.Println("vehicle has stopped")
}

type Car struct {
	GoodVehicle
	NumberOfWheels int
	NumberOfDoors  int
}

type Bike struct {
	GoodVehicle
	NumberOfWheels int
}

func GoodInheritance() {
	vehicle1 := Bike{GoodVehicle: GoodVehicle{Brand: "Ford", Model: "T90", Year: "2020"}, NumberOfWheels: 4}

	vehicle2 := Car{GoodVehicle: GoodVehicle{Brand: "BMW", Model: "S1000RR", Year: "2020"}, NumberOfWheels: 4, NumberOfDoors: 4}

	fmt.Println("\nBike Details:", vehicle2.Brand, vehicle2.Model, vehicle2.Year)
	vehicle1.Start()
	vehicle1.Stop()

	fmt.Println("Car Details:", vehicle1.Brand, vehicle1.Model, vehicle1.Year)
	vehicle2.Start()
	vehicle2.Stop()
}
