package concrete

import "fmt"

type NoQuarterState struct {
	gumballMachine *GumballMachine
}

func NewNoQuarterState(gumballMachine *GumballMachine) *NoQuarterState {
	return &NoQuarterState{gumballMachine: gumballMachine}
}

func (nq *NoQuarterState) InsertQuarter() {
	fmt.Println("You inserted a quarter")
	nq.gumballMachine.SetState(nq.gumballMachine.GetHasQuarterState())
}

func (nq *NoQuarterState) EjectQuarter() {
	fmt.Println("You haven't inserted a quarter")
}

func (nq *NoQuarterState) TurnCrank() {
	fmt.Println("You turned, but there's no quarter")
}

func (nq *NoQuarterState) Dispense() {
	fmt.Println("You need to pay first")
}

func (nq *NoQuarterState) Refill() {

}

func (nq *NoQuarterState) String() string {
	return "waiting for quarter"
}
