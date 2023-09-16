package main

import "math/rand"

func testDuck(duck Duck) {
	duck.Quack()
	duck.Fly()
}

func main() {
	duck := &MallardDuck{}
	turkey := &WildTurkey{}

	turkeyAdapter := &TurkeyAdapter{turkey: turkey}

	println("The Turkey says...")
	turkey.Gobble() // Gobble gobble
	turkey.Fly()    // I'm flying a short distance

	println("\nThe Duck says...")
	testDuck(duck)
	// Quack
	// I'm flying

	println("\nThe TurkeyAdapter says...")
	testDuck(turkeyAdapter)
	// Gobble gobble
	// I'm flying a short distance
	// I'm flying a short distance
	// I'm flying a short distance
	// I'm flying a short distance
	// I'm flying a short distance

	duckAdapter := &DuckAdapter{duck: duck, rand: rand.New(rand.NewSource(0))}

	for i := 0; i < 10; i++ {
		println("\nThe DuckAdapter says...")
		duckAdapter.Gobble()
		duckAdapter.Fly()
		// Quack
		// I'm flying
	}

	drone := &SuperDrone{}
	droneAdapter := &DroneAdapter{drone: drone}
	testDuck(droneAdapter)
	// Beep beep beep
	// Rotors are spinning
	// Taking off
}
