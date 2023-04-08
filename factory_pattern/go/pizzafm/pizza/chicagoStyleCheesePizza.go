package pizza

import "fmt"

type ChicagoStyleCheesePizza struct {
	PizzaBase
}

func NewChicagoStyleCheesePizza() *ChicagoStyleCheesePizza {
	return &ChicagoStyleCheesePizza{
		PizzaBase{
			Name:  "Chicago Style Deep Dish Cheese Pizza",
			Dough: "Extra Thick Crust Dough",
			Sauce: "Plum Tomato Sauce",
			Toppings: []string{
				"Shredded Mozzarella Cheese",
			},
		},
	}
}

func (p *ChicagoStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}
