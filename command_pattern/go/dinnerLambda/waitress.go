package main

type Waitress struct {
}

func NewWaitress() *Waitress {
	return &Waitress{}
}

func (w *Waitress) TakeOrder(order Order) {
	order()
}
