package main

import "fmt"

type Light struct {
	location string
	level    int
}

func NewLight(location string) *Light {
	return &Light{
		location: location,
		level:    0,
	}
}

func (l *Light) On() {
	l.level = 100
	fmt.Printf("Light is on\n")
}

func (l *Light) Off() {
	l.level = 0
	fmt.Printf("Light is off\n")
}

func (l *Light) Dim(level int) {
	l.level = level
	if l.level == 0 {
		l.Off()
	} else {
		fmt.Printf("Light is dimmed to %d%%\n", l.level)
	}
}

func (l *Light) GetLevel() int {
	return l.level
}
