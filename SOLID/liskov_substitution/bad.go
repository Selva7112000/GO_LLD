package liskov_substitution

import "fmt"

type BadBird interface {
	Fly()
}

type Sparrow struct{}

func (s *Sparrow) Fly() {
	fmt.Println("Sparrow is flying")
}

type Ostrich struct{}

func (o *Ostrich) Fly() {
	panic("Ostrich can't fly")
}

func BadMakeBirdFly(b BadBird) {
	b.Fly()
}

func BadLiskov() {
	s := &Sparrow{}
	BadMakeBirdFly(s)

	o := &Ostrich{}
	BadMakeBirdFly(o)
}
