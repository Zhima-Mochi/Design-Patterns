package pizzastore

import "github.com/Zhima-Mochi/Design-Patterns/factory_pattern/go/pizzafm/pizza"

type PizzaStore interface {
	CreatePizza(pizzaType string) pizza.Pizza
	OrderPizza(pizzaType string) pizza.Pizza
}
