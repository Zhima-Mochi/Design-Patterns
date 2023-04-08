package main

import (
	pizzafactory "github.com/Zhima-Mochi/Design-Patterns/factory_pattern/go/pizzas/pizzaFactory"
	pizzastore "github.com/Zhima-Mochi/Design-Patterns/factory_pattern/go/pizzas/pizzaStore"
)

func main() {
	factory := pizzafactory.NewSimplePizzaFactory()
	store := pizzastore.NewPizzaStore(factory)

	pizza := store.OrderPizza("cheese")
	println("We ordered a " + pizza.GetName())
	println(pizza.String())

	pizza = store.OrderPizza("veggie")
	println("We ordered a " + pizza.GetName())
	println(pizza.String())
}
