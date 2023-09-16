package main

var _ Turkey = (*WildTurkey)(nil)

type WildTurkey struct {
}

func (t *WildTurkey) Gobble() {
	println("Gobble gobble")
}

func (t *WildTurkey) Fly() {
	println("I'm flying a short distance")
}
