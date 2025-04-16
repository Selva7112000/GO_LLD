package polymorphism

import (
	"fmt"
	"log"
)

type MotorBike struct {
	Brand string
	Model string
	Year  int
}

func (m *MotorBike) start() {
	fmt.Println("Bike has been started")
}

func (m *MotorBike) Stop() {
	fmt.Println("Bike has been stopped")
}

type BadCar struct {
	Brand         string
	Model         string
	Year          int
	NumberOfDoors int
}

func (c *BadCar) Start() {
	fmt.Println("car has been started")
}

func (c *BadCar) Stop() {
	fmt.Println("Car has been stopeed")
}

func BadPolymorphism() {
	vehicles := []interface{}{
		&BadCar{Brand: "Ford", Model: "Focus", Year: 2008, NumberOfDoors: 5},
		&MotorBike{Brand: "Honda", Model: "Scoopy", Year: 2018},
	}

	for _, vehicle := range vehicles {
		switch v := vehicle.(type) {
		case *MotorBike:
			fmt.Printf("Inspecting %s %s (Bike)\n", v.Brand, v.Model)
			v.start()
			v.Stop()
		case *BadCar:
			fmt.Printf("Inspecting %s %s (Car)\n", v.Brand, v.Model)
			v.Start()
			v.Stop()
		default:
			log.Fatal("Error: Object is not a valid vehicle")
		}
	}
}
