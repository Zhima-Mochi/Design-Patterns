package main

import (
	"github.com/Zhima-Mochi/Design-Patterns/strategy_pattern/go/duck/behavior/fly"

	"github.com/Zhima-Mochi/Design-Patterns/strategy_pattern/go/duck/duck"
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
