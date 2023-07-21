package main

import "fmt"

type HotTub struct {
	on    bool
	heat  int
	level int
}

func NewHotTub() *HotTub {
	return &HotTub{
		on:    false,
		heat:  0,
		level: 0,
	}
}

func (ht *HotTub) On() {
	ht.on = true
}

func (ht *HotTub) Off() {
	ht.on = false
}

func (ht *HotTub) BubblesOn() {
	if ht.on {
		ht.level = 1
	}
}

func (ht *HotTub) BubblesOff() {
	if ht.on {
		ht.level = 0
	}
}

func (ht *HotTub) JetsOn() {
	if ht.on {
		ht.level = 2
	}
}

func (ht *HotTub) JetsOff() {
	if ht.on {
		ht.level = 0
	}
}

func (ht *HotTub) SetTemperature(heat int) {
	if ht.on {
		ht.heat = heat
	}
}

func (ht *HotTub) Heat() {
	if ht.on {
		ht.heat = 105
	}
}

func (ht *HotTub) Cool() {
	if ht.on {
		ht.heat = 98
	}
}

func (ht *HotTub) GetTemperature() int {
	return ht.heat
}

func (ht *HotTub) String() string {
	if ht.on {
		return fmt.Sprintf("Hot tub is on and temperature is %d degrees", ht.heat)
	}
	return "Hot tub is off"
}
