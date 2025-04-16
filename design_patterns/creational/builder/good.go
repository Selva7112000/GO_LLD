package builder

import "fmt"

type Computer struct {
	CPU string
	GPU string
	RAM string
	SSD string
}

func (c *Computer) Specs() {
	fmt.Println("Computer Specs:")
	fmt.Println("CPU:", c.CPU)
	fmt.Println("GPU:", c.GPU)
	fmt.Println("RAM:", c.RAM)
	fmt.Println("SSD:", c.SSD)
}

type ComputerBuilder interface {
	SetCPU(string) ComputerBuilder
	SetGPU(string) ComputerBuilder
	SetRAM(string) ComputerBuilder
	SetSSD(string) ComputerBuilder
	Build() *Computer
}

type GamingBuilder struct {
	computer *Computer
}

func NewGamingBuilder() *GamingBuilder {
	return &GamingBuilder{computer: &Computer{}}
}

func (b *GamingBuilder) SetCPU(cpu string) ComputerBuilder {
	b.computer.CPU = cpu
	return b
}
func (b *GamingBuilder) SetGPU(gpu string) ComputerBuilder {
	b.computer.GPU = gpu
	return b
}
func (b *GamingBuilder) SetRAM(ram string) ComputerBuilder {
	b.computer.RAM = ram
	return b
}
func (b *GamingBuilder) SetSSD(ssd string) ComputerBuilder {
	b.computer.SSD = ssd
	return b
}
func (b *GamingBuilder) Build() *Computer {
	return b.computer
}

func GoodBuilder() {
	builder := NewGamingBuilder()
	pc := builder.SetCPU("Intel i9").
		SetGPU("RTX 4090").
		SetRAM("64GB").
		SetSSD("2TB").
		Build()
	pc.Specs()
}
