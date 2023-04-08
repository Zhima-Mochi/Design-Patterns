package pizzastorefactory

import (
	"fmt"

	"github.com/Zhima-Mochi/Design-Patterns/factory_pattern/go/pizzafm/pizza"
)

type NYPizzaStore struct {
}

func NewNYPizzaStore() *NYPizzaStore {
	return &NYPizzaStore{}
}

func (n *NYPizzaStore) CreatePizza(pizzaType string) pizza.Pizza {
	switch pizzaType {
	case "cheese":
		return pizza.NewNYStyleCheesePizza()
	case "veggie":
		return pizza.NewNYStyleVeggiePizza()
	case "clam":
		return pizza.NewNYStyleClamPizza()
	case "pepperoni":
		return pizza.NewNYStylePepperoniPizza()
	default:
		return nil
	}
}

func (n *NYPizzaStore) OrderPizza(pizzaType string) pizza.Pizza {
	pizza := n.CreatePizza(pizzaType)
	fmt.Println("--- Making a ", pizza.GetName(), " ---")
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}
