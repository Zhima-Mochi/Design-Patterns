package main

import "fmt"

type Cook struct {
}

func NewCook() *Cook {
	return &Cook{}
}

func (c *Cook) MakeBurger() {
	fmt.Println("Cook: Making burger")
}

func (c *Cook) MakeFries() {
	fmt.Println("Cook: Making fries")
}
