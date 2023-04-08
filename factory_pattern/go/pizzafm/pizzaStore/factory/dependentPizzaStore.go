package pizzastorefactory

import "github.com/Zhima-Mochi/Design-Patterns/factory_pattern/go/pizzafm/pizza"

type DependentPizzaStore struct {
}

func NewDependentPizzaStore() *DependentPizzaStore {
	return &DependentPizzaStore{}
}

func (s *DependentPizzaStore) CreatePizza(pizzaStyle, pizzaType string) pizza.Pizza {
	var piz pizza.Pizza
	if pizzaStyle == "NY" {
		if pizzaType == "cheese" {
			piz = pizza.NewNYStyleCheesePizza()
		} else if pizzaType == "veggie" {
			piz = pizza.NewNYStyleVeggiePizza()
		} else if pizzaType == "clam" {
			piz = pizza.NewNYStyleClamPizza()
		} else if pizzaType == "pepperoni" {
			piz = pizza.NewNYStylePepperoniPizza()
		}
	} else if pizzaStyle == "Chicago" {
		if pizzaType == "cheese" {
			piz = pizza.NewChicagoStyleCheesePizza()
		} else if pizzaType == "veggie" {
			piz = pizza.NewChicagoStyleVeggiePizza()
		} else if pizzaType == "clam" {
			piz = pizza.NewChicagoStyleClamPizza()
		} else if pizzaType == "pepperoni" {
			piz = pizza.NewChicagoStylePepperoniPizza()
		}
	} else {
		return nil
	}
	piz.Prepare()
	piz.Bake()
	piz.Cut()
	piz.Box()
	return piz
}
