package liskov_substitution

import "fmt"

type GoodBird interface {
	Walk()
}

type FlyingBird interface {
	GoodBird
	Fly()
}

type Eagle struct{}

func (e *Eagle) Walk() {
	fmt.Println("Eagle is walking")
}

func (e *Eagle) Fly() {
	fmt.Println("Eagle is flying")
}

type Kiwi struct{}

func (k *Kiwi) Walk() {
	fmt.Println("Kiwi is walking")
}

func MakeBirdFly(data FlyingBird) {
	data.Fly()
}

func MakeBirdWalk(data GoodBird) {
	data.Walk()
}

func GoodViskov() {
	e := &Eagle{}
	MakeBirdWalk(e)
	MakeBirdFly(e)

	k := &Kiwi{}
	MakeBirdWalk(k)
	// MakeBirdFly(k)
}
