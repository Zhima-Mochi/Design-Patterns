package main

import (
	"intro/duck"
	"intro/duck/fly"
)

func main() {
	mallard := duck.NewMallardDuck()
	mallard.Display()
	mallard.PerformQuack()
	mallard.PerformFly()

	model := duck.NewModelDuck()
	model.Display()
	model.PerformFly()
	model.SetFlyingBehaviour(&fly.RocketPowered{})
	model.PerformFly()
}
