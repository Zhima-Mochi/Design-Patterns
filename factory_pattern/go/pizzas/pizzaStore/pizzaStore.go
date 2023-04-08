package pizzastore

import (
	"github.com/Zhima-Mochi/Design-Patterns/factory_pattern/go/pizzas/pizza"
	pizzafactory "github.com/Zhima-Mochi/Design-Patterns/factory_pattern/go/pizzas/pizzaStore/factory"
)

type PizzaStore struct {
	factory *pizzafactory.SimplePizzaFactory
}

func NewPizzaStore(factory *pizzafactory.SimplePizzaFactory) *PizzaStore {
	return &PizzaStore{factory: factory}
}

func (s *PizzaStore) OrderPizza(pizzaType string) pizza.Pizza {
	pizza := s.factory.CreatePizza(pizzaType)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}
