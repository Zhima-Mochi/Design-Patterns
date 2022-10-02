package quack

import "fmt"

type Quack struct {
}

func (q *Quack) Quack() {
	fmt.Println("Quack")
}

type MuteQuack struct{}

func (mq *MuteQuack) Quack() {
	fmt.Println("Silence")
}

type Squeak struct{}

func (s *Squeak) Quack() {
	fmt.Println("Squeak")
}
