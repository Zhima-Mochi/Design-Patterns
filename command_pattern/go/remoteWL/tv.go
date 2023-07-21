package main

import "fmt"

type TV struct {
	location string
	channel  int
}

func NewTV(location string) *TV {
	return &TV{
		location: location,
		channel:  0,
	}
}

func (tv *TV) On() {
	println("TV is on")
}

func (tv *TV) Off() {
	println("TV is off")
}

func (tv *TV) SetInputChannel() {
	tv.channel = 3
}

func (tv *TV) GetChannel() int {
	return tv.channel
}

func (tv *TV) String() string {
	return fmt.Sprintf("%s TV is on channel %d", tv.location, tv.channel)
}
