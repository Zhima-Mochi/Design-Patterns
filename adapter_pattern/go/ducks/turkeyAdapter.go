package main

var _ Duck = (*TurkeyAdapter)(nil)

type TurkeyAdapter struct {
	turkey Turkey
}

func (a *TurkeyAdapter) Quack() {
	a.turkey.Gobble()
}

func (a *TurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		a.turkey.Fly()
	}
}
