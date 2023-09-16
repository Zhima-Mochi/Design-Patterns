package main

var _ Drone = (*SuperDrone)(nil)

type SuperDrone struct {
}

func (d *SuperDrone) Beep() {
	println("Beep beep beep")
}

func (d *SuperDrone) SpinRotors() {
	println("Rotors are spinning")
}

func (d *SuperDrone) TakeOff() {
	println("Taking off")
}
