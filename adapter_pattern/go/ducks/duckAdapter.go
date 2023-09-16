package main

import (
	"math/rand"
)

var _ Turkey = (*DuckAdapter)(nil)

type DuckAdapter struct {
	duck Duck
	rand *rand.Rand
}

func (a *DuckAdapter) Gobble() {
	a.duck.Quack()
}

func (a *DuckAdapter) Fly() {
	if a.rand.Intn(5) == 0 {
		a.duck.Fly()
	}
}
