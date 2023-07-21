package main

import "fmt"

// Light is a struct that has a name.
type Light struct {
	Name string
}

// NewLight is a function that returns a new Light.
func NewLight(name string) *Light {
	return &Light{Name: name}
}

// On is a method of Light that prints the name of the light and the action.
func (l *Light) On() {
	fmt.Printf("%s light is on\n", l.Name)
}

// Off is a method of Light that prints the name of the light and the action.
func (l *Light) Off() {
	fmt.Printf("%s light is off\n", l.Name)
}
