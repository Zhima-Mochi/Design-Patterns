package pizzafactory

import pizza "github.com/Zhima-Mochi/Design-Patterns/factory_pattern/go/pizzas/pizza"

type SimplePizzaFactory struct {
}

func (s *SimplePizzaFactory) CreatePizza(pizzaType string) pizza.Pizza {
	switch pizzaType {
	case "cheese":
		return pizza.NewCheesePizza()
	case "pepperoni":
		return pizza.NewPepperoniPizza()
	case "clam":
		return pizza.NewClamPizza()
	case "veggie":
		return pizza.NewVeggiePizza()
	default:
		return nil
	}
}

func NewSimplePizzaFactory() *SimplePizzaFactory {
	return &SimplePizzaFactory{}
}
