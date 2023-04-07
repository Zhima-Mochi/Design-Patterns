package gumballState

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
	Refill()
}
