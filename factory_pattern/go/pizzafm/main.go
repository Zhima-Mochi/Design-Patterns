package main

import (
	"fmt"

	pizzastore "github.com/Zhima-Mochi/Design-Patterns/factory_pattern/go/pizzafm/pizzaStore"
	pizzastorefactory "github.com/Zhima-Mochi/Design-Patterns/factory_pattern/go/pizzafm/pizzaStore/factory"
)

func main() {
	var nyStore, chicagoStore pizzastore.PizzaStore
	nyStore = pizzastorefactory.NewNYPizzaStore()
	chicagoStore = pizzastorefactory.NewChicagoPizzaStore()

	pizza := nyStore.OrderPizza("cheese")
	fmt.Println("Ethan ordered a " + pizza.GetName())

	pizza = chicagoStore.OrderPizza("cheese")
	fmt.Println("Joel ordered a " + pizza.GetName())

	pizza = nyStore.OrderPizza("clam")
	fmt.Println("Ethan ordered a " + pizza.GetName())

	pizza = chicagoStore.OrderPizza("clam")
	fmt.Println("Joel ordered a " + pizza.GetName())

	pizza = nyStore.OrderPizza("pepperoni")
	fmt.Println("Ethan ordered a " + pizza.GetName())

	pizza = chicagoStore.OrderPizza("pepperoni")
	fmt.Println("Joel ordered a " + pizza.GetName())

	pizza = nyStore.OrderPizza("veggie")
	fmt.Println("Ethan ordered a " + pizza.GetName())

	pizza = chicagoStore.OrderPizza("veggie")
	fmt.Println("Joel ordered a " + pizza.GetName())
}
