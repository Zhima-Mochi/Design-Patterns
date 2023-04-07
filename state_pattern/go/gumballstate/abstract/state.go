package abstract

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
	Refill()
}
