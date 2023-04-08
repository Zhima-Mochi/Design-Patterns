package pizza

import "fmt"

type ChicagoStyleClamPizza struct {
	PizzaBase
}

func NewChicagoStyleClamPizza() *ChicagoStyleClamPizza {
	return &ChicagoStyleClamPizza{
		PizzaBase{
			Name:  "Chicago Style Clam Pizza",
			Dough: "Extra Thick Crust Dough",
			Sauce: "Plum Tomato Sauce",
			Toppings: []string{
				"Shredded Mozzarella Cheese",
				"Frozen Clams from Chesapeake Bay",
			},
		},
	}
}

func (p *ChicagoStyleClamPizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}
