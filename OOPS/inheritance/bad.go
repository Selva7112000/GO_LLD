package inheritance

// import "fmt"

// // BadVehicle struct (poorly designed base class)
// type BadVehicle struct {
// 	Brand string
// 	Model string
// 	Year  string
// }

// // Start method (generic for all vehicles)
// func (v *BadVehicle) Start() {
// 	fmt.Println("Generic vehicle is starting")
// }

// // Stop method
// func (v *BadVehicle) Stop() {
// 	fmt.Println("Generic vehicle has stopped")
// }

// // Car struct (inherits BadVehicle but adds no meaningful behavior)
// type Car struct {
// 	BadVehicle
// 	NumberOfWheels int
// 	NumberOfDoors  int
// }

// // Bike struct (inherits BadVehicle but adds no meaningful behavior)
// type Bike struct {
// 	BadVehicle
// 	NumberOfWheels int
// }

// // Bad Design: Mixing `Car` and `Bike` methods inside a common struct
// type MixedVehicle struct {
// 	Car
// 	Bike
// }

// func BadInheritance() {
// 	// Issue 1: Exposing fields directly (no encapsulation)
// 	vehicle1 := Car{
// 		BadVehicle:     BadVehicle{Brand: "Ford", Model: "T90", Year: "2020"},
// 		NumberOfWheels: 4,
// 		NumberOfDoors:  4,
// 	}

// 	vehicle2 := Bike{
// 		BadVehicle:     BadVehicle{Brand: "BMW", Model: "S1000RR", Year: "2020"},
// 		NumberOfWheels: 2,
// 	}

// 	// Issue 2: Poorly structured multiple inheritance attempt (MixedVehicle)
// 	vehicle3 := MixedVehicle{
// 		Car:  Car{BadVehicle: BadVehicle{Brand: "Toyota", Model: "Camry", Year: "2022"}, NumberOfWheels: 4, NumberOfDoors: 4},
// 		Bike: Bike{BadVehicle: BadVehicle{Brand: "Yamaha", Model: "YZF-R1", Year: "2023"}, NumberOfWheels: 2},
// 	}

// 	// Issue 3: Car and Bike are treated the same (No abstraction)
// 	fmt.Println("\nCar Details:", vehicle1.Brand, vehicle1.Model, vehicle1.Year)
// 	vehicle1.Start()
// 	vehicle1.Stop()

// 	fmt.Println("\nBike Details:", vehicle2.Brand, vehicle2.Model, vehicle2.Year)
// 	vehicle2.Start()
// 	vehicle2.Stop()

// 	// Issue 4: MixedVehicle causes ambiguity in method calls
// 	fmt.Println("\nMixedVehicle (Car & Bike) Details:", vehicle3.Car.Brand, vehicle3.Car.Model, vehicle3.Car.Year)
// 	vehicle3.Start() // Which `Start()` gets called? (Car or Bike?)
// 	vehicle3.Stop()  // Confusing behavior due to embedding both Car and Bike
// }
