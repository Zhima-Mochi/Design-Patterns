package gumballState

type GumballMachine interface {
	SetState(state State)
	GetHasQuarterState() State
}
