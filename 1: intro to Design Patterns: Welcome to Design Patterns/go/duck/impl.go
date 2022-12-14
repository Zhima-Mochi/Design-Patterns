package duck

import (
	"fmt"
	"intro/duck/behavior/fly"
	"intro/duck/behavior/quack"
)

func NewMallardDuck() *Duck {
	return &Duck{
		Display: func() {
			fmt.Println("I’m a real Mallard duck")
		},
		flyingBehavior:   &fly.WithWings{},
		quackingBehavior: &quack.Quack{},
	}
}

func NewModelDuck() *Duck {
	return &Duck{
		Display: func() {
			fmt.Println("I’m a real Model duck")
		},
		flyingBehavior:   &fly.NoWay{},
		quackingBehavior: &quack.MuteQuack{},
	}
}
