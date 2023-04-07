package duck

import (
	"fmt"

	"github.com/Zhima-Mochi/Design-Patterns/strategy_pattern/go/duck/behavior/fly"
	"github.com/Zhima-Mochi/Design-Patterns/strategy_pattern/go/duck/behavior/quack"
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
