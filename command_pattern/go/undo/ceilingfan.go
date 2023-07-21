package main

import "fmt"

const (
	HIGH   = 3
	MEDIUM = 2
	LOW    = 1
	OFF    = 0
)

type CeilingFan struct {
	location string
	speed    int
}

func NewCeilingFan(location string) *CeilingFan {
	return &CeilingFan{
		location: location,
		speed:    OFF,
	}
}

func (c *CeilingFan) High() {
	c.speed = HIGH
	fmt.Printf("%s ceiling fan is on high\n", c.location)
}

func (c *CeilingFan) Medium() {
	c.speed = MEDIUM
	fmt.Printf("%s ceiling fan is on medium\n", c.location)
}

func (c *CeilingFan) Low() {
	c.speed = LOW
	fmt.Printf("%s ceiling fan is on low\n", c.location)
}

func (c *CeilingFan) Off() {
	c.speed = OFF
	fmt.Printf("%s ceiling fan is off\n", c.location)
}

func (c *CeilingFan) GetSpeed() int {
	return c.speed
}
