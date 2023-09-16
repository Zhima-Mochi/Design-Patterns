package main

var _ Duck = (*MallardDuck)(nil)

type MallardDuck struct {
}

func (d *MallardDuck) Quack() {
	println("Quack")
}

func (d *MallardDuck) Fly() {
	println("I'm flying")
}
