package fly

import "fmt"

type NoWay struct {
}

func (fnw *NoWay) Fly() {
	fmt.Println("I can't fly")
}

type WithWings struct {
}

func (fww *WithWings) Fly() {
	fmt.Println("I'm flying")
}

type RocketPowered struct {
}

func (frp *RocketPowered) Fly() {
	fmt.Println("I'm flying with a rocket!")
}
