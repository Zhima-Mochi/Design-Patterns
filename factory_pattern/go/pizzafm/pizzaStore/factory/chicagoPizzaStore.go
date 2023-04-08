package pizzastorefactory

import (
	"fmt"

	"github.com/Zhima-Mochi/Design-Patterns/factory_pattern/go/pizzafm/pizza"
)

type ChicagoPizzaStore struct {
}

func NewChicagoPizzaStore() *ChicagoPizzaStore {
	return &ChicagoPizzaStore{}
}

func (c *ChicagoPizzaStore) CreatePizza(pizzaType string) pizza.Pizza {
	switch pizzaType {
	case "cheese":
		return pizza.NewChicagoStyleCheesePizza()
	case "veggie":
		return pizza.NewChicagoStyleVeggiePizza()
	case "clam":
		return pizza.NewChicagoStyleClamPizza()
	case "pepperoni":
		return pizza.NewChicagoStylePepperoniPizza()
	default:
		return nil
	}
}

func (c *ChicagoPizzaStore) OrderPizza(pizzaType string) pizza.Pizza {
	pizza := c.CreatePizza(pizzaType)
	fmt.Println("--- Making a ", pizza.GetName(), " ---")
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}
