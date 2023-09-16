package main

var _ Duck = (*DroneAdapter)(nil)

type DroneAdapter struct {
	drone Drone
}

func (a *DroneAdapter) Quack() {
	a.drone.Beep()
}

func (a *DroneAdapter) Fly() {
	a.drone.SpinRotors()
	a.drone.TakeOff()
}
