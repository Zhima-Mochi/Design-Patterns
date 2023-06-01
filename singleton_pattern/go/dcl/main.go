package main

import (
	"fmt"

	"github.com/Zhima-Mochi/Design-Patterns/singleton_pattern/go/dcl/singleton"
)

func main() {
	singleton := singleton.GetInstance()
	fmt.Println(singleton.GetDescription())
}
