package duck

import (
	"fmt"

	"github.com/Zhima-Mochi/Design-Patterns/strategy_pattern/go/duck/behavior"
)

type Duck struct {
	Display func()
	// behavior is variable
	flyingBehavior   behavior.Fly
	quackingBehavior behavior.Quack
}

func (d *Duck) PerformFly() {
	d.flyingBehavior.Fly()
}
func (d *Duck) PerformQuack() {
	d.quackingBehavior.Quack()
}

// all duck can swim
func (d *Duck) Swim() {
	fmt.Println("All ducks float, even decoys!")
}

func (d *Duck) SetFlyingBehaviour(b behavior.Fly) {
	d.flyingBehavior = b
}

func (d *Duck) SetQuackingBehaviour(b behavior.Quack) {
	d.quackingBehavior = b
}
