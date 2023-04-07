package concrete

import "gumballstate/abstract"

type GumballMachine struct {
	noQuarterState  abstract.State
	hasQuarterState abstract.State
	state           abstract.State
	count           int
}

func NewGumballMachine(numberGumballs int) *GumballMachine {
	gumballMachine := &GumballMachine{count: numberGumballs}
	gumballMachine.noQuarterState = NewNoQuarterState(gumballMachine)
	return gumballMachine
}

func (gm *GumballMachine) SetState(state abstract.State) {
	gm.state = state
}

func (gm *GumballMachine) GetHasQuarterState() abstract.State {
	return gm.hasQuarterState
}
