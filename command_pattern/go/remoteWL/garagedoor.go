package main

type GarageDoor struct {
	location string
}

func NewGarageDoor(location string) *GarageDoor {
	return &GarageDoor{
		location: location,
	}
}

func (gd *GarageDoor) Up() {
	println(gd.location + " garage door is open")
}

func (gd *GarageDoor) Down() {
	println(gd.location + " garage door is closed")
}

func (gd *GarageDoor) Stop() {
	println(gd.location + " garage door is stopped")
}

func (gd *GarageDoor) LightOn() {
	println(gd.location + " garage light is on")
}

func (gd *GarageDoor) LightOff() {
	println(gd.location + " garage light is off")
}

func (gd *GarageDoor) String() string {
	return gd.location + " garage door"
}
