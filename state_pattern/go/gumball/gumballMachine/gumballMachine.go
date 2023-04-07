package gumballMachine

import (
	"github.com/Zhima-Mochi/Design-Patterns/state_pattern/go/gumball/gumballState"
)

type GumballMachine struct {
	noQuarterState  gumballState.State
	hasQuarterState gumballState.State
	state           gumballState.State
	count           int
}

func NewGumballMachine(numberGumballs int) *GumballMachine {
	gumballMachine := &GumballMachine{count: numberGumballs}
	gumballMachine.noQuarterState = gumballState.NewNoQuarterState(gumballMachine)
	return gumballMachine
}

func (gm *GumballMachine) SetState(state gumballState.State) {
	gm.state = state
}

func (gm *GumballMachine) GetHasQuarterState() gumballState.State {
	return gm.hasQuarterState
}
