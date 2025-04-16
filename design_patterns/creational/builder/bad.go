package builder

import "fmt"

type BadComputer struct {
	CPU string
	GPU string
	RAM string
	SSD string
}

func NewBadComputer(cpu, gpu, ram, ssd string) *BadComputer {
	return &BadComputer{
		CPU: cpu,
		GPU: gpu,
		RAM: ram,
		SSD: ssd,
	}
}

func BadBuilder() {
	// Difficult to remember parameter order
	pc := NewBadComputer("Intel i5", "", "8GB", "512GB")
	fmt.Println("Bad Computer Specs:")
	fmt.Println("CPU:", pc.CPU)
	fmt.Println("GPU:", pc.GPU)
	fmt.Println("RAM:", pc.RAM)
	fmt.Println("SSD:", pc.SSD)
}
